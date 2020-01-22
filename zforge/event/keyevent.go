package event

import (
	"fmt"
)

type KeyPressedEvent struct {
	keyCode, repeatCount int
}

func NewKeyPressedEvent(code, count int) *Eventum {
	return newEventum(&KeyPressedEvent{keyCode: code, repeatCount: count}, KeyPressed)
}
func (kp KeyPressedEvent) GetKey() int {
	return kp.keyCode
}
func (kp KeyPressedEvent) GetRepeatCount() int {
	return kp.repeatCount
}
func (KeyPressedEvent) GetEventType() EventType {
	return KeyPressed
}
func (KeyPressedEvent) GetName() string {
	return "KeyPressed"
}
func (KeyPressedEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryKeyboard, EventCategoryInput}
}
func (kp KeyPressedEvent) String() string {
	return fmt.Sprintf("KeyPressedEvent| KEYCODE(%v (%v REPEATS))", kp.keyCode, kp.repeatCount)
}
func (KeyPressedEvent) IsInCategory(cat EventCategory) bool {
	return contains(KeyPressedEvent{}.GetCategoryFlags(), cat)
}

type KeyReleasedEvent struct {
	keyCode int
}

func NewKeyReleasedEvent(code int) *Eventum {
	return newEventum(&KeyReleasedEvent{keyCode: code}, KeyReleased)
}
func (kr *KeyReleasedEvent) GetKey() int {
	return kr.keyCode
}
func (KeyReleasedEvent) GetEventType() EventType {
	return KeyReleased
}
func (KeyReleasedEvent) GetName() string {
	return "KeyReleased"
}
func (KeyReleasedEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryKeyboard, EventCategoryInput}
}
func (kr KeyReleasedEvent) String() string {
	return fmt.Sprintf("KeyReleasedEvent| KEYCODE(%v)", kr.keyCode)
}
func (KeyReleasedEvent) IsInCategory(cat EventCategory) bool {
	return contains(KeyReleasedEvent{}.GetCategoryFlags(), cat)
}

type KeyTypedEvent struct {
	keyCode rune
}

func NewKeyTypedEvent(code rune) *Eventum {
	return newEventum(&KeyTypedEvent{keyCode: code}, KeyTyped)
}
func (kt *KeyTypedEvent) GetKey() rune {
	return kt.keyCode
}
func (kt *KeyTypedEvent) GetEventType() EventType {
	return KeyTyped
}
func (KeyTypedEvent) GetName() string {
	return "KeyTyped"
}
func (KeyTypedEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryKeyboard, EventCategoryInput}
}
func (kt *KeyTypedEvent) String() string {
	return fmt.Sprintf("KeyTypedEvent| KEYCODE(%v)", kt.keyCode)
}
func (KeyTypedEvent) IsInCategory(cat EventCategory) bool {
	return contains(KeyTypedEvent{}.GetCategoryFlags(), cat)
}
