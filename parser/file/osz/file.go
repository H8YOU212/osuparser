package osz

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/H8YOU212/osuparser/parser/file/osu"
)

type OszModel struct {
	dir      string
	baseDir  string
	Osufiles []osu.Commonosu
	Content  []string
	dst      string
}

func New() *OszModel {
	return &OszModel{
		baseDir: os.TempDir(),
	}
}

func (o *OszModel) Parse() error {
	if err := o.GetOsuFiles(); err != nil {
		return err
	}

	for i := range o.Osufiles {
		if o.Osufiles[i].Path == "" {
			continue
		}
		if err := o.Osufiles[i].ParseFile(o.Osufiles[i].Path); err != nil {
			return err
		}
	}

	return nil
}

func (o *OszModel) SetDir(dir string) {
	o.dir = dir
}

func (o *OszModel) SetBaseDir(baseDir string) {
	if baseDir != "" {
		o.baseDir = baseDir
	}
}

func (o *OszModel) GetOsuFiles() error {
	if o.dir == "" {
		return fmt.Errorf("osz path is empty")
	}

	r, err := zip.OpenReader(o.dir)
	if err != nil {
		return fmt.Errorf("open osz: %w", err)
	}
	defer r.Close()

	archiveName := strings.TrimSuffix(filepath.Base(o.dir), filepath.Ext(o.dir))
	outDir := filepath.Join(o.baseDir, archiveName)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir outdir: %w", err)
	}

	o.Osufiles = o.Osufiles[:0]
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}
		if strings.ToLower(filepath.Ext(f.Name)) != ".osu" {
			continue
		}

		src, err := f.Open()
		if err != nil {
			return fmt.Errorf("open osu in archive: %w", err)
		}

		outPath := filepath.Join(outDir, filepath.Base(f.Name))
		dst, err := os.Create(outPath)
		if err != nil {
			_ = src.Close()
			return fmt.Errorf("create extracted osu: %w", err)
		}

		if _, err := io.Copy(dst, src); err != nil {
			_ = dst.Close()
			_ = src.Close()
			return fmt.Errorf("copy osu data: %w", err)
		}

		_ = dst.Close()
		_ = src.Close()

		o.Osufiles = append(o.Osufiles, osu.Commonosu{
			Path: outPath,
		})
	}

	return nil
}
