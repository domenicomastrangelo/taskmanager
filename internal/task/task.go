package task

import (
	"log"
	"time"

	"github.com/domenicomastrangelo/taskmanager/internal/db"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type Tasks []Task

func List(maxResults int, minDate time.Time, done bool) Tasks {
	rows, err := db.DB.Query("select * from tasks where done = ? and created_at > ? limit ?", done, minDate, maxResults)

	if err != nil {
		log.Println(err.Error())
		return Tasks{}
	}

	var tasks Tasks

	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Message, &task.Done, &task.CreatedAt)

		if err != nil {
			log.Println(err.Error())
			continue
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func (t *Task) Add() bool {
	stmt, err := db.DB.Prepare("insert into tasks(title, message, done, created_at) values(?, ?, ?, ?)")

	if err != nil {
		log.Println(err.Error())
		return false
	}

	res, err := stmt.Exec(t.Title, t.Message, false, time.Now().UTC())

	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = res.LastInsertId()

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func (t *Task) Delete() {

}

func (t *Task) Save() {

}

func (t *Task) Update() {

}

func (t *Task) SetDone() bool {
	stmt, err := db.DB.Prepare("update tasks set done = ? where id = ?")

	if err != nil {
		log.Println(err.Error())
		return false
	}

	res, err := stmt.Exec(t.Done, t.ID)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	_, err = res.RowsAffected()

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
