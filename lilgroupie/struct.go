package lilGroupie

// type GroupieData struct {
// 	artists   string `json:"artists"`
// 	Locations string `json:"locations"`
// 	Dates     string `json:"dates"`
// 	Relation  string `json:"relation"`
// }

var GroupieDataArr []artistsAPI

// var generalData []GroupieData

type DataAPI struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type GroupieData struct {
	ID        int    `json:"id"`
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type index1 struct {
	Index []index `json:"index"`
}

type index struct {
	ID       int   `json:"id"`
	Relations interface{} `json:"datesLocations"`
}

// type Loc struct {
// 	dates []Date
// }

// type DateAPI2 struct {
// 	ID    int    `json:"id"`
// 	Dates string `json:"dates"`
// }

type DateAPI struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type artistsAPI struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations   interface{}   `json:"relations"`
}

type LocationsAPI struct {
	ID int `json:"id"`
	// Locations []string `json:"locations"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
