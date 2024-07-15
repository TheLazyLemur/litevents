package store

type Store interface {
	CreateEvent(eventType string, data string) (int64, error)
	CreateCursor(cursorName string, eventType string) (int64, error)
	GetNextEvent(cursorName string) (int64, string, error)
	AckEvent(cursorName string, eventId int64) error
}
