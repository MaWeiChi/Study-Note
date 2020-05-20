package a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateother(t *testing.T) {
	assert.Equal(t, Calculate(2), 3)
}
