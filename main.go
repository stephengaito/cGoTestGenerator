// A GoLang tool to go:generate CGoTests to unit test C code embedded in 
// Go code. 

package main

//go:generate esc -o templates.go templates

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "regexp"
  "strings"
  "text/template"
)

var briefDesc    = regexp.MustCompile("[\\\\\\@]brief\\s*(.*)$")
var testCase     = regexp.MustCompile("\\*(\\S+CGoTest)\\s*\\(")
var testSetup    = regexp.MustCompile("\\*(\\S+CGoTestSetup)\\s*\\(")
var testTeardown = regexp.MustCompile("(\\S+CGoTestTeardown)\\s*\\(")
var inFixture    = regexp.MustCompile("[\\\\\\@]inFixture\\s+(\\S+)")
var testFixture  = regexp.MustCompile("[\\\\\\@]testFixture\\s+(\\S+)\\s+(.*)$")
var inSuite      = regexp.MustCompile("[\\\\\\@]inSuite\\s+(\\S+)")
var testSuite    = regexp.MustCompile("[\\\\\\@]testSuite\\s+(\\S+)\\s+(.*)$")

type TestCase struct {
  Name      string
  BriefDesc string
}

type TestCases map[string]*TestCase

type TestFixture struct {
  Name         string
  BriefDesc    string
  SetupName    string
  SetupDesc    string
  TeardownName string
  TeardownDesc string
  Cases        TestCases
}

type TestFixtures map[string]*TestFixture

type TestSuite struct {
  Name      string
  BriefDesc string
  Fixtures  TestFixtures
}

type TestSuites map[string]*TestSuite

type TestRunner struct {
  Name      string
  BriefDesc string
  Suites    TestSuites
}

var testRunner *TestRunner

func newTestCase(name, briefDesc string) *TestCase {
  return &TestCase{
    Name:      name,
    BriefDesc: briefDesc,
  }
}

func newTestFixture(name, briefDesc string) *TestFixture {
  return &TestFixture{
    Name:       name,
    BriefDesc:  briefDesc,
    Cases:      make(TestCases),
  }
}

func newTestSuite(name, briefDesc string) *TestSuite {
  return &TestSuite {
    Name:      name,
    BriefDesc: briefDesc,
    Fixtures:  make(TestFixtures),
  }
}

func newTestRunner(name, briefDesc string) *TestRunner {
  return &TestRunner{
    Name:      name,
    BriefDesc: briefDesc,
    Suites:    make(TestSuites),
  }
}

