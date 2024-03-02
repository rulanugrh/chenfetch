package main

import (
	"fmt"
	"strings"

	"os/exec"
	"os/user"

	user_shell "github.com/captainsafia/go-user-shell"
	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
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

	var cpuName string
	cpus, _ := cpu.Info()
	for _, cp := range cpus {
		cpuName = cp.ModelName
	}

	data, _ := host.Info()
	vm, _ := mem.VirtualMemory()
	currentUser, _ := user.Current()

	title := "  " + xterm256.Sprintf(xterm256.Red, "こんにちは世界")
	os := "  " + "  " + xterm256.Sprintf(xterm256.Magenta, data.OS)
	kernel := "  " + "  " + xterm256.Sprintf(xterm256.Magenta, data.KernelVersion)
	uptimeCmd, _ := exec.Command("uptime", "-p").Output()
	uptime := "  " + "  " + xterm256.Sprintf(xterm256.Magenta, strings.TrimSpace(string(uptimeCmd)))
	termCmd, _ := exec.Command("/bin/sh", "-c", "echo $TERM").Output()
	terminal := "  " + "  " + xterm256.Sprintf(xterm256.Magenta, strings.TrimSpace(string(termCmd)))
	shell := "  " + "  " + xterm256.Sprintf(xterm256.Magenta, user_shell.GetUserShell())
	user := "  " + xterm256.Sprintf(xterm256.Red, fmt.Sprintf("%s@%s", currentUser.Username, data.Hostname))
	cpu := "  " + "  " + xterm256.Sprintf(xterm256.Magenta, cpuName)
	memory := "  " + "﬙  " + xterm256.Sprintf(xterm256.Magenta, fmt.Sprintf("%dMB / %dMB (%.2f%%%%)", vm.Used, vm.Total, vm.UsedPercent))

	t := table.NewWriter()
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "             ⠀⠀⠤⣄⠀⠀⠀⠀⠀"), title})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⡇⠀⠀⠀⠀"), "┌───────────────────────────────────────────────┐"})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠀⠀⠀⢀⣤⣴⣶⣶⣶⣶⣶⣦⣤⡄⠊⠀⠀⠀⠀⠀"), os})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠦⣤⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀"), kernel})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠀⠀"), uptime})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀"), terminal})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡧⠀"), shell})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠉⢹⣿⣿⢓⣺⡿⠟⠛⠛⢻⣿⣼⣿⣿⣿⣿⣿⣿⣇⠀"), "└───────────────────────────────────────────────┘"})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⢸⣿⣿⠉⠀⠀⠀⠀⠀⠀⠀⠈⢀⣿⡿⠿⢿⣿⣿⣿⠀")})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⡀⠈⣿⣿⡀⠀⠀⠠⠄⠀⠀⠀⠀⢸⣿⣗⠏⢪⣿⣿⣿⡇"), user})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⢀⡈⣿⣷⣤⣄⣀⠀⠀⠀⢀⣤⣿⣿⣷⣾⣿⣿⣿⣿⣧"), "┌───────────────────────────────────────────────┐"})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠘⠉⢹⣿⣿⣿⣿⣿⡇⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿"), cpu})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠀⠀⣾⣿⣿⠟⡿⠟⡁⠄⠚⠉⠀⠘⢿⣿⣿⣿⣿⣿⡇"), memory})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠀⠀⣿⠟⠁⠈⠉⠁⠀⠀⠀⠀⠀⠀⠀⠙⡿⠿⠏⠿⠃"), "└───────────────────────────────────────────────┘"})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.Cyan, "⠀⠀⠀⡜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⡄⠀⠀⠀"), "			" + block})
	t.AppendRow(table.Row{xterm256.Sprintf(xterm256.LightGray, "")})

	t.SetStyle(table.StyleLight)
	t.Style().Title.Align = text.AlignCenter
	t.SetTitle(xterm256.Sprintf(xterm256.Cyan, "chenfetch"))
	fmt.Println(t.Render())
}
