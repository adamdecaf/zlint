// lint_subject_dn_trailing_whitespace.go

package lints

import (
	"crypto/x509"
	"github.com/adamdecaf/zlint/util"
)

type SubjectDNTrailingSpace struct{}

func (l *SubjectDNTrailingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNTrailingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	_, trailing, err := util.CheckRDNSequenceWhiteSpace(c.RawSubject)
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
		Name:          "w_subject_dn_trailing_whitespace",
		Description:   "AttributeValue in subject RelativeDistinguishedName sequence SHOULD NOT have trailing whitespace",
		Providence:    "aswlabs certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SubjectDNTrailingSpace{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.WSubjectDnTrailingWhitespace = result },
	})
}
