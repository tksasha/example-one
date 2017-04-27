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

func (results *Results) Append(letter rune) {
  for _, elements := range *results {
    if elements[0] == int(letter) {
      elements[1]++

      return
    }
  }

  *results = append(*results, []int{ int(letter), 1 })
}

func main() {
  mask := regexp.MustCompile("[[:alpha:]]")

  var results Results

  var input string

  scanner := bufio.NewScanner(os.Stdin)

  scanner.Scan()

  input = strings.ToLower(scanner.Text())

  for _, letter := range input {
    if mask.MatchString(string(letter)) == false {
      continue
    }

    results.Append(letter)
  }

  sort.Stable(results)

  for _, elements := range results {
    fmt.Println(string(elements[0]), elements[1])
  }
}
