# defaultcasechecker
defaultcasechecker is a program for checking switch statement whether having default case.

Go does not provide exhaustive checks for switch statements. However, switch statements should always have a default case that either returns an appropriate error or performs some action. Therefore, this program identifies parts of switch statements that lack a default case.

## Install
```sh
go install github/take-2405/defaultcasechecker@latest
```

## Usage
Basic usage, specify the path of the package to check.
```sh
defaultcasechecker github.com/take-2405/defaultcasechecker
```

To check all packages
```sh
defaultcasechecker ./...
```

#### ignore test file
```sh
defaultcasechecker -check-test=true ./...
```

#### ignore generated file
```sh
defaultcasechecker -check-generate=true ./...
```

#### ignore specify switch statement
```go
	//lint:ignore defaultcasechecker test ignore comment
	switch {
	case true:
		// do something
	case false:
		// do something
	}
```
