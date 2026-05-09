package product

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/KeviinMoralees/crud-products/pkg/domains/entities"
	productRepo "github.com/KeviinMoralees/crud-products/pkg/domains/repository/product"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) productRepo.Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, product entities.Product) (entities.Product, error) {
	result, err := r.db.ExecContext(ctx,
		`INSERT INTO productos (nombre, descripcion, precio, stock) VALUES (?, ?, ?, ?)`,
		product.Nombre, product.Descripcion, product.Precio, product.Stock,
	)
	if err != nil {
		return entities.Product{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return entities.Product{}, err
	}
	product.ID = id
	return product, nil
}

func (r *repository) GetAll(ctx context.Context) ([]entities.Product, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, nombre, descripcion, precio, stock FROM productos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var p entities.Product
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if products == nil {
		products = []entities.Product{}
	}
	return products, nil
}

func (r *repository) GetByID(ctx context.Context, id int64) (entities.Product, error) {
	var p entities.Product
	err := r.db.QueryRowContext(ctx,
		`SELECT id, nombre, descripcion, precio, stock FROM productos WHERE id = ?`, id,
	).Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.Stock)
	if errors.Is(err, sql.ErrNoRows) {
		return entities.Product{}, errors.New("producto no encontrado")
	}
	return p, err
}

func (r *repository) Update(ctx context.Context, product entities.Product) (entities.Product, error) {
	result, err := r.db.ExecContext(ctx,
		`UPDATE productos SET nombre = ?, descripcion = ?, precio = ?, stock = ? WHERE id = ?`,
		product.Nombre, product.Descripcion, product.Precio, product.Stock, product.ID,
	)
	if err != nil {
		return entities.Product{}, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return entities.Product{}, err
	}
	if rows == 0 {
		return entities.Product{}, errors.New("producto no encontrado")
	}
	return product, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM productos WHERE id = ?`, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}
