package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/net"
)

func main() {

	var cpuName string
	cpus, _ := cpu.Info()
	for _, cp := range cpus {
		cpuName = fmt.Sprintf("%s", cp.ModelName)
	}

	ntw, _ := net.Interfaces()
	var ips string
	for _, nt := range ntw {
		for _, d := range nt.Addrs {
			check := strings.Contains(d.Addr, "192")
			if check {
				ips = d.Addr
				break
			}
		}
	}

	data, _ := host.Info()

	fmt.Println(" ")
	var infoOS string
	var shell string
	if data.OS == "windows" {
		infoOS = "  " + xterm256.Sprintf(xterm256.DarkCyan, "  ") + xterm256.Sprintf(xterm256.DarkCyan, data.OS)
		ex := os.Getenv("POWERLINE_COMMAND")
		shell = "  " + xterm256.Sprintf(xterm256.DarkCyan, "  ") + xterm256.Sprintf(xterm256.DarkCyan, string(ex))
	} else {
		infoOS = "  " + xterm256.Sprintf(xterm256.DarkCyan, "  ") + xterm256.Sprintf(xterm256.DarkCyan, data.OS)
		ex, _ := exec.Command("/bin/sh", "-c", "echo $TERM").Output()
		shell = "  " + xterm256.Sprintf(xterm256.DarkCyan, "  ") + xterm256.Sprintf(xterm256.DarkCyan, string(ex))
	}

	infoCPU := "  " + xterm256.Sprintf(xterm256.DarkCyan, "  ") + xterm256.Sprintf(xterm256.DarkCyan, cpuName)
	infoIP := "  " + xterm256.Sprintf(xterm256.DarkCyan, "  ") + xterm256.Sprintf(xterm256.DarkCyan, ips)

	t := table.NewWriter()
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})

	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "          ,MMM8&&&. "), xterm256.Sprintf(xterm256.DarkCyan, "┌──────────────") + xterm256.Sprintf(xterm256.Red, " ハードウェア情報 ") + xterm256.Sprintf(xterm256.DarkCyan, "─────────────┐")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "     _...MMMMM88&&&&..._ "), ""})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "  .::'''MMMMM88&&&&&&'''::. "), infoOS})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, " ::     MMMMM88&&&&&&     :: "), shell})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, " '::....MMMMM88&&&&&&....::' "), infoCPU})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "    `''''MMMMM88&&&&''''` "), infoIP})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, `    jgs   'MMM8&&&' `)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, ``), xterm256.Sprintf(xterm256.DarkCyan, "└─────────────────────────────────────────────┘")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, ` `)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})

	t.SetStyle(table.Style{
		Name: "newStyle",
		Box: table.BoxStyle{
			BottomLeft: "",
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
	})

	fmt.Println(t.Render())
}
