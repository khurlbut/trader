package quote_service

import (
	"testing";
)

func TestSayHello(t *testing.T) {
	actual := ""
	expected := "Hello, world."
    if expected != actual {
		t.Errorf("Error occured while testing sayhello: '%s' != '%s'", expected, actual);
    }
}
