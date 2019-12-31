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
	return event.NewEventum(&WindowResizeEvent{width: width, height: height}, event.WindowResize)
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
func (wr WindowResizeEvent) GetName() string {
	return "WindowResize"
}
func (wr WindowResizeEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (wr WindowResizeEvent) ToString() string {
	return fmt.Sprintf("WindowResizeEvent: %v, %v", wr.width, wr.height)
}
func (wr WindowResizeEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type WindowCloseEvent struct {
}

func NewWindowCloseEvent() *event.Eventum {
	return event.NewEventum(&WindowCloseEvent{}, event.WindowClose)
}
func (WindowCloseEvent) GetStaticType() event.EventType {
	return event.WindowClose
}
func (wc WindowCloseEvent) GetEventType() event.EventType {
	return wc.GetStaticType()
}
func (wc WindowCloseEvent) GetName() string {
	return "WindowClose"
}
func (wc WindowCloseEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (wc WindowCloseEvent) ToString() string {
	return "WindowCloseEvent: "
}
func (wc WindowCloseEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type AppTickEvent struct {
}

func NewAppTickEvent() *event.Eventum {
	return event.NewEventum(&AppTickEvent{}, event.AppTick)
}
func (AppTickEvent) GetStaticType() event.EventType {
	return event.AppTick
}
func (at AppTickEvent) GetEventType() event.EventType {
	return at.GetStaticType()
}
func (at AppTickEvent) GetName() string {
	return "AppTick"
}
func (at AppTickEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (at AppTickEvent) ToString() string {
	return "WindowTickEvent: "
}
func (at AppTickEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type AppUpdateEvent struct {
}

func NewAppUpdateEvent() *event.Eventum {
	return event.NewEventum(&AppUpdateEvent{}, event.AppUpdate)
}
func (au AppUpdateEvent) GetStaticType() event.EventType {
	return event.AppUpdate
}
func (au AppUpdateEvent) GetEventType() event.EventType {
	return au.GetStaticType()
}
func (au AppUpdateEvent) GetName() string {
	return "AppUpdate"
}
func (au AppUpdateEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (au AppUpdateEvent) ToString() string {
	return "WindowUpdateEvent: "
}
func (au AppUpdateEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}

type AppRenderEvent struct {
}

func NewAppRenderEvent() *event.Eventum {
	return event.NewEventum(&AppRenderEvent{}, event.AppRender)
}
func (ar AppRenderEvent) GetStaticType() event.EventType {
	return event.AppRender
}
func (ar AppRenderEvent) GetEventType() event.EventType {
	return ar.GetStaticType()
}
func (ar AppRenderEvent) GetName() string {
	return "AppRender"
}
func (ar AppRenderEvent) GetCategoryFlags() []event.EventCategory {
	return getCatFlags()
}
func (ar AppRenderEvent) ToString() string {
	return "WindowRenderEvent: "
}
func (ar AppRenderEvent) IsInCategory(e event.EventCategory) bool {
	return inCatCheck(e)
}
