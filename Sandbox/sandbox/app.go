package sandbox

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/application"
)

type Sandbox struct {
}

func (sb *Sandbox) Init() {

}

func (sb *Sandbox) Run() {
	for x := 0; ; x++ {
		fmt.Println(x % 3)
	}
}

func CreateApplication() *application.Application {
	result := &application.Application{App: &Sandbox{}, Name: "Sandbox"}
	return result
}
