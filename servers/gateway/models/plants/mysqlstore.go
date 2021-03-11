package plants

import (
	"database/sql"
)

// GetByType is an enumerate for GetBy* functions implemented
// by MySQLStore structs
type GetByType string

// These are the enumerates for GetByType
const (
	ID        GetByType = "ID"
	PlantName GetByType = "PlantName"
)

// MySQLStore is a plant.Store backed by MySQL
type MySQLStore struct {
	Database *sql.DB
}

// NewMySQLStore constructs a new MySQLStore, and returns an error
// if there is a problem along the way.
func NewMySQLStore(dataSourceName string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &MySQLStore{db}, nil
}

// getByProvidedType gets a specific user given the provided type.
// This requires the GetByType to be "unique" in the database.
func (ms *MySQLStore) getByProvidedType(t GetByType, arg interface{}) (*Plant, error) {
	sel := string("select ID, UserID, PlantName, WateringSchedule, LastWatered, PhotoURL from Plants where " + t + " = ?")

	rows, err := ms.Database.Query(sel, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plant := &Plant{}

	// Should never have more than one row, so only grab one
	rows.Next()
	if err := rows.Scan(
		&plant.ID,
		&plant.UserID,
		&plant.PlantName,
		&plant.WateringSchedule,
		&plant.LastWatered,
		&plant.PhotoURL); err != nil {
		return nil, err
	}
	return plant, nil
}

//GetByID returns the User with the given ID
func (ms *MySQLStore) GetByID(id int64) (*Plant, error) {
	return ms.getByProvidedType(ID, id)
}

//GetBySpecies returns the Plants of the given Species
func (ms *MySQLStore) GetByPlantName(plantName string) (*Plant, error) {
	return ms.getByProvidedType(PlantName, plantName)
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(plant *Plant) (*Plant, error) {
	ins := "insert into Plants(Species, WateringSchedule, PhotoURL) values (?,?,?)"
	res, err := ms.Database.Exec(ins, plant.PlantName, plant.WateringSchedule, plant.PhotoURL)
	if err != nil {
		return nil, err
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	plant.ID = lid
	return plant, nil
}

// //Update applies PlantUpdates to the given plant ID
// //and returns the newly-updated plant
func (ms *MySQLStore) Update(id int64, updates *Updates) (*Plant, error) {
	// Assumes updates ALWAYS includes WateringSchedule and PhotoURL
	upd := "update Plants set WateringSchedule = ?, LastWatered = ?, PhotoURL = ? where ID = ?"
	res, err := ms.Database.Exec(upd, updates.WateringSchedule, updates.LastWatered, updates.PhotoURL, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, rowsAffectedErr := res.RowsAffected()
	if rowsAffectedErr != nil {
		return nil, rowsAffectedErr
	}

	if rowsAffected != 1 {
		return nil, ErrPlantNotFound
	}

	// Get the user using GetByID
	return ms.GetByID(id)
}

// Delete deletes the plant with the given ID
func (ms *MySQLStore) Delete(id int64) error {
	del := "delete from Plants where ID = ?"
	res, err := ms.Database.Exec(del, id)
	if err != nil {
		return err
	}

	rowsAffected, rowsAffectedErr := res.RowsAffected()
	if rowsAffectedErr != nil {
		return rowsAffectedErr
	}

	if rowsAffected != 1 {
		return ErrPlantNotFound
	}

	return nil
}
