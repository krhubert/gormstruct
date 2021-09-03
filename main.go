package main

import (
	"database/sql"
	"flag"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/iancoleman/strcase"
	"github.com/takuoki/gocase"
)

var (
	sqlDriver   = flag.String("dialect", "postgres", "database driver type (support only postgres)")
	sqlSchema   = flag.String("schema", "public", "database schema")
	sqlConnStr  = flag.String("conn", "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable", "database connection string")
	sqlDatabase = flag.String("d", "postgres", "database name")

	outDir        = flag.String("out", "model", "output directory")
	pkgName       = flag.String("pkg-name", "model", "go package name")
	dbFile        = flag.String("db-file", "db.go", "db implementation main file")
	baseModelFile = flag.String("base-model-file", "base_model.go", "base model implementation file")
	useBaseModel  = flag.Bool("use-base-model-file", false, "use base model implementation")
	customSuffix  = flag.String("custom", "custom", "suffix used for custom file")
	skipTables    = flag.String("skip-tables", "db_migration_latest", "comma separated table list to skip")

	omitJson = flag.String("omit-json", "created_at,updated_at,deleted_at", "comma separated columns list to omit json (set tag to '-')")
)

func main() {
	flag.Parse()

	skipTablesMap := make(map[string]bool)
	for _, t := range strings.Split(*skipTables, ",") {
		skipTablesMap[t] = true
	}

	omitJsonColumns := make(map[string]bool)
	for _, t := range strings.Split(*omitJson, ",") {
		omitJsonColumns[t] = true
	}

	dialect := postgresDialect{}

	db, err := sql.Open(*sqlDriver, *sqlConnStr)
	die(err)
	defer db.Close()

	tables, err := dialect.tables(db, *sqlSchema)
	die(err)

	var tableInfos []tableInfo

	for _, name := range tables {
		if skipTablesMap[name] {
			continue
		}

		primaryKeys, err := dialect.primaryKeys(db, *sqlSchema, name)
		die(err)

		columns, err := dialect.columns(db, *sqlSchema, name)
		die(err)

		foreignTableNames, err := dialect.foreignTables(db, *sqlSchema, name)
		die(err)

		hasBaseModel := 0
		if *useBaseModel {
			for _, c := range columns {
				if (c.Name == "id" && primaryKeys[c.Name]) ||
					(c.Name == "updated_at" && goType(c) == TimeType) ||
					(c.Name == "created_at" && goType(c) == TimeType) ||
					(c.Name == "deleted_at" && goType(c) == TimeType) {
					hasBaseModel++
				}
			}
		}

		ti := tableInfo{
			Name:              name,
			FileName:          name + ".go",
			CustomFileName:    name + "_" + *customSuffix + ".go",
			PkgName:           *pkgName,
			StructName:        filedName(name),
			ShortStructName:   argName(strcase.ToCamel(name))[:1],
			HasPrimaryKey:     len(primaryKeys) > 0,
			HasManyPrimaryKey: len(primaryKeys) > 1,
			HasBaseModel:      *useBaseModel && hasBaseModel == 4,
		}

		for _, c := range columns {
			goTyp := goType(c)
			if goTyp == NotSupportedType {
				continue
			}

			if primaryKeys[c.Name] {
				ti.PrimaryKeys = append(ti.PrimaryKeys, primaryKeyInfo{
					ColumnName: c.Name,
					ArgName:    argName(c.Name),
					GoType:     goTyp.String(),
				})
			}

			ftn := foreignTableNames[c.Name]
			ci := columnInfo{
				Name:         c.Name,
				DBType:       c.DataTypeName,
				IsPrimaryKey: primaryKeys[c.Name],
				FieldName:    filedName(c.Name),
				Nullable:     c.Nullable,
				GoType:       goTyp.String(),
				OmitJson:     omitJsonColumns[c.Name],
			}

			// special case when file name == TableName.
			// it's a reserved method for gorm interface.
			if ci.FieldName == "TableName" {
				ci.FieldName += "_"
			}

			if !ti.HasBaseModel || (ti.HasBaseModel && (c.Name != "id" &&
				c.Name != "updated_at" &&
				c.Name != "created_at" &&
				c.Name != "deleted_at")) {

				ti.Columns = append(ti.Columns, ci)

				if pkgPath := pkgPath(goTyp); pkgPath != "" {
					if isStandardPackages[pkgPath] {
						ti.StdPkgs = append(ti.StdPkgs, pkgPath)
					} else {
						ti.NonStdPkgs = append(ti.NonStdPkgs, pkgPath)
					}
				}
			}

			if ftn != "" {
				fieldName := strings.TrimSuffix(c.Name, "_id")
				if fieldName == c.Name {
					fieldName += "_rel"
				}
				ti.ForeignTables = append(ti.ForeignTables, foreignTableInfo{
					FieldName:  filedName(fieldName),
					StructName: filedName(ftn),
					Nullable:   c.Nullable,
				})
			}
		}

		ti.HasForeignTable = len(ti.ForeignTables) > 0
		ti.StdPkgs = uniqueStrings(ti.StdPkgs)
		ti.NonStdPkgs = uniqueStrings(ti.NonStdPkgs)
		tableInfos = append(tableInfos, ti)
	}

	die(os.MkdirAll(*outDir, 0755))

	die(writeGoTmpl("db.go.tmpl", *dbFile, true, dbInfo{PkgName: *pkgName}))
	if *useBaseModel {
		die(writeGoTmpl("base_model.go.tmpl", *baseModelFile, true, dbInfo{PkgName: *pkgName}))
	}

	for _, ti := range tableInfos {
		die(writeGoTmpl("model.go.tmpl", ti.FileName, true, ti))
		die(writeGoTmpl("model_custom.go.tmpl", ti.CustomFileName, false, ti))
	}
}

func filedName(name string) string {
	if name == "type" {
		name = "typ"
	}
	return gocase.To(strcase.ToCamel(name))
}

func argName(name string) string {
	if name == "type" {
		name = "typ"
	}
	return gocase.To(strcase.ToLowerCamel(name))
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
