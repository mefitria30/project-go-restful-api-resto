package main

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/kelas.work/project-go-restful-api/internal/database"
	"github.com/kelas.work/project-go-restful-api/internal/delivery/rest"
	"github.com/kelas.work/project-go-restful-api/internal/logger"
	mRepo "github.com/kelas.work/project-go-restful-api/internal/repository/menu"
	oRepo "github.com/kelas.work/project-go-restful-api/internal/repository/order"
	uRepo "github.com/kelas.work/project-go-restful-api/internal/repository/user"
	rUsecase "github.com/kelas.work/project-go-restful-api/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const(
	dbAddress = "host=localhost port=5432 user=postgres password=root dbname=go_resto_app sslmode=disable"
)

func main() {
	logger.Init()
	//tracing.Init("http:localhost:14268/api/traces") // should run docker first
	e := echo.New()
	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		panic(err)
	}
	
	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}
	restoUsecase := rUsecase.GetUseCase(menuRepo, orderRepo, userRepo)
	h := rest.NewHandler(restoUsecase)
	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}