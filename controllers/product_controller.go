package controllers

import (
	"exhibitly/dto"
	"exhibitly/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	ProductService services.ProductService
}

// NewProductController cria uma nova instância de ProductController
func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

// CreateProduct cria um novo produto
// @Summary Criação de produto
// @Description Cria um novo produto com base nos dados fornecidos
// @Tags Products
// @Accept json
// @Produce json
// @Param        product body models.Product true "Detalhes do Produto"
// @Success 201 {object} models.Product
// @Router /products [post]
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var productDTO dto.CreateProductDto

	//Bind the body into productDto and handle possible errors
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	validate := validator.New()

	//Validate the binded data is valid, if not return formated errors
	if err := validate.Struct(productDTO); err != nil {

		validationErrors := err.(validator.ValidationErrors)

		for _, err := range validationErrors {
			fmt.Println("Campo:", err.Field())
			fmt.Println("Erro:", err.Tag())
			fmt.Println("Valor:", err.Value())
			fmt.Println("Parâmetro:", err.Param())
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return

	}

	product, err := ctrl.ProductService.CreateProduct(productDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o produto", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (ctrl *ProductController) GetProducts(c *gin.Context) {
	products, err := ctrl.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar os produtos", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (ctrl *ProductController) GetProduct(c *gin.Context) {
	productID := c.Param("id")
	product, err := ctrl.ProductService.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	c.JSON(http.StatusOK, product)
}
