package order

import "fmt"

type ProcessManager struct {
	eventHandlers map[EventType][]func(event Event)
	store         *EventStore
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		eventHandlers: make(map[EventType][]func(event Event)),
		store:         NewEventStore(),
	}
}

// イベントを発行する
func PublishEvent(pm *ProcessManager, event Event) {
	// イベントを履歴に追加
	pm.store.AddEvent(event)
}

// DBの変更をトリガーにイベントを発行する
func Listen(pm *ProcessManager) {
	for event := range pm.store.Notify() {

		fmt.Println("Received event:", event)
		if handlers, exists := pm.eventHandlers[event.EventType()]; exists {
			for _, handler := range handlers {
				go handler(event)
			}
		}
	}
}

// イベントを購読する
func SubscribeEvent[T Event](pm *ProcessManager, handler func(event T) ([]Event, error)) {
	pm.eventHandlers[getEventType[T]()] = append(pm.eventHandlers[getEventType[T]()], func(event Event) {
		if typedEvent, ok := event.(T); ok {
			events, err := handler(typedEvent)
			if err != nil {
				fmt.Println("イベント処理中にエラーが発生:", err)
			}
			for _, e := range events {
				PublishEvent(pm, e)
			}
		}
	})

}

type Event interface {
	EventType() EventType
	WorkflowID() WorkflowID
}
type EventType string

type WorkflowID string

func getEventType[T Event]() EventType {
	var instance T
	return instance.EventType()
}

type EventStore struct {
	// イベント履歴
	events []Event
	// 通知用チャンネル
	ch chan Event
}

func NewEventStore() *EventStore {
	return &EventStore{
		events: []Event{},
		ch:     make(chan Event),
	}
}
func (es *EventStore) AddEvent(event Event) {
	es.events = append(es.events, event)
}

func (es *EventStore) Notify() <-chan Event {
	ch := make(chan Event)
	go func() {
		for _, event := range es.events {
			ch <- event
		}
		close(ch)
	}()
	return ch
}
