package storage

import (
	"context"
	"database/sql"

	"github.com/valedaniel/golang-product-manager/internal/models"
)

type ProductStorage interface {
	Create(ctx context.Context, product *models.Product) error
	Get(ctx context.Context, id int) (*models.Product, error)
	List(ctx context.Context) ([]*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (storage *PostgresStorage) Create(ctx context.Context, product *models.Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, createdAt, updatedAt"

	err := storage.db.QueryRowContext(ctx, query, product.Name, product.Price).Scan(&product.Id, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (storage *PostgresStorage) Get(ctx context.Context, id int) (*models.Product, error) {
	query := "SELECT id, name, price, createdAt, updatedAt FROM products WHERE id = $1"

	product := &models.Product{}
	err := storage.db.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return product, nil
}

func (storage *PostgresStorage) List(ctx context.Context) ([]*models.Product, error) {
	query := "SELECT id, name, price, createdAt, updatedAt FROM products ORDER BY id ASC"

	rows, err := storage.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (storage *PostgresStorage) Update(ctx context.Context, product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, updatedAt = NOW() WHERE id = $3"

	rows, err := storage.db.ExecContext(ctx, query, product.Name, product.Price, product.Id)

	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (storage *PostgresStorage) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = $1"

	rows, err := storage.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
