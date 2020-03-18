package hxx

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/bits"
	"reflect"
	"unsafe"
)

var Fill = '.'

const (
	byte_bin_n = 8
	byte_oct_n = 3
	byte_dec_n = 3
	byte_hex_n = 2
	byte_bin_s = "       0       1      10      11     100     101     110     111    1000    1001    1010    1011    1100    1101    1110    1111   10000   10001   10010   10011   10100   10101   10110   10111   11000   11001   11010   11011   11100   11101   11110   11111  100000  100001  100010  100011  100100  100101  100110  100111  101000  101001  101010  101011  101100  101101  101110  101111  110000  110001  110010  110011  110100  110101  110110  110111  111000  111001  111010  111011  111100  111101  111110  111111 1000000 1000001 1000010 1000011 1000100 1000101 1000110 1000111 1001000 1001001 1001010 1001011 1001100 1001101 1001110 1001111 1010000 1010001 1010010 1010011 1010100 1010101 1010110 1010111 1011000 1011001 1011010 1011011 1011100 1011101 1011110 1011111 1100000 1100001 1100010 1100011 1100100 1100101 1100110 1100111 1101000 1101001 1101010 1101011 1101100 1101101 1101110 1101111 1110000 1110001 1110010 1110011 1110100 1110101 1110110 1110111 1111000 1111001 1111010 1111011 1111100 1111101 1111110 11111111000000010000001100000101000001110000100100001011000011010000111100010001000100110001010100010111000110010001101100011101000111110010000100100011001001010010011100101001001010110010110100101111001100010011001100110101001101110011100100111011001111010011111101000001010000110100010101000111010010010100101101001101010011110101000101010011010101010101011101011001010110110101110101011111011000010110001101100101011001110110100101101011011011010110111101110001011100110111010101110111011110010111101101111101011111111000000110000011100001011000011110001001100010111000110110001111100100011001001110010101100101111001100110011011100111011001111110100001101000111010010110100111101010011010101110101101101011111011000110110011101101011011011110111001101110111011110110111111110000011100001111000101110001111100100111001011110011011100111111010001110100111101010111010111110110011101101111011101110111111110000111100011111001011110011111101001111010111110110111101111111100011111001111110101111101111111100111111011111111011111111"
	byte_oct_s = "  0  1  2  3  4  5  6  7 10 11 12 13 14 15 16 17 20 21 22 23 24 25 26 27 30 31 32 33 34 35 36 37 40 41 42 43 44 45 46 47 50 51 52 53 54 55 56 57 60 61 62 63 64 65 66 67 70 71 72 73 74 75 76 77100101102103104105106107110111112113114115116117120121122123124125126127130131132133134135136137140141142143144145146147150151152153154155156157160161162163164165166167170171172173174175176177200201202203204205206207210211212213214215216217220221222223224225226227230231232233234235236237240241242243244245246247250251252253254255256257260261262263264265266267270271272273274275276277300301302303304305306307310311312313314315316317320321322323324325326327330331332333334335336337340341342343344345346347350351352353354355356357360361362363364365366367370371372373374375376377"
	byte_dec_s = "  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99100101102103104105106107108109110111112113114115116117118119120121122123124125126127128129130131132133134135136137138139140141142143144145146147148149150151152153154155156157158159160161162163164165166167168169170171172173174175176177178179180181182183184185186187188189190191192193194195196197198199200201202203204205206207208209210211212213214215216217218219220221222223224225226227228229230231232233234235236237238239240241242243244245246247248249250251252253254255"
	byte_hex_s = " 0 1 2 3 4 5 6 7 8 9 a b c d e f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f808182838485868788898a8b8c8d8e8f909192939495969798999a9b9c9d9e9fa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c9cacbcccdcecfd0d1d2d3d4d5d6d7d8d9dadbdcdddedfe0e1e2e3e4e5e6e7e8e9eaebecedeeeff0f1f2f3f4f5f6f7f8f9fafbfcfdfeff"
)

var (
	smalls = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
)

// Dump at memory layout representation of an object
type Dump []byte

