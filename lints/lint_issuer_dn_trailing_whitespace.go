// lint_issuer_dn_trailing_whitespace.go

package lints

import (
	"crypto/x509"
	"github.com/adamdecaf/zlint/util"
)

type IssuerDNTrailingSpace struct{}

func (l *IssuerDNTrailingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNTrailingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	_, trailing, err := util.CheckRDNSequenceWhiteSpace(c.RawIssuer)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	if trailing {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_trailing_whitespace",
		Description:   "AttributeValue in issuer RelativeDistinguishedName sequence SHOULD NOT have trailing whitespace",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerDNTrailingSpace{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.WIssuerDnTrailingWhitespace = result },
	})
}
