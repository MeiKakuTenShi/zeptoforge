package mouseEvent

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

type mouseMovedEvent struct {
	event          *event.Eventum
	mouseX, mouseY float32
}

func NewMouseMovedEvent(x, y float32) mouseMovedEvent {
	return mouseMovedEvent{event: event.NewEventum(&mouseMovedEvent{}, event.MouseMoved), mouseX: x, mouseY: y}
}
func (mm *mouseMovedEvent) GetX() float32 {
	return mm.mouseX
}
func (mm *mouseMovedEvent) GetY() float32 {
	return mm.mouseY
}
func (mouseMovedEvent) GetStaticType() event.EventType {
	return event.MouseMoved
}
func (mm mouseMovedEvent) GetEventType() event.EventType {
	return mouseMovedEvent{}.GetStaticType()
}
func (mm mouseMovedEvent) GetName() string {
	return "MouseMoved"
}
func (mm mouseMovedEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryMouse, event.EventCategoryInput}
}
func (mm mouseMovedEvent) ToString() string {
	return "MouseMovedEvent: " + fmt.Sprintf("%g, %g", mm.mouseX, mm.mouseY)
}
func (mm mouseMovedEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(mm.GetCategoryFlags(), e)
}

type mouseScrolledEvent struct {
	event            *event.Eventum
	xOffset, yOffset float32
}

func NewMouseScrolledEvent(xOffset, yOffset float32) mouseScrolledEvent {
	return mouseScrolledEvent{event: event.NewEventum(&mouseScrolledEvent{}, event.MouseScrolled), xOffset: xOffset, yOffset: yOffset}
}
func (ms mouseScrolledEvent) GetXOffset() float32 {
	return ms.xOffset
}
func (ms mouseScrolledEvent) GetYOffset() float32 {
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
	return []event.EventCategory{event.EventCategoryMouse, event.EventCategoryInput}
}
func (ms mouseScrolledEvent) ToString() string {
	return "MouseScrolledEvent: " + fmt.Sprintf("%g, %g", ms.GetXOffset(), ms.GetYOffset())
}
func (ms mouseScrolledEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(ms.GetCategoryFlags(), e)
}

type mouseButtonEvent struct {
	event  *event.Eventum
	button int
}

func NewMouseButtonEvent(b int, e event.EventType) mouseButtonEvent {
	return mouseButtonEvent{event: event.NewEventum(&mouseButtonEvent{}, event.NoneType), button: b}

}
func (mb mouseButtonEvent) GetMouseButton() int {
	return mb.button
}
func (mouseButtonEvent) GetStaticType() event.EventType {
	return event.NoneType
}
func (mb mouseButtonEvent) GetEventType() event.EventType {
	return mouseButtonEvent{}.GetStaticType()
}
func (mouseButtonEvent) GetName() string {
	return "MouseButton"
}
func (mb mouseButtonEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryMouse, event.EventCategoryInput}
}
func (mouseButtonEvent) ToString() string {
	return "MouseButtonEvent: "
}
func (mb mouseButtonEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(mb.GetCategoryFlags(), e)
}

type mouseButtonPressedEvent struct {
	mouseEvent mouseButtonEvent
}

func NewMouseButtonPressedEvent(b int) mouseButtonPressedEvent {
	return mouseButtonPressedEvent{mouseEvent: NewMouseButtonEvent(b, event.MouseButtonPressed)}
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
func (mp mouseButtonPressedEvent) ToString() string {
	return fmt.Sprintf("MouseButtonPressedEvent: %v", mp.mouseEvent.button)
}

type mouseButtonReleasedEvent struct {
	mouseEvent mouseButtonEvent
}

func NewMouseButtonReleasedEvent(b int) mouseButtonReleasedEvent {
	return mouseButtonReleasedEvent{mouseEvent: NewMouseButtonEvent(b, event.MouseButtonReleased)}
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
func (mr mouseButtonReleasedEvent) ToString() string {
	return fmt.Sprintf("MouseButtonReleasedEvent: %v", mr.mouseEvent.button)
}