// NewDump return a dump object.  o may be any primitive type or pointer to any other object type.
func NewDump(o interface{}) Dump {
	v := reflect.ValueOf(o)
	t := v.Type()
	l := t.Size()
	bs := []byte(nil)
	switch q := o.(type) {
	case int:
		bs = make([]byte, l)
		switch l {
		case 4:
			binary.BigEndian.PutUint32(bs, uint32(q))
		case 8:
			binary.BigEndian.PutUint64(bs, uint64(q))
		}
	case uint:
		bs = make([]byte, l)
		switch l {
		case 4:
			binary.BigEndian.PutUint32(bs, uint32(q))
		case 8:
			binary.BigEndian.PutUint64(bs, uint64(q))
		}
	case uint8:
		bs = make([]byte, l)
		bs[0] = q
	case int8:
		bs = make([]byte, l)
		bs[0] = uint8(q)
	case uint16:
		bs = make([]byte, l)
		binary.BigEndian.PutUint16(bs, q)
	case int16:
		bs = make([]byte, l)
		binary.BigEndian.PutUint16(bs, uint16(q))
	case uint32:
		bs = make([]byte, l)
		binary.BigEndian.PutUint32(bs, q)
	case int32:
		bs = make([]byte, l)
		binary.BigEndian.PutUint32(bs, uint32(q))
	case uint64:
		bs = make([]byte, l)
		binary.BigEndian.PutUint64(bs, q)
	case int64:
		bs = make([]byte, l)
		binary.BigEndian.PutUint64(bs, uint64(q))
	case float32:
		bs = make([]byte, l)
		binary.BigEndian.PutUint32(bs, math.Float32bits(q))
	case float64:
		bs = make([]byte, l)
		binary.BigEndian.PutUint64(bs, math.Float64bits(q))
	case string:
		bs = make([]byte, len(q))
		copy(bs, q)
	default:
		bs0 := []byte{}
		t = t.Elem()
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&bs0))
		sh.Len = int(t.Size())
		sh.Cap = sh.Len
		sh.Data = v.Pointer()
		bs = make([]byte, len(bs0))
		copy(bs, bs0)
	}
	return bs
}

// digits count how many hex digits are required for x
func digits(x uint64, base uint8) uint8 {
	if x == 0 {
		return 1
	}
	q := uint8(0)
	if bits.OnesCount8(base) == 1 {
		k := 7 - bits.LeadingZeros8(base)
		for x > 0 {
			x >>= k
			q++
		}
		return q
	}
	for x > 0 {
		x /= uint64(base)
		q++
	}
	return q
}

// itoa return x's representation as a number in base
func itoa(x uint64, fill byte, wid, base uint8) (string, uint64) {
	if wid == 0 {
		return "", 0
	}
	i := uint8(0)
	y := make([]byte, wid)
	if x == 0 {
		i = 1
		y[wid-1] = '0'
	} else if bits.OnesCount8(base) == 1 {
		k := 7 - bits.LeadingZeros8(base)
		msk := uint64(1<<k) - 1
		for x > 0 && i != wid {
			j := x & msk
			y[wid-1-i] = smalls[j]
			x >>= k
			i++
		}
	} else {
		for x > 0 && i != wid {
			j := x % uint64(base)
			y[wid-1-i] = smalls[j]
			x /= uint64(base)
			i++
		}
	}
	for i < wid {
		y[wid-1-i] = fill
		i++
	}
	return string(y), x
}

// itoaByte return the string representation of a single byte for a given base
func itoaByte(x, base uint8) string {
	if x == 0 {
		return "0"
	}
	var n uint
	var k string
	switch base {
	case 2:
		n = byte_bin_n
		k = byte_bin_s
	case 8:
		n = byte_oct_n
		k = byte_oct_s
	case 10:
		n = byte_dec_n
		k = byte_dec_s
	case 16:
		n = byte_hex_n
		k = byte_hex_s
	default:
		if x < base {
			k := x % base
			return string(smalls[k])
		}
		d := digits(uint64(x), base)
		s := make([]byte, d)
		for x > 0 {
			d--
			k := x % base
			s[d] = smalls[k]
			x /= base
		}
		return string(s)
	}
	d := uint(digits(uint64(x), base))
	j := (uint(x) + 1) * n
	return k[j-d : j]
}

// Chars produce a string that represents the dump as a string of
// characters, one character per byte.  fill is used for characters
// that do not have a graphic representation.  latin1 determines if
// high characters should be used to represent bytes.
func (d Dump) Chars(fill rune, latin1 bool) string {
	l := len(d)
	if l == 0 {
		return ""
	}
	bs := []rune{}
	// do this by number of output characters
	i := 0
	for len(bs) < l {
		if d[i] >= ' ' && d[i] <= '~' {
			bs = append(bs, rune(d[i]))
		} else if latin1 && d[i] >= '¡' && d[i] <= 'ÿ' {
			bs = append(bs, rune(d[i]))
		} else {
			bs = append(bs, fill)
		}
		i++
	}
	return string(bs)
}

