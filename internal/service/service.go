package service

import (
	"github.com/Astemirdum/lavka/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	*couriersService
	*ordersService
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		couriersService: newCouriersService(repo, log),
		ordersService:   newOrdersService(repo, log, repo.AssignBorder),
	}
}
