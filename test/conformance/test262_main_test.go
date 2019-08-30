package conformance

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

const (
	test262folder = "test262"
	test262repo   = "https://github.com/tc39/test262"
)

func TestMain(m *testing.M) {
	setup()
	defer tearDown()

	os.Exit(m.Run())
}

func setup() {
	cloneTest262Repo()
}

func tearDown() {}

func cloneTest262Repo() {
	flag.Parse()

	if testing.Short() {
		// don't clone if short testing
		return
	}

	if _, err := os.Stat(test262folder); os.IsNotExist(err) {
		log.Println("Conformance test directory does not exist, cloning it...")

		var stdout, stderr bytes.Buffer

		cmd := exec.Command("git", "clone", test262repo, test262folder)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			panic(fmt.Sprintf("Clone failed.\nStdout: '%v'\nStderr: '%v'", stdout.String(), stderr.String()))
		}
	}
}
