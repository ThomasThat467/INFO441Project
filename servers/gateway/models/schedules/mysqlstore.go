package schedules

import "database/sql"

// GetByType is an enumerate for GetBy* functions implemented
// by MySQLStore structs
type GetByType string

// These are the enumerates for GetByType
const (
	ID               GetByType = "ID"
	WateringSchedule GetByType = "WateringSchedule"
)

// MySQLStore is a schedule.Store backed by MySQL
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
func (ms *MySQLStore) getByProvidedType(t GetByType, arg interface{}) (*Schedule, error) {
	sel := string("select ID, WateringSchedule from Schedule where " + t + " = ?")

	rows, err := ms.Database.Query(sel, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schedule := &Schedule{}

	// Should never have more than one row, so only grab one
	rows.Next()
	if err := rows.Scan(
		&schedule.ID,
		&schedule.WateringSchedule); err != nil {
		return nil, err
	}
	return schedule, nil
}

//GetByID returns the User with the given ID
func (ms *MySQLStore) GetByID(id int64) (*Schedule, error) {
	return ms.getByProvidedType(ID, id)
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(schedule *Schedule) (*Schedule, error) {
	ins := "insert into Plants(Species, WateringSchedule, PhotoURL) values (?,?,?)"
	res, err := ms.Database.Exec(ins, schedule.WateringSchedule)
	if err != nil {
		return nil, err
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	schedule.ID = lid
	return schedule, nil
}

// //Update applies ScheduleUpdates to the given schedule ID
// //and returns the newly-updated schedule
func (ms *MySQLStore) Update(id int64, updates *Updates) (*Schedule, error) {
	// Assumes updates ALWAYS includes WateringSchedule and PhotoURL
	upd := "update Schedule set WateringSchedule = ?"
	res, err := ms.Database.Exec(upd, updates.WateringSchedule, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, rowsAffectedErr := res.RowsAffected()
	if rowsAffectedErr != nil {
		return nil, rowsAffectedErr
	}

	if rowsAffected != 1 {
		return nil, ErrScheduleNotFound
	}

	// Get the user using GetByID
	return ms.GetByID(id)
}
