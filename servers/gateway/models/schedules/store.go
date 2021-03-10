package schedules

import "errors"

//ErrScheduleNotFound is returned when the user can't be found
var ErrScheduleNotFound = errors.New("schedule not found")

//Store represents a store for Plants
type Store interface {
	//GetByID returns the Schedule with the given ID
	GetByID(id int64) (*Schedule, error)

	//Insert inserts the Schedule into the database, and returns
	//the newly-inserted Schedule, complete with the DBMS-assigned ID
	Insert(schedule *Schedule) (*Schedule, error)

	//Update applies ScheduleUpdates to the given schedule ID
	//and returns the newly-updated plant
	Update(id int64, updates *Updates) (*Schedule, error)
}
