package version

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Version type represents a parsed version number.
type Version []int

// Parse parses a version string into a Version type, ignoring non-integer parts.
func Parse(version string) (Version, error) {
	parts := strings.Split(version, ".")
	v := make(Version, 0, len(parts)) // Change v to dynamic size slice
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			continue // Skip non-integer parts
		}
		v = append(v, num)
	}
	if len(v) == 0 {
		return nil, nil // Return nil when v is empty
	}
	return v, nil
}

// Compare compares a Version with another.
// Returns 1 if v is greater, -1 if it is smaller, 0 if they're equal.
func (v Version) Compare(other Version) int {
	minLen := len(v)
	if len(other) < minLen {
		minLen = len(other)
	}

	for i := 0; i < minLen; i++ {
		if v[i] > other[i] {
			return 1
		} else if v[i] < other[i] {
			return -1
		}
	}

	if len(v) > len(other) {
		return 1
	} else if len(v) < len(other) {
		return -1
	}

	return 0
}

// Max returns the maximum version from a list of version strings.
type VersionPair struct {
	version Version
	str     string
}

// Max returns the maximum version from a list of version strings.
func Max(versionStrings []string) (string, error) {
	if len(versionStrings) == 0 {
		return "", fmt.Errorf("empty list of version strings")
	}

	versions := make([]VersionPair, 0, len(versionStrings))
	for _, v := range versionStrings {
		parsedVersion, _ := Parse(v) // Ignore error
		if parsedVersion != nil {
			versions = append(versions, VersionPair{version: parsedVersion, str: v})
		}
	}

	if len(versions) == 0 {
		return "", fmt.Errorf("no valid version found")
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].version.Compare(versions[j].version) == -1
	})

	return versions[len(versions)-1].str, nil
}

// Min returns the minimum version from a list of version strings.
func Min(versionStrings []string) (string, error) {
	if len(versionStrings) == 0 {
		return "", fmt.Errorf("empty list of version strings")
	}

	versions := make([]VersionPair, 0, len(versionStrings))
	for _, v := range versionStrings {
		parsedVersion, _ := Parse(v) // Ignore error
		if parsedVersion != nil {
			versions = append(versions, VersionPair{version: parsedVersion, str: v})
		}
	}

	if len(versions) == 0 {
		return "", fmt.Errorf("no valid version found")
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].version.Compare(versions[j].version) < 0
	})

	return versions[0].str, nil
}

// Locate checks if a version exists in the list of version strings.
func Locate(version Version, versionStrings []string) bool {
	for _, v := range versionStrings {
		parsedVersion, err := Parse(v)
		if err != nil {
			continue
		}
		if equal(parsedVersion, version) {
			return true
		}
	}
	return false
}

// equal checks if two Versions are equal.
func equal(a, b Version) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
