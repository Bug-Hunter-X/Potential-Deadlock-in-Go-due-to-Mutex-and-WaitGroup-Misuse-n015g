```go
func main() {
    var m sync.Mutex
    wg := sync.WaitGroup{}
    wg.Add(1)
    go func() {
        defer wg.Done()
        m.Lock()
        fmt.Println("Goroutine 1 locked the mutex")
        time.Sleep(2 * time.Second)
        m.Unlock()
    }()
    wg.Wait()
}