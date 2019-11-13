package jim

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Help(t *Task) (string, error) {
	bb := &bytes.Buffer{}

	c := exec.Command("go", "doc", fmt.Sprintf("%s.%s", t.Pkg, t.Name))
	c.Stdout = bb
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(bb.String()), nil
}
