package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmrumii/go-microservice/internal/dberrors"
	"github.com/mmrumii/go-microservice/internal/models"
)

func (s *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendorId := ctx.QueryParam("vendorId")

	vendors, err := s.DB.GetAllVendors(ctx.Request().Context(), vendorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

func (s *EchoServer) AddVendor(ctx echo.Context) error {
	vendor := new(models.Vendor)

	if err := ctx.Bind(&vendor); err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	vendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, vendor)
}

func (s *EchoServer) GetVendorById(ctx echo.Context) error {
	ID := ctx.Param("id")

	vendor, err := s.DB.GetVendorById(ctx.Request().Context(), ID)

	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, vendor)
}
