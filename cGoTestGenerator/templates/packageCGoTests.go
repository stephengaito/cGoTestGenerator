// +build cGoTests

// GoLang level wrappers of the ANSI-C tests in the {{ .Name }}
// GoLang Package. 
//
// This *should* be located in {{ .Name }}CGoTests_test.go...
// ... unfortunately `go test` forbids the use of cgo...
// ... so we need to maintain this addition level of code indirection.
//
// Package description:
//   {{ .BriefDesc }}
//
// This file is automatically (re)generated changes made to this file will 
// be lost. 
//
package {{ .Name }}

// #include "{{ .Name }}CGoTests.h"
import "C"

import (
)

{{ range $theSuite := .Suites }}
// begin suite: {{ $theSuite.Name }}
{{   range $theFixture := .Fixtures }}
  // begin fixture: {{ $theFixture.Name }}
{{     range .Cases }}
    // {{ .BriefDesc }}
    //
    // Suite:   {{ $theSuite.Name }}
    // Fixture: {{ $theFixture.Name }}
    //
    func Go_{{ .Name }}() error {
      cGoTestSuite("{{ $theSuite.Name }}", "{{ $theSuite.BriefDesc }}")
      cGoTestFixture("{{ $theFixture.Name }}", "{{ $theFixture.BriefDesc }}")
      
      cGoTestStart("{{ .Name }}", "{{ .BriefDesc }}")
      defer cGoTestFinish("{{ .Name }}")

{{       if $theFixture.SetupName }}     
      data := C.{{ $theFixture.SetupName}}()
{{       else }}
      data := C.nullSetup()
{{       end }}
{{       if $theFixture.TeardownName }}     
      defer C.{{ $theFixture.TeardownName}}(data)
{{       end }}
      
      return cGoTestPossibleError(C.{{ .Name }}(data))
    }
{{     end }}
  // end fixture: {{ $theFixture.Name }}
{{   end }}
// end suite: {{ $theSuite.Name }}
{{ end }}

