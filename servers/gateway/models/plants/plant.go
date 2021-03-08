package plants

type Plant struct {
	ID               int64  `json:"id"`
	Species          string `json:"userName"`
	WateringSchedule string `json:"photoURL"`
	PhotoURL         string `json:"photoURL"`
}

//Updates represents allowed updates to a plant
type Updates struct {
	WateringSchedule string `json:"photoURL"`
	PhotoURL         string `json:"photoURL"`
}
