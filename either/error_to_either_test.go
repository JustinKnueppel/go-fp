package either_test

import (
	"errors"
	"fmt"

	"github.com/JustinKnueppel/go-fp/either"
)

func ExampleErrorToEither0() {
	good := func() (int, error) {
		return 5, nil
	}
	bad := func() (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither0(good)())
	fmt.Println(either.ErrorToEither0(bad)())

	// Output:
	// Right 5
	// Left failed
}

func ExampleErrorToEither1() {
	good := func(x int) (int, error) {
		return x, nil
	}
	bad := func(int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither1(good)(1))
	fmt.Println(either.ErrorToEither1(bad)(1))

	// Output:
	// Right 1
	// Left failed
}

func ExampleErrorToEither2() {
	good := func(_, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither2(good)(1, 2))
	fmt.Println(either.ErrorToEither2(bad)(1, 2))

	// Output:
	// Right 2
	// Left failed
}

func ExampleErrorToEither3() {
	good := func(_, _, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither3(good)(1, 2, 3))
	fmt.Println(either.ErrorToEither3(bad)(1, 2, 3))

	// Output:
	// Right 3
	// Left failed
}

func ExampleErrorToEither4() {
	good := func(_, _, _, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int, int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither4(good)(1, 2, 3, 4))
	fmt.Println(either.ErrorToEither4(bad)(1, 2, 3, 4))

	// Output:
	// Right 4
	// Left failed
}

func ExampleErrorToEither5() {
	good := func(_, _, _, _, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int, int, int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither5(good)(1, 2, 3, 4, 5))
	fmt.Println(either.ErrorToEither5(bad)(1, 2, 3, 4, 5))

	// Output:
	// Right 5
	// Left failed
}

func ExampleErrorToEither6() {
	good := func(_, _, _, _, _, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int, int, int, int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither6(good)(1, 2, 3, 4, 5, 6))
	fmt.Println(either.ErrorToEither6(bad)(1, 2, 3, 4, 5, 6))

	// Output:
	// Right 6
	// Left failed
}

func ExampleErrorToEither7() {
	good := func(_, _, _, _, _, _, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int, int, int, int, int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither7(good)(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(either.ErrorToEither7(bad)(1, 2, 3, 4, 5, 6, 7))

	// Output:
	// Right 7
	// Left failed
}

func ExampleErrorToEither8() {
	good := func(_, _, _, _, _, _, _, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int, int, int, int, int, int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither8(good)(1, 2, 3, 4, 5, 6, 7, 8))
	fmt.Println(either.ErrorToEither8(bad)(1, 2, 3, 4, 5, 6, 7, 8))

	// Output:
	// Right 8
	// Left failed
}

func ExampleErrorToEither9() {
	good := func(_, _, _, _, _, _, _, _, x int) (int, error) {
		return x, nil
	}
	bad := func(int, int, int, int, int, int, int, int, int) (int, error) {
		return 0, errors.New("failed")
	}

	fmt.Println(either.ErrorToEither9(good)(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(either.ErrorToEither9(bad)(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// Output:
	// Right 9
	// Left failed
}
