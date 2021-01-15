package main

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

const zeros = "0000000000000000000000000000000000000000"

type Amount string

type BaseParser struct {
	AmountDecimalPoint int
}

func main() {
	var a Amount = "0.3615841"
	fmt.Println("a:", a)

	p := &BaseParser{
		AmountDecimalPoint: 8,
	}
	b, err := p.AmountToBigInt(a)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("b:", b.String())

	c := AmountToDecimalString(&b, p.AmountDecimalPoint)
	fmt.Println("c:", c)
}

// AmountToBigInt converts amount in common.JSONNumber (string) to big.Int
// it uses string operations to avoid problems with rounding
func (p *BaseParser) AmountToBigInt(n Amount) (big.Int, error) {
	var r big.Int
	s := string(n)
	i := strings.IndexByte(s, '.')
	d := p.AmountDecimalPoint
	if d > len(zeros) {
		d = len(zeros)
	}
	if i == -1 {
		s = s + zeros[:d]
	} else {
		z := d - len(s) + i + 1
		if z > 0 {
			s = s[:i] + s[i+1:] + zeros[:z]
		} else {
			s = s[:i] + s[i+1:len(s)+z]
		}
	}
	if _, ok := r.SetString(s, 10); !ok {
		return r, errors.New("AmountToBigInt: failed to convert")
	}
	return r, nil
}

// AmountToDecimalString converts amount in big.Int to string with decimal point in the place defined by the parameter d
func AmountToDecimalString(a *big.Int, d int) string {
	if a == nil {
		return ""
	}
	n := a.String()
	var s string
	if n[0] == '-' {
		n = n[1:]
		s = "-"
	}
	if d > len(zeros) {
		d = len(zeros)
	}
	if len(n) <= d {
		n = zeros[:d-len(n)+1] + n
	}
	i := len(n) - d
	ad := strings.TrimRight(n[i:], "0")
	if len(ad) > 0 {
		n = n[:i] + "." + ad
	} else {
		n = n[:i]
	}
	return s + n
}
