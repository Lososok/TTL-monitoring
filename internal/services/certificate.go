package services

import (
	"app/internal/database"
	"app/internal/models"
	"context"
	"errors"

	"github.com/google/uuid"
)

type CertificateTLSService struct {
	db *database.DBCertificateTLS
}

func NewCertificateTLSService(db *database.DBCertificateTLS) *CertificateTLSService {
	return &CertificateTLSService{db: db}
}

func (c *CertificateTLSService) AddURL(ctx context.Context, url string) error {
	if !isValidURL(url) {
		return ErrInvalidURL
	}

	cert, err := tlsHandshake(url)
	if err != nil {
		return ErrTLS
	}
	_, err = c.db.SaveCertificate(ctx, cert)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyExists) {
			return ErrAlreadyExists
		}
		return err
	}
	return nil
}

func (c *CertificateTLSService) GetAll(ctx context.Context, page, limit int) ([]models.CertificateTLS, error) {
	if page < 1 || limit < 1 {
		return nil, ErrInvalidArgs
	}
	offset := (page - 1) * limit
	return c.db.GetAllCertificates(ctx, limit, offset)
}

func (c *CertificateTLSService) GetByID(ctx context.Context, id uuid.UUID) (models.CertificateTLS, error) {
	cert, err := c.db.GetCertificateByID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return models.CertificateTLS{}, ErrNotFound
		}
		return models.CertificateTLS{}, err
	}

	return cert, nil
}

func (c *CertificateTLSService) GetByURL(url string) (models.CertificateTLS, error) {
	if !isValidURL(url) {
		return models.CertificateTLS{}, ErrInvalidURL
	}

	cert, err := tlsHandshake(url)
	if err != nil {
		return models.CertificateTLS{}, ErrTLS
	}

	return cert, nil
}
