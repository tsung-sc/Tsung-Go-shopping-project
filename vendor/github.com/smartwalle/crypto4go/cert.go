package crypto4go

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const (
	kCertificatePrefix = "-----BEGIN CERTIFICATE-----"
	kCertificateSuffix = "-----END CERTIFICATE-----"
)

var (
	ErrCertificateFailedToLoad = errors.New("crypto4go: certificate failed to load")
)

func FormatCertificate(raw string) []byte {
	return formatKey(raw, kCertificatePrefix, kCertificateSuffix, 76)
}

func ParseCertificate(b []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, ErrCertificateFailedToLoad
	}
	csr, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return csr, nil
}
