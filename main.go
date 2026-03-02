package main
import (	
	"github.com/gofiber/fiber/v2"
	"shellrean.id/Go-RestAPI/internal/api"
	"shellrean.id/Go-RestAPI/internal/config"
	"shellrean.id/Go-RestAPI/internal/connection"
	"shellrean.id/Go-RestAPI/internal/repository"
	"shellrean.id/Go-RestAPI/internal/service"
)

func main(){
	cnf:= config.Get()
	dbConnection:= connection.GetDatabase(cnf.Database)
	app:= fiber.New()
	customerRepository := repository.NewCustomer(dbConnection)
	
	customerService := service.NewCustomer(customerRepository)
	api.NewCustomer(app, customerService)
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

