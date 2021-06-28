package execbench

import (
	"fmt"
	"os/exec"
)

const (
	YES = "go"
	YESPATH = "/usr/local/bin/go"
)

func Yes(lookPath bool) error {

	y := YES
	if !lookPath {
		y = YESPATH
	}

	cmd := exec.Command(y, "version")

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%#v: output: %s", cmd.Args, out)
	}

	return nil
}


