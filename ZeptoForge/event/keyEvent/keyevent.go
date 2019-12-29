package keyEvent

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

var getCatFlags = func() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryKeyboard, event.EventCategoryInput}
}
var inCatCheck = func(cat event.EventCategory) bool {
	if cat == event.EventCategoryKeyboard || cat == event.EventCategoryInput {
		return true
	}
	return false
}

type KeyPressedEvent struct {
	KeyCode, RepeatCount int
}

func NewKeyPressedEvent(code, count int) *event.Eventum {
	return event.NewEventum(&KeyPressedEvent{KeyCode: code, RepeatCount: count}, event.KeyPressed)
}
func (kp KeyPressedEvent) GetRepeatCount() int {
	return kp.RepeatCount
}
func (kp KeyPressedEvent) GetStaticType() event.EventType {
	return event.KeyPressed
}
func (kp KeyPressedEvent) GetEventType() event.EventType {
	return kp.GetStaticType()
}
func (KeyPressedEvent) GetName() string {
	return "KeyPressed"
}
func (KeyPressedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (kp KeyPressedEvent) ToString() string {
	return fmt.Sprintf("KeyPressedEvent: %v (%v repeats)", kp.KeyCode, kp.RepeatCount)
}
func (kp KeyPressedEvent) IsInCategory(cat event.EventCategory) bool {
	return inCatCheck(cat)
}

type KeyReleasedEvent struct {
	KeyCode int
}

func NewKeyReleasedEvent(code int) *event.Eventum {
	return event.NewEventum(&KeyReleasedEvent{KeyCode: code}, event.KeyReleased)
}
func (kr KeyReleasedEvent) GetStaticType() event.EventType {
	return event.KeyReleased
}
func (kr KeyReleasedEvent) GetEventType() event.EventType {
	return kr.GetStaticType()
}
func (KeyReleasedEvent) GetName() string {
	return "KeyReleased"
}
func (KeyReleasedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (kr KeyReleasedEvent) ToString() string {
	return fmt.Sprintf("KeyReleasedEvent: %v", kr.KeyCode)
}
func (kr KeyReleasedEvent) IsInCategory(cat event.EventCategory) bool {
	return inCatCheck(cat)
}
