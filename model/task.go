package model

import (
	"encoding/json"
	"ls-kh-addtask/utils"
	"time"
)

type Task struct {
	TaskID  string   `json:"task_id"`
	Topic   string   `json:"topic"`
	User    string   `json:"user"`
	UserId  string   `json:"user_id"`
	Task    string   `json:"task"`
	Objects []string `json:"objects"`

	CreateAt string `json:"create_at"`
}

func NewTask(username, topic, userId, task string, objects []string) Task {
	return Task{
		TaskID:   utils.GetTaskId(),
		Topic:    topic,
		User:     username,
		UserId:   userId,
		Task:     task,
		Objects:  objects,
		CreateAt: time.Now().Format("2006-01-02 15:04-05"),
	}
}

func Json2Task(s string) Task {
	var t = Task{}
	_ = json.Unmarshal([]byte(s), &t)
	return t
}

func Task2Json(t Task) string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}

func (t Task) Key() string {
	return t.Topic
}
