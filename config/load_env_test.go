package config

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	err := errors.New("test")
	// trigger error
	GetConfiguration(err)
	assert.Equal(t, os.Getenv("DB_NAME"),"test_database")
}
