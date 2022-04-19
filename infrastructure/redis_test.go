package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenRedis(t *testing.T) {

	redis := OpenRedis()

	assert.NotEmpty(t, redis)
}
