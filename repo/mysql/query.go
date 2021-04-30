package mysql

const (
	QueryCreateTeam = `
		INSERT INTO teams (name)
		VALUES (?)
	`

	QueryCreatePlayer = `
		INSERT INTO players (team_id, name, position, height, weight)
		VALUES (?, ?, ?, ?, ?)
	`

	QuerySelectTeamByID = `
		SELECT * FROM teams WHERE id = ?
	`

	QuerySelectPlayerByID = `
		SELECT * FROM players WHERE id = ?
	`

	QuerySelectPlayerByTeamID = `
		SELECT * FROM players WHERE team_id = ?
	`
)
