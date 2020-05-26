package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
    msg := SayHello()
	assert.Equal(t, msg, "Hello")


}

