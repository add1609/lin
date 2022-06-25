package rn

import (
	"math"
	"testing"
)

func TestAt(t *testing.T) {
	for _, test := range []struct {
		v1   Vec
		i    int
		want float64
	}{
		{Vec{3, []float64{1, 1, 1}}, 0, 1},
		{Vec{9, []float64{
			1, 1, 1,
			1, 1, 3,
			1, 1, 1,
		}}, 5, 3},
		{Vec{9, []float64{
			1, 1, 1,
			1, 1, 1,
			1, 1, -4,
		}}, 8, -4},
		{Vec{9, []float64{
			1, 1, 1,
			3.14, 1, 1,
			1, 1, 1,
		}}, 3, 3.14},
	} {
		got := test.v1.At(test.i)
		if got != test.want {
			t.Errorf(
				"error:\ngot=%v\nwant=%v",
				got, test.want,
			)
		}
	}
}

func TestSet(t *testing.T) {
	for _, test := range []struct {
		v1, v2 Vec
		i      int
		val    float64
	}{
		{
			Vec{3, []float64{1, 1, 1}},
			Vec{3, []float64{3, 1, 1}},
			0, 3,
		},

		{
			Vec{3, []float64{1, 1, 1}},
			Vec{3, []float64{1, 0, 1}},
			1, 0,
		},

		{
			Vec{3, []float64{1, 1, 1}},
			Vec{3, []float64{-5, 1, 1}},
			0, -5,
		},

		{
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, -3}},
			8, -3,
		},

		{
			Vec{15, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			Vec{15, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1}},
			12, 0,
		},
	} {
		test.v1.Set(test.i, test.val)
		if !VecEqual(test.v1, test.v2) {
			t.Errorf(
				"error:\nv1=%v\nv2=%v",
				test.v1.X, test.v2.X,
			)
		}
	}
}

func TestSetSlice(t *testing.T) {
	for _, test := range []struct {
		v1, v2, want Vec
	}{
		{Vec{2, []float64{1, 1}}, Vec{2, []float64{0, 0}}, Vec{2, []float64{0, 0}}},
		{Vec{3, []float64{1, 1, 1}}, Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{0, 0, 0}}},
		{
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
			Vec{9, []float64{1, 4, 7, 8, 2, 5, 6, 4, 3}},
			Vec{9, []float64{1, 4, 7, 8, 2, 5, 6, 4, 3}},
		},
		{
			Vec{15, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			Vec{15, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
			Vec{15, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		},
	} {
		test.v1.SetSlice(test.v2.X)
		if !VecEqual(test.v1, test.want) {
			t.Errorf(
				"error:\nv1 = %v\nv2 = %v\nwant=%v",
				test.v1.X, test.v2.X, test.want.X,
			)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, test := range []struct {
		v1, v2, want Vec
	}{
		{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{0, 0, 0}}},
		{Vec{3, []float64{1, 0, 0}}, Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 0, 0}}},
		{Vec{3, []float64{1, 2, 3}}, Vec{3, []float64{4, 5, 6}}, Vec{3, []float64{5, 7, 9}}},
		{Vec{3, []float64{1, -3, 5}}, Vec{3, []float64{1, -6, -2}}, Vec{3, []float64{2, -9, 3}}},
		{Vec{3, []float64{1, 2, 3}}, Vec{3, []float64{-1, -2, -3}}, Vec{3, []float64{0, 0, 0}}},
		{
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			Vec{9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Vec{9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			Vec{9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			Vec{9, []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9}},
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			Vec{9, []float64{1, 2, 0.25, 0, 1, -0.0, 3, 0, 1}},
			Vec{9, []float64{0, -3, 0.25, 2, 16, 0.0, -9, 0, 0}},
			Vec{9, []float64{1, -1, 0.5, 2, 17, 0, -6, 0, 1}},
		},
	} {
		got := Add(test.v1, test.v2)
		if !VecEqual(got, test.want) {
			t.Errorf(
				"error:\n%v + %v:\ngot=%v\nwant=%v",
				test.v1.X, test.v2.X, got.X, test.want.X,
			)
		}
	}
}

func TestSub(t *testing.T) {
	for _, test := range []struct {
		v1, v2, want Vec
	}{
		{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{0, 0, 0}}},
		{Vec{3, []float64{1, 0, 0}}, Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 0, 0}}},
		{Vec{3, []float64{1, 2, 3}}, Vec{3, []float64{4, 5, 6}}, Vec{3, []float64{-3, -3, -3}}},
		{Vec{3, []float64{1, -3, 5}}, Vec{3, []float64{1, -6, -2}}, Vec{3, []float64{0, 3, 7}}},
		{Vec{3, []float64{1, 2, 3}}, Vec{3, []float64{-1, -2, -3}}, Vec{3, []float64{2, 4, 6}}},
	} {
		got := Sub(test.v1, test.v2)
		if !VecEqual(got, test.want) {
			t.Errorf(
				"error:\n%v - %v:\ngot=%v\nwant=%v",
				test.v1.X, test.v2.X, got.X, test.want.X,
			)
		}
	}
}

