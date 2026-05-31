package file

import (
	"github.com/H8YOU212/osuparser/parser/file/osr"
	"github.com/H8YOU212/osuparser/parser/file/osz"
)

type File struct {
	Osz *osz.OszModel
	Osr *osr.OsrModel
}

func New() *File {
	return &File{
		Osz: osz.New(),
		Osr: osr.New(),
	}
}
