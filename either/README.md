# Either [![Go Reference](https://pkg.go.dev/badge/github.com/JustinKnueppel/go-fp/either.svg)](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/either)

Either represents one of two specified types, but not both. By convention, `Left` represents a failure case and `Right` represents the success case. This can be used in place of an `Option` if there is information that should be forwarded on in lieu of a `None`. Additionally this can be used to make error handling in Go more functional. This package was largely inspired by Rust's [Result type](https://doc.rust-lang.org/std/result/enum.Result.html).

```go
import (
  "github.com/JustinKnueppel/go-fp/either"
)
```

## Usage

The only exported type from this package is the `Either[L, R]` generic type. An `Either` can never be instantiated directly, but rather only through the `Left(L)` or `Right(R)` constructors. An `Either` has no logical zero value. All functions in this package are curried with their final argument being the target `Either` to make composing functions as painless as possible. A few examples follow, but small examples for each function can be found in the package documentation.

### Capturing exit code

In this example we could be making some external system call that could return an error code. In some instances we may not be prepared to process the specific error code, but do not want to discard that information with something like an `Option` type.

```go
func externalCall(args []string) either.Either[int, []string] {
  result, exitCode := someSyscall(args)
  if exitCode > 0 {
    return either.Left[int, []string](exitCode)
  }
  return either.Right[int](strings.Split(result, "\n"))
}

func workWithReturnedValues([]string) string {}

func main() {
  fp.Pipe3(
    either.InspectLeft[int, []string](func(exitCode int) {
      fmt.Printf("External called with status code: %d\n", exitCode)
      os.Exit(1)
    }),
    either.Map[int, []string](workWithReturnedValues),
    either.Inspect[int](func (result string) {
      fmt.Println("Final result: %s\n", result)
    })
  )(externalCall([]string{arg1, arg2}))
}
```

In this example we can actually address the error case anywhere in the Pipe as none of the other functions will run if given a `Left`. This could allow us to pipe in `Bind` calls which could return their own errors potentially, and we could handle all of those at the end if desired.

### Dealing with errors once

Another more common use case would be a situation where we have multiple functions that can create an error, but we are handling each of those errors in the same way. Instead of having duplicate code to either propogate errors or panic on them, we could write the error handling once. This example will show a before and after of how the code could change.

Before:

```go
func getConfigFromContext(*http.Request) (Config, error) {}
func getClientset(Config) (Clientset, error) {}
func openConnection(Clientset) (Connection, error) {}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  config, err := getConfigFromContext(r)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  clientset, err := getClientset(config)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  conn, err := openConnection(clientset)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  // Now we can handle the request with our open connection
}
```

After:

```go
func getConfigFromContext(*http.Request) either.Either[error, Config] {}
func getClientset(Config) either.Either[error, Clientset] {}
func openConnection(Clientset) either.Either[error, Connection] {}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  fp.Pipe5(
    getConfigFromContext,
    getClientset,
    openConnection,
    either.InspectLeft[error, Connection](func(err error) {
      w.WriteHeader(http.StatusInternalServerError)
    }),
    either.Inspect[error](func(conn Connection) {
      //Handle the request here
    })
  )(r)
}
```

In this example, we are able to handle the error once and allow the errors to propogate without having to handle them explicitly. In order to keep our main business logic from being nested, we could also call a function inside the `Inspect` function which takes the request, response, and connection and leave the error handling as separate from the main logic.
