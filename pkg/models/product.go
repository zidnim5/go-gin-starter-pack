package models

import (
	"Starter/config"
	"errors"

	"github.com/jinzhu/gorm"
)

type Product struct {
	Slug         string
	Nama         string
	Quantity     int
	StatusBarang string
}

func (product *Product) TableName() string {
	return "product"
}

func CreateProduct(product *Product) (err error) {
	if err = config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (s *Product) UpdateProduct(barang Product, slug string) (err error) {
	check := config.DB.Where("slug =?", slug).First(&barang).Error

	if gorm.IsRecordNotFoundError(check) {
		return errors.New("Data not found")
	}

	err = config.DB.Model(s).Where("slug = ?", slug).Updates(s).Error
	return err
	// return
}
