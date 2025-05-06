package main

import (
	"fmt"
	"sync"
	"time"
)

// what if i did not take buffered channel

// what if i did not send job as an argument in the go routine

type Job int

func ProcessJob(job Job) error {
	time.Sleep(1 * time.Second)
	if job == 1 || job == 3 {
		return fmt.Errorf("error in job :%d", job)
	}
	fmt.Printf("Job Done %d\n", job)
	return nil
}
func mainx() {
	jobs := []Job{1, 2, 3, 4, 5}
	errChan := make(chan error, len(jobs)) // buffered because no one is reading because of wg.wait at line 35
	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		go func(job Job) {
			defer wg.Done()
			err := ProcessJob(job)
			errChan <- err
		}(job) // why sending job ? closure ?
	}
	wg.Wait()
	close(errChan) // why this
	for err := range errChan {
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("----------")
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("----------")
}