func TestAbs(t *testing.T) {
	for _, test := range []struct {
		v1, want Vec
	}{
		{
			Vec{3, []float64{0, 0, 0}},
			Vec{3, []float64{0, 0, 0}},
		},
		{
			Vec{3, []float64{1, 1, 1}},
			Vec{3, []float64{1, 1, 1}},
		},
		{
			Vec{3, []float64{-1, -1, -1}},
			Vec{3, []float64{1, 1, 1}},
		},
		{
			Vec{3, []float64{-5, 3, 0}},
			Vec{3, []float64{5, 3, 0}},
		},

		{
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
		},
		{
			Vec{9, []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9}},
			Vec{9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	} {
		got := Abs(test.v1)
		if !VecEqual(got, test.want) {
			t.Errorf(
				"error:\n|%v|\ngot=%v\nwant=%v",
				test.v1.X, got.X, test.want.X,
			)
		}
	}
}

func TestScale(t *testing.T) {
	for _, test := range []struct {
		v1, want Vec
		r        float64
	}{
		{
			Vec{3, []float64{0, 0, 0}},
			Vec{3, []float64{0, 0, 0}},
			5,
		},
		{
			Vec{3, []float64{1, 1, 1}},
			Vec{3, []float64{1, 1, 1}},
			1,
		},
		{
			Vec{3, []float64{1, 1, 1}},
			Vec{3, []float64{0, 0, 0}},
			0,
		},
		{
			Vec{3, []float64{-5, 3, 0}},
			Vec{3, []float64{5, -3, -0}},
			-1,
		},
		{
			Vec{3, []float64{-1, -1, -1}},
			Vec{3, []float64{2, 2, 2}},
			-2,
		},

		{
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			5,
		},
		{
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
			1,
		},
		{
			Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}},
			Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			0,
		},
		{
			Vec{9, []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9}},
			Vec{9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			-1,
		},
	} {
		got := Scale(test.r, test.v1)
		if !VecEqual(got, test.want) {
			t.Errorf(
				"error:\n%v * %v\ngot=%v\nwant=%v",
				test.r, test.v1.X, got.X, test.want.X,
			)
		}
	}
}

