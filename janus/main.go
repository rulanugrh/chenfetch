package main

import (
	"fmt"
	"strings"

	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

func main() {

	var cpuName string
	cpus, _ := cpu.Info()
	for _, cp := range cpus {
		cpuName = cp.ModelName
	}
	
	data, _ := host.Info()
  var infoOS string
  if data.OS == "windows" {
    infoOS = "  " +xterm256.Sprintf(xterm256.Yellow, "  ")+ xterm256.Sprintf(xterm256.Cyan, strings.Title("windows"))
  } else{
    infoOS = "  "+xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, strings.Title(data.OS))
  }

	infoKernel := "  "+xterm256.Sprintf(xterm256.Yellow, "󰀄 ") + xterm256.Sprintf(xterm256.Cyan, strings.Title(data.Hostname))
	infoPlatform := "  " + "🦪 " + xterm256.Sprintf(xterm256.Cyan, strings.Title(data.Platform))
	infoCPU := "  " + xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, strings.Title(cpuName))

	t := table.NewWriter()
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "          ,MMM8&&&. "), xterm256.Sprintf(xterm256.Cyan, "┌─────────────")+xterm256.Sprintf(xterm256.Red, " Hardware Information ")+ xterm256.Sprintf(xterm256.Cyan, "────────────┐") })
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "     _...MMMMM88&&&&..._ ")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "  .::'''MMMMM88&&&&&&'''::. "), infoOS})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, " ::     MMMMM88&&&&&&     :: "), infoKernel})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, " '::....MMMMM88&&&&&&....::' "), infoPlatform})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "    `''''MMMMM88&&&&''''` "), infoCPU})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, `    jgs   'MMM8&&&' `)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, ``), xterm256.Sprintf(xterm256.Cyan, "└───────────────────────────────────────────────┘")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, ` `)})

  t.SetStyle(table.StyleLight)
  t.Style().Title.Align = text.AlignCenter
  t.SetTitle(xterm256.Sprintf(xterm256.Cyan, "chenfetch"))

	fmt.Println(t.Render())
}
