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

	var block string
	var colorList = []xterm256.Color{
		xterm256.Red,
		xterm256.Green,
		xterm256.Blue,
		xterm256.Cyan,
		xterm256.Yellow,
		xterm256.Magenta,
		xterm256.White,
		xterm256.LightGray,
		xterm256.Black,
	}

	for _, color := range colorList {
		block += xterm256.Sprintf(color, "  ")
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

	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ""), xterm256.Sprintf(xterm256.DarkCyan, "┌──────────────") + xterm256.Sprintf(xterm256.Red, " ハードウェア情報 ") + xterm256.Sprintf(xterm256.DarkCyan, "─────────────┐")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGray, " _ ___  _ "), ""})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Orange, " .oooooooooo. "), infoOS})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGray, " |"+xterm256.Sprintf(xterm256.Orange, "'oooooooooo'")+xterm256.Sprintf(xterm256.DarkGray, "| ")), shell})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGray, " |            | "), infoCPU})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGray, "  '.________.' "), infoIP})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ``)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ``), xterm256.Sprintf(xterm256.DarkCyan, "└─────────────────────────────────────────────┘")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ` `)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ""), block})

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
