package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("`+name+`") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHelloNameInvalidDueNumber(t *testing.T) {
	msg, err := Hello("Test1")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("Test1") = %q, %v, want "", error`, msg, err)
	}
}

func TestHelloNameInvalidDueSpecialChar(t *testing.T) {
	msg, err := Hello("Test#")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("Test#") = %q, %v, want "", error`, msg, err)
	}
}
