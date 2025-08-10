package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, tasks <-chan string, results *[]string, mutex *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range tasks {
        fmt.Printf("Worker %d started task: %s\n", id, task)
        time.Sleep(500 * time.Millisecond) // Simulate delay
        result := fmt.Sprintf("Processed by worker %d: %s", id, task)
        mutex.Lock()
        *results = append(*results, result)
        mutex.Unlock()
        fmt.Printf("Worker %d finished task: %s\n", id, task)
    }
}

func main() {
    taskChan := make(chan string)
    var results []string
    var mutex sync.Mutex
    var wg sync.WaitGroup

    // Start workers
    numWorkers := 4
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, taskChan, &results, &mutex, &wg)
    }

    // Send tasks
    for i := 1; i <= 10; i++ {
        taskChan <- fmt.Sprintf("Task-%d", i)
    }
    close(taskChan) // Close the channel

    wg.Wait()

    // Output results
    fmt.Println("\n=== Results ===")
    for _, result := range results {
        fmt.Println(result)
    }
}
