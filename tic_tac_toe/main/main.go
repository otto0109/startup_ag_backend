package main

import (
	"animals/tic_tac_toe"
)

func main() {
	field := tic_tac_toe.TicTacToeField{}
	field.Row0.Field0.SetValueX()
	field.Row1.Field1.SetValueX()
	field.Row2.Field2.SetValueX()
	field.PrintTicTacToe()
}
