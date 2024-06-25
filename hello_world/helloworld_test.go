package hello_world

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	msg := "Murshid"
	want := regexp.MustCompile(`\b`+msg+`\b`)
	res,err := Hello("Murshid")
	if !want.MatchString(res) || err!=nil {
		t.Fatalf(`Hello("Murshid") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}