package ops

import "litevents/store"

func GetOrCreateCursor(s store.Store, cursorName string, eventType string) (int64, error) {
	cursorId, err := s.CreateCursor(cursorName, eventType)
	if err != nil {
		return 0, err
	}

	return cursorId, nil
}
