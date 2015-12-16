package db

//import (
//	"testing"
//)

//func TestConnection(t *testing.T) {
//	cInfo := &ConnInfo {
//		Username:POSTGRES_USERNAME,
//		Password:POSTGRES_PASSWORD,
//		Host:POSTGRES_HOST,
//		Port:POSTGRES_PORT,
//		Dbname:POSTGRES_DBNAME,
//		Sslmode:POSTGRES_SSLMODE,
//	}
//	db, err := Connection(cInfo)
//	if err != nil {
//		t.Fatal(err)
//	}
//	_, err = db.Exec("select 1")
//	if err != nil {
//		t.Fatal(err)
//	}
//}

//func TestOpenDB(t *testing.T) {
//	db, err := OpenDB()
//	if err != nil {
//		t.Fatal(err)
//	}
	
//	_, err = db.Exec("select 1")
//	if err != nil {
//		t.Fatal(err)
//	}
//}