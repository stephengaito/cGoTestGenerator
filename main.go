// A GoLang tool to go:generate CGoTests to unit test C code embedded in 
// Go code. 

package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "regexp"
  "strings"
)

var testCase    = regexp.MustCompile("CGoTest\\s*\\(")
var inFixture   = regexp.MustCompile("[\\\\\\@]inFixture\\s+(\\S+)")
var testFixture = regexp.MustCompile("[\\\\\\@]testFixture\\s+(\\S+)\\s+(.*)$")
var inSuite     = regexp.MustCompile("[\\\\\\@]inSuite\\s+(\\S+)")
var testSuite   = regexp.MustCompile("[\\\\\\@]testSuite\\s+(\\S+)\\s+(.*)$")

func processTestFile(testFile string) error {
  fmt.Printf("looking at: [%s]\n", testFile)
  file, err := os.Open(testFile)
  if err != nil { return err }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    if testCase.MatchString(text) {
      fmt.Printf("testCase: [%s]\n", text)
    }
    matches := inFixture.FindStringSubmatch(text)
    if matches != nil {
      fmt.Printf("inFixture: [%s]\n", matches[1])
    }
    matches = testFixture.FindStringSubmatch(text)
    if matches != nil {
      fmt.Printf("testFixture: [%s](%s)\n", matches[1], matches[2])
    }
    matches = inSuite.FindStringSubmatch(text)
    if matches != nil {
      fmt.Printf("inSuite: [%s]\n", matches[1])
    }
    matches = testSuite.FindStringSubmatch(text)
    if matches != nil {
      fmt.Printf("testSuite: [%s](%s)\n", matches[1], matches[2])
    }
  }
  if err := scanner.Err(); err != nil { return err }
  return nil
}

func main() {

  err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if ! strings.HasSuffix(path, "CGoTest.c") {
      return nil
    }
    processTestFile(path)
    return nil
  })
  if err != nil {
    fmt.Printf("error: %v", )
  }
}