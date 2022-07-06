package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {
	db := Connect()

	assert.NotNil(t, db)
}
