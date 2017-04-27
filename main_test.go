package main_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/one"

  "fmt"
)

var _ = fmt.Println // DELME

func NewResults() (results Results) {
  for _, letter := range "Hello" {
    results.Append(letter)
  }

  return
}

func TestResultsLen(t *testing.T) {
  results := NewResults()

  assert.Equal(t, 4, results.Len())
}

func TestResultsLess(t *testing.T) {
  results := NewResults()

  assert.True(t, results.Less(2, 3))
}

func TestResultsSwap(t *testing.T) {
  results := NewResults()

  results.Swap(1, 2)

  expectation := Results{
    []int{ 108, 2 },
    []int{ 101, 1 },
  }

  assert.Equal(t, expectation, results[1:3])
}

func TestResultsAppend(t *testing.T) {
  var results Results

  for _, letter := range "Hello World! and hello world again" {
    results.Append(letter)
  }

  assert.Equal(t, 11, results.Len())
}
