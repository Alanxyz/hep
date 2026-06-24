package pdf

import (
	"fmt"
	"runtime"
	"os/exec"
)

func OpenFile(path string) error {
	if runtime.GOOS != "linux" {
		return fmt.Errorf("Not in Linux? Not my problem.")
	}
	cmd := exec.Command("xdg-open", path)
	return cmd.Start()
}
