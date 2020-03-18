package hxx

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func makebs(size int) []byte {
	bs := make([]byte, size)
	for i, _ := range bs {
		bs[i] = byte(i % 256)
	}
	return bs
}

var DUMP_TestStringify_Hex = `00: 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f
10: 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f
20: 20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f
30: 30 31 32 33 34 35 36 37 38 39 3a 3b 3c 3d 3e 3f
40: 40 41 42 43 44 45 46 47 48 49 4a 4b 4c 4d 4e 4f
50: 50 51 52 53 54 55 56 57 58 59 5a 5b 5c`

var DUMP_TestStringify_Hex2 = `00: 0001 0203 0405 0607 0809 0a0b 0c0d 0e0f
10: 1011 1213 1415 1617 1819 1a1b 1c1d 1e1f
20: 2021 2223 2425 2627 2829 2a2b 2c2d 2e2f
30: 3031 3233 3435 3637 3839 3a3b 3c3d 3e3f
40: 4041 4243 4445 4647 4849 4a4b 4c4d 4e4f
50: 5051 5253 5455 5657 5859 5a5b 5c`

var DUMP_TestStringify_Hex3 = `00: 0100 0302 0504 0706 0908 0b0a 0d0c 0f0e
10: 1110 1312 1514 1716 1918 1b1a 1d1c 1f1e
20: 2120 2322 2524 2726 2928 2b2a 2d2c 2f2e
30: 3130 3332 3534 3736 3938 3b3a 3d3c 3f3e
40: 4140 4342 4544 4746 4948 4b4a 4d4c 4f4e
50: 5150 5352 5554 5756 5958 5b5a 5c`

var DUMP_TestStringify_Bin = `0000000: 00000000 00000001 00000010 00000011 00000100 00000101 00000110 00000111 00001000 00001001 00001010 00001011 00001100 00001101 00001110 00001111
0010000: 00010000 00010001 00010010 00010011 00010100 00010101 00010110 00010111 00011000 00011001 00011010 00011011 00011100 00011101 00011110 00011111
0100000: 00100000 00100001 00100010 00100011 00100100 00100101 00100110 00100111 00101000 00101001 00101010 00101011 00101100 00101101 00101110 00101111
0110000: 00110000 00110001 00110010 00110011 00110100 00110101 00110110 00110111 00111000 00111001 00111010 00111011 00111100 00111101 00111110 00111111
1000000: 01000000 01000001 01000010 01000011 01000100 01000101 01000110 01000111 01001000 01001001 01001010 01001011 01001100 01001101 01001110 01001111
1010000: 01010000 01010001 01010010 01010011 01010100 01010101 01010110 01010111 01011000 01011001 01011010 01011011 01011100`

var DUMP_TestStringify_Oct = `000: 000 001 002 003 004 005 006 007 010 011 012 013 014 015 016 017
020: 020 021 022 023 024 025 026 027 030 031 032 033 034 035 036 037
040: 040 041 042 043 044 045 046 047 050 051 052 053 054 055 056 057
060: 060 061 062 063 064 065 066 067 070 071 072 073 074 075 076 077
100: 100 101 102 103 104 105 106 107 110 111 112 113 114 115 116 117
120: 120 121 122 123 124 125 126 127 130 131 132 133 134`

var DUMP_TestStringify_Dec = `00: 000 001 002 003 004 005 006 007 008 009 010 011 012 013 014 015
16: 016 017 018 019 020 021 022 023 024 025 026 027 028 029 030 031
32: 032 033 034 035 036 037 038 039 040 041 042 043 044 045 046 047
48: 048 049 050 051 052 053 054 055 056 057 058 059 060 061 062 063
64: 064 065 066 067 068 069 070 071 072 073 074 075 076 077 078 079
80: 080 081 082 083 084 085 086 087 088 089 090 091 092`

func TestStringify(t *testing.T) {
	tests := []struct {
		c, littlendian, latin1, zerofill bool
		bcount, group, base              int
		expect                           string
	}{
		{false, true, true, true, 16, 1, 16, DUMP_TestStringify_Hex},
		{false, true, true, true, 16, 2, 16, DUMP_TestStringify_Hex2},
		{false, false, true, true, 16, 2, 16, DUMP_TestStringify_Hex3},
		{false, true, true, true, 16, 1, 2, DUMP_TestStringify_Bin},
		{false, true, true, true, 16, 1, 8, DUMP_TestStringify_Oct},
		{false, true, true, true, 16, 1, 10, DUMP_TestStringify_Dec},
	}
	for i, tst := range tests {
		s := Dump(makebs(93)).Stringify(tst.c, tst.littlendian, tst.latin1, tst.zerofill, tst.bcount, tst.group, tst.base)
		if s != tst.expect {
			t.Errorf("%d: unexpected value: %s", i, s)
		}
	}

}

func TestItoa(t *testing.T) {
	tests := []struct {
		in     uint64
		fill   byte
		digits uint8
		base   uint8
		out    string
	}{
		{
			100,
			'0',
			digits(100, 16),
			16,
			"64",
		},
		{
			16,
			'0',
			digits(16, 16),
			16,
			"10",
		},
		{
			13635,
			'0',
			digits(13635, 16),
			16,
			"3543",
		},
		{
			3930,
			'0',
			digits(3930, 16) + 1,
			16,
			"0f5a",
		},
	}
	for i, tst := range tests {
		s, _ := itoa(tst.in, tst.fill, tst.digits, tst.base)
		if s != tst.out {
			t.Errorf("%d: unexpected value: %q", i, s)
		}
	}
}

