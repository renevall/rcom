package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	workers = 10
)

type EmailCleanerService struct {
	cleaner EmailCleaner
}

type EmailCleaner interface {
	BatchBlackList(email []string) error
}

func NewEmailCleanerService(postgresClient EmailCleaner) *EmailCleanerService {
	return &EmailCleanerService{
		cleaner: postgresClient,
	}
}

func (e *EmailCleanerService) CleanByFile(path string, mode string, batch int, line int) error {
	// open text file and stream over a channel each line
	file, err := e.OpenFile(path)
	if err != nil {
		return err
	}

	// use concurrency to process the data in batches of n lines

	if mode == "concurrent" {
		return e.ProcessDataConcurrent(file, batch, line)
	}

	return e.ProcessData(file, batch, line)
}

func (e *EmailCleanerService) OpenFile(path string) (*os.File, error) {
	// open the file and return an error if it fails

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil

}

// ProcessData will process the data in the file one line at a time.
func (e *EmailCleanerService) ProcessData(file io.Reader, batch int, line int) error {
	buf := bufio.NewReader(file)
	var emailBatch []string
	var linesProcessed int

	now := time.Now()

	for {
		line, err := buf.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		if line != "" {
			emailBatch = append(emailBatch, line)
			if len(emailBatch) == batch {
				err := e.cleaner.BatchBlackList(emailBatch)
				if err != nil {
					fmt.Printf("Error processing batch: %v", err)
					return err
				}
				linesProcessed += len(emailBatch)
				emailBatch = nil // reset the batch
			}

		}

		if err == io.EOF {
			break
		}
	}

	if len(emailBatch) > 0 {
		err := e.cleaner.BatchBlackList(emailBatch)
		if err != nil {
			fmt.Printf("Error processing batch: %v", err)
			return err
		}
		linesProcessed += len(emailBatch)
	}

	elapsed := time.Since(now)

	fmt.Printf("Processed %d lines in %s", linesProcessed, elapsed)

	return nil
}

func (e *EmailCleanerService) ProcessDataConcurrent(file io.Reader, batchSize int, line int) error {
	now := time.Now()
	// batchesCh will contains the batches of emails to be processed
	batchesCh := e.StreamFile(file, batchSize)
	results := make(chan int)

	var wg sync.WaitGroup
	var totalProcessed int

	// Start workers
	for i := 0; i < workers; i++ {
		fmt.Printf("Starting worker %d\n", i)
		wg.Add(1)
		go e.Worker(i, batchesCh, results, &wg)
	}

	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Close results channel to finish the results collection loop

	}()

	// Fan-in: Collect results and aggregate them

	for result := range results {
		totalProcessed += result
	}

	elapsed := time.Since(now)
	fmt.Printf("Processed %d lines concurrently in %s\n", totalProcessed, elapsed)

	return nil
}

// StreamFile will stream the content of a file over a channel of batches of emails
func (e *EmailCleanerService) StreamFile(file io.Reader, batchSize int) <-chan []string {
	buf := bufio.NewReader(file)
	out := make(chan []string)

	go func() {
		defer close(out)
		var emailBatch []string
		for {
			line, err := buf.ReadString('\n')
			if err != nil && err != io.EOF {
				fmt.Printf("Error reading file: %v\n", err)
				return
			}

			if line != "" {

				emailBatch = append(emailBatch, line)
				if len(emailBatch) == batchSize {
					out <- emailBatch
					emailBatch = []string{}
				}
			}

			if err == io.EOF {
				break
			}

		}

		// Send any remaining emails as a batch
		if len(emailBatch) > 0 {
			out <- emailBatch
		}

	}()

	return out
}

// Worker will process a batch of emails
func (e *EmailCleanerService) Worker(id int, batches <-chan []string, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for emails := range batches {
		err := e.cleaner.BatchBlackList(emails)
		if err != nil {
			fmt.Printf("Worker %d: Error processing batch: %v\n", id, err)
			continue // skip this batch on error, but continue processing
		}
		results <- len(emails) // send the count of processed emails back
	}
	fmt.Printf("Worker %d: Finished\n", id)
}
