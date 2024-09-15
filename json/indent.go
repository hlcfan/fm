package json

import (
	"bufio"
	"os"
	"strings"
)

func Indent(tokens []Token, indent string) {
	level := 0
	isJsonKey := false
	buf := bufio.NewWriter(os.Stdout)

	for _, token := range tokens {
		if token.Kind == JsonBeginObject {
			buf.WriteString(strings.Repeat(indent, level))
			buf.WriteRune('{')
			buf.WriteByte('\n')
			level++
			isJsonKey = true

			continue
		}

		if token.Kind == JsonEndObject {
			level--
			buf.WriteByte('\n')
			buf.WriteString(strings.Repeat(indent, level))
			buf.WriteRune('}')

			continue
		}

		if token.Kind == JsonBeginArray {
			buf.WriteRune('[')
			buf.WriteByte('\n')
			level++
			isJsonKey = true

			continue
		}

		if token.Kind == JsonEndArray {
			level--
			buf.WriteByte('\n')
			buf.WriteString(strings.Repeat(indent, level))
			buf.WriteRune(']')

			continue
		}

		if token.Kind == JsonSyntax && token.Value == "," {
			buf.WriteString(token.Value)
			buf.WriteByte('\n')
			isJsonKey = true

			continue
		}

		if token.Kind == JsonSyntax && token.Value == ":" {
			buf.WriteString(token.Value)
			buf.WriteRune(' ')

			continue
		}

		if token.Kind == JsonSyntax {
			buf.WriteString(token.Value)

			continue
		}

		if token.Kind == JsonString {
			if isJsonKey {
				// buf.WriteString(strings.Repeat(indent, level))
				currentLineIndent(buf, indent, level)
				isJsonKey = false
			}
			buf.WriteRune('"')
			buf.WriteString(token.Value)
			buf.WriteRune('"')

			continue
		}

		if token.Kind == JsonNumber {
			buf.WriteString(token.Value)

			continue
		}

		if token.Kind == JsonBoolean {
			buf.WriteString(token.Value)

			continue
		}

		if token.Kind == JsonNull {
			buf.WriteString(token.Value)

			continue
		}
	}

	writeNewline(buf)
	buf.Flush()
}

func writeNewline(w *bufio.Writer) {
	w.WriteByte('\n')
}

func currentLineIndent(w *bufio.Writer, indent string, level int) {
	w.WriteString(strings.Repeat(indent, level))
}

func nextLineIndent(w *bufio.Writer, indent string, level int) {
	w.WriteString(strings.Repeat(indent, level+1))
}
