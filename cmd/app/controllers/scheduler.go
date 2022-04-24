package controllers

import (
	"time"

	"github.com/go-co-op/gocron"
)

var s = gocron.NewScheduler(time.UTC)

func notifyUser() {

	// fmt.Println("Task!")
}

func startScheduler() {
	// stime := time.Date(2022, time.April, 22, 7, 19, 0, 0, time.UTC)
	stime := time.Now()
	afterTime := stime.Add(time.Second * 10)
	job, _ := s.Every(1).Day().StartAt(afterTime).Do(notifyUser)
	job.LimitRunsTo(1)
	// s.Every(5).Seconds().Do(task)
	s.StartAsync()
}
