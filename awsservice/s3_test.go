package awsservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewS3(t *testing.T) {

	s := NewS3()
	assert.NotNil(t, s)

}
