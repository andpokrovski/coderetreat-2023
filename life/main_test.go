package main

import (
	"strconv"
	"testing"
)

func TestGame_CalculateState(t *testing.T) {
	testCases := []struct {
		preset    Preset
		fieldSize int
		expField  func() Field
	}{
		{
			preset:    Preset{{0, 0}},
			fieldSize: 1,
			expField: func() Field {
				return CreateField(1)
			},
		},
		{
			preset:    Preset{{0, 0}},
			fieldSize: 1,
			expField: func() Field {
				return CreateField(1)
			},
		},
		{
			preset:    Preset{{1, 0}, {1, 1}, {1, 2}},
			fieldSize: 3,
			expField: func() Field {
				f := CreateField(3)

				f[0][1] = true
				f[1][1] = true
				f[2][1] = true

				return f
			},
		},
		// Если только одна из клеток живая - на следующей итерации умрет, таких случая - 4

	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			game := New(tc.preset, tc.fieldSize)
			game.CalculateState()
			res := game.State()

			if tc.expField()[0][0] != res[0][0] {
				t.Fail()
			}
		})
	}
}

func compareFields(a, b Field) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

// 16 состояний
