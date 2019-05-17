package hxx

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"unicode"
	"unsafe"
)

var (
	smalls = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	bigs   = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
)

type Dump []byte

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func byte2rune(b byte) rune {
	if unicode.IsPrint(rune(b)) {
		return rune(b)
	}
	return '.'
}

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
		bs[0] = uint8(q)
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
		copy(bs, []byte(q))
	default:
		bs = []byte{}
		t = t.Elem()
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
		sh.Len = int(t.Size())
		sh.Cap = sh.Len
		sh.Data = v.Pointer()
	}
	return Dump(bs)
}

// digits count how many hex digits are required for x
func digits(x uint64) int {
	y := 0
	for x != 0 {
		y++
		x = x >> 4
	}
	return y
}

// hex return a hex representation for x
func hex(x uint64, fill byte, w int) string {
	if w == 0 {
		return ""
	}
	i := 0
	y := make([]byte, w)
	for x != 0 && i != w {
		y[w-1-i] = smalls[15&x]
		x = x >> 4
		i++
	}
	for i < w {
		y[w-1-i] = fill
		i++
	}
	if len(y) <= 1 {
		return string(y)
	}
	return string(y)
}

// Chars produce a string that represents the dump as a string of
// characters, one character per byte.  fill is used for characters
// that do not have a graphic representaiton.  latin1 determines if
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
			bs = append(bs, '.')
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
	bs := []byte{}
	// do this by number of output characters
	l := d.Digits()
	if w >= 0 {
		l = min(w, l)
	}
	i := 0
	// while there are still chars left
	for len(bs) < l {
		b := d[i]
		bs = append(bs, smalls[(b>>4)&15])
		bs = append(bs, smalls[b&15])
		i++
	}
	return string(bs[:l])
}

// Format support for "fmt" package Printf functions.
func (d Dump) Format(f fmt.State, c rune) {
	w, ok := f.Width()
	if !ok {
		w = d.Digits()
	}
	p, ok := f.Precision()
	if !ok {
		p = 16
	}
	fsharp := f.Flag('#')
	switch c {
	case 's':
		f.Write([]byte(d.Chars('.', fsharp)))
	case 'x':
		f.Write([]byte(d.Hex(w)))
	case 'v':
		f.Write([]byte(d.Stringify(true, false, p)))
	}
}

func (d Dump) Stringify(c, latin1 bool, e int) (s string) {
	l := len(d)
	i := 0
	x := digits(uint64(l))
	for i < l {
		m := i + e
		if m >= l {
			m = l
		}
		j := i
		// address
		s = s + hex(uint64(j), '0', x) + ":"
		// hex representation
		r := ""
		for j < m {
			r = r + " " + hex(uint64(d[j]), '0', 2)
			j++
		}
		if c {
			for len(r) < (3 * e) {
				r = r + "   "
			}
			s = s + r + "  " + d[i:m].Chars('.', false) + "\n"
		} else {
			s = s + r + "\n"
		}
		i += e
	}
	return s
}

// Digits returns the number of digits required
// for a hex representaiton of the dump
func (d Dump) Digits() int {
	return len(d) * 2
}
