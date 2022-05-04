package main

import (
	"github.com/eltaljohn/go-db-gorm/model"
	"github.com/eltaljohn/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	invoice := model.InvoiceHeader{
		Client: "Alvaro Felipe",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.DB().Create(&invoice)
}
