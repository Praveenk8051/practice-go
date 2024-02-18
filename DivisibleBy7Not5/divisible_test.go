package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	assert.Equal(t, int(2002), int(2002))
	assert.Equal(t, 2009, int(2009))
	assert.Equal(t, 2023, int(2023))
	assert.Equal(t, 2051, int(2051))

	assert.NotEqual(t, 2024, "Number doesn't satify the needs")
	assert.NotEqual(t, 2050, "Number doesn't satify the needs")
	assert.NotEqual(t, 2055, "Number doesn't satify the needs")
}
