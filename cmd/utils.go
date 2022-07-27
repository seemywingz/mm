package cmd

import (
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

func sleepSeconds(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}

func randRC() {
	cliclick("rc", randCoords())
	sleepSeconds(3)
	cliclick("c", "-10,+0")
}

func randCoords() string {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(xmax-xmin) + xmin
	y := rand.Intn(ymax-ymin) + ymin
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func cliclick(command, coords string) error {
	execCMD := exec.Command("/usr/local/bin/cliclick", command+":"+coords)
	if err := execCMD.Run(); err != nil {
		return err
	}
	return nil
}
