package webrtc

import (
	"crypto"
	"crypto/x509"
	"time"
)

type Certificate struct {
	privateKey crypto.PrivateKey
	x509Cert   *x509.Certificate
	statsID    string
}

func NewCertificate(key crypto.PrivateKey, tpl x509.Certificate) (*Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Certificate) Equals(cert Certificate) bool { _ = "STUB: not implemented"; return false }

func (c Certificate) Expires() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c Certificate) GetFingerprints() ([]DTLSFingerprint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateCertificate(secretKey crypto.PrivateKey) (*Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CertificateFromX509(privateKey crypto.PrivateKey, certificate *x509.Certificate) Certificate {
	_ = "STUB: not implemented"
	return *new(Certificate)
}

func (c Certificate) collectStats(report *statsReportCollector) error {
	_ = "STUB: not implemented"
	return nil
}

func CertificateFromPEM(pems string) (*Certificate, error) {
	_ = "STUB: not implemented" //nolint: cyclop
	return nil, nil
}

func (c Certificate) PEM() (string, error) { _ = "STUB: not implemented"; return "", nil }
