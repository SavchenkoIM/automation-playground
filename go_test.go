package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {

	t.Run("Main Test", func(t *testing.T) {
		main()
		assert.Equal(t, "", "f")
	})
}
