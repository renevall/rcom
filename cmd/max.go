/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	k = 3
)

// maxCmd represents the max command
var maxCmd = &cobra.Command{
	Use:   "max",
	Short: "determine the max sum of k elements in an array of n elements",
	Long: `This command takes an array of n elements and a number k, and returns the maximum sum of k consecutive elements in the array.
If no array is provided, the default array is [2, 1, 5, 1, 3, 2] and the default k is 3.
Example usage:
max --k 3
max 2,1,5,1,3,2 --k 3
max -k 4 "34,22,56,1,32,12,1,22,7"`,
	Run: func(cmd *cobra.Command, args []string) {
		arr := []int{2, 1, 5, 1, 3, 2}
		var err error
		if len(args) == 1 {
			arr, err = parseArray(args[0])
			if err != nil {
				fmt.Println("Error parsing array:", err)
				return
			}
		}

		max := findMaxSumSubarray(arr, k)
		fmt.Printf("Maximum sum of a subarray of size %d for %v array is: %d \n", k, arr, max)

	},
}

func init() {
	rootCmd.AddCommand(maxCmd)

	maxCmd.Flags().IntVarP(&k, "k", "k", 3, "The number of elements to sum")
}

func parseArray(input string) ([]int, error) {
	var arr []int
	valuesStr := strings.Split(input, ",")
	for _, v := range valuesStr {
		v = strings.TrimSpace(v)
		value, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		arr = append(arr, value)
	}

	return arr, nil
}

func findMaxSumSubarray(arr []int, k int) int {
	var maxSum, windowSum int
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum = windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}
