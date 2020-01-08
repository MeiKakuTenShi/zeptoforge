package input

type input interface {
	IsKeyPressedImpl(int) bool
	IsMouseButtonPressedImpl(int) bool
	GetMousePositionImpl() (x, y float32)
	GetMouseXImpl() float32
	GetMouseYImpl() float32
}

type inputSingleton struct {
	input input
}

var (
	instance = &inputSingleton{}
)

func Singleton(in input) *inputSingleton {
	if instance.input == nil {
		instance = &inputSingleton{input: in}
	}
	return instance
}
func IsKeyPressed(keycode int) bool {
	return instance.input.IsKeyPressedImpl(keycode)
}
func IsMouseButtonPressed(button int) bool {
	return instance.input.IsMouseButtonPressedImpl(button)
}
func GetMousePosition() (x, y float32) {
	xPos, yPos := instance.input.GetMousePositionImpl()
	return xPos, yPos
}
func GetMouseX() float32 {
	return instance.input.GetMouseXImpl()
}
func GetMouseY() float32 {
	return instance.input.GetMouseYImpl()
}
