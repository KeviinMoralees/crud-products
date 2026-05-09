package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	productRepo "github.com/KeviinMoralees/crud-products/pkg/repositories/product"
	productService "github.com/KeviinMoralees/crud-products/pkg/domains/product"
	"github.com/KeviinMoralees/crud-products/pkg/services/rest/routes"
)

func main() {
	db, err := sql.Open("sqlite3", "./productos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS productos (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre      TEXT    NOT NULL,
		descripcion TEXT,
		precio      REAL    NOT NULL,
		stock       INTEGER NOT NULL DEFAULT 0
	)`)
	if err != nil {
		log.Fatal(err)
	}

	repo := productRepo.NewRepository(db)
	svc := productService.NewService(repo)
	handler := routes.NewHandler(svc)

	handler.Run(":8080")
}
