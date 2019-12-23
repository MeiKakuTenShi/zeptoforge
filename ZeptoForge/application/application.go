package application

type App interface {
	Init()
	Run()
}

type Application struct {
	App  App
	Name string
}

func (Application) Run() {
	for x := 0; ; x++ {
		// fmt.Println(x % 2)
	}
}
