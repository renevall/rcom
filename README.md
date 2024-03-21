## Problem #1 - Adapter Pattern

### Situation

Laughing Man System has a notification system that allows the clients to define how they want to be notified of relevant events. The supported channels of notification would be:

- Email
- SMS
- Slack

### Task

We need to design and implement a prototype notification service that allows us to send the notification to the target as defined by the client. Because we expect to keep adding delivery channels, we need to make sure the code will be easily extended with future integrations. We also need to track which channel drives better engagement.

### Action

For this task, we will create a CLI to simulate both a request and a server response to that request. This would allow us to quickly simulate what the real implementation would be doing.

We will use a notification package to group up the implementation for each kind of delivery channel.

There will be a service package in charge of dealing with the use case at hand, it will adopt a general input and use the notification channel for final delivery

### Result

The expectation is to be able to run the following commands to send notifications over the desired channels:

- `rcom send -t foo@bar.com -m "your account was just approved!" -c email`
- `rcom send -t 505-9045845 -m "someone is waiting for you in the lobby" -c sms`
- `rcom send -t c3245/databases -m "your database is over 80%" -c slack`

### Commentary

Like many real-world problems, you may have 2 or more classes/structs/integrations that essentially do the same thing but their internals make them incompatible. Making a god class would make the code harder to maintain and extend. This is where the adapter pattern is handly and Go's powerful interfaces allow us to model the feature easily. In this example, we abstracted the send method which is the essence of a message delivery system.

The code I provide is a simulation. In a production environment, we would be looking more into error handling, essential things such as retry, more code coverage, and integration tests. The logging also could be done better, and we still need to add monitoring codes such as Prometheus and open tracing.

Regardless of the improvements needed, the simulator is a valid approach to understanding the value of the adapter and the power of implementing it in an easy and manageable way.

## Problem #2 - Concurrency

### Situation

Our email-sending provider has gathered 500,000 emails that are bouncing for different reasons. To protect our public record in the context of email sending, we should blacklist those emails/alerts so that we don't continue to send them

### Task

We need to update our subscription service so that we no longer send the alerts to those emails that have been bouncing. Keep should avoid or minimize any downtime.

### Action

For this task, we created 2 implementations of the handling of the case. One sequential and one concurrent. Normally a task like this would be done using batches of of updates since one by one would be too slow.

The sequential by default uses batches of 2000, taking a total time of 12s. The concurrent uses the same batches of 2000 but uses 10 workers to take advantage of the natural db parallelism. The run time for the concurrent was 1.5 seconds

### Result

500,000 emails were blacklisted from our subscription system, protecting our email reputation since bounces count as a negative towards systems tracking reputation, say google itself.

For the test, we show the effect concurrency can have on performance. However, it should be done when it makes sense.

The CLI could be run by:

- `rcom blacklist # sequencial`
- `rcom blacklist -m concurrent`

### Commentary

Choosing concurrency for this task was a stretch since the difference between the 2 approaches is immaterial for a one-time task.

But, there are many cases in the real world where you better use this approach(fan-out -> fan-in). For example, a system where you have a 1-hour window for maintenance once a month. In this case, maybe sequential would take 12 hours to run, but by parallelizing it takes you 40 minutes.

Another example is when you have external latency affecting the execution of a long-running process. A web crawler or spider would be a common use case. Instead of crawling one site at a time, you can speed up the main process by processing two sites at a time.

As far as testing, and concerns such as deadlocks, the reality is it is hard to test concurrent code. For this reason, the biggest win is to separate the concurrent construct from specific executions such as writing to the database. That way you can test the database part without any concurrency concern.

Even before testing concurrency, some techniques help you avoid the feared deadlocks and leaks:

- always know when a routine will end
- close channels by producers and not consumers
- wrap use of mutexes in dedicated functions instead of in between functional code

Finally when you do a test for concurrency:

- avoid using sleep to fake a time-lapse. This makes the test flaky
- don't use buffered channels in your test, this more often than not would hide a problem.
- test with higher level functions than the specific one. Chances are it would be easier to verify behavior if you focus on a bigger piece of code
- Run your test with a race detection flag

Overall, practice and focus on good ones will help you write safe concurrent code. But never be afraid of leaving it sequential. There is nothing wrong with that and at times it can even outperform concurrent one!

> The implemented code could break down the implementation a bit more to improve testability.
> The code could print the last line correctly processed to be able to resume in case of a problem updating the database. We avoid starting from scratch in such cases.

## Problem #3 - Open Ended

For this last test I solved the common code challenge "find the max sum values of k consecutive numbers in a given array of n elements". This problem is a rolling window pattern problem and a real-world use case of it is for example calculating the hourly moving average expense reported by a z system, based on the average an alarm would trigger if the threshold is reached.

### To run

The CLI could be run by:

- `rcom max # defaults`
- `rcom max -k 4 "34,22,56,1,32,12,1,22,7"`

