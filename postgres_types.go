package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
)

var (
	NotSupportedType = reflect.TypeOf(nil)
	BoolType         = reflect.TypeOf(false)
	ByteType         = reflect.TypeOf(byte(0))
	StringType       = reflect.TypeOf(string(""))
	IntType          = reflect.TypeOf(int(0))
	Int16Type        = reflect.TypeOf(int16(0))
	Int32Type        = reflect.TypeOf(int32(0))
	Int64Type        = reflect.TypeOf(int64(0))
	Float32Type      = reflect.TypeOf(float32(0))
	Float64Type      = reflect.TypeOf(float64(0))
	RawJSONType      = reflect.TypeOf(json.RawMessage{})
	TimeType         = reflect.TypeOf(time.Time{})
	DecimalType      = reflect.TypeOf(decimal.Decimal{})
	UUIDType         = reflect.TypeOf(uuid.UUID{})
)

var pgtypes = map[reflect.Type]struct {
	Basic     reflect.Type
	Null      reflect.Type
	Array     reflect.Type
	NullArray reflect.Type
}{
	BoolType: {
		Basic:     BoolType,
		Null:      reflect.TypeOf(null.Bool{}),
		Array:     reflect.TypeOf(pq.BoolArray{}),
		NullArray: reflect.TypeOf(pq.BoolArray{}),
	},
	ByteType: {
		Basic:     ByteType,
		Null:      reflect.TypeOf(null.Byte{}),
		Array:     reflect.TypeOf(null.Bytes{}),
		NullArray: reflect.TypeOf(null.Bytes{}),
	},
	StringType: {
		Basic:     StringType,
		Null:      reflect.TypeOf(null.String{}),
		Array:     reflect.TypeOf(pq.StringArray{}),
		NullArray: reflect.TypeOf(pq.StringArray{}),
	},
	IntType: {
		Basic:     IntType,
		Null:      reflect.TypeOf(null.Int{}),
		Array:     reflect.TypeOf(pq.Int64Array{}),
		NullArray: reflect.TypeOf(pq.Int64Array{}),
	},
	Int16Type: {
		Basic:     Int16Type,
		Null:      reflect.TypeOf(null.Int16{}),
		Array:     reflect.TypeOf(pq.Int32Array{}),
		NullArray: reflect.TypeOf(pq.Int32Array{}),
	},
	Int32Type: {
		Basic:     Int32Type,
		Null:      reflect.TypeOf(null.Int32{}),
		Array:     reflect.TypeOf(pq.Int32Array{}),
		NullArray: reflect.TypeOf(pq.Int32Array{}),
	},
	Int64Type: {
		Basic:     Int64Type,
		Null:      reflect.TypeOf(null.Int64{}),
		Array:     reflect.TypeOf(pq.Int64Array{}),
		NullArray: reflect.TypeOf(pq.Int64Array{}),
	},
	Float32Type: {
		Basic:     Float32Type,
		Null:      reflect.TypeOf(null.Float32{}),
		Array:     reflect.TypeOf(pq.Float32Array{}),
		NullArray: reflect.TypeOf(pq.Float32Array{}),
	},
	Float64Type: {
		Basic:     Float64Type,
		Null:      reflect.TypeOf(null.Float64{}),
		Array:     reflect.TypeOf(pq.Float64Array{}),
		NullArray: reflect.TypeOf(pq.Float64Array{}),
	},
	RawJSONType: {
		Basic:     RawJSONType,
		Null:      reflect.TypeOf(null.JSON{}),
		Array:     NotSupportedType,
		NullArray: NotSupportedType,
	},
	TimeType: {
		Basic:     TimeType,
		Null:      reflect.TypeOf(null.Time{}),
		Array:     NotSupportedType,
		NullArray: NotSupportedType,
	},
	DecimalType: {
		Basic:     DecimalType,
		Null:      reflect.TypeOf(decimal.NullDecimal{}),
		Array:     NotSupportedType,
		NullArray: NotSupportedType,
	},
	UUIDType: {
		Basic:     UUIDType,
		Null:      NotSupportedType,
		Array:     NotSupportedType,
		NullArray: NotSupportedType,
	},
}

func goType(ci postgresColumnInfo) reflect.Type {
	pgtype := func(typ reflect.Type) reflect.Type {
		if ci.IsArray && ci.Nullable {
			return pgtypes[typ].NullArray
		} else if ci.IsArray {
			return pgtypes[typ].Array
		} else if ci.Nullable {
			return pgtypes[typ].Null
		} else {
			return pgtypes[typ].Basic
		}
	}

	// remove array and keep only data type
	dataTypeName := strings.TrimRight(ci.DataTypeName, "[]")
	switch dataTypeName {
	case "bool":
		return pgtype(BoolType)
	case "character":
		return pgtype(ByteType)
	case "bit", "interval", "bit varying", "money", "character varying", "cidr", "inet", "macaddr", "macaddr8", "text", "xml", "varchar":
		return pgtype(StringType)
	case "smallint":
		return pgtype(Int16Type)
	case "integer", "serial":
		return pgtype(Int32Type)
	case "bigint", "bigserial":
		return pgtype(Int64Type)
	case "real":
		return pgtype(Float32Type)
	case "double precision":
		return pgtype(Float64Type)
	case "json", "jsonb":
		return pgtype(RawJSONType)
	case "date", "timestamp without time zone", "timestamp with time zone", "time without time zone", "time with time zone":
		return pgtype(TimeType)
	case "decimal", "numeric":
		return pgtype(DecimalType)
	case "uuid":
		return pgtype(UUIDType)
	}

	return NotSupportedType
}

/*
TODO:

 bigint
 bit
 bit varying
 boolean
 box
 bytea
 character
 character varying
 cid
 cidr
 circle
 citext
 cube
 date
 daterange
 double precision
 ean13
 gtsvector
 hstore
 inet
 int4range
 int8range
 integer
 interval
 isbn
 isbn13
 ismn
 ismn13
 issn
 issn13
 json
 jsonb
 jsonpath
 line
 lseg
 ltree
 macaddr
 macaddr8
 money
 name
 numeric
 numrange
 oid
 path
 pg_lsn
 point
 polygon
 real
 smallint
 text
 tid
 time with time zone
 time without time zone
 timestamp with time zone
 timestamp without time zone
 tsquery
 tsrange
 tstzrange
 tsvector
 txid_snapshot
 upc
 uuid
 xid
 xid8
 xml
*/
