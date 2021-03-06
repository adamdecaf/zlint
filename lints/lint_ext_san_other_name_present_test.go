// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSANOtherNamePresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANOtherName.pem"
	desEnum := Error
	out, _ := Lints["e_ext_san_other_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANOtherNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEDIParty.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_san_other_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
