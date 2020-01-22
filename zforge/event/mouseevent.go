package event

import (
	"fmt"
)

type MouseMovedEvent struct {
	mouseX, mouseY float64
}

func NewMouseMovedEvent(x, y float64) *Eventum {
	return newEventum(&MouseMovedEvent{mouseX: x, mouseY: y}, MouseMoved)
}
func (mm MouseMovedEvent) GetX() float64 {
	return mm.mouseX
}
func (mm MouseMovedEvent) GetY() float64 {
	return mm.mouseY
}
func (MouseMovedEvent) GetStaticType() EventType {
	return MouseMoved
}
func (mm MouseMovedEvent) GetEventType() EventType {
	return mm.GetStaticType()
}
func (MouseMovedEvent) GetName() string {
	return "MouseMoved"
}
func (MouseMovedEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryMouse, EventCategoryInput}
}
func (mm MouseMovedEvent) String() string {
	return fmt.Sprintf("MouseMovedEvent| XPOS(%g) YPOS(%g)", mm.mouseX, mm.mouseY)
}
func (MouseMovedEvent) IsInCategory(cat EventCategory) bool {
	return contains(MouseMovedEvent{}.GetCategoryFlags(), cat)
}

type MouseScrolledEvent struct {
	xOffset, yOffset float64
}

func NewMouseScrolledEvent(xOff, yOff float64) *Eventum {
	return newEventum(&MouseScrolledEvent{xOffset: xOff, yOffset: yOff}, MouseScrolled)
}
func (ms MouseScrolledEvent) GetXOffset() float64 {
	return ms.xOffset
}
func (ms MouseScrolledEvent) GetYOffset() float64 {
	return ms.yOffset
}
func (MouseScrolledEvent) GetStaticType() EventType {
	return MouseScrolled
}
func (ms MouseScrolledEvent) GetEventType() EventType {
	return ms.GetStaticType()
}
func (MouseScrolledEvent) GetName() string {
	return "MouseScrolled"
}
func (MouseScrolledEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryMouse, EventCategoryInput}
}
func (ms MouseScrolledEvent) String() string {
	return fmt.Sprintf("MouseScrolledEvent| XOFFSET(%g) YOFFSET(%g)", ms.xOffset, ms.yOffset)
}
func (MouseScrolledEvent) IsInCategory(cat EventCategory) bool {
	return contains(MouseScrolledEvent{}.GetCategoryFlags(), cat)
}

type MouseButtonPressedEvent struct {
	button int
}

func NewMouseButtonPressedEvent(b int) *Eventum {
	return newEventum(&MouseButtonPressedEvent{button: b}, MouseButtonPressed)
}
func (mp MouseButtonPressedEvent) GetButton() int {
	return mp.button
}
func (MouseButtonPressedEvent) GetStaticType() EventType {
	return MouseButtonPressed
}
func (mp MouseButtonPressedEvent) GetEventType() EventType {
	return mp.GetStaticType()
}
func (MouseButtonPressedEvent) GetName() string {
	return "MouseButtonPressed"
}
func (MouseButtonPressedEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryMouse, EventCategoryInput}
}
func (mp MouseButtonPressedEvent) String() string {
	return fmt.Sprintf("MouseButtonPressedEvent| BUTTON(%v)", mp.button)
}
func (MouseButtonPressedEvent) IsInCategory(cat EventCategory) bool {
	return contains(MouseButtonPressedEvent{}.GetCategoryFlags(), cat)
}

type MouseButtonReleasedEvent struct {
	button int
}

func NewMouseButtonReleasedEvent(b int) *Eventum {
	return newEventum(&MouseButtonReleasedEvent{button: b}, MouseButtonReleased)
}
func (mr MouseButtonReleasedEvent) GetButton() int {
	return mr.button
}
func (MouseButtonReleasedEvent) GetStaticType() EventType {
	return MouseButtonReleased
}
func (mr MouseButtonReleasedEvent) GetEventType() EventType {
	return mr.GetStaticType()
}
func (MouseButtonReleasedEvent) GetName() string {
	return "MouseButtonReleased"
}
func (MouseButtonReleasedEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryMouse, EventCategoryInput}
}
func (mr MouseButtonReleasedEvent) String() string {
	return fmt.Sprintf("MouseButtonReleasedEvent| BUTTON(%v)", mr.button)
}
func (MouseButtonReleasedEvent) IsInCategory(cat EventCategory) bool {
	return contains(MouseButtonReleasedEvent{}.GetCategoryFlags(), cat)
}
