package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	log := NewTestLogger()
	log.Info("Test")

	assert.EqualValues(t, []string{"Test"}, log.Messages)
}
