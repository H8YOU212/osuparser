package parser

import (
	"fmt"
	"path/filepath"

	"github.com/H8YOU212/osuparser/parser/file"
	"github.com/H8YOU212/osuparser/parser/file/osr"
	"github.com/H8YOU212/osuparser/parser/file/osz"
	"github.com/H8YOU212/osuparser/parser/osunet"
)

type Parser struct {
	File *file.File
	Net  *osunet.Client
	dst  string
}

// ParseResult содержит результат парсинга
type ParseResult struct {
	Osz *osz.OszModel
	Osr *osr.OsrModel
}

func NewParser() *Parser {
	return &Parser{
		File: file.New(),
		Net:  osunet.NewClient(),
	}
}

func (p *Parser) Parse(path string) (*ParseResult, error) {
	if path == "" {
		return nil, fmt.Errorf("path is empty")
	}

	ext := filepath.Ext(path)

	result := &ParseResult{}

	switch ext {
	case ".osz":
		p.File.Osz.SetDir(path)
		if err := p.File.Osz.Parse(); err != nil {
			return nil, err
		}
		result.Osz = p.File.Osz
	case ".osr":
		p.File.Osr.SetDir(path)
		if err := p.File.Osr.Decode(); err != nil {
			return nil, err
		}
		result.Osr = p.File.Osr
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	return result, nil
}
