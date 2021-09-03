package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"

	_ "github.com/lib/pq"
)

//go:embed templates/*
var gotemplates embed.FS

type dbInfo struct {
	PkgName string
}

type tableInfo struct {
	FileName        string
	CustomFileName  string
	PkgName         string
	StructName      string
	ShortStructName string
	StdPkgs         []string
	NonStdPkgs      []string

	Name    string
	Columns []columnInfo

	PrimaryKeys       []primaryKeyInfo
	HasPrimaryKey     bool
	HasManyPrimaryKey bool

	ForeignTables   []foreignTableInfo
	HasForeignTable bool
}

type primaryKeyInfo struct {
	ColumnName string
	ArgName    string
	GoType     string
}

type foreignTableInfo struct {
	FieldName  string
	StructName string
	Nullable   bool
}

type columnInfo struct {
	Name         string
	DBType       string
	IsPrimaryKey bool
	FieldName    string
	GoType       string
	Nullable     bool
}

func writeGoTmpl(tmplFile, file string, overwrite bool, data interface{}) error {
	path := filepath.Join(*outDir, file)
	if !overwrite {
		if _, err := os.Stat(path); err == nil {
			return nil
		}
	}

	content, err := gotemplates.ReadFile("templates/" + tmplFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New(tmplFile).Parse(string(content))
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	source, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("%s\n\n%w", buf.String(), err)
	}

	_, err = f.Write(source)
	return err
}
