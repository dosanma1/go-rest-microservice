package test

import (
	"log"

	"github.com/dosanma1/go-rest-microservice/config"
	"github.com/dosanma1/go-rest-microservice/internal/postgresql"
	"github.com/dosanma1/go-rest-microservice/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func setup() (*gin.Engine, error) {
	var logger = logrus.New()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgxPool, err := postgresql.NewPgxConn(cfg)
	if err != nil {
		logger.Fatalf("NewPgxConn: %+v", err)
	}
	logger.Printf("PostgreSQL connected: %+v", pgxPool.Stat().TotalConns())

	s := server.NewServer(logger, cfg, pgxPool)

	router := s.SetupRouter()

	return router, err
}