func TestDot(t *testing.T) {
	for _, test := range []struct {
		v1, v2 Vec
		want   float64
	}{
		{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 1, 1}}, 0},
		{Vec{3, []float64{1, 2, 3}}, Vec{3, []float64{1, 2, 3}}, 14},
		{Vec{3, []float64{1, 2, 3}}, Vec{3, []float64{3, 4, 5}}, 26},
		{Vec{3, []float64{1, 2, 3}}, Vec{3, []float64{-1, -2, -3}}, -14},
		{Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}}, Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}}, 0},
		{Vec{9, []float64{9, 8, 7, 6, 5, 4, 3, 2, 1}}, Vec{9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 165},
		{Vec{9, []float64{9, 2, 7, 4, 5, 6, 3, 8, 1}}, Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}}, 45},
		{Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}}, Vec{9, []float64{-1, -1, -1, -1, -1, -1, -1, -1, -1}}, -9},
		{Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}}, Vec{9, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}}, 0},
	} {
		{
			got := Dot(test.v1, test.v2)
			if got != test.want {
				t.Errorf(
					"error\n%v · %v:\ngot=%v\nwant=%v",
					test.v1.X, test.v2.X, got, test.want,
				)
			}
		}
		{
			got := Dot(test.v2, test.v1)
			if got != test.want {
				t.Errorf(
					"error\n%v · %v:\ngot=%v\nwant=%v",
					test.v2.X, test.v1.X, got, test.want,
				)
			}
		}
	}
}

func TestCross(t *testing.T) {
	for _, test := range []struct {
		v1, v2, want Vec
	}{
		{Vec{3, []float64{0, 0, 0}}, Vec{3, []float64{1, 1, 1}}, Vec{3, []float64{0, 0, 0}}},
		{Vec{3, []float64{1, 1, 1}}, Vec{3, []float64{1, 1, 1}}, Vec{3, []float64{0, 0, 0}}},
		{Vec{3, []float64{1, 0, 0}}, Vec{3, []float64{0, 1, 0}}, Vec{3, []float64{0, 0, 1}}},
		{Vec{3, []float64{8, 2, 5}}, Vec{3, []float64{5, 1, 3}}, Vec{3, []float64{1, 1, -2}}},
		{Vec{3, []float64{5, 1, 3}}, Vec{3, []float64{8, 2, 5}}, Vec{3, []float64{-1, -1, 2}}},
	} {
		got := Cross(test.v1, test.v2)
		if !VecEqual(got, test.want) {
			t.Errorf(
				"error:\n%v × %v:\ngot = %v\nwant = %v",
				test.v1.X, test.v2.X, got.X, test.want.X,
			)
		}
	}
}

func TestNorm(t *testing.T) {
	for _, test := range []struct {
		v1   Vec
		want float64
	}{
		{Vec{3, []float64{1, 0, 0}}, 1},
		{Vec{3, []float64{1, 2, 1}}, math.Sqrt(6)},
		{Vec{3, []float64{-2, 5, 3}}, math.Sqrt(38)},
		{Vec{3, []float64{-5, 2, -3}}, math.Sqrt(38)},
		{Vec{3, []float64{-0.1, 1.1, 3.2}}, 3.385262175962152},
		{Vec{9, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}}, 0},
		{Vec{9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, math.Sqrt(285)},
		{Vec{9, []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9}}, math.Sqrt(285)},
		{Vec{9, []float64{5, 1, 0, 35, 5, 16, 3, 0, 9}}, math.Sqrt(1622)},
	} {
		got := Norm(test.v1)
		if got != test.want {
			t.Errorf(
				"error:\n|%v| = %v, want %v",
				test.v1.X, got, test.want,
			)
		}
	}
}

func TestVecN(t *testing.T) {
	for _, test := range []struct {
		n    int
		val  float64
		want Vec
	}{
		{1, 0, Vec{1, []float64{0}}},
		{2, 5, Vec{2, []float64{5, 5}}},
		{3, 9, Vec{3, []float64{9, 9, 9}}},
		{9, 1, Vec{9, []float64{
			1, 1, 1,
			1, 1, 1,
			1, 1, 1,
		}}},
		{27, 1, Vec{27, []float64{
			1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1,
		}}},
	} {
		got := VecN(test.n, test.val)
		if !VecEqual(got, test.want) {
			t.Errorf(
				"error:\nN:%v, value: %v\ngot = %v\nwant = %v",
				test.n, test.val, got.X, test.want.X,
			)
		}
	}
}
