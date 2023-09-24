package database

import (
	"context"

	"github.com/mmrumii/go-microservice/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context, vendorId string) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Where(models.Vendor{VendorID: vendorId}).
		Find(&vendors)
	return vendors, result.Error
}
