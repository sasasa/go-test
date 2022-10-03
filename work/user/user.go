package user

import(
	"time"
	"encoding/json"
	. "local.packages/utility"
)


type encodable interface{
	Encode() (string, error)
}
type User struct {
	Id int
	Name string
	Email string
	Created time.Time
}

func (user *User) Encode() (string, error) {
		/* JSON エン コード */
		bs, err := json.Marshal(user)
		return string(bs), err
}

func NewUser(dict Dict) (encodable) {
	/* User 構造 体 の 初期化 */
	u := new(User)
	u.Id = dict["Id"].(int)
	u.Name = dict["Name"].(string)
	u.Email = dict["Email"].(string)
	u.Created = time.Now()
	return u
}