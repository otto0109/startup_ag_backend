package tic_tac_toe

import "fmt"

type TicTacToeField struct {
	Row0 Row
	Row1 Row
	Row2 Row
}

func (t *TicTacToeField) PrintTicTacToe() {
	t.Row0.PrintRow()
	fmt.Print("\n- - -\n")
	t.Row1.PrintRow()
	fmt.Print("\n- - -\n")
	t.Row2.PrintRow()
}

type Row struct {
	Field0 Field
	Field1 Field
	Field2 Field
}

func (r *Row) PrintRow() {
	r.Field0.PrintField()
	fmt.Print("|")
	r.Field1.PrintField()
	fmt.Print("|")
	r.Field2.PrintField()
}

type Field struct {
	value string
}

func (f *Field) PrintField() {
	if f.value == "" {
		fmt.Print(" ")
	} else {
		fmt.Print(f.value)
	}
}

func (f *Field) SetValueX() {
	f.value = "X"
}

func (f *Field) SetValueO() {
	f.value = "O"
}
