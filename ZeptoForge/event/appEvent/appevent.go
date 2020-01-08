package appEvent

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

var getCatFlags = func() []event.EventCategory {
	return []event.EventCategory{event.EventCategoryApplication}
}
var inCatCheck = func(cat event.EventCategory) bool {
	return event.Contains([]event.EventCategory{event.EventCategoryApplication}, cat)
}

type WindowResizeEvent struct {
	width, height int
}

func NewWindowResizeEvent(width, height int) *event.Eventum {
	ev := new(WindowResizeEvent)
	ev.width = width
	ev.height = height

	return event.NewEventum(ev, event.WindowResize)
}
func (wr WindowResizeEvent) GetWidth() int {
	return wr.width
}
func (wr WindowResizeEvent) GetHeight() int {
	return wr.height
}
func (WindowResizeEvent) GetStaticType() event.EventType {
	return event.WindowResize
}
func (wr WindowResizeEvent) GetEventType() event.EventType {
	return wr.GetStaticType()
}
func (WindowResizeEvent) GetName() string {
	return "WindowResize"
}
func (WindowResizeEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (wr WindowResizeEvent) String() string {
	return fmt.Sprintf("WindowResizeEvent| WIDTH(%v) HEIGHT(%v)", wr.width, wr.height)
}
func (WindowResizeEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type WindowCloseEvent struct {
}

func NewWindowCloseEvent() *event.Eventum {
	return event.NewEventum(new(WindowCloseEvent), event.WindowClose)
}
func (WindowCloseEvent) GetStaticType() event.EventType {
	return event.WindowClose
}
func (wc WindowCloseEvent) GetEventType() event.EventType {
	return wc.GetStaticType()
}
func (WindowCloseEvent) GetName() string {
	return "WindowClose"
}
func (WindowCloseEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (WindowCloseEvent) String() string {
	return "WindowCloseEvent"
}
func (WindowCloseEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type AppTickEvent struct {
}

func NewAppTickEvent() *event.Eventum {
	return event.NewEventum(new(AppTickEvent), event.AppTick)
}
func (AppTickEvent) GetStaticType() event.EventType {
	return event.AppTick
}
func (at AppTickEvent) GetEventType() event.EventType {
	return at.GetStaticType()
}
func (AppTickEvent) GetName() string {
	return "AppTick"
}
func (AppTickEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (AppTickEvent) String() string {
	return "WindowTickEvent"
}
func (AppTickEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type AppUpdateEvent struct {
}

func NewAppUpdateEvent() *event.Eventum {
	return event.NewEventum(new(AppUpdateEvent), event.AppUpdate)
}
func (AppUpdateEvent) GetStaticType() event.EventType {
	return event.AppUpdate
}
func (au AppUpdateEvent) GetEventType() event.EventType {
	return au.GetStaticType()
}
func (AppUpdateEvent) GetName() string {
	return "AppUpdate"
}
func (AppUpdateEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (AppUpdateEvent) String() string {
	return "WindowUpdateEvent"
}
func (AppUpdateEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type AppRenderEvent struct {
}

func NewAppRenderEvent() *event.Eventum {
	return event.NewEventum(new(AppRenderEvent), event.AppRender)
}
func (AppRenderEvent) GetStaticType() event.EventType {
	return event.AppRender
}
func (ar AppRenderEvent) GetEventType() event.EventType {
	return ar.GetStaticType()
}
func (AppRenderEvent) GetName() string {
	return "AppRender"
}
func (AppRenderEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (AppRenderEvent) String() string {
	return "WindowRenderEvent"
}
func (AppRenderEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}
