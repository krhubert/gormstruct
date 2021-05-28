package main

import "database/sql"

type postgresColumnInfo struct {
	Name         string
	DataType     string
	Nullable     bool
	IsArray      bool
	DataTypeName string
}

type postgresDialect struct{}

func (postgresDialect) tables(db *sql.DB, tableSchema string) ([]string, error) {
	const query = `
		SELECT
			table_name
		FROM
			information_schema.tables
		WHERE
			(
				table_type = 'BASE TABLE' OR
				table_type = 'VIEW'
			) AND table_schema = $1 
		ORDER BY
			table_name`

	rows, err := db.Query(query, tableSchema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

func (postgresDialect) primaryKeys(db *sql.DB, tableSchema, tableName string) ([]string, error) {
	const query = `
		SELECT
			kcu.column_name
		FROM
			information_schema.table_constraints tco
		JOIN
			information_schema.key_column_usage kcu
		ON kcu.constraint_name = tco.constraint_name AND
			kcu.constraint_schema = tco.constraint_schema AND
			kcu.constraint_name = tco.constraint_name
		WHERE
			tco.constraint_type = 'PRIMARY KEY' AND
			kcu.table_schema = $1 AND
			kcu.table_name = $2
		ORDER BY
			kcu.ordinal_position`

	rows, err := db.Query(query, tableSchema, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

func (postgresDialect) columns(db *sql.DB, tableSchema, tableName string) ([]postgresColumnInfo, error) {
	const query = `
		SELECT
			column_name,
			data_type,
			is_nullable = 'YES' as nullable,
			data_type = 'ARRAY' as is_array,
			udt_name::regtype as udt_name
		FROM information_schema.columns
		WHERE 
			table_schema = $1 AND
			table_name = $2
		ORDER BY
			column_name`

	rows, err := db.Query(query, tableSchema, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cols []postgresColumnInfo

	for rows.Next() {
		var name, dataType, dataTypeName string
		var nullable, isArray bool

		if err = rows.Scan(&name, &dataType, &nullable, &isArray, &dataTypeName); err != nil {
			return nil, err
		}
		cols = append(cols, postgresColumnInfo{
			Name:         name,
			DataType:     dataType,
			Nullable:     nullable,
			IsArray:      isArray,
			DataTypeName: dataTypeName,
		})
	}
	return cols, nil
}

func (postgresDialect) foreignTables(db *sql.DB, tableSchema, tableName string) (map[string]string, error) {
	const query = `
    SELECT
      kcu.column_name,
      ccu.table_name AS foreign_table_name
    FROM
      information_schema.table_constraints AS tc
        JOIN information_schema.key_column_usage AS kcu
          ON tc.constraint_name = kcu.constraint_name
          AND tc.table_schema = kcu.table_schema
        JOIN information_schema.constraint_column_usage AS ccu
          ON ccu.constraint_name = tc.constraint_name
          AND ccu.table_schema = tc.table_schema
    WHERE
      tc.constraint_type = 'FOREIGN KEY' AND
      tc.table_schema = $1 AND
      ccu.table_schema = $1 AND
      tc.table_name= $2`

	rows, err := db.Query(query, tableSchema, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foreignTableNames := make(map[string]string, 0)
	for rows.Next() {
		var col, ftable string
		if err := rows.Scan(&col, &ftable); err != nil {
			return nil, err
		}
		foreignTableNames[col] = ftable
	}
	return foreignTableNames, nil
}
