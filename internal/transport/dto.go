package transport

import (
	"app/internal/models"
	"time"

	"github.com/google/uuid"
)

type AddURLRequest struct {
	URL string `json:"url" example:"example.com"`
}

type CreateCertificateResponse struct {
	ID uuid.UUID `json:"id"`
}

type PaginationQuery struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CertificateResponse struct {
	ID           uuid.UUID `json:"id"`
	URL          string    `json:"url"`
	SerialNumber string    `json:"serial_number"`
	NotAfter     time.Time `json:"not_after"`
	TTLDays      int64     `json:"ttl_days"`
}

type CertificatesListResponse struct {
	Data []CertificateResponse `json:"data"`
	Meta struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}

func NewCertificatesListResponse(certs []models.CertificateTLS, page, limit int) CertificatesListResponse {
	response := CertificatesListResponse{
		Data: make([]CertificateResponse, 0, len(certs)),
		Meta: struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
		}{
			Page:  page,
			Limit: limit,
		},
	}

	for _, cert := range certs {
		response.Data = append(response.Data, toResponse(cert))
	}

	return response
}

func toResponse(cert models.CertificateTLS) CertificateResponse {
	ttl := time.Until(cert.NotAfter)

	return CertificateResponse{
		ID:           cert.ID,
		URL:          cert.URL,
		SerialNumber: cert.SerialNumber,
		NotAfter:     cert.NotAfter,
		TTLDays:      int64(ttl.Hours() / 24),
	}
}
