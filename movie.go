package cinema_diary

type Movie struct {
	Id       int    `json:"id" db:"m_id"`
	Name     string `json:"name" db:"m_name"`
	Director string `json:"director" db:"director"`
	Year     int    `json:"year" db:"year"`
}
