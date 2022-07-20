package server

import (
	"log"
	"net/http"
	"time"

	v1 "github.com/dosanma1/go-rest-microservice/domain/controller/v1"
	"github.com/dosanma1/go-rest-microservice/domain/repository"
	"github.com/dosanma1/go-rest-microservice/domain/service"
	"github.com/dosanma1/go-rest-microservice/middleware"
	"github.com/gin-gonic/gin"
)

func (s *server) runHttpServer() {
	server := &http.Server{
		Addr:           s.cfg.HTTP.Port,
		Handler:        s.SetupRouter(),
		ReadTimeout:    s.cfg.HTTP.ReadTimeout * time.Second,
		WriteTimeout:   s.cfg.HTTP.WriteTimeout * time.Second,
		IdleTimeout:    s.cfg.HTTP.MaxConnectionIdle * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       log.New(s.log.Writer(), "", 0),
	}
	if err := server.ListenAndServe(); err != nil {
		s.log.Fatalln(err)
	}
}

func (s *server) SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	v1Route := r.Group("/api/v1")
	{
		productRoute := v1Route.Group("/product")
		{
			productRepository := repository.NewProductRepository(s.pgxPool, s.log)
			productService := service.NewProductService(productRepository, s.log)
			productController := v1.NewProductController(productService, s.log)

			productRoute.GET("/:productID", productController.GetProduct)
			productRoute.POST("/", productController.CreateProduct)
			productRoute.PUT("/", productController.UpdateProduct)
			productRoute.DELETE("/:productID", productController.DeleteProduct)
		}
	}
	return r
}
