package hxx

import (
	"fmt"
	"regexp"
	"testing"
)

func makebs(size int) []byte {
	bs := make([]byte, size)
	for i, _ := range bs {
		bs[i] = byte(i % 256)
	}
	return bs
}

var DUMP_TestStringify = `00: 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f
10: 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f
20: 20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f
30: 30 31 32 33 34 35 36 37 38 39 3a 3b 3c 3d 3e 3f
40: 40 41 42 43 44 45 46 47 48 49 4a 4b 4c 4d 4e 4f
50: 50 51 52 53 54 55 56 57 58 59 5a 5b 5c
`

func TestStringify(t *testing.T) {
	s := Dump(makebs(93)).Stringify(false, false, 16)
	if s != DUMP_TestStringify {
		t.Errorf("unexpected value")
	}

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
	s0 := Dump(s).Hex(-1)
	if s0 != "48656c6c6f21" {
		t.Errorf("unexpected value: %q", s0)
	}
	s1 := Dump(s).Hex(0)
	if s1 != "" {
		t.Errorf("unexpected value: %q", s1)
	}
	s2 := Dump(s).Hex(1)
	if s2 != "4" {
		t.Errorf("unexpected value: %q", s2)
	}
	s3 := Dump(s).Hex(5)
	if s3 != "48656" {
		t.Errorf("unexpected value: %q", s3)
	}
	s4 := Dump(s).Hex(15)
	if s4 != "48656c6c6f21" {
		t.Errorf("unexpected value: %q", s0)
	}
	s5 := Dump(s).Hex(20)
	if s5 != "48656c6c6f21" {
		t.Errorf("unexpected value: %q", s0)
	}
}

func TestChars(t *testing.T) {
	s := "Eĥoŝanĝo Hello ĉiuĵaŭde"
	s0 := Dump(s).Chars(' ', false)
	if s0 != "E..o..an..o Hello ..iu..a..de" {
		t.Errorf("unexpected value: %q", s0)
	}
	// turn utf8 sequences into latin1 chars
	s1 := Dump(s).Chars(' ', true)
	if s1 != "EÄ¥oÅ.anÄ.o Hello Ä.iuÄµaÅ­de" {
		t.Errorf("unexpected value: %q", s1)
	}
	// turn utf8 sequences into latin1 chars
	s2 := Dump(s[:6]).Chars(' ', true)
	if s2 != "EÄ¥oÅ." {
		t.Errorf("unexpected value: %q", s2)
	}
	s3 := Dump(s).Chars(' ', false)
	if s3 != "E..o..an..o Hello ..iu..a..de" {
		t.Errorf("unexpected value: %q", s3)
	}
}

var DUMP_TestFormatHex = `000: 4c 6f 72 65 6d 20 49 70 73 75 6d 20 69 73 20 73  Lorem Ipsum is s
010: 69 6d 70 6c 79 20 64 75 6d 6d 79 20 74 65 78 74  imply dummy text
020: 20 6f 66 20 74 68 65 20 70 72 69 6e 74 69 6e 67   of the printing
030: 20 61 6e 64 20 74 79 70 65 73 65 74 74 69 6e 67   and typesetting
040: 20 69 6e 64 75 73 74 72 79 2e 20 4c 6f 72 65 6d   industry. Lorem
050: 20 49 70 73 75 6d 20 68 61 73 20 62 65 65 6e 20   Ipsum has been 
060: 74 68 65 20 69 6e 64 75 73 74 72 79 27 73 20 73  the industry's s
070: 74 61 6e 64 61 72 64 20 64 75 6d 6d 79 20 74 65  tandard dummy te
080: 78 74 20 65 76 65 72 20 73 69 6e 63 65 20 74 68  xt ever since th
090: 65 20 31 35 30 30 73 2c 20 77 68 65 6e 20 61 6e  e 1500s, when an
0a0: 20 75 6e 6b 6e 6f 77 6e 20 70 72 69 6e 74 65 72   unknown printer
0b0: 20 74 6f 6f 6b 20 61 20 67 61 6c 6c 65 79 20 6f   took a galley o
0c0: 66 20 74 79 70 65 20 61 6e 64 20 73 63 72 61 6d  f type and scram
0d0: 62 6c 65 64 20 69 74 20 74 6f 20 6d 61 6b 65 20  bled it to make 
0e0: 61 20 74 79 70 65 20 73 70 65 63 69 6d 65 6e 20  a type specimen 
0f0: 62 6f 6f 6b 2e 20 49 74 20 68 61 73 20 73 75 72  book. It has sur
100: 76 69 76 65 64 20 6e 6f 74 20 6f 6e 6c 79 20 66  vived not only f
110: 69 76 65 20 63 65 6e 74 75 72 69 65 73 2c 20 62  ive centuries, b
120: 75 74 20 61 6c 73 6f 20 74 68 65 20 6c 65 61 70  ut also the leap
130: 20 69 6e 74 6f 20 65 6c 65 63 74 72 6f 6e 69 63   into electronic
140: 20 74 79 70 65 73 65 74 74 69 6e 67 2c 20 72 65   typesetting, re
150: 6d 61 69 6e 69 6e 67 20 65 73 73 65 6e 74 69 61  maining essentia
160: 6c 6c 79 20 75 6e 63 68 61 6e 67 65 64 2e 20 49  lly unchanged. I
170: 74 20 77 61 73 20 70 6f 70 75 6c 61 72 69 73 65  t was popularise
180: 64 20 69 6e 20 74 68 65 20 31 39 36 30 73 20 77  d in the 1960s w
190: 69 74 68 20 74 68 65 20 72 65 6c 65 61 73 65 20  ith the release 
1a0: 6f 66 20 4c 65 74 72 61 73 65 74 20 73 68 65 65  of Letraset shee
1b0: 74 73 20 63 6f 6e 74 61 69 6e 69 6e 67 20 4c 6f  ts containing Lo
1c0: 72 65 6d 20 49 70 73 75 6d 20 70 61 73 73 61 67  rem Ipsum passag
1d0: 65 73 2c 20 61 6e 64 20 6d 6f 72 65 20 72 65 63  es, and more rec
1e0: 65 6e 74 6c 79 20 77 69 74 68 20 64 65 73 6b 74  ently with deskt
1f0: 6f 70 20 70 75 62 6c 69 73 68 69 6e 67 20 73 6f  op publishing so
200: 66 74 77 61 72 65 20 6c 69 6b 65 20 41 6c 64 75  ftware like Aldu
210: 73 20 50 61 67 65 4d 61 6b 65 72 20 69 6e 63 6c  s PageMaker incl
220: 75 64 69 6e 67 20 76 65 72 73 69 6f 6e 73 20 6f  uding versions o
230: 66 20 4c 6f 72 65 6d 20 49 70 73 75 6d 2e        f Lorem Ipsum.
`

