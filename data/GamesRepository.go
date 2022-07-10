package data

import (
	"database/sql"
	"time"
)

type GamesSQLRepository struct {
	db *sql.DB
}

func NewGamesSQLRepository(db *sql.DB) *GamesSQLRepository {
	if db == nil {
		panic("Missing db")
	}

	return &GamesSQLRepository{db: db}
}

func (r GamesSQLRepository) GetAllGames() ([]Game, error) {
	//Get number of rows
	count := 0
	err := r.db.QueryRow("SELECT count(*) FROM Games").Scan(&count)
	if err != nil {
		return nil, err
	}

	games := make([]Game, 0, count)
	rows, err := r.db.Query("SELECT * FROM Games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var name string
	var publisher string
	var releaseDate time.Time
	var description string

	for rows.Next() {
		rows.Scan(&id, &name, &publisher, &releaseDate, &description)
		game := Game{
			Name:        name,
			Publisher:   publisher,
			ReleaseDate: releaseDate,
			Description: description}
		games = append(games, game)
	}
	err = rows.Err()
	return games, err
}

func (r GamesSQLRepository) GetGameByID(id int) (Game, error) {
	row := r.db.QueryRow("SELECT [name], [publisher], [releaseDate], [description] FROM Games WHERE [id] = ?", id)
	if row.Err() != nil {
		return Game{}, row.Err()
	}
	var name string
	var publisher string
	var releaseDate time.Time
	var description string
	row.Scan(&name, &publisher, &releaseDate, &description)
	return Game{
		Name:        name,
		Publisher:   publisher,
		ReleaseDate: releaseDate,
		Description: description}, row.Err()
}

func (r GamesSQLRepository) GetAllGamesByPublisher(publisher string) ([]Game, error) {
	//Get number of rows
	count := 0
	err := r.db.QueryRow("SELECT count(*) FROM [Games] WHERE [publisher] = ?", publisher).Scan(&count)
	if err != nil {
		return nil, err
	}

	games := make([]Game, 0, count)
	rows, err := r.db.Query("SELECT [name], [releaseDate], [description] FROM Games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var name string
	var releaseDate time.Time
	var description string

	for rows.Next() {
		rows.Scan(&name, &releaseDate, &description)
		games = append(games, Game{
			Name:        name,
			Publisher:   publisher,
			ReleaseDate: releaseDate,
			Description: description})
	}
	err = rows.Err()
	return games, err
}

func (r GamesSQLRepository) AddNewGame(game *Game) error {
	stmt, err := r.db.Prepare("INSERT INTO [Games](name, publisher, releaseDate, description) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(game.Name, game.Publisher, game.ReleaseDate, game.Description)
	return err
}

func (r GamesSQLRepository) EditGame(id int, updatedGame *Game) error {
	stmt, err := r.db.Prepare("UPDATE [Games] SET name = ?, publisher = ?, releaseDate = ?, description = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(updatedGame.Name, updatedGame.Publisher, updatedGame.ReleaseDate, updatedGame.Description, id)
	return err
}

func (r GamesSQLRepository) RemoveGame(id int) error {
	_, err := r.db.Exec("DELETE FROM [Games] WHERE [id] = ?)", id)
	return err
}
