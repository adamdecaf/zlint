// lint_sub_ca_name_constraints_not_critical.go
/************************************************
Change this to match providence TEXT
************************************************/

package lints

import (
	"crypto/x509"
	"github.com/adamdecaf/zlint/util"
)

type SubCANameConstraintsNotCritical struct {
	// Internal data here
}

func (l *SubCANameConstraintsNotCritical) Initialize() error {
	return nil
}

func (l *SubCANameConstraintsNotCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsSubCA(cert) && util.IsExtInCert(cert, util.NameConstOID)
}

func (l *SubCANameConstraintsNotCritical) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if ski := util.GetExtFromCert(cert, util.NameConstOID); ski.Critical {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_ca_name_constraints_not_critical",
		Description:   "Subordinate CA certificate nameConstraints extension should be marked critical if present",
		Providence:    "CAB: 7.1.2.2",
		EffectiveDate: util.CABV102Date,
		Test:          &SubCANameConstraintsNotCritical{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.WSubCaNameConstraintsNotCritical = result },
	})
}
