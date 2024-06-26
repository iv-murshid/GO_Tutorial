package if_else

import (
	"regexp"
	"testing"
)

func TestIfelse(t *testing.T){
	msg := 8
	want := regexp.MustCompile(`\b`+"even"+`\b`)
	res,err := If_else(msg)

	if !want.MatchString(res) || err!=nil {
		t.Fatalf(`Ifelse(8) = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}