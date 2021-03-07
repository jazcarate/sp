// Package sourcing contains ways to report a state
package sourcing

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

const markdownTemplate string = `# {{.Name}}
{{if .Participants}}{{else}}
🗅 Starting out a new Split Chain and don't know "what now?".

No problem! Check the [docs](https://github.com/jazcarate/sp/blob/master/docs/new_sp_now_what.md)
{{end}}
## Operations
Current trust configuration: **{{ .Configuration }}** [(ℹ)](https://github.com/jazcarate/sp/blob/master/docs/understanding_a_report.md.md{{ .Configuration | ToMarkdownAnchor }})

### Log
🌈 Fresh new 🌈
`

func toMarkdownAnchor(s fmt.Stringer) string {
	return "#" + strings.ToLower(s.String())
}

// Markdown converts a state to a markdown report.
func (s *State) Markdown(wr io.Writer) error {
	funcMap := template.FuncMap{
		"ToMarkdownAnchor": toMarkdownAnchor,
	}

	if s == nil {
		s = NewState()
	}

	tmpl, tmplErr := template.New("markdown").Funcs(funcMap).Parse(markdownTemplate)
	if tmplErr != nil {
		return fmt.Errorf("template parsing: %w", tmplErr)
	}

	execErr := tmpl.Execute(wr, s)
	if execErr != nil {
		return fmt.Errorf("template executing: %w", execErr)
	}

	return nil
}
