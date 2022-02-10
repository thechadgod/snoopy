package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

func isStarted() (bool, error) {

	cmd := exec.Command("sudo", "torctl", "status")

	stdout, err := cmd.Output()

	if err != nil {
		return false, err
	}

	stringOutput := string(stdout)

	if strings.Contains(stringOutput, "torctl is started") {
		return true, nil
	} else if strings.Contains(stringOutput, "torctl is stopped") {
		return false, nil
	} else {
		return false, fmt.Errorf("torctl status returned unexpected output: %s", stringOutput)
	}

}

func StartTor() {

	cmd := exec.Command("sudo", "torctl", "start")

	stdout, err := cmd.Output()

	if err != nil {

		hasStarted, err := isStarted()

		if err != nil {
			panic(err)
		}

		if hasStarted {

			fmt.Println("Tor is already running")
			return

		}
	}

	stringStdout := string(stdout)

	if strings.Contains(stringStdout, "ERROR") {

		println(stringStdout)
		return
	} else {
		fmt.Println(string("Started Tor"))
		return
	}

}

func StopTor() {

	cmd := exec.Command("sudo", "torctl", "stop")

	_, err := cmd.Output()

	if err != nil {

		hasStarted, err := isStarted()

		if err != nil {
			panic(err)
		}

		if !hasStarted {

			fmt.Println("Tor is not running")
			return

		}
	}

	fmt.Println(string("Stopped Tor"))

}

func StatusTor() {

	hasStarted, err := isStarted()

	if err != nil {
		panic(err)
	}

	if hasStarted {
		fmt.Println("Tor is running")
	} else {
		fmt.Println("Tor is not running")
	}
}

func ChangeID() {

	hasStarted, err := isStarted()

	if err != nil {
		panic(err)
	}

	if hasStarted {

		cmd := exec.Command("sudo", "torctl", "chngid")

		stdout, err := cmd.Output()

		if err != nil {

			fmt.Println("An error occured:")
			fmt.Println(string(stdout))
			return

		}

		if strings.Contains(string(stdout), "tor identity changed") {

			fmt.Println("Changed your tor identity")
			return

		} else {

			fmt.Println(string(stdout))
			return

		}

	} else {

		fmt.Println("Tor is not running")

	}

}

func ChangeMac() {

	hasStarted, err := isStarted()

	if err != nil {
		panic(err)
	}

	if hasStarted {

		cmd := exec.Command("sudo", "torctl", "chngmac")

		stdout, err := cmd.Output()

		if err != nil {

			fmt.Println("An error occured:")
			fmt.Println(string(stdout))
			return

		}

		if strings.Contains(string(stdout), "changed mac addresses") {

			fmt.Println("Changed your mac addresses")
			return

		} else {

			fmt.Println(string(stdout))
			return

		}

	} else {

		fmt.Println("Tor is not running")

	}

}
