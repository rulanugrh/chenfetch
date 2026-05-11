package main

import (
	"fmt"
	"net"
	"strings"

	"os/exec"
	"os/user"

	user_shell "github.com/captainsafia/go-user-shell"
	"github.com/gilliek/go-xterm256/xterm256"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shirou/gopsutil/v3/host"
)

var asciiLines = []string{
	`⠀⠀⠀⠀⠀⠀⢀⣤⣶⣶⣤⡀⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⣴⠟⠉⠀⠀⠉⠻⣦⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⣼⠃⠀⢀⣤⣤⡀⠀⠘⣧⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⣿⠀⠀⣿⣿⣿⣿⠀⠀⣿⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⢿⡄⠀⠘⠿⠿⠃⠀⢠⡿⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠈⢿⣦⣄⠀⠀⣠⣴⡿⠁⠀⠀⠀⠀`,
	`⠀⠀⠀⠀⠀⠀⠉⠛⠿⠿⠛⠉⠀⠀⠀⠀⠀⠀`,
	`⠀⠀⠀⢀⣀⣀⣀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀`,
	`⠀⠀⢰⠟⠉⠀ HIKARI ⠀⠉⠻⡆⠀⠀`,
	`⠀⠀⠘⣧⣀⠀⠀⠀⠀⠀⠀⣀⣼⠃⠀⠀`,
	`⠀⠀⠀⠀⠉⠛⠶⠶⠶⠛⠉⠀⠀⠀⠀⠀`,
}

var (
	colorRosewater = xterm256.Color{ForegroundColor: 224, BackgroundColor: -1}
	colorFlamingo  = xterm256.Color{ForegroundColor: 217, BackgroundColor: -1}
	colorPink      = xterm256.Color{ForegroundColor: 212, BackgroundColor: -1}
	colorMauve     = xterm256.Color{ForegroundColor: 183, BackgroundColor: -1}
	colorRed       = xterm256.Color{ForegroundColor: 203, BackgroundColor: -1}
	colorPeach     = xterm256.Color{ForegroundColor: 216, BackgroundColor: -1}
	colorYellow    = xterm256.Color{ForegroundColor: 222, BackgroundColor: -1}
	colorGreen     = xterm256.Color{ForegroundColor: 114, BackgroundColor: -1}
	colorTeal      = xterm256.Color{ForegroundColor: 80, BackgroundColor: -1}
	colorSky       = xterm256.Color{ForegroundColor: 117, BackgroundColor: -1}
	colorSapphire  = xterm256.Color{ForegroundColor: 111, BackgroundColor: -1}
	colorBlue      = xterm256.Color{ForegroundColor: 75, BackgroundColor: -1}
	colorLavender  = xterm256.Color{ForegroundColor: 189, BackgroundColor: -1}
	colorSubtext   = xterm256.Color{ForegroundColor: 245, BackgroundColor: -1}
	colorOverlay   = xterm256.Color{ForegroundColor: 240, BackgroundColor: -1}
)

func c(col xterm256.Color, s string) string {
	return xterm256.Sprintf(col, s)
}

const boxW = 38

func boxTop() string {
	return "╭" + strings.Repeat("─", boxW) + "╮"
}

func boxBottom() string {
	return "╰" + strings.Repeat("─", boxW) + "╯"
}

func boxDiv() string {
	return "├" + strings.Repeat("─", boxW) + "┤"
}

func boxRow(ico, key, val string) string {
	return fmt.Sprintf(
		"│  %s  %s  %s",
		ico,
		key,
		val,
	)
}

func titleRow(s string) string {
	return "│  " + s
}

func localIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "unknown"
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 ||
			iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, _ := iface.Addrs()

		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			if ip4 := ip.To4(); ip4 != nil {
				return ip4.String()
			}
		}
	}

	return "unknown"
}

func main() {

	hostInfo, _ := host.Info()
	currentUser, _ := user.Current()

	uptimeRaw, _ := exec.Command("uptime", "-p").Output()
	termRaw, _ := exec.Command("/bin/sh", "-c", "echo $TERM").Output()

	uptimeStr := strings.TrimSpace(string(uptimeRaw))
	termStr := strings.TrimSpace(string(termRaw))
	shellStr := user_shell.GetUserShell()
	ipStr := localIP()
	getVersion := strings.Split(hostInfo.KernelVersion, "-")[0]
	title := fmt.Sprintf("%s: %s", hostInfo.Platform, getVersion)

	swatchColors := []xterm256.Color{
		colorRosewater,
		colorFlamingo,
		colorPink,
		colorMauve,
		colorBlue,
		colorSky,
		colorGreen,
		colorLavender,
	}

	var block string
	for _, col := range swatchColors {
		block += c(col, "██")
	}

	infoLines := []string{
		c(colorOverlay, boxTop()),

		titleRow(
			c(colorPink, "") +
				c(colorLavender,
					fmt.Sprintf("%s@%s",
						currentUser.Username,
						hostInfo.Hostname,
					),
				) +
				c(colorPink, ""),
		),

		c(colorOverlay, boxDiv()),

		boxRow(
			c(colorPeach, ""),
			c(colorFlamingo, "os      "),
			c(colorRosewater, hostInfo.OS),
		),

		boxRow(
			c(colorPeach, ""),
			c(colorFlamingo, "version "),
			c(colorRed, title),
		),

		boxRow(
			c(colorSky, ""),
			c(colorFlamingo, "term    "),
			c(colorSky, termStr),
		),

		boxRow(
			c(colorGreen, "󱑈"),
			c(colorFlamingo, "uptime  "),
			c(colorGreen, uptimeStr),
		),

		boxRow(
			c(colorBlue, ""),
			c(colorFlamingo, "shell   "),
			c(colorBlue, shellStr),
		),

		boxRow(
			c(colorMauve, ""),
			c(colorFlamingo, "ip      "),
			c(colorLavender, ipStr),
		),

		c(colorOverlay, boxDiv()),

		"│  " + block,

		c(colorOverlay, boxBottom()),
	}

	for len(asciiLines) < len(infoLines) {
		asciiLines = append(asciiLines, strings.Repeat(" ", 20))
	}

	for len(infoLines) < len(asciiLines) {
		infoLines = append(infoLines, "")
	}

	t := table.NewWriter()

	t.SetStyle(table.Style{
		Name: "catppuccin-fetch",

		Box: table.BoxStyle{
			BottomLeft:       " ",
			BottomRight:      " ",
			BottomSeparator:  " ",
			Left:             " ",
			LeftSeparator:    " ",
			MiddleHorizontal: " ",
			MiddleSeparator:  " ",
			MiddleVertical:   " ",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			Right:            " ",
			RightSeparator:   " ",
			TopLeft:          " ",
			TopRight:         " ",
			TopSeparator:     " ",
			UnfinishedRow:    " ",
		},

		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
	})

	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number:   1,
			Align:    text.AlignLeft,
			WidthMin: 22,
		},
		{
			Number: 2,
			Align:  text.AlignLeft,
		},
	})

	t.AppendRow(table.Row{"", ""})

	for i := range asciiLines {

		infoCol := ""

		if i < len(infoLines) {
			infoCol = infoLines[i]
		}

		t.AppendRow(table.Row{
			c(colorTeal, asciiLines[i]),
			infoCol,
		})
	}

	t.AppendRow(table.Row{"", ""})

	fmt.Println(t.Render())
}
