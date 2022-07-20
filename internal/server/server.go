package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dosanma1/go-rest-microservice/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

const SERVICE_NAME = "stock"

type server struct {
	log     *logrus.Logger
	cfg     *config.Config
	pgxPool *pgxpool.Pool
}

func NewServer(log *logrus.Logger, cfg *config.Config, pgxPool *pgxpool.Pool) *server {
	return &server{
		log:     log,
		cfg:     cfg,
		pgxPool: pgxPool,
	}
}

func (s *server) Run() {

	go func() {
		s.log.Infof("Server is listening on PORT: %s", s.cfg.HTTP.Port)
		s.runHttpServer()
	}()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		done <- true
	}()

	<-done
	exit(0)
}

func exit(status int) {
	if status == 0 {
		fmt.Printf("%s service exit....\n", SERVICE_NAME)
	} else {
		fmt.Printf("%s  service with error code %d....\n", SERVICE_NAME, status)
	}
	os.Exit(status)
}
