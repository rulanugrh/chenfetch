package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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
  for _, nt := range ntw {
    fmt.Println(nt.Addrs)
  }
	data, _ := host.Info()
  host := fmt.Sprintf("%s@%s", strings.ToLower(os.Getenv("USERNAME")), data.Hostname)

	var infoOS string
	var shell string
  var user string
	if data.OS == "windows" {
		infoOS = "  " + xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, data.OS)
		ex := os.Getenv("POWERLINE_COMMAND")
		shell = "  " + xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, string(ex))
    user = "  " + xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, host)
 	} else {
		infoOS = "  " + xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, data.OS)
		ex, _ := exec.Command("/bin/sh", "-c", "echo $TERM").Output()
		shell = "  " + xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, string(ex))
	}
  
	// infoKernel := "  " + xterm256.Sprintf(xterm256.Yellow, "󰀄 ") + xterm256.Sprintf(xterm256.Cyan, data.OS)
  
	infoCPU := "  " + xterm256.Sprintf(xterm256.Yellow, "  ") + xterm256.Sprintf(xterm256.Cyan, cpuName)

	t := table.NewWriter()
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "          ,MMM8&&&. "), xterm256.Sprintf(xterm256.Cyan, "┌────────────") + xterm256.Sprintf(xterm256.Red, " Hardware Information ") + xterm256.Sprintf(xterm256.Cyan, "───────────┐")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "     _...MMMMM88&&&&..._ ")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "  .::'''MMMMM88&&&&&&'''::. "), user})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, " ::     MMMMM88&&&&&&     :: "), infoOS})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, " '::....MMMMM88&&&&&&....::' "), shell})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, "    `''''MMMMM88&&&&''''` "), infoCPU})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, `    jgs   'MMM8&&&' `)})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, ``), xterm256.Sprintf(xterm256.Cyan, "└─────────────────────────────────────────────┘")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Yellow, ` `)})

	t.SetStyle(table.StyleLight)
	t.Style().Title.Align = text.AlignCenter
	t.SetTitle(xterm256.Sprintf(xterm256.Cyan, "chenfetch"))

	fmt.Println(t.Render())
}
