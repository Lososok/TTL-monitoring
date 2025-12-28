package database

import (
	"app/internal/models"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type SertificateInterface interface {
	SaveCertificate(ctx context.Context, cert models.CertificateTLS) (uuid.UUID, error)
	GetAllCertificates(ctx context.Context, limit, offset int) ([]models.CertificateTLS, error)
	GetCertificateByID(ctx context.Context, id uuid.UUID) (models.CertificateTLS, error)
}

type DataBase struct {
	SertificateInterface
}

func NewDataBase(db *sql.DB) *DataBase {
	return &DataBase{
		SertificateInterface: NewDBCertificateTLS(db),
	}
}
