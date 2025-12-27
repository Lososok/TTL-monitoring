package models

import (
	"time"

	"github.com/google/uuid"
)

type CertificateTLS struct {
	ID           uuid.UUID
	URL          string
	Version      int
	NotBefore    time.Time
	NotAfter     time.Time
	Subject      string
	Issuer       string
	SerialNumber string
}