// Hex print w hex digits from d and fill with zeros to w
func (d Dump) Hex(w int) string {
	if w == 0 {
		return ""
	}
	ql := byte_hex_n * len(d)
	if w > 0 && w < ql {
		ql = w
	}
	bs := make([]byte, ql)
	// while there are still chars left
	i := 0
	for i*2 < ql {
		s, _ := itoa(uint64(d[i]), '0', 2, 16)
		copy(bs[i*2:], s)
		i++
	}
	return string(bs[:ql])
}

// Format support for "fmt" package Printf functions.
//
// supported formatter verbs:
// ```
// s  output bytes as ASCII characters, one character per byte.  ASCII control characters or characters outside the 7 bit
//    range are displayed as '.'.
// x  print dump as compact string of hex digits, width supported.
// b  represent bytes in the dump with eight binary digits.
// o  represent bytes in the dump with three octal digits.
// d  represent bytes in the dump with three decimal digits.
// x  represent bytes in the dump with two hex digits.
// ```
//
// supported formatter flags:
// ```
// #  include latin1 characters in the output
// -  remove ascii display from hexdump
// 0  zero-fill bytes
// ```
//
// The format precision for `%b`, `%o`, `%d` and `%x` specifies the number of bytes output per line.
//
func (d Dump) Format(f fmt.State, c rune) {
	w, ok := f.Width()
	if !ok {
		w = -1
	}
	p, ok := f.Precision()
	if !ok {
		p = 16
	}
	fsharp := f.Flag('#')
	fminus := f.Flag('-')
	fzero := f.Flag('0')
	switch c {
	case 's':
		f.Write([]byte(d.Chars('.', fsharp)))
	case 'x':
		f.Write([]byte(d.Hex(w)))
	case 'b':
		f.Write([]byte(d.Stringify(!fminus, false, fsharp, fzero, p, 1, 2)))
	case 'o':
		f.Write([]byte(d.Stringify(!fminus, false, fsharp, fzero, p, 1, 8)))
	case 'd':
		f.Write([]byte(d.Stringify(!fminus, false, fsharp, fzero, p, 1, 10)))
	case 'v':
		f.Write([]byte(d.Stringify(!fminus, false, fsharp, fzero, p, 1, 16)))
	}
}

func bstoi(bs []byte, littledian bool) uint64 {
	x := uint64(0)
	if littledian {
		for i := 0; i < len(bs); i++ {
			x *= 256
			x += uint64(bs[i])
		}
		return x
	}
	for i := len(bs) - 1; i > -1; i-- {
		x *= 256
		x += uint64(bs[i])
	}
	return x
}

// Stringify produce a hex-dump string from a dump object, d.  Each line will being with an address
// followed by each byte.  if c is true, the ascii representation of the preceding bytes will follow.
// latin1 will allow latin1 characters in the output string.  littledian will output grouped bytes in
// little-endian order.  bcount controls the number of bytes output on each line of the hex-dump.
// if zerofill is true, each byte will be filled in with leading zeros.
func (d Dump) Stringify(c, littlendian, latin1, zerofill bool, bcount, group, base int) (s string) {
	bse := uint8(base)
	l := len(d)
	i := 0
	digitfill := byte(' ')
	if zerofill {
		digitfill = '0'
	}
	for i < l {
		if i > 0 {
			s = s + "\n"
		}
		m := i + bcount
		if m >= l {
			m = l
		}
		j := i
		// address
		ss, _ := itoa(uint64(j), digitfill, digits(uint64(l), bse), bse)
		s = s + ss + ":"
		// bytes representation
		r := ""
		for j < m {
			k := group
			// don't process more groups than there is data
			if j+k >= m {
				k = m - j
			}
			width := digits((1<<(k*8))-1, bse)
			s, _ := itoa(bstoi(d[j:j+k], littlendian), digitfill, width, bse)
			r = r + " " + s
			j += group
		}
		if c {
			for len(r) < (3 * bcount) {
				r = r + "   "
			}
			s = s + r + "  " + d[i:m].Chars('.', latin1)
		} else {
			s = s + r
		}
		i += bcount
	}
	return s
}
