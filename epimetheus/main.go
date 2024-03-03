package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jaypipes/ghw"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/net"
)

func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}
func ByteFormat(inputNum float64, precision int) string {

	if precision <= 0 {
		precision = 1
	}

	var unit string
	var returnVal float64

	if inputNum >= 1000000000 {
		returnVal = RoundUp((inputNum / 1073741824), precision)
		unit = " GB" // gigabyte
	} else if inputNum >= 1000000 {
		returnVal = RoundUp((inputNum / 1048576), precision)
		unit = " MB" // megabyte
	} else if inputNum >= 1000 {
		returnVal = RoundUp((inputNum / 1024), precision)
		unit = " KB" // kilobyte
	} else {
		returnVal = inputNum
		unit = " bytes" // byte
	}

	return strconv.FormatFloat(returnVal, 'f', precision, 64) + unit

}

func main() {

	var cpuName string
	cpus, _ := cpu.Info()
	for _, cp := range cpus {
		cpuName = fmt.Sprintf("%s", cp.ModelName)
	}

	bi, _ := ghw.BIOS()
	vendor := "  " + xterm256.Sprintf(xterm256.DarkCyan, "󰌢 ") + xterm256.Sprintf(xterm256.DarkCyan, bi.Vendor)

	// mem, _ := ghw.Memory()
	// fmt.Println(mem.String())

	bl, _ := ghw.Block()

	disk := "  " + xterm256.Sprintf(xterm256.DarkCyan, "󰨣 ") + xterm256.Sprintf(xterm256.DarkCyan, ByteFormat(float64(bl.TotalPhysicalBytes), 1)+" Available")

	var block string
	var colorList = []xterm256.Color{
		xterm256.Red,
		xterm256.Green,
		xterm256.Blue,
		xterm256.Cyan,
		xterm256.Yellow,
		xterm256.Magenta,
		xterm256.White,
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
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ""), xterm256.Sprintf(xterm256.DarkCyan, "┌─────────────") + xterm256.Sprintf(xterm256.Magenta, " E P I M E T H E U S ") + xterm256.Sprintf(xterm256.DarkCyan, "────────────┐")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.White, "    _ ___  _  "), infoOS})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Orange, "  .oooooooooo.  "), shell})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.White, " |"+xterm256.Sprintf(xterm256.Orange, "'oooooooooo'")+xterm256.Sprintf(xterm256.White, "|  ")), infoCPU})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.White, " |            |  "), infoIP})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.White, "  '.________.'   "), vendor})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ``), disk})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, `    `)})
	t.AppendRow(table.Row{})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ""), "  " + block + "   " + block})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, ` `)})

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
