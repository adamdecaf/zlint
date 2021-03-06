// lint_br_ian_bare_wildcard.go

package lints

import (
	"crypto/x509"
	"github.com/adamdecaf/zlint/util"
	"strings"
)

type brIANBareWildcard struct {
	// Internal data here
}

func (l *brIANBareWildcard) Initialize() error {
	return nil
}

func (l *brIANBareWildcard) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *brIANBareWildcard) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.IANDNSNames {
		if strings.HasSuffix(dns, "*") {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ian_bare_wildcard",
		Description:   "A wildcard MUST be accompanied by other data to its right (Only checks DNSName)",
		Providence:    "awslabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &brIANBareWildcard{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.EIanBareWildcard = result },
	})
}
