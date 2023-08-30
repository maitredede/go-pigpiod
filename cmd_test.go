package pigpiod

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReadBoth(t *testing.T) {
	buff := &bytes.Buffer{}
	if err := binary.Write(buff, endianess, int32(-1)); err != nil {
		t.Fatal(err)
	}

	i, u, err := readBoth(buff)
	if err != nil {
		t.Fatal(err)
	}
	if i != -1 {
		t.Logf("int32 wrong value: %v, expected -1", i)
		t.Fail()
	}
	if u != 0xFFFFFFFF {
		t.Logf("uint32 wrong value: %04x, expected 0xFFFFFFFF", u)
		t.Fail()
	}
}
