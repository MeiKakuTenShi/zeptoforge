package keyEvent

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

type keyEvent struct {
	event   *event.Eventum
	keyCode int
}

func (k keyEvent) GetKeyCode() int {
	return k.keyCode
}
func (keyEvent) GetEventType() event.EventType {
	return event.NoneType
}
func (keyEvent) GetName() string {
	return ""
}
func (k keyEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryKeyboard, event.EventCategoryInput}
}
func (keyEvent) ToString() string {
	return ""
}
func (keyEvent) IsInCategory(cat event.EventCategory) bool {
	return false
}

type keyPressedEvent struct {
	event       keyEvent
	repeatCount int
}

func NewKeyPressedEvent(code, count int) keyPressedEvent {
	return keyPressedEvent{event: keyEvent{event: event.NewEventum(&keyEvent{}, event.KeyPressed), keyCode: code}, repeatCount: count}
}
func (kp keyPressedEvent) GetRepeatCount() int {
	return kp.repeatCount
}
func (kp keyPressedEvent) GetStaticType() event.EventType {
	return event.KeyPressed
}
func (kp keyPressedEvent) GetEventType() event.EventType {
	return kp.GetStaticType()
}
func (keyPressedEvent) GetName() string {
	return "KeyPressed"
}
func (kp keyPressedEvent) ToString() string {
	return "KeyPressedEvent: " + string(kp.event.keyCode) + " ( " + string(kp.repeatCount) + " repeats)"
}

type keyReleasedEvent struct {
	event keyEvent
}

func NewkeyReleasedEvent(code int) keyReleasedEvent {
	return keyReleasedEvent{event: keyEvent{event: event.NewEventum(&keyEvent{}, event.KeyPressed), keyCode: code}}
}

func (kr keyReleasedEvent) GetStaticType() event.EventType {
	return event.KeyReleased
}
func (kr keyReleasedEvent) GetEventType() event.EventType {
	return kr.GetStaticType()
}
func (kr keyReleasedEvent) GetName() string {
	return "KeyReleased"
}
func (kr keyReleasedEvent) ToString() string {
	return "KeyReleasedEvent: " + string(kr.event.keyCode)
}
