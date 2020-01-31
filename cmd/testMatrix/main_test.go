package main_test

import (
	"bytes"
	"os/exec"
	"testing"
)

var cmdPath = "github.com/iuryfukuda/ibcc/cmd/testMatrix"

func TestIsTriangular(t *testing.T) {
	cmd := exec.Command("go", "run", cmdPath, "-matrix", "1,2\n0,0", "-t", "-v")
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			t.Logf("exit code: [%d]\n%s", exitError.ExitCode(), err)
		} else {
			t.Logf("another err: %s", err)
		}
	}
	t.Logf("stdout:\n%q\nstderr:\n%q", out.String(), outErr.String())
	t.Log(outErr.String())
}
