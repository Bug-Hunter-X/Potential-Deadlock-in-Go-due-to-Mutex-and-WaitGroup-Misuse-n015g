# Potential Deadlock in Go

This repository demonstrates a potential deadlock scenario in Go using `sync.Mutex` and `sync.WaitGroup`. The main goroutine waits for a child goroutine to complete, but this can lead to a deadlock if the mutex is not handled correctly.

## Bug Description
The bug arises due to a race condition.  The `WaitGroup` ensures the main goroutine waits for the child goroutine to complete; however, the child goroutine never gets a chance to execute due to the mutex being locked indefinitely. The child goroutine holds the mutex but it can't release it because the main goroutine never finishes waiting for the WaitGroup to be done, resulting in a deadlock situation.