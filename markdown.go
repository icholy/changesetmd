package main

import (
	"fmt"
	"strings"
)

// MarkdownTable generates markdown tables
type MarkdownTable struct {
	EmptyValue string
	MaxLen     int
	strings.Builder
}

// WriteHeaders writes the header names.
func (t *MarkdownTable) WriteHeaders(names ...string) {
	if len(names) == 0 {
		return
	}
	for i, h := range names {
		if i == 0 {
			t.WriteByte('|')
		}
		fmt.Fprintf(t, " %s |", h)
	}
	t.WriteByte('\n')
	for i := range names {
		if i == 0 {
			t.WriteByte('|')
		}
		t.WriteString(" --- |")
	}
	t.WriteByte('\n')
}

// WriteRow writes row values to the table.
func (t *MarkdownTable) WriteRow(values ...string) {
	if len(values) == 0 {
		return
	}
	for i, v := range values {
		if i == 0 {
			t.WriteByte('|')
		}
		if t.MaxLen > 0 {
			v = BreakLine(v, t.MaxLen)
		}
		if v == "" {
			v = t.EmptyValue
		}
		fmt.Fprintf(t, " %s |", v)
	}
	t.WriteByte('\n')
}

// BreakLine adds line breaks when the string exceeds the max length.
func BreakLine(s string, maxlen int) string {
	if len(s) <= maxlen {
		return s
	}
	var b strings.Builder
	for i, r := range s {
		if i != 0 && i%maxlen == 0 {
			b.WriteString("<br>")
		}
		b.WriteRune(r)
	}
	return b.String()
}
