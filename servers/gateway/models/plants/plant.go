package plants

//gravatarBasePhotoURL is the base URL for Gravatar image requests.
//See https://id.gravatar.com/site/implement/images/ for details
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

// PlantInventory ...
type PlantInventory struct {
	Plants []Plant `json:"Plants"`
}

// Plant ...
type Plant struct {
	ID               int64  `json:"id"`
	UserID           int64  `json:"userId"`
	PlantName        string `json:"plantName"`
	WateringSchedule string `json:"wateringSchedule"`
	LastWatered      string `json:"lastWatered"`
	PhotoURL         string `json:"photoURL"`
}

//Updates represents allowed updates to a plant
type Updates struct {
	WateringSchedule string   `json:"schedule"`
	LastWatered      string   `json:"lastWatered"`
	PhotoURL         string   `json:"photoURL"`
}

//NewPlant represents a new plant that added by a user
type NewPlant struct {
	PlantName        string `json:"plantName"`
	WateringSchedule string `json:"wateringSchedule"`
	LastWatered      string `json:"lastWatered"`
	PhotoURL         string `json:"photoURL"`
}

//ToPlant converts the NewPlant to a Plant
func (np *NewPlant) ToPlant(userID int64) (*Plant, error) {

	newPlant := &Plant{
		UserID:           userID,
		PlantName:        np.PlantName,
		WateringSchedule: np.WateringSchedule,
		LastWatered:      np.LastWatered,
		PhotoURL:         np.PhotoURL,
	}
	return newPlant, nil
}
