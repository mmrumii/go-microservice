package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmrumii/go-microservice/internal/database"
	"github.com/mmrumii/go-microservice/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllCustomers(ctx echo.Context) error
	AddCustomer(ctx echo.Context) error
	GetCustomerById(ctx echo.Context) error
	UpdateCustomer(ctx echo.Context) error
	DeleteCustomer(ctx echo.Context) error

	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error

	GetAllVendors(ctx echo.Context) error
	AddVendor(ctx echo.Context) error
	GetVendorById(ctx echo.Context) error

	GetAllServices(ctx echo.Context) error
	AddService(ctx echo.Context) error
	GetServiceById(ctx echo.Context) error
	UpdateService(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRouters()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occured: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRouters() {
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	cg := s.echo.Group("/customers")
	cg.GET("", s.GetAllCustomers)
	cg.POST("", s.AddCustomer)
	cg.GET("/:id", s.GetCustomerById)
	cg.PUT("/:id", s.UpdateCustomer)
	cg.DELETE("/:id", s.DeleteCustomer)

	pg := s.echo.Group("/products")
	pg.GET("", s.GetAllProducts)
	pg.POST("", s.AddProduct)
	pg.GET("/:id", s.GetProductById)

	vg := s.echo.Group("/vendors")
	vg.GET("", s.GetAllVendors)
	vg.POST("", s.AddVendor)
	vg.GET("/:id", s.GetVendorById)

	sg := s.echo.Group("/services")
	sg.GET("", s.GetAllServices)
	sg.POST("", s.AddService)
	sg.GET("/:id", s.GetServiceById)
	sg.PUT("/:id", s.UpdateService)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "Ok"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
