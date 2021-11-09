/*
给定一系列的task,要求在规定的timeout里跑完,不然就报错,如果操作系统给了中断信号,也报错
*/
package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("cannot finish tasks within the timeout")
	ErrInterrupt = errors.New("received interrupt from OS")
)

type Runner struct {
	interrupt chan os.Signal // 模拟系统中断
	complete  chan error
	timeout   <-chan time.Time // 计时
	tasks     []func(int)      // task列表
}

func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t),
		tasks:     make([]func(int), 0),
	}
}

func (r *Runner) AddTasks(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil
}

func (r *Runner) Start() error {
	// relay interrupt from OS
	signal.Notify(r.interrupt, os.Interrupt)
	// run the task
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("Task complete #%d\n", id)
	}
}

func main() {
	r := New(4 * time.Second)

	r.AddTasks(createTask(), createTask(), createTask())

	err := r.Start()
	switch err {
	case ErrInterrupt:
		fmt.Println("tasks interrupt")
	case ErrTimeout:
		fmt.Println("tasks timeout")
	default:
		fmt.Println("all tasks finished")
	}
}
