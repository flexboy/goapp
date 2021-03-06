package config

import (
	"flag"
	"fmt"
	"github.com/urfave/cli"
)

func glogFlagShim(fakeVals map[string]string) {
	flag.VisitAll(func(fl *flag.Flag) {
		if val, ok := fakeVals[fl.Name]; ok {
			fl.Value.Set(val)
		}
	})
}

func glogGangstaShim(c *cli.Context) {
	_ = flag.CommandLine.Parse([]string{})
	glogFlagShim(map[string]string{
		"v":                fmt.Sprint(c.Int("v")),
		"logtostderr":      fmt.Sprint(c.Bool("logtostderr")),
		"stderrthreshold":  fmt.Sprint(c.Int("stderrthreshold")),
		"alsologtostderr":  fmt.Sprint(c.Bool("alsologtostderr")),
		"vmodule":          c.String("vmodule"),
		"log_dir":          c.String("log_dir"),
		"log_backtrace_at": c.String("log_backtrace_at"),
	})
}

var GlogFlags = []cli.Flag{
	cli.IntFlag{
		Name: "v", Value: 0, Usage: "log level for V logs",
	},
	cli.BoolFlag{
		Name: "logtostderr", Usage: "log to standard error instead of files",
	},
	cli.IntFlag{
		Name:  "stderrthreshold",
		Usage: "logs at or above this threshold go to stderr",
	},
	cli.BoolFlag{
		Name: "alsologtostderr", Usage: "log to standard error as well as files",
	},
	cli.StringFlag{
		Name:  "vmodule",
		Usage: "comma-separated list of pattern=N settings for file-filtered logging",
	},
	cli.StringFlag{
		Name: "log_dir", Usage: "If non-empty, write log files in this directory",
	},
	cli.StringFlag{
		Name:  "log_backtrace_at",
		Usage: "when logging hits line file:N, emit a stack trace",
		Value: ":0",
	},
}
