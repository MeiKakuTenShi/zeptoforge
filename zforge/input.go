package zforge

var (
	inputInstance *inputSingleton
)

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

func initInputSingleton(in input) *inputSingleton {
	if inputInstance == nil {
		x := new(inputSingleton)
		x.input = in

		inputInstance = x
	}
	return inputInstance
}
func IsKeyPressed(keycode int) bool {
	return inputInstance.input.IsKeyPressedImpl(keycode)
}
func IsMouseButtonPressed(button int) bool {
	return inputInstance.input.IsMouseButtonPressedImpl(button)
}
func GetMousePosition() (x, y float32) {
	xPos, yPos := inputInstance.input.GetMousePositionImpl()
	return xPos, yPos
}
func GetMouseX() float32 {
	return inputInstance.input.GetMouseXImpl()
}
func GetMouseY() float32 {
	return inputInstance.input.GetMouseYImpl()
}
