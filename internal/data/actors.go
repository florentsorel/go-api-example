package data

import (
	"database/sql"
	"errors"
)

type Actor struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Active         Bool     `json:"active"`
	CreationDate   Time     `json:"created_at"`
	LastUpdateDate NullTime `json:"updated_at"`
}

type ActorModel struct {
	DB *sql.DB
}

func (m ActorModel) Insert(actor *Actor) error {
	return nil
}

func (m ActorModel) Get(id int64) (*Actor, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT idActor, name, isActive, creationDate, lastUpdateDate
        FROM Actor
    WHERE idActor = ?`

	var actor Actor

	err := m.DB.QueryRow(query, id).Scan(
		&actor.ID,
		&actor.Name,
		&actor.Active,
		&actor.CreationDate,
		&actor.LastUpdateDate,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err

		}
	}
	return &actor, nil
}

func (m ActorModel) Update(actor *Actor) error {
	return nil
}

func (m ActorModel) Delete(id int64) error {
	return nil
}
