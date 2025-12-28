package services

import (
	"app/internal/database"
	"app/internal/models"
	"context"

	"github.com/google/uuid"
)

type SertificateInterface interface {
	AddURL(ctx context.Context, url string) (uuid.UUID, error)
	GetAll(ctx context.Context, page, limit int) ([]models.CertificateTLS, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.CertificateTLS, error)
	GetByURL(url string) (models.CertificateTLS, error)
}

type Service struct {
	SertificateInterface
}

func NewService(db *database.DataBase) *Service {
	return &Service{
		SertificateInterface: NewCertificateTLSService(db),
	}
}
