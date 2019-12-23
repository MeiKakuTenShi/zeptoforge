package appEvent

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

type windowResizeEvent struct {
	event         *event.Eventum
	width, height uint
}

func NewWindowResizeEvent(width, height uint) windowResizeEvent {
	return windowResizeEvent{event: event.NewEventum(&windowResizeEvent{}, event.WindowResize), width: width, height: height}
}
func (wr windowResizeEvent) GetWidth() uint {
	return wr.width
}
func (wr windowResizeEvent) GetHeight() uint {
	return wr.height
}
func (wr windowResizeEvent) GetStaticType() event.EventType {
	return event.WindowResize
}
func (wr windowResizeEvent) GetEventType() event.EventType {
	return wr.GetStaticType()
}
func (wr windowResizeEvent) GetName() string {
	return "WindowResize"
}
func (wr windowResizeEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryApplication}
}
func (wr windowResizeEvent) ToString() string {
	return "WindowResizeEvent: " + string(wr.width) + ", " + string(wr.height)
}
func (wr windowResizeEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(wr.GetCategoryFlags(), e)
}

type windowCloseEvent struct {
	event *event.Eventum
}

func NewWindowCloseEvent() windowCloseEvent {
	return windowCloseEvent{event: event.NewEventum(&windowCloseEvent{}, event.WindowClose)}
}
func (wc windowCloseEvent) GetStaticType() event.EventType {
	return event.WindowClose
}
func (wc windowCloseEvent) GetEventType() event.EventType {
	return wc.GetStaticType()
}
func (wc windowCloseEvent) GetName() string {
	return "WindowClose"
}
func (wc windowCloseEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryApplication}
}
func (wc windowCloseEvent) ToString() string {
	return "Window Close Event"
}
func (wc windowCloseEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(wc.GetCategoryFlags(), e)
}

type appTickEvent struct {
	event *event.Eventum
}

func NewAppTickEvent() appTickEvent {
	return appTickEvent{event: event.NewEventum(&appTickEvent{}, event.AppTick)}
}
func (at appTickEvent) GetStaticType() event.EventType {
	return event.AppTick
}
func (at appTickEvent) GetEventType() event.EventType {
	return at.GetStaticType()
}
func (at appTickEvent) GetName() string {
	return "AppTick"
}
func (at appTickEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryApplication}
}
func (at appTickEvent) ToString() string {
	return "Window Close Event"
}
func (at appTickEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(at.GetCategoryFlags(), e)
}

type appUpdateEvent struct {
	event *event.Eventum
}

func NewAppUpdateEvent() appUpdateEvent {
	return appUpdateEvent{event: event.NewEventum(&appUpdateEvent{}, event.AppUpdate)}
}
func (au appUpdateEvent) GetStaticType() event.EventType {
	return event.AppTick
}
func (au appUpdateEvent) GetEventType() event.EventType {
	return au.GetStaticType()
}
func (au appUpdateEvent) GetName() string {
	return "AppUpdate"
}
func (au appUpdateEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryApplication}
}
func (au appUpdateEvent) ToString() string {
	return "Window Update Event"
}
func (au appUpdateEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(au.GetCategoryFlags(), e)
}

type appRenderEvent struct {
	event *event.Eventum
}

func NewAppRenderEvent() appRenderEvent {
	return appRenderEvent{event: event.NewEventum(&appRenderEvent{}, event.AppRender)}
}
func (ar appRenderEvent) GetStaticType() event.EventType {
	return event.AppTick
}
func (ar appRenderEvent) GetEventType() event.EventType {
	return ar.GetStaticType()
}
func (ar appRenderEvent) GetName() string {
	return "AppRender"
}
func (ar appRenderEvent) GetCategoryFlags() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryApplication}
}
func (ar appRenderEvent) ToString() string {
	return "Window Render Event"
}
func (ar appRenderEvent) IsInCategory(e event.EventCategory) bool {
	return event.Contains(ar.GetCategoryFlags(), e)
}
