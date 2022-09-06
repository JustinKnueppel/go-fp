package fp_test

import (
	"testing"

	"github.com/JustinKnueppel/go-fp"
)

func TestPipe2(t *testing.T) {
	t.Run("composes left to right", func(t *testing.T) {
		initial := 2
		expected := 6

		if fp.Pipe2(add1, double)(initial) != expected {
			t.Fail()
		}
	})

	t.Run("composes across types", func(t *testing.T) {
		initial := "hello"
		expected := 10

		if fp.Pipe2(strlen, double)(initial) != expected {
			t.Fail()
		}
	})
}

func TestPipe3(t *testing.T) {
	if fp.Pipe3(double, double, add1)(2) != 9 {
		t.Fatal("should compose 3 functions left to right")
	}
}

func TestPipe4(t *testing.T) {
	if fp.Pipe4(double, double, add1, add1)(2) != 10 {
		t.Fatal("should compose 4 functions left to right")
	}
}

func TestPipeNested(t *testing.T) {
	add1 := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	triple := func(x int) int { return x * 3 }

	result4 := fp.Pipe4(triple, add1, double, add1)(5)
	result2 := fp.Pipe2(fp.Pipe2(triple, add1), fp.Pipe2(double, add1))(5)

	if result2 != result4 {
		t.Fatal("nesting compose should be equal to larger compose")
	}
}

func TestPipe5(t *testing.T) {
	if fp.Pipe5(double, double, add1, add1, add1)(2) != 11 {
		t.Fatal("should compose 5 functions left to right")
	}
}

func TestPipe6(t *testing.T) {
	if fp.Pipe6(double, double, add1, add1, add1, add1)(2) != 12 {
		t.Fatal("should compose 6 functions left to right")
	}
}

func TestPipe7(t *testing.T) {
	if fp.Pipe7(double, double, add1, add1, add1, add1, add1)(2) != 13 {
		t.Fatal("should compose 7 functions left to right")
	}
}

func TestPipe8(t *testing.T) {
	if fp.Pipe8(double, double, add1, add1, add1, add1, add1, add1)(2) != 14 {
		t.Fatal("should compose 8 functions left to right")
	}
}

func TestPipe9(t *testing.T) {
	if fp.Pipe9(double, double, add1, add1, add1, add1, add1, add1, add1)(2) != 15 {
		t.Fatal("should compose 9 functions left to right")
	}
}
