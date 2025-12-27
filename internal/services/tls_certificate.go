package services

import (
	"app/internal/models"
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

func tls_handshake(url string) (models.CertificateTLS, error) {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	config := &tls.Config{
		// InsecureSkipVerify: *insecure,
		RootCAs: rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.CertificateTLS{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.CertificateTLS{}, err
	}
	defer resp.Body.Close()
	return models.CertificateTLS{
		URL:          url,
		Version:      resp.TLS.PeerCertificates[0].Version,
		NotBefore:    resp.TLS.PeerCertificates[0].NotBefore,
		NotAfter:     resp.TLS.PeerCertificates[0].NotAfter,
		Subject:      resp.TLS.PeerCertificates[0].Subject.CommonName,
		Issuer:       resp.TLS.PeerCertificates[0].Issuer.CommonName,
		SerialNumber: resp.TLS.PeerCertificates[0].SerialNumber.String(),
	}, nil
}
