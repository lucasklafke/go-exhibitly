package repositories

import (
	"exhibitly/models"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll() ([]models.Product, error)
	FindByID(productID string) (*models.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

// NewProductRepository cria uma nova instância de productRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

// Create insere um produto no banco de dados
func (r *productRepository) Create(product *models.Product) error {
	fmt.Print(*product)
	return r.DB.Create(product).Error
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByID(productID string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.First(&product, productID).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
