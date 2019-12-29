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

type mouseMovedEvent struct {
	mouseX, mouseY float64
}

func NewMouseMovedEvent(x, y float64) *event.Eventum {
	return event.NewEventum(&mouseMovedEvent{mouseX: x, mouseY: y}, event.MouseMoved)
}
func (mm *mouseMovedEvent) GetX() float64 {
	return mm.mouseX
}
func (mm *mouseMovedEvent) GetY() float64 {
	return mm.mouseY
}
func (mm mouseMovedEvent) GetStaticType() event.EventType {
	return event.MouseMoved
}
func (mm mouseMovedEvent) GetEventType() event.EventType {
	return mouseMovedEvent{}.GetStaticType()
}
func (mm mouseMovedEvent) GetName() string {
	return "MouseMoved"
}
func (mm mouseMovedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (mm mouseMovedEvent) ToString() string {
	return "MouseMovedEvent: " + fmt.Sprintf("%g, %g", mm.mouseX, mm.mouseY)
}
func (mm mouseMovedEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type mouseScrolledEvent struct {
	xOffset, yOffset float64
}

func NewMouseScrolledEvent(xOff, yOff float64) *event.Eventum {
	return event.NewEventum(&mouseScrolledEvent{xOffset: xOff, yOffset: yOff}, event.MouseScrolled)
}
func (ms mouseScrolledEvent) GetXOffset() float64 {
	return ms.xOffset
}
func (ms mouseScrolledEvent) GetYOffset() float64 {
	return ms.yOffset
}
func (mouseScrolledEvent) GetStaticType() event.EventType {
	return event.MouseScrolled
}
func (ms mouseScrolledEvent) GetEventType() event.EventType {
	return mouseScrolledEvent{}.GetStaticType()
}
func (ms mouseScrolledEvent) GetName() string {
	return "MouseScrolled"
}
func (ms mouseScrolledEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (ms mouseScrolledEvent) ToString() string {
	return "MouseScrolledEvent: " + fmt.Sprintf("%g, %g", ms.GetXOffset(), ms.GetYOffset())
}
func (ms mouseScrolledEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type mouseButtonPressedEvent struct {
	button int
}

func NewMouseButtonPressedEvent(b int) *event.Eventum {
	return event.NewEventum(&mouseButtonPressedEvent{button: b}, event.MouseButtonPressed)
}
func (mouseButtonPressedEvent) GetStaticType() event.EventType {
	return event.MouseButtonPressed
}
func (mp mouseButtonPressedEvent) GetEventType() event.EventType {
	return mouseButtonPressedEvent{}.GetStaticType()
}
func (mp mouseButtonPressedEvent) GetName() string {
	return "MouseButtonPressed"
}
func (mp mouseButtonPressedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (mp mouseButtonPressedEvent) ToString() string {
	return fmt.Sprintf("MouseButtonPressedEvent: %v", mp.button)
}
func (mp mouseButtonPressedEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type mouseButtonReleasedEvent struct {
	button int
}

func NewMouseButtonReleasedEvent(b int) *event.Eventum {
	return event.NewEventum(&mouseButtonReleasedEvent{button: b}, event.MouseButtonReleased)
}
func (mouseButtonReleasedEvent) GetStaticType() event.EventType {
	return event.MouseButtonReleased
}
func (mr mouseButtonReleasedEvent) GetEventType() event.EventType {
	return mr.GetStaticType()
}
func (mr mouseButtonReleasedEvent) GetName() string {
	return "MouseButtonReleased"
}
func (mr mouseButtonReleasedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (mr mouseButtonReleasedEvent) ToString() string {
	return fmt.Sprintf("MouseButtonReleasedEvent: %v", mr.button)
}
func (mr mouseButtonReleasedEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}
