package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/beego/beego/v2/adapter/toolbox"
)

type schedule struct {
	Name     string
	Tasklist map[string]*toolbox.Task
	SpecStr  string
	Run      bool
}

func newScheduleMgmt(name string) *schedule {
	return &schedule{
		Name:     name,
		Tasklist: make(map[string]*toolbox.Task),
		Run:      false,
	}
}

func (s *schedule) NewSchedule(taskName, schedule string, taskFunc toolbox.TaskFunc) {
	s.SpecStr = schedule
	task := toolbox.NewTask(taskName, s.Parser(schedule), taskFunc)
	s.Tasklist[taskName] = task
	toolbox.AddTask(taskName, task)
	fmt.Println(task.GetNext())
}

func (s *schedule) StartSchedule() {
	s.Run = true
	toolbox.StartTask()
}

func (s *schedule) StopSchedule() {
	s.Run = false
	toolbox.StopTask()
}

func (s *schedule) UpdateSchedule(taskName, schedule string, taskFunc toolbox.TaskFunc, enable bool) {
	if s.Tasklist[taskName] == nil {
		return
	}

	if s.SpecStr == schedule && &s.Tasklist[taskName].DoFunc == &taskFunc && s.Run == enable {
		return
	}

	if s.Run {
		toolbox.StopTask()
	}

	if s.SpecStr != schedule || &s.Tasklist[taskName].DoFunc != &taskFunc {
		toolbox.DeleteTask(taskName)
		s.NewSchedule(taskName, schedule, taskFunc)

	}

	if enable {
		toolbox.StartTask()
	}
}

func (s *schedule) Parser(task string) string {
	splits := strings.Split(task, " ")
	switch len(splits) {
	case 4:
		return fmt.Sprintf("0 0 %s %s %s %s ", splits[0], splits[1], splits[2], splits[3])
	case 5:
		return fmt.Sprintf("0 %s %s %s %s %s ", splits[0], splits[1], splits[2], splits[3], splits[4])
	case 6:
		return fmt.Sprintf("%s %s %s %s %s %s ", splits[0], splits[1], splits[2], splits[3], splits[4], splits[5])
	default:
		return "0/5 * * * * * "
	}
}

func (s *schedule) Print1() error {
	fmt.Print(time.Now())
	fmt.Println("  exec tasking 1")
	return nil
}

func (s *schedule) Print2() error {
	fmt.Print(time.Now())
	fmt.Println("  exec tasking 2")
	return nil
}
