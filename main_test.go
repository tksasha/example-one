package main_test

import (
  "testing"
  "sort"

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
    []int{ int('l'), 2 },
    []int{ int('e'), 1 },
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

func TestResultsSorting(t *testing.T) {
  results := NewResults()

  sort.Stable(results)

  expectation := Results{
    []int{ int('l'), 2 },
    []int{ int('h'), 1 },
    []int{ int('e'), 1 },
    []int{ int('o'), 1 },
  }

  assert.Equal(t, expectation, results)
}
