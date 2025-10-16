package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` "COMMAND"

Injects a %TIME:format% environment variable in a command.
COMMAND must be quoted and contain a %TIME:format% substring.

The format can be a golang Time.Format layout (e.g. 20060102) or a
YYYYMMDD style format. If format is empty, time.UnixDate is used.
(e.g. "Mon Jan _2 15:04:05 MST 2006")

  YYYY -> 2006 (Year)
  MM   -> 01   (Month)
  DD   -> 02   (Day)
  HH   -> 15   (Hour)
  mm   -> 04   (Minute)
  ss   -> 05   (Second)
  ms   -> 000  (Millisecond)

OPTIONS:

  -?, --help
          display this help message
  -v, --version
          print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` "echo Time is %TIME:%"
  Time is Thu Oct 16 17:42:12 CEST 2025`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` "echo Time is %TIME:YYYYMMDD-HHmmss%"
  Time is 20251016-174217`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help || flag.NArg() == 0 {
		flag.Usage()
		return
	}
	process(flag.Arg(0))
}

var layoutReplacer = strings.NewReplacer(
	"YYYY", "2006",
	"MM", "01",
	"DD", "02",
	"HH", "15",
	"mm", "04",
	"ss", "05",
	"ms", "000",
)

func process(command string) {
	re := regexp.MustCompile(`%TIME:(.*?)%`)
	now := time.Now()

	processedCommand := re.ReplaceAllStringFunc(command, func(placeholder string) string {
		// Extract format from placeholder, e.g., "%TIME:2006%" -> "2006"
		submatches := re.FindStringSubmatch(placeholder)
		if len(submatches) > 1 {
			format := submatches[1]
			if format == "" {
				return now.Format(time.UnixDate)
			}
			goLayout := layoutReplacer.Replace(format)
			return now.Format(goLayout)
		}
		// Should not happen with the given regex, but as a fallback:
		return placeholder
	})

	cmd := exec.Command("cmd", "/c", processedCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
