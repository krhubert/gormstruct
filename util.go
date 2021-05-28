package main

import (
	"reflect"

	"golang.org/x/tools/go/packages"
)

var isStandardPackages = make(map[string]bool)

func init() {
	pkgs, err := packages.Load(nil, "std")
	if err != nil {
		panic(err)
	}

	for _, p := range pkgs {
		isStandardPackages[p.PkgPath] = true
	}
}

func pkgPath(typ reflect.Type) string {
	if path := typ.PkgPath(); path != "main" {
		return path
	}
	return ""
}

func uniqueStrings(s []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, v := range s {
		if _, ok := keys[v]; !ok {
			keys[v] = true
			list = append(list, v)
		}
	}
	return list
}
