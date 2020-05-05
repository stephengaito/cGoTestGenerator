// This GoLang file provides a standard GoLang interface to the underlying 
// ANSI-C cGoTest framework for the {{ .Name }} GoLang Package. 

package {{ .Name }}
//
// Package description:
//   {{ .BriefDesc }}

// #include "{{ .Name }}CGoTests.h"
import "C"

import (
  "errors"
  "fmt"
  "testing"
)

var curTestSuite = ""

// Record to stdout when a given testSuite starts
//
func cGoTestSuite(cGoSuiteName, cGoSuiteDesc string){
  if curTestSuite == cGoSuiteName { return } 
  if curTestSuite != "" {
    fmt.Printf("\nFinishing Suite: %s\n", curTestSuite)
  }
  curTestSuite = cGoSuiteName
  fmt.Printf("\n=====================================================\n")
  fmt.Printf("Test Suite: %s\n", curTestSuite)
  fmt.Printf("   %s\n", cGoSuiteDesc)
  fmt.Printf("=====================================================\n")
}

var curTestFixture = ""

// Record to stdout when a given testFixture starts
//
func cGoTestFixture(cGoFixtureName, cGoFixtureDesc string) {
  if curTestFixture == cGoFixtureName { return } 
  if curTestFixture != "" {
    fmt.Printf("\nFinishing Fixture: %s\n", curTestFixture)
  }
  curTestFixture = cGoFixtureName
  fmt.Printf("\n+++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
  fmt.Printf("Test Fixture: %s\n", curTestFixture)
  fmt.Printf("   %s\n", cGoFixtureDesc)
  fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
}

// Record to stdout when a given test starts
//
func cGoTestStart(cGoTestName, cGoTestDesc string) {
  fmt.Printf("\n-----------------------------------------------------\n")
  fmt.Printf("Test Case: %s\n", cGoTestName)
  fmt.Printf("  %s\n", cGoTestDesc)
  fmt.Printf("-----------------------------------------------------\n")
}


// Record to stdout a message from the cTest system
//
//export cGoTestLog
func cGoTestLog(cGoTestMessage *C.char) {
  if cGoTestMessage != nil {
    fmt.Printf("  %s\n", C.GoString(cGoTestMessage))
  }
}

// Record to stdout when a given test finishes
//
func cGoTestFinish(cGoTestName string) {
  fmt.Printf("-----------------------------------------------------\n\n")
}

// Turn a possible error message from the cTest ANSI-C component into 
// GoLang Errors. 
//
func cGoTestPossibleError(possibleErrorMessage *C.char) error {
  if possibleErrorMessage != nil {
    return errors.New(C.GoString(possibleErrorMessage))
  }
  return nil
}

// Deal with possible error reports from the GoLang component of cTest 
//
func cGoTestMayBeError(t *testing.T, message string, aPossibleError error) {
  if aPossibleError != nil {
    t.Errorf("%s\nerror: %s", message, aPossibleError.Error())
  }
}