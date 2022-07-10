package data

type IGamesRepository interface {
	GetAllGames() ([]Game, error)
	GetGameByID(id int) (Game, error)
	GetAllGamesByPublisher(publisher string) ([]Game, error)
	AddNewGame(game *Game) error
	EditGame(id int, updatedGame *Game) error
	RemoveGame(id int) error
}
