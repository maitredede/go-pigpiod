package pigpiod

import "testing"

func TestErrorValidation(t *testing.T) {
	for code, val := range pigsErrors {
		if code != val.value {
			t.Logf("code %v mismatch val %+v", code, val)
			t.Fail()
		}
	}
}
