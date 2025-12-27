package services

import (
	"app/internal/database"
	"app/internal/models"
	"context"

	"github.com/google/uuid"
)

type CertificateTLSService struct {
	db *database.DBCertificateTLS
}

func NewCertificateTLSService(db *database.DBCertificateTLS) *CertificateTLSService {
	return &CertificateTLSService{db: db}
}

func (c *CertificateTLSService) AddURL(ctx context.Context, url string) error {
	cert, err := tls_handshake(url)
	if err != nil {
		return err
	}
	_, err = c.db.SaveCertificate(ctx, cert)
	if err != nil {
		return err
	}
	return nil
}

func (c *CertificateTLSService) GetAll(ctx context.Context, page, limit int) ([]models.CertificateTLS, error) {
	offset := (page - 1) * limit
	return c.db.GetAllCertificates(ctx, limit, offset)
}

func (c *CertificateTLSService) GetByID(ctx context.Context, id uuid.UUID) (models.CertificateTLS, error) {
	return c.db.GetCertificateByID(ctx, id)
}

func (c *CertificateTLSService) GetByURL(url string) (models.CertificateTLS, error) {
	return tls_handshake(url)
}
