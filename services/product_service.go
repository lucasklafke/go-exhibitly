package services

import (
	"exhibitly/dto"
	"exhibitly/models"
	"exhibitly/repositories"
)

// ProductService é a interface para o serviço de produto
type ProductService interface {
	CreateProduct(productDTO dto.CreateProductDto) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	GetProduct(productID string) (*models.Product, error)
}

// productService é a implementação do ProductService
type productService struct {
	Repository repositories.ProductRepository
}

// NewProductService cria uma nova instância do productService
func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{Repository: repo}
}

// CreateProduct cria um novo produto no banco de dados
func (s *productService) CreateProduct(productDTO dto.CreateProductDto) (*models.Product, error) {
	product := models.Product{
		Title:       productDTO.Title,
		Price:       productDTO.Price,
		Description: productDTO.Description,
		Image:       productDTO.Image,
		// Preencha com outros campos conforme necessário
	}

	// Inserir o produto no banco de dados
	if err := s.Repository.Create(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *productService) GetProducts() ([]models.Product, error) {
	products, err := s.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) GetProduct(productID string) (*models.Product, error) {
	product, err := s.Repository.FindByID(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
