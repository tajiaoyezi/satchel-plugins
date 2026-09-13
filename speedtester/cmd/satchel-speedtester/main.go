// satchel-speedtester 是家用测速端的入口，只做装配。M0 只有 --version，测速端本体从 M7 起填。
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/satchel/satchel-plugins/speedtester/internal/buildinfo"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr *os.File) int {
	fs := flag.NewFlagSet("satchel-speedtester", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {} // 用法由下面按情形打印：--help 打到 stdout，敲错时 flag 已把错误写到 stderr
	version := fs.Bool("version", false, "打印版本、commit 与构建时间后退出")
	if err := fs.Parse(args); err != nil {
		// 人主动要帮助不是用法错误：用法打到 stdout、退出码 0，与主控 satchel --help 一致；敲错了才是 2。
		if errors.Is(err, flag.ErrHelp) {
			fmt.Fprintln(stdout, "用法：satchel-speedtester [--version]")
			fs.SetOutput(stdout)
			fs.PrintDefaults()
			return 0
		}
		fmt.Fprintln(stderr, "用法：satchel-speedtester [--version]")
		return 2
	}
	if *version {
		info := buildinfo.Get()
		fmt.Fprintf(stdout, "satchel-speedtester %s (commit %s, built %s)\n", info.Version, info.Commit, info.Date)
		return 0
	}
	fmt.Fprintln(stderr, "satchel-speedtester 还没有可运行的功能（M0 骨架）；用 --version 查看版本")
	return 1
}
