package service

import (
	"time"

	"github.com/dosanma1/go-rest-microservice/domain/model"
	"github.com/dosanma1/go-rest-microservice/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProductService interface {
	GetProduct(c *gin.Context, productID uuid.UUID) (*model.ProductResponse, error)
	CreateProduct(c *gin.Context, productRequest model.ProductRequest) (*model.ProductResponse, error)
	UpdateProduct(c *gin.Context, productRequest model.ProductRequest) (*model.ProductResponse, error)
	DeleteProduct(c *gin.Context, productID uuid.UUID) (*model.ProductResponse, error)
}

type productService struct {
	productRepository repository.ProductRepository
	log               *logrus.Logger
}

func NewProductService(productRepo repository.ProductRepository, log *logrus.Logger) *productService {
	return &productService{
		productRepository: productRepo,
		log:               log,
	}
}

func (p productService) GetProduct(c *gin.Context, productID uuid.UUID) (*model.ProductResponse, error) {
	product, err := p.productRepository.Get(c, productID)
	if err != nil {
		return &model.ProductResponse{
			Status: "error",
		}, err
	}

	return &model.ProductResponse{
		Status:   "success",
		Response: &product,
	}, err
}

func (p productService) CreateProduct(c *gin.Context, productRequest model.ProductRequest) (*model.ProductResponse, error) {

	newProduct := model.Product{
		Name:      productRequest.Name,
		Quantity:  productRequest.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	productID, err := p.productRepository.Create(c, newProduct)
	if err != nil {
		return &model.ProductResponse{
			Status: "error",
		}, err
	}

	newProduct.ProductID = productID

	return &model.ProductResponse{
		Status:   "success",
		Response: &newProduct,
	}, err
}

func (p productService) UpdateProduct(c *gin.Context, productRequest model.ProductRequest) (*model.ProductResponse, error) {

	updatedProduct := model.Product{
		ProductID: productRequest.ProductID,
		Name:      productRequest.Name,
		Quantity:  productRequest.Quantity,
		UpdatedAt: time.Now(),
	}

	err := p.productRepository.Update(c, updatedProduct)
	if err != nil {
		return &model.ProductResponse{
			Status: "error",
		}, err
	}

	return &model.ProductResponse{
		Status:   "success",
		Response: &updatedProduct,
	}, err
}

func (p productService) DeleteProduct(c *gin.Context, productID uuid.UUID) (*model.ProductResponse, error) {
	err := p.productRepository.Delete(c, productID)
	if err != nil {
		return &model.ProductResponse{
			Status: "error",
		}, err
	}

	return &model.ProductResponse{
		Status: "success",
	}, err
}
