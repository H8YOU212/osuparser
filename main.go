package osuparser

import (
	"osuparser/parser"
)

// NewParser создаёт новый парсер
func NewParser() *parser.Parser {
	return parser.NewParser()
}

// Parse парсит файл и возвращает результат
func Parse(path string) (*parser.ParseResult, error) {
	p := NewParser()
	return p.Parse(path)
}
