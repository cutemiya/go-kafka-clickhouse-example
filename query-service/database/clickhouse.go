package database

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func NewClickhouse(ctx context.Context, connectionString string) (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr:     []string{connectionString},
		Protocol: clickhouse.HTTP,
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//TLS: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
	})
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, err
}
