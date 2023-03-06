package cinema_diary

type Movie struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}
