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
	Name            string
	FileName        string
	CustomFileName  string
	PkgName         string
	StructName      string
	ShortStructName string
	PrimaryKey      string
	PkGoType        string
	Columns         []columnInfo
	ForeignTables   []foreignTableInfo

	// imports
	StdPkgs    []string
	NonStdPkgs []string

	// foreign table
	HasForeignTable bool
	HasPrimaryKey   bool
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

func writeGoTmpl(tmplFile, file string, force bool, data interface{}) error {
	content, err := gotemplates.ReadFile("templates/" + tmplFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New(tmplFile).Parse(string(content))
	if err != nil {
		return err
	}

	perms := os.O_RDWR | os.O_CREATE
	if force {
		perms |= os.O_TRUNC
	}

	f, err := os.OpenFile(filepath.Join(*outDir, file), perms, 0666)
	if !force && os.IsNotExist(err) {
		return nil
	}
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
