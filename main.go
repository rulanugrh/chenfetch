package main

import (
	"fmt"
	"strings"

	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

func main() {
	var s []string
	var sA []string
	s = append(s, ""+xterm256.Sprintf(xterm256.DarkYellow, "████")+"  "+xterm256.Sprintf(xterm256.DarkMagenta, "████")+"  "+xterm256.Sprintf(xterm256.Red, "████")+"  "+xterm256.Sprintf(xterm256.DarkBlue, "████")+"  "+xterm256.Sprintf(xterm256.DarkCyan, "████"))
	sA = append(sA, ""+xterm256.Sprintf(xterm256.Black, "████")+"  "+xterm256.Sprintf(xterm256.Magenta, "████")+"  "+xterm256.Sprintf(xterm256.LightGray, "████")+"  "+xterm256.Sprintf(xterm256.Blue, "████")+"  "+xterm256.Sprintf(xterm256.DarkGray, "████"))

	var block string
	for _, blk := range s {
		block = blk
	}

	var block2 string
	for _, blks := range sA {
		block2 = blks
	}

	data, _ := host.Info()
	infoOS := "🍔 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(data.OS))
	infoKernel := "🥙 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(data.KernelArch))
	infoHostname := "🍱 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(data.Hostname))
	infoPlatform := "🍚 " + xterm256.Sprintf(xterm256.Yellow, strings.Title(data.PlatformFamily))

	t := table.NewWriter()
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, "   _____")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, "  | ___ |")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, "  ||   ||  J.O"), infoOS})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, "  ||___||"), infoKernel})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, "  |   _ |"), infoHostname})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, "  |_____|"), infoPlatform})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, ` /_/_|_\_\----.`)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, `/_/__|__\_\   )`)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, `             (`), block})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.DarkGreen, `             []`), block2})

	fmt.Println(t.Render())
}
