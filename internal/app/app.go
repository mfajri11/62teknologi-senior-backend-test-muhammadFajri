package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/handler"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/repository"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/service"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/postgres"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func Run() {
	router := gin.Default()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	pg := postgres.New("postgresql://postgres:secret@172.17.0.1:5432/62teknologi?sslmode=disable")
	repo := repository.NewBusinessRepository(pg)
	svc := service.NewBusinessService(repo)
	handler := handler.NewBusinessHandler(svc)

	router.DELETE("/business_search/:id", handler.Delete())
	router.PUT("/business_search/:id", handler.Update())
	router.POST("/business_search", handler.Create())
	router.GET("/business_search", handler.Search())
	// TODO: add unit test if possible
	log.Fatal().Err(router.Run())
}
