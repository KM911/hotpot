package format

import (
	"strings"

	"github.com/gookit/color"
)

// ╭───────┬─────╮
// │ msg   │ str │
// ╰───────┴─────╯

func BoxMessage(msg string, str string, l1 int, l2 int) {
	FormatStringBuilder.WriteString("╭")
	FormatStringBuilder.WriteString(strings.Repeat("─", l1+2))
	FormatStringBuilder.WriteString("┬")
	FormatStringBuilder.WriteString(strings.Repeat("─", l2+2))
	FormatStringBuilder.WriteString("╮\n│ ")
	FormatStringBuilder.WriteString(msg)
	FormatStringBuilder.WriteString(" │ ")
	FormatStringBuilder.WriteString(str)
	FormatStringBuilder.WriteString(" │\n╰")
	FormatStringBuilder.WriteString(strings.Repeat("─", l1+2))
	FormatStringBuilder.WriteString("┴")
	FormatStringBuilder.WriteString(strings.Repeat("─", l2+2))
	FormatStringBuilder.WriteString("╯\n")
	println(FormatStringBuilder.String())
	FormatStringBuilder.Reset()
}

func BoxError(msg string, str string) {

	BoxMessage(color.Error.Render(msg), str, len(msg), len(str))
}

func BoxWarning(msg string, str string) {
	BoxMessage(color.Warn.Render(msg), str, len(msg), len(str))
}

func BoxInfo(msg string, str string) {
	BoxMessage(color.Info.Render(msg), str, len(msg), len(str))
}

func BoxSuccess(msg string, str string) {
	BoxMessage(color.Success.Render(msg), str, len(msg), len(str))
}

func BoxNote(msg string, str string) {
	BoxMessage(color.Note.Render(msg), str, len(msg), len(str))
}

func BoxExample() {
	BoxError("error", "error")
	BoxWarning("warning", "warning")
	BoxInfo("info", "info")
	BoxSuccess("success", "success")
	BoxNote("note", "note")
}
