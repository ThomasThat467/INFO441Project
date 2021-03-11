package plants

import (
	"errors"
)

//ErrPlantNotFound is returned when the user can't be found
var ErrPlantNotFound = errors.New("plant not found")

//Store represents a store for Plants
type Store interface {
	//GetByID returns the Plant with the given ID
	GetByID(id int64) (*Plant, error)

	//GetByPlantName returns the Plants of the given PlantName
	GetByPlantName(plantName string) (*Plant, error)

	//Insert inserts the plant into the database, and returns
	//the newly-inserted Plant, complete with the DBMS-assigned ID
	Insert(plant *Plant) (*Plant, error)

	//Update applies PlantUpdates to the given plant ID
	//and returns the newly-updated plant
	Update(id int64, updates *Updates) (*Plant, error)

	//Delete deletes the plant with the given ID
	//Delete(id int64) error
	Delete(id int64) error
}
