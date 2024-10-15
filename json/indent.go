package json

import (
	"bufio"
	"bytes"
	"slices"
	"strings"
)

func Indent(tokens []Token, indent string) bytes.Buffer {
	level := 0
	isJsonKey := false
	var out bytes.Buffer
	buf := bufio.NewWriter(&out)
	// fmt.Println("==Len", len(tokens))
	for _, token := range tokens {
		// json key and next token isn't end object'
		// var nextToken Token
		// if i+1 < len(tokens)-1 {
		// 	nextToken = tokens[i+1]
		// }

		// fmt.Printf("Cur token: %#v, next token: %#v, jsonKey: %v, index: %v\n", token, nextToken, isJsonKey, i)
		if token.Kind == JsonBeginObject {
			if isJsonKey {
				writeNewline(buf)
				currentLineIndent(buf, indent, level)
			}
			buf.WriteRune('{')
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
			level++
			isJsonKey = true

			continue
		}

		if token.Kind == JsonEndArray {
			level--
			writeNewline(buf)
			currentLineIndent(buf, indent, level)
			// isJsonKey = true
			buf.WriteRune(']')

			continue
		}

		if token.Kind == JsonSyntax && token.Value == "," {
			buf.WriteString(token.Value)
			isJsonKey = true

			continue
		}

		if token.Kind == JsonSyntax && token.Value == ":" {
			buf.WriteString(token.Value)
			buf.WriteRune(' ')
			isJsonKey = false

			continue
		}

		if token.Kind == JsonSyntax {
			buf.WriteString(token.Value)
			isJsonKey = false

			continue
		}

		if token.Kind == JsonString {
			if isJsonKey {
				writeNewline(buf)
				currentLineIndent(buf, indent, level)
				buf.WriteString(token.Value)
			} else {
				if isDataToken(token) {
					newStr := token.Value
					buf.WriteString(newStr)
				}
			}

			continue
		}

		if token.Kind == JsonNumber {
			if isJsonKey {
				writeNewline(buf)
				currentLineIndent(buf, indent, level)
			}

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

func isDataToken(token Token) bool {
	return slices.Contains([]int{
		int(JsonBoolean),
		int(JsonNull),
		int(JsonNumber),
		int(JsonString),
	}, int(token.Kind))
}
