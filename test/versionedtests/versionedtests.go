package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

func parseVersion(version string) (int, error) {
	versions := strings.Split(version, ".")
	if len(versions) != 3 {
		return 0, fmt.Errorf("version does not have 3 numbers seperated by dots: %s", version)
	}
	n := 0
	for _, v := range versions {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		n = n*10 + i
	}
	return n, nil
}

func less(this, that string) bool {
	thisNum, err := parseVersion(this)
	if err != nil {
		panic(err)
	}
	thatNum, err := parseVersion(that)
	if err != nil {
		panic(err)
	}
	return thisNum < thatNum
}

func getVersion() string {
	versionBytes, _ := exec.Command("protoc", "--version").CombinedOutput()
	version := strings.TrimSpace(string(versionBytes))
	versions := strings.Split(version, " ")
	if len(versions) != 2 {
		panic("version string returned from protoc is seperated with a space: " + version)
	}
	return versions[1]
}

func findVersion(line string) string {
	words := strings.Split(line, " ")
	for _, word := range words {
		_, err := parseVersion(word)
		if err == nil {
			return word
		}
	}
	panic("no version found in line: " + line)
}

func main() {
	currentVersion := getVersion()
	data, err := ioutil.ReadFile("thetest.proto")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	inside := false
	for i, line := range lines {
		if strings.HasPrefix(line, "//begin versioned") {
			minVersion := findVersion(line)
			if less(currentVersion, minVersion) {
				inside = true
			}
			continue
		}
		if inside {
			if strings.HasPrefix(line, "//end versioned") {
				inside = false
				continue
			}
			lines[i] = `// ` + line
		}
	}
	if err := ioutil.WriteFile("thetest.proto", []byte(strings.Join(lines, "\n")), 0666); err != nil {
		panic(err)
	}
}
