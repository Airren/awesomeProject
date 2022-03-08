package awesomeProject

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"testing"
	"time"
)

func TestRun(t *testing.T) {

	cmd := exec.Command("./sleep.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("test failed: %v", err)
	}
}

func TestDate(t *testing.T) {
	fmt.Println(time.Now().String()[:19])
}

func TestExecDaemon(t *testing.T) {

	logFile, _ := os.Create(fmt.Sprintf("./ipsec-%v.log", time.Now().String()[:19]))
	var err error
	cmd := exec.Command("./sleep.sh")
	cmd.Stdout = logFile
	cmd.Stderr = os.Stderr

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(cmd *exec.Cmd) {
		err = cmd.Start()
		if err != nil {
			log.Fatalf("command: %v start failed: %v ", cmd.Path, err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Fatalf("command: %v run failed: %v", cmd.Path, err)
		}
		wg.Done()
	}(cmd)

	t.Logf("wait the goroutine")
	wg.Wait()
	t.Logf("test finished")

	if err != nil {
		t.Fatal("run failed")
	}
}
