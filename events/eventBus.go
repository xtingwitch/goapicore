package events

type EventBus struct {
	listeners map[string][]func(Event)
	eventChan chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[string][]func(Event)),
		eventChan: make(chan Event),
	}
}

func (eb *EventBus) AddListener(eventType string, listener func(Event)) {
	eb.listeners[eventType] = append(eb.listeners[eventType], listener)
}

func (eb *EventBus) Broadcast(event Event) {
	eb.eventChan <- event
}

func (eb *EventBus) Start() {
	for event := range eb.eventChan {
		for _, listener := range eb.listeners[event.Name] {
			listener(event)
		}
	}
}

var globalEventBus *EventBus

func GetGlobalEventBus() *EventBus {
	if globalEventBus == nil {
		globalEventBus = NewEventBus()
		go globalEventBus.Start()
	}

	return globalEventBus
}
