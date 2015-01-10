package game

import (
	"fmt"

	r "github.com/dancannon/gorethink"
	store "github.com/gabesullice/foos/storage"
)

type Game struct {
	Id string `gorethink:"id,omitempty" json:"id"`

	TeamOne Team `gorethink:"players" json:"players"`
	TeamTwo Team `gorethink:"players" json:"players"`

	Status string `gorethink:"status" json:"status"`

	// Started  int64 `gorethink:"started" json:"started"`
	// Finished int64 `gorethink:"finished" json:"finished"`

	Created int64 `gorethink:"created" json:"created"`
	Updated int64 `gorethink:"updated" json:"updated"`
}

type Team struct {
	PlayerOne, PlayerTwo Player
}

type Player struct {
	User     string
	OffGoals uint8
	DefGoals uint8
	OwnGoals uint8
}

func GetGame(id string, s store.Session) (Game, error) {
	var game Game

	res, err := r.Db("foos").Table("games").GetAllByIndex("id", id).Run(s)

	if err != nil {
		return game, err
	}

	err = res.One(&game)

	return game, err
}

func GetGames(s store.Session) ([]Game, error) {
	var games []Game

	res, err := r.Db("foos").Table("games").Run(s)

	if err != nil {
		return games, err
	}

	err = res.All(&games)

	return games, err
}

func (g *Game) Save(s store.Session) error {
	if g.Status == "" {
		g.Status = "inProgress"
		res, err := r.Db("foos").Table("games").Insert(g).RunWrite(s)
		g.Id = res.GeneratedKeys[0]
		if err != nil {
			return err
		}
	} else {
		_, err := r.Db("foos").Table("games").Get(g.Id).Update(g).Run(s)

		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Delete(s store.Session) error {
	if g.Id == "" {
		return fmt.Errorf("The game could not be deleted without an id.")
	}

	_, err := r.Db("foos").Table("games").Get(g.Id).Delete().Run(s)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Finish(s store.Session) error {
	g.Status = "finished"
	return g.Save(s)
}
