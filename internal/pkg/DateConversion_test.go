package pkg

import (
	"testing"
)

func TestDateTime(t *testing.T) {
	//result from "$(wget -qSO- google.com 2>&1 | sed -n 's/^ *Date: *//p' | head -n 1)"
	datetime := "Mon, 22 Apr 2024 15:10:40 GMT"
	want := "202404221510.40"
	msg, err := ConvertDateTime(datetime)

	if msg != want || err != nil {
		t.Fatalf(`ConvertDateTime(%q) = %q, %v, want match for %#q, nil`, datetime, msg, err, want)
	}
}

func TestDateTimeEmpty(t *testing.T) {
	msg, err := ConvertDateTime("")
	if msg != "" || err == nil {
		t.Fatalf(`ConvertDateTime("") = %q, %v, want "", error`, msg, err)
	}
}
