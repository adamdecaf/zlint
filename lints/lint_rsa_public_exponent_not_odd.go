// lint_rsa_public_exponent_not_odd.go
/*******************************************************************************************************
"CAB: 6.1.6"
RSA: The CA SHALL confirm that the value of the public exponent is an odd number equal to 3 or more. Additionally, the public exponent SHOULD be in the range between 2^16+1 and 2^256-1. The modulus SHOULD also have the following characteristics: an odd number, not the power of a prime, and have no factors smaller than 752. [Source: Section 5.3.3, NIST SP 800-89].
*******************************************************************************************************/

package lints

import (
	"crypto/rsa"
	"crypto/x509"
	"github.com/adamdecaf/zlint/util"
)

type rsaParsedTestsKeyExpOdd struct {
	// Internal data here
}

func (l *rsaParsedTestsKeyExpOdd) Initialize() error {
	return nil
}

func (l *rsaParsedTestsKeyExpOdd) CheckApplies(c *x509.Certificate) bool {
	_, ok := c.PublicKey.(*rsa.PublicKey)
	return ok && c.PublicKeyAlgorithm == x509.RSA
}

func (l *rsaParsedTestsKeyExpOdd) RunTest(c *x509.Certificate) (ResultStruct, error) {
	key := c.PublicKey.(*rsa.PublicKey)
	if key.E%2 == 1 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_rsa_public_exponent_not_odd",
		Description:   "RSA public key has to be an odd number",
		Providence:    "CAB: 6.1.6",
		EffectiveDate: util.CABV113Date,
		Test:          &rsaParsedTestsKeyExpOdd{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ERsaPublicExponentNotOdd = result },
	})
}
