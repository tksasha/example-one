package main

import (
  "fmt"
  "sort"
  "strings"
  "regexp"
  "bufio"
  "os"
)

type Results [][]int

func (results Results) Len() int {
  return len(results)
}

func (results Results) Less(i, j int) bool {
  return results[j][1] < results[i][1]
}

func (results Results) Swap(i, j int) {
  results[i], results[j] = results[j], results[i]
}

func (results *Results) Append(input rune) {
  //
  // only letters(a-zA-Z) allowed
  //
  mask := regexp.MustCompile("[[:alpha:]]")

  //
  // convert `rune` to `string`
  //
  letter := string(input)

  //
  // ignore non-letters characters
  //
  if mask.MatchString(letter) == false {
    return
  }

  //
  // downcase
  //
  letter = strings.ToLower(letter)

  //
  // get ASCII code of letter
  //
  code := int(letter[0])

  //
  // if `letter` is already in `results` increase counters
  // else add `letter` to `results` with counter's value 1
  //
  for _, elements := range *results {
    if elements[0] == code {
      elements[1]++

      return
    }
  }

  //
  // update `results`
  //
  *results = append(*results, []int{ code, 1 })
}

func main() {
  var results Results

  var input string

  scanner := bufio.NewScanner(os.Stdin)

  scanner.Scan()

  input = scanner.Text()

  for _, letter := range input {
    results.Append(letter)
  }

  sort.Stable(results)

  for _, elements := range results {
    fmt.Println(string(elements[0]), elements[1])
  }
}
