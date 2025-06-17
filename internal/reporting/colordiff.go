package reporting

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

func ColorizeDiff(diff string) string {
	coloredDiff := &strings.Builder{}
	lines := strings.Split(diff, "\n")
	colorBold := color.New(color.Bold)
	colorBold.EnableColor()
	colorCyan := color.New(color.FgCyan)
	colorCyan.EnableColor()
	colorRed := color.New(color.FgRed)
	colorRed.EnableColor()
	colorGreen := color.New(color.FgGreen)
	colorGreen.EnableColor()

	for i, line := range lines {
		if strings.HasPrefix(line, "---") || strings.HasPrefix(line, "+++") {
			colorBold.Fprint(coloredDiff, line)
		} else if strings.HasPrefix(line, "@@") {
			colorCyan.Fprint(coloredDiff, line)
		} else if strings.HasPrefix(line, "-") {
			colorRed.Fprint(coloredDiff, line)
		} else if strings.HasPrefix(line, "+") {
			colorGreen.Fprint(coloredDiff, line)
		} else {
			coloredDiff.WriteString(line)
		}
		if i < len(lines)-1 { // Avoid adding an extra newline at the very end
			coloredDiff.WriteString("\n")
		}
	}
	return coloredDiff.String()
}

func ShouldUseColorDiff(colorMode string, paramName string) (bool, error) {
	switch colorMode {
	case "never":
		return false, nil
	case "auto":
		return os.Getenv("TERM") != "dumb" &&
				(isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())),
			nil
	case "always":
		return true, nil
	default:
		return false, fmt.Errorf("--%v must be 'always', 'auto' or 'never', not '%v'", paramName, colorMode)
	}
}
