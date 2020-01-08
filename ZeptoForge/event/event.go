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
	String() string
	IsInCategory(EventCategory) bool
}

type EventDispatcher struct {
	event *Eventum
}
type Eventum struct {
	eventum   Event
	handled   bool
	eventType EventType
}
type EventFn struct {
	Event Event
	Fn    func(Eventum) bool
}

func NewEventum(ev Event, e EventType) *Eventum {
	even := new(Eventum)
	even.eventum = ev
	even.handled = false
	even.eventType = e

	return even
}
func (e *Eventum) GetEvent() Event {
	return e.eventum
}

func NewEventDispatcher(e *Eventum) *EventDispatcher {
	dis := new(EventDispatcher)
	dis.event = e
	return dis
}
func (ed EventDispatcher) Dispatch(fn EventFn) bool {
	if ed.event.eventType == fn.Event.GetEventType() {
		ed.event.handled = fn.Fn(*ed.event)
		return true
	}
	return false
}
func (e *Eventum) String() string {
	return e.eventum.String()
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
