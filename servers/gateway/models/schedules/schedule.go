package schedules

//gravatarBasePhotoURL is the base URL for Gravatar image requests.
//See https://id.gravatar.com/site/implement/images/ for details
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

type Schedule struct {
	ID               int64  `json:"id"`
	WateringSchedule string `json:"photoURL"`
}

//Updates represents allowed updates to a schedule
type Updates struct {
	WateringSchedule string `json:"photoURL"`
}

//NewPlant represents a new plant that added by a user
type NewSchedule struct {
	WateringSchedule string `json:"photoURL"`
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (ns *NewSchedule) ToSchedule() (*Schedule, error) {

	newSchedule := &Schedule{
		WateringSchedule: ns.WateringSchedule,
	}

	//GetGravitar(newSchedule)
	return newSchedule, nil
}

// GetGravitar calculates the gravitar hash based on the string given and
// stores it for the plant  -- copy from user.go
// not sure for schedule
/* func GetGravitar(user *Schedule) {
	photoURLHash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(str))))
	photoURLHashString := hex.EncodeToString(photoURLHash[:])
	user.PhotoURL = gravatarBasePhotoURL + photoURLHashString
} */
