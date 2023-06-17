## version comparison in go

Small Go module to parse, compare, and find the minimum and maximum from a list of version strings while attempting to be as performant as possible. This module also provides a function to locate a version in a list of version strings.

### Features

- Parse a version string into a slice of integers, ignoring non-integer parts.
- Find the minimum and maximum version from a list of version strings.
- Locate a version in a list of version strings.
- Compare versions: returns 1 if a version is greater than another, -1 if smaller, 0 if they're equal.


### Usage

Import the package into your Go application:

```go
import "github.com/drew-mcl/version"
```

Use the Max and Min functions to find the maximum and minimum version in a list of version strings:

```go
max, err := version.Max([]string{"1.2.3", "2.3.4", "3.4.5"})
min, err := version.Min([]string{"1.2.3", "2.3.4", "3.4.5"})
```

Use the Locate function to check if a version exists in a list of version strings:

```go
exists := version.Locate(version, []string{"1.2.3", "2.3.4", "3.4.5"})
```

Use the Parse function to convert a version string into a Version type:
```go
v, err := version.Parse("1.2.3")
```

Use the Compare method to compare two versions:
```go
comparison := v.Compare(otherVersion) // Returns 1, -1, or 0
```

### Tests

Run the tests using the go test command:

```sh
Copy code
go test ./...
```
### License

This project is licensed under the MIT License.