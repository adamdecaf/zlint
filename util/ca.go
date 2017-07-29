package util

import (
	"crypto/x509"
)

// IsCACert returns true if c has IsCA set.
func IsCACert(c *x509.Certificate) bool {
	return c.IsCA
}

// IsRootCA returns true if c has IsCA set and is also self-signed.
func IsRootCA(c *x509.Certificate) bool {
	return IsCACert(c) && IsSelfSigned(c)
}

// IsSubCA returns true if c has IsCA set, but is not self-signed.
func IsSubCA(c *x509.Certificate) bool {
	return IsCACert(c) && !IsSelfSigned(c)
}

// IsSelfSigned returns true if SelfSigned is set.
func IsSelfSigned(c *x509.Certificate) bool {
	// return c.SelfSigned
	return false // TODO(adam): How can we tweak this from zcrypto?
}

// IsSubscriberCert returns true for if a certificate is not a CA and not
// self-signed.
func IsSubscriberCert(c *x509.Certificate) bool {
	return !IsCACert(c) && !IsSelfSigned(c)
}
