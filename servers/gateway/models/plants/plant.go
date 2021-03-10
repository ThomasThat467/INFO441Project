package plants

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//gravatarBasePhotoURL is the base URL for Gravatar image requests.
//See https://id.gravatar.com/site/implement/images/ for details
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

type Plant struct {
	ID               int64  `json:"id"`
	Species          string `json:"species"`
	WateringSchedule string `json:"schedule"`
	PhotoURL         string `json:"photoURL"`
}

//Updates represents allowed updates to a plant
type Updates struct {
	WateringSchedule string `json:"schedule"`
	PhotoURL         string `json:"photoURL"`
}

//NewPlant represents a new plant that added by a user
type NewPlant struct {
	Species          string `json:"species"`
	WateringSchedule string `json:"schedule"`
	PhotoURL         string `json:"photoURL"`
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (np *NewPlant) ToPlant() (*Plant, error) {

	newPlant := &Plant{
		Species:          np.Species,
		WateringSchedule: np.WateringSchedule,
	}

	GetGravitar(newPlant, np.Species)
	return newPlant, nil
}

// GetGravitar calculates the gravitar hash based on the string given and
// stores it for the plant  -- copy from user.go
func GetGravitar(plant *Plant, str string) {
	photoURLHash := md5.Sum([]byte(strings.ToLower(strings.TrimSpace(str))))
	photoURLHashString := hex.EncodeToString(photoURLHash[:])
	plant.PhotoURL = gravatarBasePhotoURL + photoURLHashString
}
