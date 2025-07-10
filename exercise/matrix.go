package exercise

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func NewMatrix(s string) (Matrix, error) {
	if len(s) == 0 {
		return Matrix{}, errors.New("empty string")
	}
	row := strings.Split(s, "\n")
	var prevColL int
	var mat Matrix
	for i, r := range row {
		if r == "" {
			return Matrix{}, errors.New("error")
		}
		var x []int
		col := strings.Split(strings.TrimSpace(r), " ")
		if i > 0 && prevColL != len(col) {
			return Matrix{}, errors.New("mismatch row")
		}
		prevColL = len(col)
		for _, c := range col {
			c = strings.TrimSpace(c)
			num, err := strconv.Atoi(c)
			if err != nil {
				return Matrix{}, errors.New("not an integer")
			}
			x = append(x, num)
		}
		mat = append(mat, x)
	}
	return mat, nil
}

func (m Matrix) Cols() [][]int {
	var colMat [][]int
	for i := 0; i < len(m[0]); i++ {
		var c []int
		for j := 0; j < len(m); j++ {
			c = append(c, m[j][i])
		}
		colMat = append(colMat, c)
	}
	return colMat
}

func (m Matrix) Rows() [][]int {
	copyMat := make([][]int, len(m))
	for i := range m {
		copyMat[i] = make([]int, len(m[i]))
		copy(copyMat[i], m[i])
	}
	return copyMat
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
