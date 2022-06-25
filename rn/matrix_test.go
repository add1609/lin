package rn

import "testing"

func TestMatAt(t *testing.T) {
	for _, test := range []struct {
		m    Mat
		i, j int
		want float64
	}{
		//  [1]
		//  [2]
		//  [3]
		{Mat{3, 1, COL, []float64{1, 2, 3}}, 0, 0, 1},
		{Mat{3, 1, COL, []float64{1, 2, 3}}, 0, 1, 2},
		{Mat{3, 1, COL, []float64{1, 2, 3}}, 0, 2, 3},

		{Mat{3, 1, ROW, []float64{1, 2, 3}}, 0, 0, 1},
		{Mat{3, 1, ROW, []float64{1, 2, 3}}, 1, 0, 2},
		{Mat{3, 1, ROW, []float64{1, 2, 3}}, 2, 0, 3},

		//  [1, 2, 3]
		{Mat{1, 3, COL, []float64{1, 2, 3}}, 0, 0, 1},
		{Mat{1, 3, COL, []float64{1, 2, 3}}, 1, 0, 2},
		{Mat{1, 3, COL, []float64{1, 2, 3}}, 2, 0, 3},

		{Mat{1, 3, ROW, []float64{1, 2, 3}}, 0, 0, 1},
		{Mat{1, 3, ROW, []float64{1, 2, 3}}, 0, 1, 2},
		{Mat{1, 3, ROW, []float64{1, 2, 3}}, 0, 2, 3},

		//  [1, 2, 3]
		//  [4, 5, 6]
		{Mat{2, 3, COL, []float64{1, 2, 3, 4, 5, 6}}, 0, 0, 1},
		{Mat{2, 3, COL, []float64{1, 2, 3, 4, 5, 6}}, 1, 0, 2},
		{Mat{2, 3, COL, []float64{1, 2, 3, 4, 5, 6}}, 2, 0, 3},
		{Mat{2, 3, COL, []float64{1, 2, 3, 4, 5, 6}}, 0, 1, 4},
		{Mat{2, 3, COL, []float64{1, 2, 3, 4, 5, 6}}, 1, 1, 5},
		{Mat{2, 3, COL, []float64{1, 2, 3, 4, 5, 6}}, 2, 1, 6},
	} {
		got := test.m.At(test.i, test.j)
		if got != test.want {
			t.Errorf(
				"error:\nmat:  %v\nwant: mat[%v][%v] %v\ngot:  %v",
				test.m, test.i, test.j, test.want, got,
			)
		}
	}
}

func TestMatT(t *testing.T) {
	for _, test := range []struct {
		m, want Mat
	}{
		{
			Mat{1, 1, COL, []float64{1}},
			Mat{1, 1, ROW, []float64{1}},
		},
		{
			Mat{1, 3, COL, []float64{1, 2, 3}},
			Mat{3, 1, ROW, []float64{1, 2, 3}},
		},
		{
			Mat{2, 3, COL, []float64{1, 2, 3, 4, 5, 6}},
			Mat{3, 2, ROW, []float64{1, 2, 3, 4, 5, 6}},
		},
		{
			Mat{3, 3, COL, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			Mat{3, 3, ROW, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	} {
		got := test.m.T()
		if !MatEqual(got, test.want) {
			t.Errorf(
				"error:\nmat:  %v\nwant: %v\ngot:  %v",
				test.m, test.want, got,
			)
		}

	}
}
