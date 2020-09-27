package monitor

import (
	"fmt"
	"os"
	"os/exec"
)

//ExecCmd TODO
func ExecCmd() string {
	out, err := exec.Command("vcgencmd", "measure_temp").CombinedOutput()
	fmt.Print(string(out))

	ret := string(out)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return ret
}
