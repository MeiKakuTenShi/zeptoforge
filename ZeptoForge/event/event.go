package event

type EventType int
type EventCategory uint8

const (
	NoneType EventType = iota
	WindowClose
	WindowResize
	WindowFocus
	WindowLostFocus
	WindowMoved
	AppTick
	AppUpdate
	AppRender
	KeyPressed
	KeyReleased
	MouseButtonPressed
	MouseButtonReleased
	MouseMoved
	MouseScrolled

	NoneCategory = EventCategory(iota)
	EventCategoryApplication
	EventCategoryInput
	EventCategoryKeyboard
	EventCategoryMouse
	EventCategoryMouseButton
)

type Event interface {
	GetEventType() EventType
	GetName() string
	GetCategoryFlags() []EventCategory
	ToString() string
	IsInCategory(EventCategory) bool
}

type eventDispatcher struct {
	event *Eventum
}
type Eventum struct {
	eventum   Event
	handled   bool
	eventType EventType
}
type eventFn struct {
	event Event
	fn    func(*Event) bool
}

func NewEventum(ev Event, e EventType) *Eventum {
	return &Eventum{eventum: ev, handled: false, eventType: e}
}

func NewEventDispatcher(e *Eventum) eventDispatcher {
	return eventDispatcher{event: e}
}
func (ed eventDispatcher) Dispatch(fn eventFn) bool {
	if ed.event.eventum.GetEventType() == fn.event.GetEventType() {
		ed.event.handled = fn.fn(&ed.event.eventum)
		return true
	}
	return false
}

var Contains = func(s []EventCategory, e EventCategory) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
