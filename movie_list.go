package cinema_diary

type MoviesList struct {
	UserId    int  `json:"u_id" db:"u_id"`
	MovieId   int  `json:"m_id" db:"m_id"`
	IsWatched bool `json:"is_watched" db:"is_watched"`
	IsLiked   bool `json:"is_liked" db:"is_liked"`
	Score     int  `json:"score" db:"score"`
}
