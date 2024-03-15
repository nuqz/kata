package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExpression(t *testing.T) {
	t.Run("It should work", func(t *testing.T) {
		cases := map[string]struct {
			Expected Expression
		}{
			"2+3": {
				Expected: Expression{
					Operation: "+",
					Op1:       &Operand{Arabic, 2},
					Op2:       &Operand{Arabic, 3},
				},
			},
			"X-VII": {
				Expected: Expression{
					Operation: "-",
					Op1:       &Operand{Roman, 10},
					Op2:       &Operand{Roman, 7},
				},
			},
		}

		for input, tc := range cases {
			t.Run(input, func(t *testing.T) {
				a := assert.New(t)
				out := NewExpression(input)
				a.Equal(tc.Expected.Operation, out.Operation)
				a.Equal(tc.Expected.Op1.c, out.Op1.c)
				a.Equal(tc.Expected.Op1.V, out.Op1.V)
				a.Equal(tc.Expected.Op2.c, out.Op2.c)
				a.Equal(tc.Expected.Op2.V, out.Op2.V)
			})
		}
	})
	t.Run("Must panic", func(t *testing.T) {
		cases := []string{
			"11+10",
			"I + 9",
			"7 + V",
			"1",
			"1 + 2 + 3",
			`1 +
			2`,
		}

		for _, input := range cases {
			t.Run(input, func(t *testing.T) {
				require.Panics(t, func() {
					expr := NewExpression(input)
					expr.Result()
				})
			})
		}
	})
}

func TestExpressionResult(t *testing.T) {
	cases := map[string]string{
		"1+1":     "2",
		"1+3":     "4",
		"2+4":     "6",
		"2+5":     "7",
		"2+6":     "8",
		"3+7":     "10",
		"4+8":     "12",
		"4+9":     "13",
		"4+10":    "14",
		"8-9":     "-1",
		"I+I":     "II",
		"I+II":    "III",
		"I+III":   "IV",
		"II+III":  "V",
		"III+III": "VI",
		"X+VI":    "XVI",

		"1 + 2":  "3",
		"VI/III": "II",
		"X*X":    "C",
		"V*X":    "L",
	}

	for input, expected := range cases {
		t.Run(fmt.Sprintf("%s=%s", input, expected),
			func(t *testing.T) {
				expr := NewExpression(input)
				assert.Equal(t, expected, expr.Result())
			})
	}
}
