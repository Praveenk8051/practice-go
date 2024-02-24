package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)
	actual := DivisibleBy7(2100)
	asserts.Equal(true, actual)
	actual = DivisibleBy7(2101)
	asserts.Equal(false, actual)
	actual = DivisibleBy7(2105)
	asserts.Equal(false, actual)

	actual = NotDivisibleBy5(3001)
	asserts.Equal(true, actual)
	actual = NotDivisibleBy5(3003)
	asserts.Equal(true, actual)
	actual = NotDivisibleBy5(3005)
	asserts.Equal(false, actual)

}
