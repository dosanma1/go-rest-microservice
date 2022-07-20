package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/dosanma1/go-rest-microservice/domain/model"
)

type ProductRepository interface {
	Get(c *gin.Context, productID uuid.UUID) (model.Product, error)
	Create(c *gin.Context, product model.Product) (uuid.UUID, error)
	Update(c *gin.Context, product model.Product) error
	Delete(c *gin.Context, productID uuid.UUID) error
}

type productRepository struct {
	db  *pgxpool.Pool
	log *logrus.Logger
}

func NewProductRepository(db *pgxpool.Pool, log *logrus.Logger) *productRepository {
	return &productRepository{
		db:  db,
		log: log,
	}
}

func (p productRepository) Get(c *gin.Context, productID uuid.UUID) (model.Product, error) {
	var product model.Product

	err := p.db.QueryRow(c.Request.Context(), "SELECT product_id, name, quantity FROM products WHERE product_id = $1", productID).Scan(&product.ProductID, &product.Name, &product.Quantity)
	if err != nil {
		p.log.Println(err)
		return model.Product{}, err
	}

	return product, err
}

func (p productRepository) Create(c *gin.Context, product model.Product) (uuid.UUID, error) {
	var productId uuid.UUID

	err := p.db.QueryRow(c.Request.Context(), "INSERT INTO products (name, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING product_id", product.Name, product.Quantity, product.CreatedAt, product.UpdatedAt).Scan(&productId)
	if err != nil {
		p.log.Println(err)
	}

	return productId, err
}

func (p productRepository) Update(c *gin.Context, product model.Product) error {
	_, err := p.db.Exec(c.Request.Context(), "UPDATE products SET name = $1, quantity = $2, updated_at = $3 WHERE product_id = $4", product.Name, product.Quantity, product.UpdatedAt, product.ProductID)
	if err != nil {
		p.log.Println(err)
	}

	return err
}

func (p productRepository) Delete(c *gin.Context, productID uuid.UUID) error {
	_, err := p.db.Exec(c.Request.Context(), "DELETE FROM products WHERE product_id = $1", productID)
	if err != nil {
		p.log.Println(err)
	}

	return err
}
