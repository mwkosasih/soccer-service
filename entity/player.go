package entity

type Player struct {
	ID        int32  `json:"id" db:"id"`
	TeamID    int32  `json:"team_id" db:"team_id"`
	Name      string `json:"name" db:"name"`
	Position  string `json:"position" db:"position"`
	Height    int32  `json:"height" db:"height"`
	Weight    int32  `json:"weight" db:"weight"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
