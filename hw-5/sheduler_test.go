package sheduler

import (
	"testing"
	"time"
)

func makeTaskFunc(waitFor int, err error) func() error {
	return func() error {
		time.Sleep(time.Duration(waitFor) * time.Second)
		return err
	}
}

const countTasks = 5

func TestSheduleTasks(t *testing.T) {
	tasks := make([]func() error, 0, countTasks)
	for i := 0; i < countTasks; i++ {
		tasks = append(tasks, makeTaskFunc(i+1, nil))
	}

	SheduleTasks(tasks, 3, 0)
}
