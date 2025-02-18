package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface { // Job 인터페이스
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index) // 각 작업
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ { // 작업 할당
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i] // 각 작업을 Go 루틴으로 실행
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
