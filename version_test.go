package version

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMax(t *testing.T) {
	versionStrings := []string{"29.20", "29.1", "1", "23.4345.3", "271.13.0", "invalid.version"}
	expectedMax := "271.13.0"

	max, err := Max(versionStrings)
	if err != nil {
		t.Errorf("Failed to get max version: %v", err)
	}

	if err != nil {
		t.Errorf("Failed to parse expected max version: %v", err)
	}

	if max != expectedMax {
		t.Errorf("Max version is %v, expected %v", max, expectedMax)
	}
}

func TestMin(t *testing.T) {
	versionStrings := []string{"29.20", "29.1", "1", "23.4345.3", "271.13.0", "invalid.version"}
	expectedMin := "1"

	min, err := Min(versionStrings)
	if err != nil {
		t.Errorf("Failed to get min version: %v", err)
	}

	if min != expectedMin {
		t.Errorf("Min version is %v, expected %v", min, expectedMin)
	}
}

func TestLocate(t *testing.T) {
	versionStrings := []string{"29.20", "29.1", "1", "23.4345.3", "271.13.0", "invalid.version"}
	versionsToLocate := []string{"29.1", "23.4345.3"}

	for _, v := range versionsToLocate {
		parsedVersion, err := Parse(v)
		if err != nil {
			t.Errorf("Failed to parse version %s: %v", v, err)
			continue
		}

		if !Locate(parsedVersion, versionStrings) {
			t.Errorf("Locate function returned false for version %s", v)
		}
	}

	versionNotInSet := "100.200.300"
	parsedVersionNotInSet, err := Parse(versionNotInSet)
	if err != nil {
		t.Errorf("Failed to parse version %s: %v", versionNotInSet, err)
	}

	if Locate(parsedVersionNotInSet, versionStrings) {
		t.Errorf("Locate function returned true for version not in the set: %s", versionNotInSet)
	}
}

func TestPerformance(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	versionStrings := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		v := Version{rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)}
		versionStrings[i] = strconv.Itoa(v[0]) + "." + strconv.Itoa(v[1]) + "." + strconv.Itoa(v[2])
	}

	start := time.Now()
	_, err := Max(versionStrings)
	if err != nil {
		t.Fatalf("Error finding max version: %v", err)
	}
	duration := time.Since(start)
	t.Logf("Time to find max of 1,000 versions: %v", duration)

	start = time.Now()
	_, err = Min(versionStrings)
	if err != nil {
		t.Fatalf("Error finding min version: %v", err)
	}
	duration = time.Since(start)
	t.Logf("Time to find min of 1,000 versions: %v", duration)

	testVersion := Version{rand.Intn(1000), rand.Intn(1000), rand.Intn(1000)}
	start = time.Now()
	found := Locate(testVersion, versionStrings)
	duration = time.Since(start)
	t.Logf("Time to locate version %v in 1,000 versions: %v. Found: %v", testVersion, duration, found)
}

func TestParseWithInvalidVersion(t *testing.T) {
	versionStrings := []string{"29.20", "29.1", "13", "200-SNAPSHOT"}

	expectedMax := "29.20"

	max, err := Max(versionStrings)
	if err != nil {
		t.Errorf("Failed to parse version %s: %v", expectedMax, err)
	}

	if !(max == expectedMax) {
		t.Errorf("Max version is %v, expected %v", max, expectedMax)
	}
}
