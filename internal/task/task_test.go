package task

import (
	"os"
	"testing"
	"time"

	"github.com/domenicomastrangelo/taskmanager/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		t.FailNow()
	}

	db.Connect(homeDir+"/.taskmanager/", "dbTest.sqlite")
	tasks := List(10, time.Now().Add(time.Duration(-10*(int(time.Hour)*24))).UTC(), false)

	assert.IsType(t, Tasks{}, tasks)
}

func TestAdd(t *testing.T) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		t.FailNow()
	}

	db.Connect(homeDir+"/.taskmanager/", "dbTest.sqlite")
	task := Task{
		Title:     "This is the title",
		Message:   "This is the message",
		Done:      false,
		CreatedAt: time.Now().UTC(),
	}
	res := task.Add()

	assert.Equal(t, res, true)
}
