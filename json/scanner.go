package json

import (
	"fmt"
	"unicode"
)

type TokenKind int

const (
	JsonBoolean TokenKind = iota
	JsonNull
	JsonNumber
	JsonString
	JsonSyntax
	JsonBeginArray
	JsonEndArray
	JsonBeginObject
	JsonEndObject
)

type Token struct {
	Kind  TokenKind
	Value string
	// lineNo int
	// colNo  int
}

type Scanner struct {
	Source  string
	Start   int
	Current int
	End     int
}

func NewScanner(input string) *Scanner {
	return &Scanner{
		Source: input,
		End:    len(input) - 1,
	}
}

func (s *Scanner) Scan() []Token {
	tokens := []Token{}
	i := 0

	for s.Current < len(s.Source) {
		i++
		if i > 40 {
			break
		}
		c := s.Source[s.Current]
		// fmt.Printf("===Char: %s\n", string(c))

		var token Token

		switch c {
		case '{':
			tokens = append(tokens, Token{
				Kind:  JsonBeginObject,
				Value: "{",
			})
			s.Current++
		case '}':
			tokens = append(tokens, Token{
				Kind:  JsonEndObject,
				Value: "}",
			})
			s.Current++
		case '[':
			tokens = append(tokens, Token{
				Kind:  JsonBeginArray,
				Value: "[",
			})
			s.Current++
		case ']':
			tokens = append(tokens, Token{
				Kind:  JsonEndArray,
				Value: "]",
			})
			s.Current++
		case ':':
			tokens = append(tokens, Token{
				Kind:  JsonSyntax,
				Value: ":",
			})
			s.Current++
		case '"':
			s.Current++
			token = s.stringLiteral()
			tokens = append(tokens, token)
			// fmt.Printf("===Token: %#v\n", token)
			s.Current++
			continue
		case ',':
			tokens = append(tokens, Token{
				Kind:  JsonSyntax,
				Value: ",",
			})
			s.Current++
			continue
		case 't', 'f':
			token = s.boolLiteral()
			tokens = append(tokens, token)
			continue
		case 'n':
			token = s.nullLiteral()
			tokens = append(tokens, token)
			continue
		case ' ':
			s.Current++
			continue
		default:
			isDigit := unicode.IsDigit(rune(s.Source[s.Current]))
			if isDigit {
				token = s.numberLiteral()
				tokens = append(tokens, token)
			} else {
				fmt.Printf("===Unknown c: %#v\n", string(s.Source[s.Current]))
			}
			continue
		}
	}

	return tokens
}

func (s *Scanner) stringLiteral() Token {
	curr := s.Current
	for s.Current < len(s.Source) && s.Source[s.Current] != '"' {
		s.Current++
	}

	return Token{
		Kind:  JsonString,
		Value: s.Source[curr:s.Current],
	}
}

func (s *Scanner) numberLiteral() Token {
	curr := s.Current
	for unicode.IsDigit(rune(s.Source[s.Current])) {
		s.Current++
	}

	return Token{
		Kind:  JsonNumber,
		Value: s.Source[curr:s.Current],
	}
}

func (s *Scanner) boolLiteral() Token {
	curr := s.Current
	t := s.Source[s.Current : s.Current+4]
	f := s.Source[s.Current : s.Current+5]

	if t == "true" {
		s.Current += 4
		return Token{
			Kind:  JsonNumber,
			Value: s.Source[curr:s.Current],
		}
	}

	if f == "false" {
		s.Current += 5
		return Token{
			Kind:  JsonBoolean,
			Value: s.Source[curr:s.Current],
		}
	}

	return Token{}
}

func (s *Scanner) nullLiteral() Token {
	curr := s.Current
	null := s.Source[s.Current : s.Current+4]

	if null == "null" {
		s.Current += 4
		return Token{
			Kind:  JsonNull,
			Value: s.Source[curr:s.Current],
		}
	}

	return Token{}
}
