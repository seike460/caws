package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCaws(t *testing.T) {
	caws := newCaws()
	setS3Service(caws)
	assert.NotNil(t, caws)
}
