package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	db *sql.DB
}

func (sqlitestore *SqliteStore) CreateEvent(eventType string, data string) (int64, error) {
	sql := `
		INSERT INTO events (event_type, event_data)
		VALUES (?, ?) RETURNING event_id;
	`

	var id int64
	err := sqlitestore.db.QueryRow(sql, eventType, data).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (sqlitestore *SqliteStore) CreateCursor(cursorName string, eventType string) (int64, error) {
	sql := `
		INSERT OR IGNORE INTO cursors (cursor_name, event_type)
		VALUES (?, ?) RETURNING cursor_id;
	`

	var id int64
	err := sqlitestore.db.QueryRow(sql, cursorName, eventType).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (sqlitestore *SqliteStore) GetNextEvent(cursorName string) (int64, string, error) {
	sql := `
		SELECT last_event_id, event_type FROM cursors
		WHERE cursor_name = ?;
	`

	var lastEventId int64
	var eventType string

	err := sqlitestore.db.QueryRow(sql, cursorName).Scan(&lastEventId, &eventType)
	if err != nil {
		return 0, "", err
	}

	sql = `
		SELECT event_id, event_data FROM events
		WHERE event_id > ? AND event_type = ? ORDER BY event_id ASC LIMIT 1;
	`

	var eventData string
	var eventId int64

	err = sqlitestore.db.QueryRow(sql, lastEventId, eventType).Scan(&eventId, &eventData)
	if err != nil {
		return 0, "", err
	}

	return eventId, eventData, nil
}

func (sqlitestore *SqliteStore) AckEvent(cursorName string, eventId int64) error {
	sql := `
		UPDATE cursors
		SET last_event_id = ?
		WHERE cursor_name = ?;
	`

	_, err := sqlitestore.db.Exec(sql, eventId, cursorName)
	if err != nil {
		return err
	}

	return nil
}

func NewSqliteStore(dbString string) *SqliteStore {
	db, err := sql.Open("sqlite3", dbString)
	if err != nil {
		panic(err)
	}

	schema := `
		pragma journal_mode = WAL;

		CREATE TABLE IF NOT EXISTS events (
			event_id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_type TEXT NOT NULL CHECK (event_type <> ''),
			event_data TEXT NOT NULL CHECK (event_data <> '')
		);

		CREATE TABLE IF NOT EXISTS cursors (
			cursor_id INTEGER PRIMARY KEY AUTOINCREMENT,
			cursor_name TEXT UNIQUE NOT NULL CHECK (cursor_name <> ''),
			event_type TEXT NOT NULL CHECK (event_type <> ''),
			last_event_id INTEGER NOT NULL DEFAULT 0
		);
	`

	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}

	return &SqliteStore{
		db: db,
	}
}
