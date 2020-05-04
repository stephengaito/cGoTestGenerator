// +build tests

// This GoLang file provides a standard GoLang interface to the underlying 
// ANSI-C cTest framework. 

package cJoyLoL

// #include "cTests.h"
import "C"

import (
  "errors"
  "fmt"
  "testing"
)

// Record to stdout when a given test starts
//
func cTestStart(cTestName string) {
  fmt.Printf("\n Starting: %s\n", cTestName)
}


// Record to stdout a message from the cTest system
//
//export cTestLog
func cTestLog(cTestMessage *C.char) {
  if cTestMessage != nil {
    fmt.Printf("  %s\n", C.GoString(cTestMessage))
  }
}

// Record to stdout when a given test finishes
//
func cTestFinish(cTestName string) {
  fmt.Printf("Finishing: %s\n\n", cTestName)
}

// Turn a possible error message from the cTest ANSI-C component into 
// GoLang Errors. 
//
func cTestPossibleError(possibleErrorMessage *C.char) error {
  if possibleErrorMessage != nil {
    return errors.New(C.GoString(possibleErrorMessage))
  }
  return nil
}

// Deal with possible error reports from the GoLang component of cTest 
//
func cTestMayBeError(t *testing.T, message string, aPossibleError error) {
  if aPossibleError != nil {
    t.Errorf("%s\nerror: %s", message, aPossibleError.Error())
  }
}