# cGoTestGenerator

 A GoLang tool to generate unit tests of C code embedded in Go 
 
We roughly follow the standard 
[xUnit](https://en.wikipedia.org/wiki/XUnit) setup. 

The most basic cGoTest primitives are the **cGoTest assertions**. All 
assertions test their respective assertion and if it fails report an 
error. Failing XXX_MayFail assertions will stop the test case at that 
point. All other XXX assertions will allow the test case to proceed. 

These assertions are placed in **cGoTestCases**. cGoTestCases are meant to 
test one "action" (and any associated structures). A cGoTestCase is an 
ANSI-C function which returns `char*` and takes a single `void*` argument 
and whose name ends in `CGoTest`. The *most recent* `\inFixture 
<fixtureName>` (or `@inFixture <fixtureName>`) marker is used to 
associated a given cGoTestCase to a cGoTestFixture. 

cGoTestCases are run as part of a **cGoTestFixture**. cGoTestFixtures 
provide the context for one or more cGoTestCases. A cGoTestFixture is 
defined by the `\testFixture <fixtureName>` (or `@testFixture 
<fixtureName>`) marker. 

Typically a cGoTestFixture will define a **cGoTestSetup** and/or 
**cGoTestTeardown** function used to respectively setup and teardown all 
required structures for each cGoTestCase in the cGoTestFixture. A 
cGoTestSetup is an ANSI-C function which returns a `void*` and takes no 
arguments. A cGoTestTeardown is an ANSI-C function which takes a single 
`void*` argument. The *most recent* `\inFixture <fixtureName>` (or 
`@inFixture <fixtureName>`) marker is used to assocaite the setup or 
teardown method with a cGoTestFixture. 

One or more cGoTestFixtures are run as part of a **cGoTestSuite**.

**NOTE**: Since we are using line based [regular 
expressions](https://golang.org/pkg/regexp/) to find the cGoTestCases, 
cGoTestSetup, cGoTestTeardown functions as well as cGoTestFixture and 
cGoTestSuite definitions, the respective function prototypes and/or 
markers **MUST** be on *one* *line*. Function prototypes and/or markers 
which occurs over more than one line will not be recognized. 

## Typical use:
 
```
     //go:generate cGoTestGenerator <packageName> <packageDesc>
 ```
 
Will run the cGoTestGenerator on any files ending in "CGotTest.c" in the 
current directory and any of its subdirectories. (See 
[go:generate](https://blog.golang.org/generate)) 

## Requirements

We use Matt Jibson's '[mjibson/esc](https://github.com/mjibson/esc) to 
embed all "static" ANSI-C code files required by the cGoTests into the 
single cGoTestGenerator binary.

To install `esc` type:

```
go get -u github.com/mjibson/esc
```

## Gitignore

The automatically generated files satisfy the following `.gitignore` 
patterns:

```
*GoTests.go
*GoTests_test.go
*GoTests.h
*GoTestsUtils.h
```

## Doxygen

The `\testSuite`, `\inSuite`, `\testFixture`, and `\inFixture` markers are 
all meant to work with Doxygen as embedded commands. To do this you need to 
add the following to your Doxyfile:

```
ALIASES += testSuite="\par Suite:\n"
ALIASES += inSuite="\par Suite:\n"
ALIASES += testFixture="\par Fixture:\n"
ALIASES += inFixture="\par Fixture:\n"
```

(See [Doxygen Custon 
Commands](http://www.doxygen.nl/manual/custcmd.html)). 

