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
	KeyTyped
	MouseButtonPressed
	MouseButtonReleased
	MouseMoved
	MouseScrolled

	NoneCategory EventCategory = iota
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
	String() string
	IsInCategory(EventCategory) bool
}
type EventFn struct {
	Event    EventType
	Receiver interface{}
	Fn       func(interface{}, Eventum) bool
}

type Eventum struct {
	event     Event
	handled   bool
	eventType EventType
}

func newEventum(e Event, et EventType) *Eventum {
	r := new(Eventum)
	r.event = e
	r.handled = false
	r.eventType = et

	return r
}
func (e *Eventum) GetEvent() Event {
	return e.event
}
func (e *Eventum) String() string {
	return e.event.String()
}
func (e *Eventum) Done() bool {
	return e.handled
}

type EventDispatcher struct {
	event *Eventum
}

func NewEventDispatcher(e *Eventum) *EventDispatcher {
	r := new(EventDispatcher)
	r.event = e
	return r
}
func (ed EventDispatcher) Dispatch(fn EventFn) bool {
	if ed.event.eventType == fn.Event {
		ed.event.handled = fn.Fn(fn.Receiver, *ed.event)
		return true
	}
	return false
}

func contains(s []EventCategory, e EventCategory) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
