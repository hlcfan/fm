package json

import (
	"bufio"
	"bytes"
	"strings"
)

func Indent(tokens []Token, indent string) bytes.Buffer {
	level := 0
	isJsonKey := false
	var out bytes.Buffer
	buf := bufio.NewWriter(&out)

	for _, token := range tokens {
		if token.Kind == JsonBeginObject {
			currentLineIndent(buf, indent, level)
			buf.WriteRune('{')
			writeNewline(buf)
			level++
			isJsonKey = true

			continue
		}

		if token.Kind == JsonEndObject {
			level--
			writeNewline(buf)
			currentLineIndent(buf, indent, level)
			buf.WriteRune('}')

			continue
		}

		if token.Kind == JsonBeginArray {
			buf.WriteRune('[')
			writeNewline(buf)
			level++
			isJsonKey = true

			continue
		}

		if token.Kind == JsonEndArray {
			level--
			writeNewline(buf)
			currentLineIndent(buf, indent, level)
			buf.WriteRune(']')

			continue
		}

		if token.Kind == JsonSyntax && token.Value == "," {
			buf.WriteString(token.Value)
			writeNewline(buf)
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
				currentLineIndent(buf, indent, level)
				isJsonKey = false
			}

			buf.WriteString(token.Value)

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

	return out
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
