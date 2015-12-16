package dbtest

import (
	"github.com/gotools/logs"
)

type PgTableFieldLength struct {
	TypeOid Oid
	FieldLength Oid
}

var PostgresTabFieldLen = [] PgTableFieldLength{
	{T_varchar, 1024},
}

const (
	POSTGRES_VALUE_TYPE_BINARY byte = iota + 1
	POSTGRES_VALUE_TYPE_STRING 
)

type PgValueType struct {
	TypeOid Oid
	TypeValueFormat byte
}

var PostgresValueType = [] PgValueType{
	{T_varchar, POSTGRES_VALUE_TYPE_STRING},
}

func ShowAllType() {
	for index, value := range PostgresType {
		logs.Debug("index: %d, name: %s, id: %d", index, value.TypeName, value.TypeOid)
		if value.TypeOid == 0 {
			break
		}
	}
}