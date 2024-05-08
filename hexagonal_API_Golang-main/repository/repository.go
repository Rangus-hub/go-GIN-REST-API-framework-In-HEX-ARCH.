/*
	"github.com/jackc/pgx/v4/pgxpool" for ustilizing
    	return r.pool.Exec(ctx, query, args...)

    doing go get github.com/jackc/pgx/v4/pgxpool@v4.18.3
*/

package repository

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EmployeeRepository interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

type PostgresEmployeeRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresEmployeeRepository(pool *pgxpool.Pool) *PostgresEmployeeRepository {
	return &PostgresEmployeeRepository{pool}
}

func (r *PostgresEmployeeRepository) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return r.pool.Query(ctx, query, args...) //sending *pgx.Row instead of pgx.Rows gives errors in handler
	//WOULD NOT ABLE TO PERFORM rows.Close() in handler
}

func (r *PostgresEmployeeRepository) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return r.pool.QueryRow(ctx, query, args...)
}

func (r *PostgresEmployeeRepository) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return r.pool.Exec(ctx, query, args...)
}
