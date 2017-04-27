package main_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/one"
)

func TestResultsLen(t *testing.T) {
  results := make(Results, 10, 59)

  assert.Equal(t, 10, results.Len())
}
