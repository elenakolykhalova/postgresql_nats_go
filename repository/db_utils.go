package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"runtime"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	connstr := "postgres://mcarry:pass@localhost:5432/testmcarry"
	if connstr == "" {
		log.Fatal("не указан адрес подключения к БД")
	}
	db, err := pgxpool.Connect(context.Background(), connstr)
	return db, err
}

func ConnectWithMoreConn(connsToDBPerCore int) (*pgxpool.Pool, error) {
	connstr := "postgres://mcarry:pass@localhost:5432/testmcarry"
	config, err := pgxpool.ParseConfig(connstr)
	if err != nil {
		return nil, err
	}
	config.MaxConns = int32(runtime.NumCPU() * connsToDBPerCore)
	db, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("Connection to DB is not succeed %w", err)
	}
	return db, err
}
