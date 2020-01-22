package zforge

import (
	"bytes"
	"fmt"
	"log"

	"github.com/fatih/color"
)

const (
	Lcore = 1 << iota
	Lclient
)

var (
	core, client       bytes.Buffer
	coreLog, clientLog *log.Logger

	basicText = color.New(color.FgBlack, color.BgWhite).SprintFunc()
	errorText = color.New(color.FgMagenta).SprintFunc()
	warnText  = color.New(color.FgYellow).SprintFunc()
	infoText  = color.New(color.FgCyan).SprintFunc()
	traceText = color.New(color.FgHiBlue).SprintFunc()
	fatalText = color.New(color.FgRed).SprintFunc()

	// CORE log macros
	ZF_CORE_ERROR = func(msg ...interface{}) {
		coreLog.Panic(fmt.Sprint(errorText("CORE_ERROR"), ": ", basicText(msg)))
		line, _ := core.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_CORE_WARN = func(msg ...interface{}) {
		coreLog.Print(fmt.Sprint(warnText("CORE_WARNING"), ": ", basicText(msg)))
		line, _ := core.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_CORE_INFO = func(msg ...interface{}) {
		coreLog.Print(fmt.Sprint(infoText("CORE_INFO"), ": ", basicText(msg)))
		// fmt.Print(&core)
		line, _ := core.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_CORE_TRACE = func(msg ...interface{}) {
		coreLog.Print(fmt.Sprint(traceText("CORE_TRACE"), ": ", basicText(msg)))
		line, _ := core.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_CORE_FATAL = func(msg ...interface{}) {
		coreLog.Fatal(fmt.Sprint(fatalText("CORE_FATAL"), ": ", basicText(msg)))
		line, _ := core.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}

	// CLIENT log macros
	ZF_ERROR = func(msg ...interface{}) {
		clientLog.Panic(fmt.Sprint(errorText("CLIENT_ERROR"), ": ", basicText(msg)))
		line, _ := client.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_WARN = func(msg ...interface{}) {
		clientLog.Print(fmt.Sprint(warnText("CLIENT_WARN"), ": ", basicText(msg)))
		line, _ := client.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_INFO = func(msg ...interface{}) {
		clientLog.Print(fmt.Sprint(infoText("CLIENT_INFO"), ": ", basicText(msg)))
		line, _ := client.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_TRACE = func(msg ...interface{}) {
		clientLog.Print(fmt.Sprint(traceText("CLIENT_TRACE"), ": ", basicText(msg)))
		line, _ := client.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
	ZF_FATAL = func(msg ...interface{}) {
		clientLog.Fatal(fmt.Sprint(fatalText("CLIENT_FATAL"), ": ", basicText(msg)))
		line, _ := client.ReadString([]byte("\n")[0])
		fmt.Print(line)
	}
)

func initLogSys() {
	core = bytes.Buffer{}
	client = bytes.Buffer{}

	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	coreLog = log.New(&core, fmt.Sprint(blue("ZF_CORE"), ": "), log.Ldate|log.Ltime|log.Lshortfile)
	clientLog = log.New(&client, fmt.Sprint(green("APP"), ": "), log.Ldate|log.Ltime|log.Lshortfile)
}

func PrintLog(rec int) {
	if rec == Lcore {
		fmt.Print(&core)
	} else if rec == Lclient {
		fmt.Print(&client)
	} else {
		fmt.Printf("LogSystem::PrintMessage(): 'rec' = %v - value undefined", rec)
	}
}
