package Switch

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	msg := 13
	want := regexp.MustCompile(`\b`+"It's after noon"+`\b`)
	res,err := Switch(msg)
	if !want.MatchString(res) || err!=nil {
		t.Fatalf(`Switch(13) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}