package mouseEvent

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

var getCatFlags = func() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryMouse, event.EventCategoryInput}
}
var inCatCheck = func(cat event.EventCategory) bool {
	return event.Contains([]event.EventCategory{event.EventCategoryMouse, event.EventCategoryInput}, cat)
}

type MouseMovedEvent struct {
	mouseX, mouseY float64
}

func NewMouseMovedEvent(x, y float64) *event.Eventum {
	return event.NewEventum(&MouseMovedEvent{mouseX: x, mouseY: y}, event.MouseMoved)
}
func (mm MouseMovedEvent) GetX() float64 {
	return mm.mouseX
}
func (mm MouseMovedEvent) GetY() float64 {
	return mm.mouseY
}
func (MouseMovedEvent) GetStaticType() event.EventType {
	return event.MouseMoved
}
func (mm MouseMovedEvent) GetEventType() event.EventType {
	return mm.GetStaticType()
}
func (MouseMovedEvent) GetName() string {
	return "MouseMoved"
}
func (MouseMovedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (mm MouseMovedEvent) String() string {
	return fmt.Sprintf("MouseMovedEvent| XPOS(%g) YPOS(%g)", mm.mouseX, mm.mouseY)
}
func (MouseMovedEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type MouseScrolledEvent struct {
	xOffset, yOffset float64
}

func NewMouseScrolledEvent(xOff, yOff float64) *event.Eventum {
	return event.NewEventum(&MouseScrolledEvent{xOffset: xOff, yOffset: yOff}, event.MouseScrolled)
}
func (ms MouseScrolledEvent) GetXOffset() float64 {
	return ms.xOffset
}
func (ms MouseScrolledEvent) GetYOffset() float64 {
	return ms.yOffset
}
func (MouseScrolledEvent) GetStaticType() event.EventType {
	return event.MouseScrolled
}
func (ms MouseScrolledEvent) GetEventType() event.EventType {
	return ms.GetStaticType()
}
func (MouseScrolledEvent) GetName() string {
	return "MouseScrolled"
}
func (MouseScrolledEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (ms MouseScrolledEvent) String() string {
	return fmt.Sprintf("MouseScrolledEvent| XOFFSET(%g) YOFFSET(%g)", ms.xOffset, ms.yOffset)
}
func (MouseScrolledEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type MouseButtonPressedEvent struct {
	button int
}

func NewMouseButtonPressedEvent(b int) *event.Eventum {
	return event.NewEventum(&MouseButtonPressedEvent{button: b}, event.MouseButtonPressed)
}
func (mp MouseButtonPressedEvent) GetButton() int {
	return mp.button
}
func (MouseButtonPressedEvent) GetStaticType() event.EventType {
	return event.MouseButtonPressed
}
func (mp MouseButtonPressedEvent) GetEventType() event.EventType {
	return mp.GetStaticType()
}
func (MouseButtonPressedEvent) GetName() string {
	return "MouseButtonPressed"
}
func (MouseButtonPressedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (mp MouseButtonPressedEvent) String() string {
	return fmt.Sprintf("MouseButtonPressedEvent| BUTTON(%v)", mp.button)
}
func (MouseButtonPressedEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type MouseButtonReleasedEvent struct {
	button int
}

func NewMouseButtonReleasedEvent(b int) *event.Eventum {
	return event.NewEventum(&MouseButtonReleasedEvent{button: b}, event.MouseButtonReleased)
}
func (mr MouseButtonReleasedEvent) GetButton() int {
	return mr.button
}
func (MouseButtonReleasedEvent) GetStaticType() event.EventType {
	return event.MouseButtonReleased
}
func (mr MouseButtonReleasedEvent) GetEventType() event.EventType {
	return mr.GetStaticType()
}
func (MouseButtonReleasedEvent) GetName() string {
	return "MouseButtonReleased"
}
func (MouseButtonReleasedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (mr MouseButtonReleasedEvent) String() string {
	return fmt.Sprintf("MouseButtonReleasedEvent| BUTTON(%v)", mr.button)
}
func (MouseButtonReleasedEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}
