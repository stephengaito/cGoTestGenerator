/// \file
/// \brief This ANSI-C Header file provides the ANSI-C based cTest testing 
/// framework for the {{ .Name }} GoLang Package. 
///
/// Package description:
///   {{ .BriefDesc }}

#ifndef {{ .Name }}_CGO_TESTS_H
#define {{ .Name }}_CGO_TESTS_H

{{ range .Suites }}
// begin suite: {{ .Name }}
{{   range .Fixtures }}
  // begin fixture: {{ .Name }}
  {{ if .SetupName }}  extern void *{{ .SetupName }} (void); {{ end }}
  {{ if .TeardownName }}  extern void {{ .TeardownName }} (void *data); {{ end }}
{{     range .Cases }}
    extern char *{{ .Name }}(void *data);
{{     end }}
  // end fixture: {{ .Name }}
{{   end }}
// end suite: {{ .Name }}
{{ end }}

#endif