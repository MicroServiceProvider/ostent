// +build linux

package ostent

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getDistrib() string {
	// https://unix.stackexchange.com/q/35183
	std, err := exec.Command("lsb_release", "-i", "-r").CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "lsb_release: %s\n", err)
		return ""
	}
	id, release := "", ""
	// strings.TrimRight(string(std), "\n\t ")
	for _, s := range strings.Split(string(std), "\n") {
		b := strings.Split(s, "\t")
		if len(b) == 2 {
			if b[0] == "Distributor ID:" {
				id = b[1]
				continue
			} else if b[0] == "Release:" {
				release = b[1]
				continue
			}
		}
	}
	if id == "" {
		fmt.Fprintf(os.Stderr, "Could not identify Distributor ID")
	}
	if release == "" {
		fmt.Fprintf(os.Stderr, "Could not identify Release")
	}
	if id == "" || release == "" {
		return ""
	}
	return id + " " + release
}

func procname(_ int, proc_name string) string {
	return proc_name // from /proc/_/stat, may be shortened
}