func processTestFile(testFile string) error {
  fmt.Printf("\nscanning cGoTest file: [%s]\n", testFile)

  file, err := os.Open(testFile)
  if err != nil { return err }
  defer file.Close()

  curSuite     := testRunner.Suites["main"]
  curFixture   := curSuite.Fixtures["main"]
  curBriefDesc := "no desc"
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    matches := briefDesc.FindStringSubmatch(text)
    if matches != nil {
      curBriefDesc = matches[1]
      fmt.Printf("  briefDesc:    [%s]\n", curBriefDesc)
    }
    matches = testCase.FindStringSubmatch(text)
    if matches != nil {
      caseName := matches[1]
      curFixture.Cases[caseName] = newTestCase(caseName, curBriefDesc)
      fmt.Printf("  testCase:     [%s] (%s)\n", caseName, curBriefDesc)
    }
    matches = testSetup.FindStringSubmatch(text)
    if matches != nil {
      setupName := matches[1]
      curFixture.SetupName = setupName
      curFixture.SetupDesc = curBriefDesc
      fmt.Printf("  testSetup:    [%s] (%s)\n", setupName, curBriefDesc)
    }
    matches = testTeardown.FindStringSubmatch(text)
    if matches != nil {
      teardownName := matches[1]
      curFixture.TeardownName = teardownName
      curFixture.TeardownDesc = curBriefDesc
      fmt.Printf("  testTeardown: [%s] (%s)\n", teardownName, curBriefDesc)
    }
    matches = inFixture.FindStringSubmatch(text)
    if matches != nil {
      fixtureName := matches[1]
      curFixture   = curSuite.Fixtures[fixtureName]
      if curFixture == nil {
        curFixture = newTestFixture(fixtureName, "unknown")
        curSuite.Fixtures[fixtureName] = curFixture
      }
      fmt.Printf("  inFixture:    [%s]\n", fixtureName)
    }
    matches = testFixture.FindStringSubmatch(text)
    if matches != nil {
      fixtureName := matches[1]
      fixtureDesc := matches[2]
      aFixture := curSuite.Fixtures[fixtureName]
      if aFixture == nil {
        aFixture = newTestFixture(fixtureName, fixtureDesc)
        curSuite.Fixtures[fixtureName] = aFixture
      }
      aFixture.BriefDesc = fixtureDesc
      fmt.Printf("  testFixture:  [%s] (%s)\n", fixtureName, fixtureDesc)
    }
    matches = inSuite.FindStringSubmatch(text)
    if matches != nil {
      suiteName := matches[1]
      curSuite = testRunner.Suites[suiteName]
      if curSuite == nil {
        curSuite = newTestSuite(suiteName, "unknown")
        testRunner.Suites[suiteName] = curSuite
      }
      fmt.Printf("  inSuite:      [%s]\n", suiteName)
    }
    matches = testSuite.FindStringSubmatch(text)
    if matches != nil {
      suiteName := matches[1]
      suiteDesc := matches[2]
      aSuite := testRunner.Suites[suiteName]
      if aSuite == nil {
        aSuite = newTestSuite(suiteName, suiteDesc)
        testRunner.Suites[suiteName] = aSuite
      }
      aSuite.BriefDesc = suiteDesc
      fmt.Printf("  testSuite:    [%s] (%s)\n", suiteName, suiteDesc)
    }
  }
  if err := scanner.Err(); err != nil { return err }
  return nil
}

func createFileFrom(filePath, templatePath string) {
  fmt.Printf("\ncreating   file: [%s]\n", filePath)
  fmt.Printf("  from template: [%s]\n", templatePath)
  
  aTemplateStr, err := FSString(false, templatePath)
  if err != nil {
    fmt.Printf(
      "Could not load the template: [%s] Error: %s\n",
      templatePath,
      err,
    )
    os.Exit(-1)
  }
  aTemplate, err := template.New("default").Parse(aTemplateStr)
  if err != nil {
    fmt.Printf(
      "Could not parse the template: [%s] Error: %s\n",
      templatePath,
      err,
    )
    os.Exit(-1)
  }
  aFile, err := os.Create(filePath)
  if err != nil {
    fmt.Printf(
      "Could not create the file: [%s] Error: %s\n",
      filePath,
      err,
    )
    os.Exit(-1)
  }
  err = aTemplate.Execute(aFile, testRunner)
  if err != nil {
    fmt.Printf(
      "Could not run the template: [%s]\n  on the file: [%s]\n  Error: %s\n",
      templatePath,
      filePath,
      err,
    )
    os.Exit(-1)
  }
}

func main() {

  // setup the default test suites
  //
  packageName := "main"
  if 1 < len(os.Args) { packageName = os.Args[1] }
  //
  packageDesc := "Main package"
  if 2 < len(os.Args) { packageDesc = strings.Join(os.Args[2:], " ") }
  //
  testRunner = newTestRunner(packageName, packageDesc)
  testRunner.Suites["main"] = newTestSuite("main", "Main (default) TestSuite")
  testRunner.Suites["main"].Fixtures["main"] =
    newTestFixture("main", "Main (default) Fixture in Main Suite")
  fmt.Printf("package: [%s] (%s)\n", packageName, packageDesc)
  
  // Walk through all of the "CGoTest" files ...
  // ... building up the testSuites structure
  //
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
  
  // Now create the required cGoTest files
  //
  createFileFrom("cGoTests.h",                       "/templates/cGoTests.h")
  createFileFrom("cGoTests.go",                      "/templates/cGoTests.go")
  createFileFrom(testRunner.Name+"CGoTests.h",       "/templates/packageCGoTests.h")
  createFileFrom(testRunner.Name+"CGoTests.go",      "/templates/packageCGoTests.go")
  createFileFrom(testRunner.Name+"CGoTests_test.go", "/templates/packageCGoTests_test.go")

  fmt.Printf("\n")
}
