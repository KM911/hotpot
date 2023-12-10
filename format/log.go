package format

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
)

// Current Log Level
// Error Warning Info Note Success

var ()

func FileLogger(src string) {
	logFile, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}
func Error(msg string) {
	ErrorMessage("Error", msg)
}
func LogErorr(msg string) {
	log.Println(color.BgRed.Render("[Error]"), ": ", color.Error.Render(msg))
}

func ErrorMessage(error string, msg string) {
	fmt.Println(color.BgRed.Render(error), ": ", color.Error.Render(msg))
}

func Warning(msg string) {
	WarningMessage("Warning", msg)
}
func LogWarning(msg string) {
	log.Println(color.BgYellow.Render("[Warning]"), ": ", color.Warn.Render(msg))
}

func WarningMessage(warning string, msg string) {
	fmt.Println(color.BgYellow.Render(warning), ": ", color.Warn.Render(msg))
}

func Info(msg string) {
	InfoMessage("Info", msg)
}

func LogInfo(msg string) {
	log.Println(color.BgBlue.Render("[Info]"), ": ", color.Info.Render(msg))
}

func InfoMessage(info string, msg string) {
	fmt.Println(color.BgBlue.Render(info), ": ", color.Info.Render(msg))
}

func Note(msg string) {
	NoteMessage("Note", msg)
}
func LogNote(msg string) {
	log.Println(color.BgHiBlue.Render("[Note]"), ": ", color.Note.Render(msg))
}

func NoteMessage(note string, msg string) {
	fmt.Println(color.BgHiBlue.Render(note), ": ", color.Note.Render(msg))
}

func Success(msg string) {
	SuccessMessage("Success", msg)
}

func LogSuccess(msg string) {
	log.Println(color.BgGreen.Render("[Success]"), ": ", color.Success.Render(msg))
}

func SuccessMessage(success string, msg string) {
	fmt.Println(color.BgGreen.Render(success), ": ", color.Success.Render(msg))
}

func LogExample() {
	LogErorr("error")
	LogWarning("warning")
	LogInfo("info")
	LogNote("note")
	LogSuccess("success")
}

func Example() {
	Error("error")
	Warning("warning")
	Info("info")
	Note("note")
	Success("success")
}
