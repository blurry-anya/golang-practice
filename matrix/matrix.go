package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	matrix := Matrix{}

	stringRows := strings.Split(s, "\n")
	removeSpacesFromItems(stringRows)

	for _, v := range stringRows {
		if v == "" {
			return nil, errors.New("empty row")
		}
		stringItems := strings.Split(v, " ")

		row := []int{}

		for _, item := range stringItems {
			intItem, err := strconv.Atoi(item)
			if err != nil {
				return nil, errors.New("not a valid item")
			}
			row = append(row, intItem)
		}

		matrix = append(matrix, row)
	}

	if !matrix.isEvenRows() {
		return nil, errors.New("uneven rows")
	}

	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	rowLength := len(m)
	colLength := len(m[0])

	cols := make([][]int, colLength)

	for i := 0; i < colLength; i++ {
		cols[i] = make([]int, rowLength)
		for j := 0; j < rowLength; j++ {
			cols[i][j] = m[j][i]
		}
	}

	return cols
}

func (m Matrix) Rows() [][]int {
	rows := Matrix{}

	for i := 0; i < len(m); i++ {
		row := []int{}
		for j := 0; j < len(m[i]); j++ {
			row = append(row, m[i][j])
		}
		rows = append(rows, row)
	}

	return rows
}

func (m Matrix) Set(row, col, val int) bool {

	rowLength := len(m)
	colLength := len(m[0])

	if !(row < rowLength && col < colLength && row >= 0 && col >= 0) {
		return false
	}

	m[row][col] = val

	return true
}

func (m Matrix) isEvenRows() bool {
	for i := 0; i < len(m); i += 2 {
		if i == len(m)-1 {
			break
		}
		if len(m[i]) != len(m[i+1]) {
			return false
		}
	}
	return true
}

func removeSpacesFromItems(array []string) {
	for i := range array {
		array[i] = strings.TrimSpace(array[i])
	}
}
