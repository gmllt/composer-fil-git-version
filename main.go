package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Composer struct {
	Name               string      `json:"name,omitempty"`
	Description        interface{} `json:"description,omitempty"`
	Version            string      `json:"version,omitempty"`
	Type               interface{} `json:"type,omitempty"`
	Keywords           interface{} `json:"keywords,omitempty"`
	HomePage           interface{} `json:"homepage,omitempty"`
	Readme             interface{} `json:"readme,omitempty"`
	Time               interface{} `json:"time,omitempty"`
	License            interface{} `json:"license,omitempty"`
	Authors            interface{} `json:"authors,omitempty"`
	Support            interface{} `json:"support,omitempty"`
	Require            interface{} `json:"require,omitempty"`
	RequireDev         interface{} `json:"require-dev,omitempty"`
	Conflict           interface{} `json:"conflict,omitempty"`
	Replace            interface{} `json:"replace,omitempty"`
	Provide            interface{} `json:"provide,omitempty"`
	Suggest            interface{} `json:"suggest,omitempty"`
	Autoload           interface{} `json:"autoload,omitempty"`
	AutoloadDev        interface{} `json:"autoload-dev,omitempty"`
	IncludePath        interface{} `json:"include-path,omitempty"`
	TargetDir          interface{} `json:"target-dir,omitempty"`
	MinimumStability   interface{} `json:"minimum-stability,omitempty"`
	PreferStable       interface{} `json:"prefer-stable,omitempty"`
	Repositories       interface{} `json:"repositories,omitempty"`
	Config             interface{} `json:"config,omitempty"`
	Scripts            interface{} `json:"scripts,omitempty"`
	Extra              interface{} `json:"extra,omitempty"`
	Bin                interface{} `json:"bin,omitempty"`
	Archive            interface{} `json:"archive,omitempty"`
	Abandoned          interface{} `json:"abandoned,omitempty"`
	NonFeatureBranches interface{} `json:"non-feature-branches,omitempty"`
}

func main() {
	var args = os.Args[1:]

	var version = ""

	if len(args) >= 1 {
		version = args[0]
	} else {
		version = getGitVersion()
	}

	fmt.Println("Setting version to '"+version+"'")

	jsonFile, err := os.Open("composer.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	var composer Composer
	err = json.Unmarshal(byteValue, &composer)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	composer.Version = version

	jsonByte, err := json.Marshal(composer)
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	composer.Version = version

	err = ioutil.WriteFile("composer.json", jsonByte, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(7)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	err = jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(8)
	}
}

func getGitVersion() string {
	var currentVersion = "0.0.0"
	bytes, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		fmt.Println("No version found, using default '0.0.0'.")
	} else {
		currentVersion = strings.TrimSpace(string(bytes))
		fmt.Println("Found version '"+currentVersion+"'")
	}

	var commitCount = "1"
	countBytes, err := exec.Command("git", "rev-list", "--count", currentVersion).Output()
	if err != nil {
		fmt.Println("Using default count of 1.")
	} else {
		count, _ := strconv.Atoi(strings.TrimSpace(string(countBytes)))
		commitCount = strconv.Itoa(count)
		fmt.Println("Found "+commitCount+" commits.")
	}

	return currentVersion + "-alpha" + commitCount
}
