package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		t.FailNow()
	}

	db := Connect(homeDir+"/.taskmanager/", "dbTest.sqlite")

	assert.NotNil(t, db)
}