func TestFormatHex(t *testing.T) {
	s := fmt.Sprintf("%x", NewDump(' '))
	if s != "00000020" {
		t.Errorf("unexpected value: %q", s)
	}
	s = fmt.Sprintf("%x", NewDump("hello world!"))
	if s != "68656c6c6f20776f726c6421" {
		t.Errorf("unexpected value: %q", s)
	}
	s = fmt.Sprintf("%v", NewDump("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."))
	if s != DUMP_TestFormatHex {
		t.Errorf("unexpected value: %q", s)
	}
	s = fmt.Sprintf("%v", NewDump("hello world!"))
	if s != "0: 68 65 6c 6c 6f 20 77 6f 72 6c 64 21              hello world!\n" {
		t.Errorf("unexpected value: %q", s)
	}
	s = fmt.Sprintf("%s", NewDump("hello world!"))
	if s != "hello world!" {
		t.Errorf("unexpected value: %q", s)
	}
	s = fmt.Sprintf("%x", NewDump(17389))
	if s != "00000000000043ed" {
		t.Errorf("unexpected value: %q", s)
	}
	s = fmt.Sprintf("%x", NewDump(uint16(1<<15)))
	if s != "8000" {
		t.Errorf("unexpected value: %q", s)
	}
}

func TestDumpStruct0(t *testing.T) {
	q := &struct {
		j rune
		f int
		g int
	}{
		'%',
		50,
		60,
	}
	s := fmt.Sprintf("%v", NewDump(q))
	if s != `00: 25 00 00 00 00 00 00 00 32 00 00 00 00 00 00 00  %.......2.......
10: 3c 00 00 00 00 00 00 00                          <.......
` {
		t.Errorf("unexpected value: %q", s)
	}
}

func TestDumpStruct1(t *testing.T) {
	q := &struct {
		j []byte
		f int32
		g [][]string
		k [4]int16
	}{
		[]byte("hello there!"),
		50,
		[][]string{{"one", "two", "three"}},
		[...]int16{3, 5, 7, 13},
	}
	s := fmt.Sprintf("%v", NewDump(q))
	if ok, _ := regexp.Match(`00: .. .. .. .. .. .. .. .. 0c 00 00 00 00 00 00 00  .<..............
        10: 0c 00 00 00 00 00 00 00 32 00 00 00 00 00 00 00  ........2.......
        20: .. .. .. .. .. .. .. .. 01 00 00 00 00 00 00 00   ...............
        30: 01 00 00 00 00 00 00 00 03 00 05 00 07 00 0d 00  ................`, []byte(`00: 10 3c 0d 00 c0 00 00 00 0c 00 00 00 00 00 00 00  .<..............
        10: 0c 00 00 00 00 00 00 00 32 00 00 00 00 00 00 00  ........2.......
        20: 20 a1 0c 00 c0 00 00 00 01 00 00 00 00 00 00 00   ...............
        30: 01 00 00 00 00 00 00 00 03 00 05 00 07 00 0d 00  ................`)); !ok {
		t.Errorf("unexpected value: %s", s)
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
		_ = ds.Hex(-1)
	}
}