func TestItoaByte(t *testing.T) {
	tests := []struct {
		in   uint8
		base uint8
		out  string
	}{
		{
			100,
			16,
			"64",
		},
		{
			16,
			16,
			"10",
		},
		{
			1,
			16,
			"1",
		},
		{
			255,
			16,
			"ff",
		},
		{
			255,
			10,
			"255",
		},
		{
			255,
			8,
			"377",
		},
		{
			255,
			2,
			"11111111",
		},
		{
			255,
			3,
			"100110",
		},
	}
	for i, tst := range tests {
		s := itoaByte(tst.in, tst.base)
		if s != tst.out {
			t.Errorf("%d: unexpected value: %q", i, s)
		}
	}
}

func TestDigits(t *testing.T) {
	tests := []struct {
		in     uint64
		base   uint8
		expect uint8
	}{
		{
			0,
			3,
			1,
		},
		{
			1,
			3,
			1,
		},
		{
			2,
			3,
			1,
		},
		{
			0,
			16,
			1,
		},
		{
			1,
			16,
			1,
		},
		{
			15,
			16,
			1,
		},
		{
			100,
			16,
			2,
		},
		{
			100,
			16,
			2,
		},
		{
			100,
			2,
			7,
		},
		{
			16,
			16,
			2,
		},
		{
			3,
			2,
			2,
		},
		{
			10,
			10,
			2,
		},
		{
			999999999999999,
			10,
			15,
		},
		{
			13635,
			16,
			4,
		},
		{
			3930,
			16,
			3,
		},
	}
	for i, tst := range tests {
		x := digits(tst.in, tst.base)
		if x != tst.expect {
			t.Errorf("%d: unexpected value: %d", i, x)
		}
	}
}

func TestHex(t *testing.T) {
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
		t.Errorf("unexpected value: %q", s4)
	}
	s5 := Dump(s).Hex(20)
	if s5 != "48656c6c6f21" {
		t.Errorf("unexpected value: %q", s5)
	}
}

func TestChars(t *testing.T) {
	s := "Eĥoŝanĝo Hello ĉiuĵaŭde"
	s0 := Dump(s).Chars('.', false)
	if s0 != "E..o..an..o Hello ..iu..a..de" {
		t.Errorf("unexpected value: %q", s0)
	}
	// turn utf8 sequences into latin1 chars
	s1 := Dump(s).Chars('.', true)
	if s1 != "EÄ¥oÅ.anÄ.o Hello Ä.iuÄµaÅ­de" {
		t.Errorf("unexpected value: %q", s1)
	}
	// turn utf8 sequences into latin1 chars
	s2 := Dump(s[:6]).Chars('.', true)
	if s2 != "EÄ¥oÅ." {
		t.Errorf("unexpected value: %q", s2)
	}
	s3 := Dump(s).Chars('.', false)
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
230: 66 20 4c 6f 72 65 6d 20 49 70 73 75 6d 2e        f Lorem Ipsum.`

func TestFormatHex(t *testing.T) {
	tests := []struct {
		input    interface{}
		fmt      string
		expected string
	}{
		{
			input:    ' ',
			fmt:      "%x",
			expected: "00000020",
		},
		{
			input:    "hello world!",
			fmt:      "%x",
			expected: "68656c6c6f20776f726c6421",
		},
		{
			input:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			fmt:      "%0v",
			expected: DUMP_TestFormatHex,
		},
		{
			input:    "hello world!",
			fmt:      "%0v",
			expected: "0: 68 65 6c 6c 6f 20 77 6f 72 6c 64 21              hello world!",
		},
		{
			input:    "hello\nworld!",
			fmt:      "%v",
			expected: "0: 68 65 6c 6c 6f  a 77 6f 72 6c 64 21              hello.world!",
		},
		{
			input:    "hello world!",
			fmt:      "%s",
			expected: "hello world!",
		},
		{
			input:    17389,
			fmt:      "%x",
			expected: "00000000000043ed",
		},
		{
			input:    uint16(1 << 15),
			fmt:      "%x",
			expected: "8000",
		},
	}
	for i, tst := range tests {
		s := fmt.Sprintf(tst.fmt, NewDump(tst.input))
		if s != tst.expected {
			t.Errorf("%d: unexpected value: %q", i, s)
		}
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
	s := fmt.Sprintf("%0v", NewDump(q))
	if s != `00: 25 00 00 00 00 00 00 00 32 00 00 00 00 00 00 00  %.......2.......
10: 3c 00 00 00 00 00 00 00                          <.......` {
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
	s := fmt.Sprintf("%0v", NewDump(q))
	if ok, _ := regexp.Match(`00: .. .. .. .. .. .. .. .. 0c 00 00 00 00 00 00 00  ................
10: 0c 00 00 00 00 00 00 00 32 00 00 00 00 00 00 00  ........2.......
20: .. .. .. .. .. .. .. .. 01 00 00 00 00 00 00 00  ................
30: 01 00 00 00 00 00 00 00 03 00 05 00 07 00 0d 00  ................`, []byte(s)); !ok {
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

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range []uint8{3, 4, 11} {
			_, _ = itoa(uint64(i), '0', 8, k)
		}
	}
}

func BenchmarkItoaByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range []uint8{3, 4, 11} {
			_ = itoaByte(uint8(i%256), k)
		}
	}
}

func BenchmarkSysItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, k := range []int{3, 4, 11} {
			_ = strconv.FormatInt(int64(i), k)
		}
	}
}

func BenchmarkPrintHex(b *testing.B) {
	ds := Dump(makebs(93))
	for i := uint64(0); i < uint64(b.N); i++ {
		_ = ds.Hex(-1)
		_ = ds.Hex(93 * 2)
	}
}
