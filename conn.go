package dbtest

import(
	"fmt"
	"errors"
	"database/sql"
	
	"github.com/gotools/logs"
	
	_"github.com/lib/pq"
)

type ConnInfo struct {
	Username string
	Password string
	Host string
	Port string
	Dbname string
	Sslmode string
}

func Connection(cInfo *ConnInfo) (*sql.DB, error) {
	if len(cInfo.Username) <= 0 {
		return nil, errors.New("username is empty")
	}
	
	if len(cInfo.Host) <= 0 {
		cInfo.Host = "127.0.0.1"
		logs.Debug("host is empty, default use 127.0.0.1")
	}
	
	if len(cInfo.Port) <= 0 {
		cInfo.Port = "5432"
		logs.Debug("port is empty, default use 5432")
	}
	
	if len(cInfo.Dbname) <= 0 {
		cInfo.Dbname = cInfo.Username
		logs.Debug("dbname is empty, defualt use user name: %s", cInfo.Username)
	}
	
	if len(cInfo.Sslmode) <= 0 {
		cInfo.Sslmode = "disable"
		logs.Debug("sslmode is empty, default use disable")
	}
	
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cInfo.Username, 
		cInfo.Password, cInfo.Host, cInfo.Port, cInfo.Dbname, cInfo.Sslmode)
	return sql.Open("postgres", connStr)
}

func OpenDB() (*sql.DB, error) {
	connStr := &ConnInfo{
		POSTGRES_USERNAME,
		POSTGRES_PASSWORD,
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_DBNAME,
		POSTGRES_SSLMODE, 
		}
	return Connection(connStr)
}