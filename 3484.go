package main

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	dict map[string]int
}

func SpreadConstructor(rows int) Spreadsheet {
	return Spreadsheet{
		dict: make(map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.dict[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.dict[cell] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	vals := strings.Split(formula[1:], "+")
	v1, v2 := vals[0], vals[1]
	return this.getValue(v1) + this.getValue(v2)
}
func (this *Spreadsheet) getValue(val string) int {
	if val[0] >= 'A' && val[0] <= 'Z' {
		return this.dict[val]
	}
	v, _ := strconv.Atoi(val)
	return v
}
