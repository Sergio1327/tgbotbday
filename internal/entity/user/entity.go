package user

import "time"

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	BirthDate time.Time `db:"bday"`
}

func (u *User) SetNextYear() {
	u.BirthDate = u.BirthDate.AddDate(1, 0, 0)
}
