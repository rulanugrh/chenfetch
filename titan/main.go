package main

import (
	"fmt"
	"strings"

	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

func main() {
	var s []string
	var sA []string
	s = append(s, "  "+xterm256.Sprintf(xterm256.Black, "████")+""+xterm256.Sprintf(xterm256.DarkRed, "████")+""+xterm256.Sprintf(xterm256.DarkGreen, "████")+""+xterm256.Sprintf(xterm256.DarkBlue, "████")+""+xterm256.Sprintf(xterm256.DarkCyan, "████")+""+xterm256.Sprintf(xterm256.DarkYellow, "████")+""+xterm256.Sprintf(xterm256.DarkMagenta, "████")+""+xterm256.Sprintf(xterm256.DarkGray, "████"))
	sA = append(sA, "  "+xterm256.Sprintf(xterm256.LightGray, "████")+""+xterm256.Sprintf(xterm256.Red, "████")+""+xterm256.Sprintf(xterm256.Green, "████")+""+xterm256.Sprintf(xterm256.Blue, "████")+""+xterm256.Sprintf(xterm256.Cyan, "████")+""+xterm256.Sprintf(xterm256.Yellow, "████")+""+xterm256.Sprintf(xterm256.Magenta, "████")+""+xterm256.Sprintf(xterm256.LightGray, "████"))

	var block string
	for _, blk := range s {
		block = blk
	}

	var block2 string
	for _, blks := range sA {
		block2 = blks
	}

	var cpuName string
	cpus, _ := cpu.Info()
	for _, cp := range cpus {
		cpuName = cp.ModelName
	}

	data, _ := host.Info()
	Host := "  " + xterm256.Sprintf(xterm256.Red, "Hello, Everyone 👋")

	infoOS := "  " + "🍔 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(data.KernelArch))
	infoKernel := "  " + "🥙 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(data.PlatformFamily))
	infoPlatform := "  " + "🥪 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(data.Platform))
	infoCPU := "  " + "🥗 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(cpuName))

	t := table.NewWriter()
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, "     _____ "), Host})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, "    | ___ | ")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, "    ||   ||  J.O "), infoOS})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, "    ||___|| "), infoKernel})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, "    |   _ | "), infoPlatform})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, `    |_____| `), infoCPU})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, `   /_/_|_\_\----. `)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, `  /_/__|__\_\   ) `), block})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, `               ( `), block2})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Blue, `               [] `)})

	fmt.Println(t.Render())
}
