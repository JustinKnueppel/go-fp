package fp_test

import (
	"testing"

	"github.com/JustinKnueppel/go-fp"
)

func add1(x int) int {
	return x + 1
}

func double(x int) int {
	return x * 2
}

func strlen(s string) int {
	return len(s)
}

func TestCompose2(t *testing.T) {
	t.Run("composes right to left", func(t *testing.T) {
		initial := 2
		expected := 6

		if fp.Compose2(double, add1)(initial) != expected {
			t.Fail()
		}
	})

	t.Run("composes across types", func(t *testing.T) {
		initial := "hello"
		expected := 10

		if fp.Compose2(double, strlen)(initial) != expected {
			t.Fail()
		}
	})
}

func TestCompose3(t *testing.T) {
	if fp.Compose3(add1, double, double)(2) != 9 {
		t.Fatal("should compose 3 functions right to left")
	}
}

func TestCompose4(t *testing.T) {
	if fp.Compose4(add1, add1, double, double)(2) != 10 {
		t.Fatal("should compose 4 functions right to left")
	}
}

func TestCompose5(t *testing.T) {
	if fp.Compose5(add1, add1, add1, double, double)(2) != 11 {
		t.Fatal("should compose 5 functions right to left")
	}
}

func TestCompose6(t *testing.T) {
	if fp.Compose6(add1, add1, add1, add1, double, double)(2) != 12 {
		t.Fatal("should compose 6 functions right to left")
	}
}

func TestCompose7(t *testing.T) {
	if fp.Compose7(add1, add1, add1, add1, add1, double, double)(2) != 13 {
		t.Fatal("should compose 7 functions right to left")
	}
}

func TestCompose8(t *testing.T) {
	if fp.Compose8(add1, add1, add1, add1, add1, add1, double, double)(2) != 14 {
		t.Fatal("should compose 8 functions right to left")
	}
}

func TestCompose9(t *testing.T) {
	if fp.Compose9(add1, add1, add1, add1, add1, add1, add1, double, double)(2) != 15 {
		t.Fatal("should compose 9 functions right to left")
	}
}
