package main

import (
	"fmt"
	"strconv"
	"strings"
)

type cipher uint8

const (
	Roman cipher = iota
	Arabic
)

var (
	Operations     = []string{"+", "-", "*", "/"}
	ArabicByRomans = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	RomansByArabic = map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
		19: "IXX",
		20: "XX",
	}
)

type Operand struct {
	c cipher
	V int
}

func NewOperand(op string) (*Operand, error) {
	if v, ok := ArabicByRomans[op]; ok {
		return &Operand{Roman, v}, nil
	}

	trimmed := strings.Trim(op, " ")
	v, err := strconv.Atoi(trimmed)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %s", trimmed)
	}

	return &Operand{Arabic, v}, nil
}

type Expression struct {
	Operation string
	Op1       *Operand
	Op2       *Operand
}

func NewExpression(exprStr string) *Expression {
	for _, op := range Operations {
		split := strings.Split(exprStr, op)
		if len(split) > 2 {
			panic("format of input is not for this task")
		}

		if len(split) == 2 {
			a, err := NewOperand(split[0])
			if err != nil {
				panic(err)
			}

			b, err := NewOperand(split[1])
			if err != nil {
				panic(err)
			}

			if a.c != b.c {
				panic("mixed Arabic and Roman expression")
			}

			if a.V > 10 || b.V > 10 {
				panic("one of the operands is bigger than 10")
			}

			return &Expression{
				Operation: op,
				Op1:       a,
				Op2:       b,
			}
		}
	}

	panic("invalid expression " + exprStr)
}

var (
	i2rC = []string{"", "C", "CC", "CCC"}
	i2rX = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	i2rI = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func ArabicToRoman(i int) string {
	return i2rC[i/100] + i2rX[(i%100)/10] + i2rI[i%10]
}

func (expr *Expression) Result() string {
	out := -999
	switch expr.Operation {
	case "+":
		out = expr.Op1.V + expr.Op2.V
	case "-":
		out = expr.Op1.V - expr.Op2.V
		if out < 0 && expr.Op1.c == Roman {
			panic("negative result when using Roman numbers")
		}
	case "*":
		out = expr.Op1.V * expr.Op2.V
	case "/":
		out = expr.Op1.V / expr.Op2.V
	default:
		panic("invalid expression")
	}

	if expr.Op1.c == Roman && expr.Op2.c == Roman {
		return ArabicToRoman(out)
	}

	return strconv.Itoa(out)
}
