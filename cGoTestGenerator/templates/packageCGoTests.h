// +build cGoTests

/// \file
/// \brief This ANSI-C Header file provides the ANSI-C based cTest testing 
/// framework for the {{ .Name }} GoLang Package. 
///
/// Package description:
///   {{ .BriefDesc }}
///
/// This file is automatically (re)generated changes made to this file will 
/// be lost. 

#ifndef {{ .Name }}_CGO_TESTS_H
#define {{ .Name }}_CGO_TESTS_H

#ifndef NULL
#define NULL 0
#endif

extern void *nullSetup(void);

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