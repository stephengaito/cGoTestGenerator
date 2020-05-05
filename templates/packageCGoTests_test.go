// GoLang level tests for the {{ .Name }} ANSI-C code

package {{ .Name }}
//
// Package description:
//   {{ .BriefDesc }}

import (
  "testing"
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
    func Test_{{ .Name }}(t *testing.T) {      
      cGoTestMayBeError(t, "{{ .Name }}", Go_{{.Name }}())
    }
{{     end }}
  // end fixture: {{ $theFixture.Name }}
{{   end }}
// end suite: {{ $theSuite.Name }}
{{ end }}

