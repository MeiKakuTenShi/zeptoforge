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
type EventFn struct {
	Event Event
	Fn    func(Event) bool
}

func NewEventum(ev Event, e EventType) *Eventum {
	return &Eventum{eventum: ev, handled: false, eventType: e}
}
func (e *Eventum) GetEvent() Event {
	return e.eventum
}

func NewEventDispatcher(e *Eventum) eventDispatcher {
	return eventDispatcher{event: e}
}
func (ed eventDispatcher) Dispatch(fn EventFn) bool {
	if ed.event.eventum.GetEventType() == fn.Event.GetEventType() {
		ed.event.handled = fn.Fn(ed.event.eventum)
		return true
	}
	return false
}
func (e *Eventum) ToString() string {
	return e.eventum.ToString()
}
func (e *Eventum) Done() bool {
	return e.handled
}

var Contains = func(s []EventCategory, e EventCategory) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
