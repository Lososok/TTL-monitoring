package database

import (
	"app/internal/models"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

type DBCertificateTLS struct {
	db *sql.DB
}

func NewDBCertificateTLB(db *sql.DB) *DBCertificateTLS {
	return &DBCertificateTLS{db: db}
}

func (db *DBCertificateTLS) SaveCertificate(ctx context.Context, cert models.CertificateTLS) (uuid.UUID, error) {
	var id uuid.UUID
	err := db.db.QueryRowContext(
		ctx,
		`INSERT INTO certificate_tls (
			url,
			version,
			not_before,
			not_after,
			c_subject,
			c_issuer,
			serial_number
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
		RETURNING id;`,
		cert.URL,
		cert.Version,
		cert.NotBefore,
		cert.NotAfter,
		cert.Subject,
		cert.Issuer,
		cert.SerialNumber,
	).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return uuid.Nil, ErrAlreadyExists
			}
		}
		return uuid.Nil, err
	}
	return id, nil
}

func (db *DBCertificateTLS) GetAllCertificates(ctx context.Context, limit, offset int) ([]models.CertificateTLS, error) {
	rows, err := db.db.QueryContext(
		ctx,
		`SELECT id, url, version, not_before, not_after, c_subject, c_issuer, serial_number
		FROM certificate_tls
		ORDER BY not_after ASC
		LIMIT $1 OFFSET $2`,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var certs []models.CertificateTLS
	for rows.Next() {
		var cert models.CertificateTLS
		if err := rows.Scan(
			&cert.ID,
			&cert.URL,
			&cert.Version,
			&cert.NotBefore,
			&cert.NotAfter,
			&cert.Subject,
			&cert.Issuer,
			&cert.SerialNumber,
		); err != nil {
			return nil, err
		}
		certs = append(certs, cert)
	}

	return certs, nil
}

func (db *DBCertificateTLS) GetCertificateByID(ctx context.Context, id uuid.UUID) (models.CertificateTLS, error) {
	var cert models.CertificateTLS
	err := db.db.QueryRowContext(
		ctx,
		`SELECT id, url, version, not_before, not_after, c_subject, c_issuer, serial_number
		FROM certificate_tls
		WHERE id = $1`,
		id,
	).Scan(
		&cert.ID,
		&cert.URL,
		&cert.Version,
		&cert.NotBefore,
		&cert.NotAfter,
		&cert.Subject,
		&cert.Issuer,
		&cert.SerialNumber,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.CertificateTLS{}, ErrNotFound
		}
		return models.CertificateTLS{}, err
	}

	return cert, nil
}
