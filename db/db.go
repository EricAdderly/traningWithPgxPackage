package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateDbPool(url string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// rows, err := dbpool.Query(context.Background(), "select uuid, expiration_date from public.auth_tokens where uuid = '86793620-1e64-11ee-be56-0242ac120002'")
	// if err != nil {
	// 	fmt.Println(123)
	// 	return nil, err
	// }
	// sessionDb := Session{}

	// for rows.Next() {

	// 	err = rows.Scan(&sessionDb.UUID, &sessionDb.ExpirationDate)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("Authrisation - CheckSessionsData - rows.Scan: %w", err)
	// 	}

	// }

	return dbpool, nil
}

// type Session struct {
// 	UUID           string
// 	AccessToken    string
// 	ExpirationDate time.Time
// 	CreatedDate    time.Time
// }
