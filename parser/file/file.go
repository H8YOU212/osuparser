package file

import (
	"osuparser/parser/file/osr"
	"osuparser/parser/file/osz"
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
