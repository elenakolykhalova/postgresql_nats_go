package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Order struct {
	Order_uid string `json:"order_uid"`
	Data      string
}

func AddOrderTx(db *pgxpool.Pool, data Order) {
	ctx := context.Background()
	query := `insert into orders (order_uid, data) values ($1, $2)`
	_, err := db.Exec(
		ctx,
		query,
		data.Order_uid,
		data.Data,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func FillCache(InMap map[string]string) {
	restoreCacheConn, err := ConnectToDB()
	defer restoreCacheConn.Close()
	if err != nil {
		log.Fatal(err)
	}
	query := `select * from orders`
	rows, err := restoreCacheConn.Query(context.Background(), query)
	for rows.Next() {
		var key, data string
		err = rows.Scan(
			&key,
			&data,
		)
		if err != nil {
			log.Fatal(err)
		}
		InMap[string(key)] = string(data)
	}
	fmt.Printf("Cache is restored from DB. It has %d orders...\n", len(InMap))
}
