package syscmd

import "testing"

func TestRunSysCmd(t *testing.T) {
	t.Log(string(RunSysCmd("ls")))
	t.Log(string(RunBashCmd("mktemp -t")))
}
