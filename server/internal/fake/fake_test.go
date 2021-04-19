package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFakeData(t *testing.T) {
	num := 10
	capstones := GenerateFakeCapstones(num)

	assert.Len(t, capstones, num, "List length should equal given num")

}
