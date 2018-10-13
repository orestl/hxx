package hxx

import (
	"fmt"
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
	if v.Kind() != reflect.Ptr {
		panic(fmt.Errorf("cannot dump non pointer"))
	}
	t := v.Type()
	t = t.Elem()
	bs := []byte{}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sh.Len = int(t.Size())
	sh.Cap = sh.Len
	sh.Data = v.Pointer()
	return Dump(bs)
}

func digits(x uint64) int {
	y := 0
	for x != 0 {
		y++
		x = x >> 4
	}
	return y
}

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

func (d Dump) PrintChars(fill rune, latin1 bool) string {
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
	for k := i; k < l; k++ {
		bs = append(bs, fill)
	}
	return string(bs)
}

// PrintHex print w hex digits from d.
func (d Dump) PrintHex(w int) string {
	if w == 0 {
		return ""
	}
	bs := []byte{}
	// do this by number of output characters
	l := len(d) * 2
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

func (d Dump) Format(f fmt.State, c rune) {
	switch c {
	case 's':
		bs := []byte{}
		f.Write(bs)
	}
}

func (d Dump) Stringify(e int) (s string) {
	l := len(d)
	i := 0
	x := digits(uint64(l))
	for i < l {
		j := i
		s = s + hex(uint64(i), '0', x) + ":"
		for j < i+e && j < l {
			s = s + " " + hex(uint64(d[j]), '0', 2)
			j++
		}
		for j < i+e {
			s = s + " " + "  "
			j++
		}
		s = s + "\n"
		i = j
	}
	return s
}
