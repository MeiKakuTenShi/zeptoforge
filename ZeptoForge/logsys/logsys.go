package logsys

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

const (
	Lcore = 1 << iota
	Lclient
)

var (
	core, client       bytes.Buffer
	coreLog, clientLog *log.Logger

	// CORE log macros
	ZF_CORE_ERROR = func(msg ...interface{}) {
		coreLog.Panicf("CORE_ERROR: %s", msg)
		fmt.Print(&core)
	}
	ZF_CORE_WARN = func(msg ...interface{}) {
		coreLog.Printf("CORE_WARNING: %s", msg)
		fmt.Print(&core)
	}
	ZF_CORE_INFO = func(msg ...interface{}) {
		coreLog.Printf("CORE_INFO: %s", msg)
		fmt.Print(&core)
	}
	ZF_CORE_TRACE = func(msg ...interface{}) {
		coreLog.Printf("CORE_TRACE: %s", msg)
		fmt.Print(&core)
	}
	ZF_CORE_FATAL = func(msg ...interface{}) {
		coreLog.Fatalf("CORE_FATAL: %s", msg)
		fmt.Print(&core)
	}

	// CLIENT log macros
	ZF_ERROR = func(msg ...interface{}) {
		clientLog.Panicf("CLIENT_ERROR: %s", msg)
		fmt.Print(&client)
	}
	ZF_WARN = func(msg ...interface{}) {
		clientLog.Printf("CLIENT_WARNING: %s", msg)
		fmt.Print(&client)
	}
	ZF_INFO = func(msg ...interface{}) {
		clientLog.Printf("CLIENT_INFO: %s", msg)
		fmt.Print(&client)
	}
	ZF_TRACE = func(msg ...interface{}) {

	}
	ZF_FATAL = func(msg ...interface{}) {
		clientLog.Fatalf("CLIENT_FATAL: %s", msg)
		fmt.Print(&client)
	}
)

func Init() {
	core = bytes.Buffer{}
	client = bytes.Buffer{}
	coreLog = log.New(&core, "Z_FORGE: ", log.Ldate|log.Ltime|log.Lshortfile)
	clientLog = log.New(&client, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func PrintMessage(rec int, message string) {
	if rec == Lcore {
		coreLog.Output(2, message)
		fmt.Println(&core)
	} else if rec == Lclient {
		clientLog.Output(2, message)
		fmt.Println(&client)
	} else {
		fmt.Printf("LogSystem::PrintMessage(): 'rec' = %v - value undefined", rec)
	}
}

func GetCoreLog() (*log.Logger, error) {
	if coreLog != nil {
		return coreLog, nil
	}
	return nil, errors.New("coreLog does not exist, check if logsys has been initialized; 'clue': Init()")
}

func GetClientLog() (*log.Logger, error) {
	if clientLog != nil {
		return clientLog, nil
	}
	return nil, errors.New("clientLog does not exist, check if logsys has been initialized; 'clue': Init()")
}
