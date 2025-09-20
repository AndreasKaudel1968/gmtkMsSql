package mssql

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strconv"

	_ "github.com/microsoft/go-mssqldb"
)

type ParamDirection int

const (
	DirectionInput ParamDirection = iota
	DirectionOutput
)

type ProcParam[T any] struct {
	Name       string
	Direction  ParamDirection
	Value      T
	OutPointer any
}

//var db *sql.DB = nil

//var cnt int = 0

func RunProc(storedProcedure string, params *[]ProcParam[any]) (*sql.Rows, error) {

	var err error
	/*
		if db == nil {
			db, err = connect()
		} else {
			err = db.Ping()

			if err != nil {

				db, err = connect()
			}
		}
	*/
	db, err := connect()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var args []any
	for i, p := range *params {

		switch p.Direction {
		case DirectionInput:
			args = append(args, sql.Named(p.Name, p.Value))
		case DirectionOutput:
			//args = append(args, sql.Named(p.Name, sql.Out{Dest: &p.OutPointer}))
			args = append(args, sql.Named(p.Name, sql.Out{Dest: &(*params)[i].OutPointer}))
		default:
			fmt.Println("undefined direction in sql param list")
		}
	}

	rows, err := db.Query(storedProcedure, args...)

	//args = make([]any, 0)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func connect() (*sql.DB, error) {

	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbServer := os.Getenv("DATABASE_SERVER")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_DBNAME")
	dbAppName := os.Getenv("DATABASE_APPNAME")

	info := fmt.Sprintf("connect: try to start connection: Server=%s, Port=%s, DB=%s, User=%s", dbServer, dbPort, dbName, dbUser)

	fmt.Println(info)

	iDbPort, _ := strconv.Atoi(dbPort)

	query := url.Values{}
	query.Add("database", dbName)
	//query.Add("TrustServerCertificate", "true")
	query.Add("encrypt", "disable")

	if len(dbAppName) > 0 {
		//query.Add("app name", fmt.Sprintf("%s-%d", dbAppName, cnt))
		query.Add("app name", dbAppName)
	}

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(dbUser, dbPassword),
		Host:   fmt.Sprintf("%s:%d", dbServer, int(iDbPort)),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	//db.SetMaxIdleConns(50)
	//db.SetMaxOpenConns(50)

	return db, nil
}
