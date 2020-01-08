package keyEvent

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

var getCatFlags = func() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryKeyboard, event.EventCategoryInput}
}
var inCatCheck = func(cat event.EventCategory) bool {
	return event.Contains(getCatFlags(), cat)
}

type KeyPressedEvent struct {
	keyCode, repeatCount int
}

func NewKeyPressedEvent(code, count int) *event.Eventum {
	ev := new(KeyPressedEvent)
	ev.keyCode = code
	ev.repeatCount = count

	return event.NewEventum(ev, event.KeyPressed)
}
func (kp KeyPressedEvent) GetKey() int {
	return kp.keyCode
}
func (kp KeyPressedEvent) GetRepeatCount() int {
	return kp.repeatCount
}
func (KeyPressedEvent) GetStaticType() event.EventType {
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
func (kp KeyPressedEvent) String() string {
	return fmt.Sprintf("KeyPressedEvent| KEYCODE(%v (%v REPEATS))", kp.keyCode, kp.repeatCount)
}
func (KeyPressedEvent) IsInCategory(cat event.EventCategory) bool {
	return inCatCheck(cat)
}

type KeyReleasedEvent struct {
	keyCode int
}

func NewKeyReleasedEvent(code int) *event.Eventum {
	ev := new(KeyReleasedEvent)
	ev.keyCode = code

	return event.NewEventum(ev, event.KeyReleased)
}
func (kr *KeyReleasedEvent) GetKey() int {
	return kr.keyCode
}
func (KeyReleasedEvent) GetStaticType() event.EventType {
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
func (kr KeyReleasedEvent) String() string {
	return fmt.Sprintf("KeyReleasedEvent| KEYCODE(%v)", kr.keyCode)
}
func (KeyReleasedEvent) IsInCategory(cat event.EventCategory) bool {
	return inCatCheck(cat)
}

type KeyTypedEvent struct {
	keyCode rune
}

func NewKeyTypedEvent(code rune) *event.Eventum {
	ev := new(KeyTypedEvent)
	ev.keyCode = code

	return event.NewEventum(ev, event.KeyTyped)
}
func (kt *KeyTypedEvent) GetKey() rune {
	return kt.keyCode
}
func (KeyTypedEvent) GetStaticType() event.EventType {
	return event.KeyTyped
}
func (kt *KeyTypedEvent) GetEventType() event.EventType {
	return kt.GetStaticType()
}
func (KeyTypedEvent) GetName() string {
	return "KeyTyped"
}
func (KeyTypedEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (kt *KeyTypedEvent) String() string {
	return fmt.Sprintf("KeyTypedEvent| KEYCODE(%v)", kt.keyCode)
}
func (KeyTypedEvent) IsInCategory(cat event.EventCategory) bool {
	return inCatCheck(cat)
}
