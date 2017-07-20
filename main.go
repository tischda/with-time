package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const PROG_NAME string = "gonow"

var version string
var flag_help = flag.Bool("help", false, "displays this help message")
var flag_version = flag.Bool("version", false, "print version and exit")

func init() {
	flag.BoolVar(flag_help, "h", false, "")
	flag.BoolVar(flag_version, "v", false, "")
}

func main() {
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-h] [-v] command\n\nOPTIONS:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flag_version {
		fmt.Println("gonow version", version)
	} else {
		if *flag_help || flag.NArg() != 1 {
			flag.Usage()
			os.Exit(1)
		}
		process(flag.Arg(0))
	}
}

func process(command string) {
	cmd := exec.Command("cmd", "/c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("NOW=%s", time.Now().Format("20060102-150405")))
	cmd.Run()
}
