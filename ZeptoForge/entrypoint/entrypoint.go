package main

import (
	"github.com/MeiKakuTenShi/zeptoforge/Sandbox/sandbox"
)

func main() {
	sb := sandbox.CreateApplication()
	sb.App.Init()
	sb.Run()
}
