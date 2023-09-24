package database

import (
	"context"

	"github.com/mmrumii/go-microservice/internal/models"
)

func (c Client) GetAllServices(ctx context.Context, serviceId string) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Where(models.Service{ServiceID: serviceId}).
		Find(&services)
	return services, result.Error
}
