package infrastructure

import (
	"database/sql"
	"sensor/src/domain/entities"
	"sensor/src/domain/repositories"
)

type MySQLEventRepository struct {
	db *sql.DB
}

func NewMySQLEventRepository(db *sql.DB) repositories.EventRepository {
	return &MySQLEventRepository{db: db}
}

func (repo *MySQLEventRepository) Create(event *entities.Event) error {
	query := "INSERT INTO events (zone) VALUES (?)"
	_, err := repo.db.Exec(query, event.Zone)
	return err
}

func (repo *MySQLEventRepository) GetAll() ([]entities.Event, error) {
	query := "SELECT id, zone, detected_at FROM events"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []entities.Event
	for rows.Next() {
		var event entities.Event
		if err := rows.Scan(&event.ID, &event.Zone, &event.DetectedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (repo *MySQLEventRepository) GetByID(id int) (*entities.Event, error) {
	query := "SELECT id, zone, detected_at FROM events WHERE id = ?"
	var event entities.Event
	err := repo.db.QueryRow(query, id).Scan(&event.ID, &event.Zone, &event.DetectedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}
