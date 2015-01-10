package user

import (
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
	store "github.com/gabesullice/foos/storage"
)

type User struct {
	Id        string `gorethink:"id,omitempty" json:"id"`
	Mail      string `gorethink:"mail" json:"mail"`
	FirstName string `gorethink:"firstName" json:"firstName"`
	LastName  string `gorethink:"lastName" json:"lastName"`

	Created int64 `gorethink:"created" json:"created"`
	Updated int64 `gorethink:"updated" json:"updated"`
}

func GetUser(name string, s store.Session) (User, error) {
	var user User

	res, err := r.Db("foos").Table("users").GetAllByIndex("name", name).Run(s)

	if err != nil {
		return user, err
	}

	err = res.One(&user)

	return user, err
}

func GetUsers(s store.Session) ([]User, error) {
	var users []User

	res, err := r.Db("foos").Table("users").Run(s)

	if err != nil {
		return users, err
	}

	err = res.All(&users)

	return users, err
}

func (u *User) Save(s store.Session) error {
	if u.Created == 0 {
		u.Created = time.Now().Unix()
		res, err := r.Db("foos").Table("users").Insert(u).RunWrite(s)
		u.Id = res.GeneratedKeys[0]
		if err != nil {
			return err
		}
	} else {
		u.Updated = time.Now().Unix()
		_, err := r.Db("foos").Table("users").Get(u.Id).Update(u).Run(s)

		if err != nil {
			return err
		}
	}
	return nil
}

func (u User) Check(s store.Session) error {
	res, err := r.Db("foos").Table("users").Field("mail").Count(u.Mail).Run(s)

	if err != nil {
		return err
	}

	var count int

	if err = res.One(&count); err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("The given email, %s, is not unique.", u.Mail)
	}

	return nil
}
