package task

import (
	"testing"
	"time"

	"github.com/domenicomastrangelo/taskmanager/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	db.Connect()
	tasks := List(10, time.Now().Add(time.Duration(-10*(int(time.Hour)*24))).UTC(), false)

	assert.IsType(t, Tasks{}, tasks)
}
