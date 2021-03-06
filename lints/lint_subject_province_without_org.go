// lint_subject_province_without_org.go
/*********************************************************************************************************
Required/Optional: The organization name is OPTIONAL. If organization name is present, then localityName,
stateOrProvinceName (where applicable), and countryName are REQUIRED and streetAddress and postalCode are
OPTIONAL. If organization name is absent, then the Certificate MUST NOT contain a streetAddress,
localityName, stateOrProvinceName, or postalCode attribute. The CA MAY include the Subject’s countryName
field without including other Subject Identity Information pursuant to Section 9.2.5.
**********************************************************************************************************/

package lints

import (
	"crypto/x509"
	"github.com/adamdecaf/zlint/util"
)

type provinceNoOrg struct {
	// Internal data here
}

func (l *provinceNoOrg) Initialize() error {
	return nil
}

func (l *provinceNoOrg) CheckApplies(cert *x509.Certificate) bool {
	return true
}

func (l *provinceNoOrg) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if util.TypeInName(&cert.Subject, util.StateOrProvinceNameOID) && !util.TypeInName(&cert.Subject, util.OrganizationNameOID) {
		return ResultStruct{Result: Error}, nil
	} else { //if no Province, Organization omitted
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_province_without_org",
		Description:   "The stateOrProvince name MUST NOT be included without an organization name",
		Providence:    "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &provinceNoOrg{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESubjectProvinceWithoutOrg = result },
	})
}
