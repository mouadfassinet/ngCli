package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func formatExceptFlag(rawExceptRepositories string) []string {
	return strings.Split(rawExceptRepositories, "-")
}

func main() {
	futurBranchName := flag.String("branch", "", "Futur branch name")
	rawExceptRepositories := flag.String("except", "", "Except repositores")
	resetRepositories := flag.Bool("reset", false, "Reset repositores")
	fetchRepositories := flag.Bool("fetch", false, "Fetch repositores")
	pullRepositories := flag.Bool("pull", false, "Pull origin repositores")
	flag.Parse()

	exceptRepositories := formatExceptFlag(*rawExceptRepositories)
	exceptRepositories = append(exceptRepositories, "", "", "", "", "")
	directoriesPath := map[string]string{
		"intranet":     "./",
		"production":   "./node_modules/iad-ng-production",
		"people":       "./node_modules/iad-ng-people",
		"clubBusiness": "./node_modules/iad-ng-club-business",
		"dashboard":    "./node_modules/iad-ng-dashboard",
		"core":         "./node_modules/iad-ng-core",
		"css":          "./node_modules/iad-ng-css",
	}

	var filterDirectoriesPath []string
	for key, value := range directoriesPath {
		if (key != exceptRepositories[0]) &&
			(key != exceptRepositories[1]) &&
			(key != exceptRepositories[2]) &&
			(key != exceptRepositories[3]) &&
			(key != exceptRepositories[4]) &&
			(key != exceptRepositories[5]) {
			filterDirectoriesPath = append(filterDirectoriesPath, value)
		}
	}

	for _, element := range filterDirectoriesPath {
		if *resetRepositories {
			cmd := exec.Command("git", "reset --hard")
			cmd.Run()
		}

		if *fetchRepositories {
			cmd := exec.Command("git", "fetch")
			cmd.Run()
		}

		if *pullRepositories {
			cmd := exec.Command("git", "pull origin")
			cmd.Run()
		}
		cmd := exec.Command("git", "checkout", *futurBranchName)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Dir = element
		cmd.Run()
		fmt.Print(out.String())
	}
}
