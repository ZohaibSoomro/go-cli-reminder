package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	// "time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func main() {
	cmdEnvKey := "GO_REMINDER"
	cmdEnvValue := "1"
	if len(os.Args) < 3 {
		fmt.Printf("Usage:%s <hh:mm> <message>\n", os.Args[0])
		os.Exit(2)
	}
	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)
	cmd := exec.Command("prog", os.Args[2])
	parsedTime, err := w.Parse(os.Args[1], now)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	diff := parsedTime.Time.Sub(now)
	//if not a future tme set then close app
	if now.After(parsedTime.Time) {
		os.Exit(1)
	}
	if os.Getenv(cmdEnvKey) == cmdEnvValue {
		time.Sleep(diff)
		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
		cmd.Run()
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", cmdEnvKey, cmdEnvValue))
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Reminder displaying in", diff.Round(time.Second))
	}
	// fmt.Println(parsedTime.Time)
}
