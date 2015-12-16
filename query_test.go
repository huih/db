package db

import (
	"testing"
//	"testcase"
)

//func TestSimpleQuery(t *testing.T){
//	for _, sql := range testcase.SimpleSql {
//		err := SimpleQuery(sql, SkipQueryResult)
//		if err != nil {
//			t.Errorf("%s: %s", err.Error(), sql)
//		}
//	}
//}

func TestSimpleQuery_1(t *testing.T){
//	sql := "set client_encoding=utf8;select * from pgbench_tellers;"
	sql := "set client_encoding=utf8;select * from bigtable;"
	err := SimpleQuery(sql, OutputQueryResult)
	if err != nil {
		t.Errorf("%s: %s", err.Error(), sql)
	}
}

//func TestSimpleQuery_trans(t *testing.T){
//	for _, sql := range testcase.SimpleSql_trans {
//		err := TxSimpleQuery(sql, OutputQueryResult)
//		if err != nil {
//			t.Errorf("%s: %s", err.Error(), sql)
//		}
//	}
//}

//func TestPreparedSimpleQuery(t *testing.T){
//	for _, sql := range testcase.Prepared_simpleQuery {
//		err := PreparedQuery(sql, SkipQueryResult)
//		if err != nil {
//			t.Errorf("%s: %s", err.Error(), sql)
//		}
//	}
//}