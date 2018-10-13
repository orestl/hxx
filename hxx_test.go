package hxx

import (
	"fmt"
	"testing"
)

func makebs(size int) []byte {
	bs := make([]byte, size)
	for i, _ := range bs {
		bs[i] = byte(i % 256)
	}
	return bs
}

func TestStringify(t *testing.T) {
	t.Logf("\n%s", Dump(makebs(93)).Stringify(16))
}

func TestNewDumpInterface(t *testing.T) {
	s := "Hi There!"
	t.Logf("%s", NewDump(&s).Stringify(16))
}

func TestHex(t *testing.T) {
	s0 := hex(100, '0', digits(100))
	if s0 != "64" {
		t.Errorf("unexpected value: %q", s0)
	}
	s1 := hex(16, '0', digits(16))
	if s1 != "10" {
		t.Errorf("unexpected value: %q", s1)
	}
	s2 := hex(13635, '0', digits(13635))
	if s2 != "3543" {
		t.Errorf("unexpected value: %q", s2)
	}
	s3 := hex(3930, '0', digits(3930)+1)
	if s3 != "0f5a" {
		t.Errorf("unexpected value: %q", s3)
	}
}

func TestPrintHex(t *testing.T) {
	s := "Hello!"
	s0 := Dump(s).PrintHex(-1)
	if s0 != "48656c6c6f21" {
		t.Errorf("unexpected value: %q", s0)
	}
	s1 := Dump(s).PrintHex(0)
	if s1 != "" {
		t.Errorf("unexpected value: %q", s1)
	}
	s2 := Dump(s).PrintHex(1)
	if s2 != "4" {
		t.Errorf("unexpected value: %q", s2)
	}
	s3 := Dump(s).PrintHex(5)
	if s3 != "48656" {
		t.Errorf("unexpected value: %q", s3)
	}
	s4 := Dump(s).PrintHex(15)
	if s4 != "48656c6c6f21" {
		t.Errorf("unexpected value: %q", s0)
	}
}

func TestPrintChars(t *testing.T) {
	s := "Eĥoŝanĝo Hello ĉiuĵaŭde"
	s0 := Dump(s).PrintChars(' ', false)
	if s0 != "E..o..an..o Hello ..iu..a..de" {
		t.Errorf("unexpected value: %q", s0)
	}
	s1 := Dump(s).PrintChars(' ', true)
	if s1 != "EÄ¥oÅ.anÄ.o Hello Ä.iuÄµaÅ­de" {
		t.Errorf("unexpected value: %q", s1)
	}
	s2 := Dump(s[:6]).PrintChars(' ', true)
	if s2 != "EÄ¥oÅ." {
		t.Errorf("unexpected value: %q", s2)
	}
	s3 := Dump(s).PrintChars(' ', false)
	if s3 != "E..o..an..o Hello ..iu..a..de" {
		t.Errorf("unexpected value: %q", s3)
	}
}

func BenchmarkFormatterX(b *testing.B) {
	d := Dump(makebs(64))
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%x", d)
		_ = fmt.Sprintf("%5x", d)
	}
}

func BenchmarkBaselineFormatterX(b *testing.B) {
	d := makebs(64)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%x", d)
		_ = fmt.Sprintf("%5x", d)
	}
}

func BenchmarkHex(b *testing.B) {
	for i := uint64(0); i < uint64(b.N); i++ {
		_ = hex(i, '0', 8)
	}
}

func BenchmarkPrintHex(b *testing.B) {
	ds := Dump(makebs(93))
	for i := uint64(0); i < uint64(b.N); i++ {
		_ = ds.PrintHex(-1)
	}
}
