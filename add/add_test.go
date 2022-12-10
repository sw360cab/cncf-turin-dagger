package add_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sw360cab/dagger/add"
)

func TestAdd(t *testing.T) {
	t.Run("Adding two positive number", func(t *testing.T) {
		assert.Equal(t, 3, add.Adder(2, 1))
		assert.GreaterOrEqual(t, 1, add.Adder(0, 1))
		assert.Less(t, -1, add.Adder(1, -1))
		assert.Less(t, add.Adder(-1, -2), -1)
	})
}
