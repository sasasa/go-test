package user

import(
	"time"
)

type User struct {
	Id int
	Name string
	Email string
	Created time.Time
}

func NewUser() (*User) {
	/* User 構造 体 の 初期化 */
	u := new(User)
	u.Id = 1
	u.Name = "山田 太郎"
	u.Email = "yamada@example.com"
	u.Created = time.Now()
	return u
}