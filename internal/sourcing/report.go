// Package sourcing contains ways to report a state
package sourcing

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"
)

const markdownTemplate string = `# {{ .Name }}{{ $st := . }}
{{if .Participants}}
## At a glance [(ℹ)]((https://github.com/jazcarate/sp/blob/master/docs/understanding_a_report.md.md#at-a-glance))
| | {{ range .Participants }}[{{ .Name }}]({{ .Name | ToMarkdownAnchor }}) | {{ end }}
| --- | {{ range .Participants }} ---: | {{ end }}{{ range $x, $ := .Participants }}
| [{{ .Name }}]({{ .Name | ToMarkdownAnchor }}) | {{ range $y, $ := $st.Participants }} {{ Glance $x $y }} |{{ end }}{{ end }}

## Participants
| Name | Split | %age |
| --- | ---: | ---: |{{ range .Participants }}
| [{{ .Name }}]({{ .Name | ToMarkdownAnchor }}) | {{ .Split }} | {{ .SplitPercentage }}% |{{ end }}

## Details
{{ range $i, $ := .Participants }}
### {{ .Name }}
<!-- Public key: {{ .PublicKey }} -->
You currently split the expenditures at **{{ .SplitPercentage }}%**.
{{ range $ip, $val := $st.Balance.Iterate $i }}{{ BalanceDetail $ip $val }}{{ end }}{{ end }}{{ else }}
🗅 Starting out a new Split Chain and don't know "what now?".

No problem! Check the [docs](https://github.com/jazcarate/sp/blob/master/docs/new_sp_now_what.md)
{{end}}
## Operations
Current trust configuration: **{{ .Configuration }}** [(ℹ)](https://github.com/jazcarate/sp/blob/master/docs/understanding_a_report.md.md{{ .Configuration | ToMarkdownAnchor }})

### Log [(go to the last ⬇)](#op-{{ .LastOp }})
{{if .Log}}| # |  Operation | On | Note | Status |
| ---: | --- | --- | --- | ---: |{{ range $i, $op := .Log }}
| [{{ $i }}](#op-{{ $i }})<a id="op-{{ $i }}"></a> | {{ $op.Operation | ToOpMarkdown }}<!-- Sign &{{ $op.By}} {{ $op.Signature }}-->  | {{ $op.On | ToTime }} | {{ $op.Note }} | {{if $op.Valid }}✅{{ else }}❓{{ end }} |{{ else }}
🌈 Fresh new 🌈{{ end }}{{ end }}
`

func participant(name string) string {
	return fmt.Sprintf("[%s](%s)", name, toMarkdownAnchor(name))
}

func toOpMarkdown(op StateChanger) string {
	switch e := op.(type) {
	case AddParticipant:
		return fmt.Sprintf("➕ Add %s<!-- Public key: %s -->", participant(e.Name), e.PublicKey)
	case SplitParticipant:
		return fmt.Sprintf("🪓 Split %s for `%d`", participant(e.Name), e.NewSplit)
	case Configure:
		return fmt.Sprintf("💻 Configure to `%s`", e.NewConfig)
	case Spend:
		return fmt.Sprintf("💸 %s **spent $%d**", participant(e.Who), e.Amount)
	case MultiOp:
		var result string
		for _, o := range e.Ops {
			result += toOpMarkdown(o) + "<br>"
		}

		return result
	default:
		panic(fmt.Sprint("Unknown event type", op))
	}
}

func toMarkdownAnchor(s string) string {
	return "#" + strings.ToLower(s)
}

func toTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// Markdown converts a state to a markdown report.
func (s *State) Markdown(wr io.Writer) error {
	glance := func(x int, y int) string {
		p, err := s.Balance.Get(x, y)
		if err != nil {
			panic(fmt.Sprintf("can't get the participant at (%d, %d)", x, y))
		}

		if p == 0 {
			return ""
		}

		if p < 0 {
			return fmt.Sprint("$", -p, "◀")
		}

		return fmt.Sprint("$", p, "🔼")
	}

	balanceDetail := func(pIndex int, i int) string {
		p := s.Participants[pIndex]

		if i == 0 {
			return ""
		}

		if i > 0 {
			return fmt.Sprintf("- You **owe $%d** to %v\n", i, participant(p.Name))
		}

		return fmt.Sprintf("- You are **owed $%d** from %v\n", -i, participant(p.Name))
	}

	funcMap := template.FuncMap{
		"ToMarkdownAnchor": toMarkdownAnchor,
		"ToOpMarkdown":     toOpMarkdown,
		"ToTime":           toTime,
		"Glance":           glance,
		"BalanceDetail":    balanceDetail,
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
