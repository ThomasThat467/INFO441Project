package plants

import (
	"errors"
)

//ErrUserNotFound is returned when the user can't be found
var ErrPlantNotFound = errors.New("plant not found")

//Store represents a store for Users
type Store interface {
	//GetByID returns the Plant with the given ID
	GetByID(id int64) (*Plant, error)

	//GetBySpecies returns the Plants of the given Species
	GetBySpecies(species string) (*Plant, error)

	//Insert inserts the plant into the database, and returns
	//the newly-inserted Plant, complete with the DBMS-assigned ID
	Insert(plant *Plant) (*Plant, error)

	//Update applies PlantUpdates to the given plant ID
	//and returns the newly-updated plant
	Update(id int64, updates *Updates) (*Plant, error)

	//Delete deletes the plant with the given ID
	//Delete(id int64) error

	// Jisu changes to this
	Delete(id int64) (*Plant, error)
}
