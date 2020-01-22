package event

import (
	"fmt"
)

type WindowResizeEvent struct {
	width, height int
}

func NewWindowResizeEvent(w, h int) *Eventum {
	return newEventum(&WindowResizeEvent{width: w, height: h}, WindowResize)
}
func (wr WindowResizeEvent) GetWidth() int {
	return wr.width
}
func (wr WindowResizeEvent) GetHeight() int {
	return wr.height
}
func (wr WindowResizeEvent) GetEventType() EventType {
	return WindowResize
}
func (WindowResizeEvent) GetName() string {
	return "WindowResize"
}
func (WindowResizeEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryApplication}
}
func (wr WindowResizeEvent) String() string {
	return fmt.Sprintf("WindowResizeEvent| WIDTH(%v) HEIGHT(%v)", wr.width, wr.height)
}
func (WindowResizeEvent) IsInCategory(cat EventCategory) bool {
	return contains(WindowResizeEvent{}.GetCategoryFlags(), cat)
}

type WindowCloseEvent struct {
}

func NewWindowCloseEvent() *Eventum {
	return newEventum(&WindowCloseEvent{}, WindowClose)
}
func (wc WindowCloseEvent) GetEventType() EventType {
	return WindowClose
}
func (WindowCloseEvent) GetName() string {
	return "WindowClose"
}
func (WindowCloseEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryApplication}
}
func (WindowCloseEvent) String() string {
	return "WindowCloseEvent"
}
func (WindowCloseEvent) IsInCategory(cat EventCategory) bool {
	return contains(WindowCloseEvent{}.GetCategoryFlags(), cat)
}

type AppTickEvent struct {
}

func NewAppTickEvent() *Eventum {
	return newEventum(&AppTickEvent{}, AppTick)
}
func (at AppTickEvent) GetEventType() EventType {
	return AppTick
}
func (AppTickEvent) GetName() string {
	return "AppTick"
}
func (AppTickEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryApplication}
}
func (AppTickEvent) String() string {
	return "WindowTickEvent"
}
func (AppTickEvent) IsInCategory(cat EventCategory) bool {
	return contains(AppTickEvent{}.GetCategoryFlags(), cat)
}

type AppUpdateEvent struct {
}

func NewAppUpdateEvent() *Eventum {
	return newEventum(&AppUpdateEvent{}, AppUpdate)
}
func (au AppUpdateEvent) GetEventType() EventType {
	return AppUpdate
}
func (AppUpdateEvent) GetName() string {
	return "AppUpdate"
}
func (AppUpdateEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryApplication}
}
func (AppUpdateEvent) String() string {
	return "WindowUpdateEvent"
}
func (AppUpdateEvent) IsInCategory(cat EventCategory) bool {
	return contains(AppUpdateEvent{}.GetCategoryFlags(), cat)
}

type AppRenderEvent struct {
}

func NewAppRenderEvent() *Eventum {
	return newEventum(&AppRenderEvent{}, AppRender)
}
func (ar AppRenderEvent) GetEventType() EventType {
	return AppRender
}
func (AppRenderEvent) GetName() string {
	return "AppRender"
}
func (AppRenderEvent) GetCategoryFlags() []EventCategory {
	return []EventCategory{EventCategoryApplication}
}
func (AppRenderEvent) String() string {
	return "WindowRenderEvent"
}
func (AppRenderEvent) IsInCategory(cat EventCategory) bool {
	return contains(AppRenderEvent{}.GetCategoryFlags(), cat)
}
