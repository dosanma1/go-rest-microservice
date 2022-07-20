package v1

import (
	"net/http"

	"github.com/dosanma1/go-rest-microservice/domain/model"
	"github.com/dosanma1/go-rest-microservice/domain/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	productService service.ProductService
	log            *logrus.Logger
}

func NewProductController(productService service.ProductService, log *logrus.Logger) *ProductController {
	return &ProductController{
		productService: productService,
		log:            log,
	}
}

func (p ProductController) GetProduct(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("productID"))
	if err != nil {
		p.log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productResponse, err := p.productService.GetProduct(c, productID)
	if err != nil {
		p.log.Println(err)

		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, productResponse)
			return
		}

		c.JSON(http.StatusInternalServerError, productResponse)
		return
	}

	c.JSON(http.StatusOK, productResponse)
}

func (p ProductController) CreateProduct(c *gin.Context) {
	var productRequest model.ProductRequest
	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		p.log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid json",
		})
		return
	}

	productResponse, err := p.productService.CreateProduct(c, productRequest)
	if err != nil {
		p.log.Println(err)
		c.JSON(http.StatusInternalServerError, productResponse)
		return
	}

	c.JSON(http.StatusOK, productResponse)
}

func (p ProductController) UpdateProduct(c *gin.Context) {
	var productRequest model.ProductRequest
	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		p.log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid json",
		})
		return
	}

	productResponse, err := p.productService.UpdateProduct(c, productRequest)
	if err != nil {
		p.log.Println(err)
		c.JSON(http.StatusInternalServerError, productResponse)
		return
	}

	c.JSON(http.StatusOK, productResponse)
}

func (p ProductController) DeleteProduct(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("productID"))
	if err != nil {
		p.log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productResponse, err := p.productService.DeleteProduct(c, productID)
	if err != nil {
		p.log.Println(err)
		c.JSON(http.StatusInternalServerError, productResponse)
		return
	}

	c.JSON(http.StatusOK, productResponse)
}
