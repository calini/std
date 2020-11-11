package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBefore(t *testing.T) {
	str := "A^2+B^2=C^2"
	exp := "A^2"
	actual := Before(str, "+")

	assert.Equal(t, exp, actual)
}

func TestAfter(t *testing.T) {
	str := "A^2+B^2=C^2"
	exp := "C^2"
	actual := After(str, "=")

	assert.Equal(t, exp, actual)
}

func TestBetween(t *testing.T) {
	str := "A^2+B^2=C^2"
	exp := "B^2"
	actual := Between(str, "+", "=")

	assert.Equal(t, exp, actual)
}
