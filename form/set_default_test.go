package form

import (
	"net/url"
	"testing"
	"time"
)

type Third struct {
	Name string
	Aget uint
}

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `sql:"index"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	DeletedBy string    `json:"deleted_by"`
	Operator  string    `gorm:"-" json:"operator"` //记录本次操作的user_code
	Third
}
type Client struct {
	//todo uint time 不可以
	ID        uint   `gorm:"primary_key"`
	UserCode  string `json:"user_code"` //who
	Operation string `json:"operation"` //did what
	Result    string `json:"result"`
	M         string `json:"m"`
	C         string `json:"c"`
	A         string `json:"a"`
	P         string `json:"p" gorm:"type:TEXT"`
	IP        string `json:"ip"`
	BaseModel
}

func TestSetDefault(t *testing.T) {
	a := Client{}

	ss := url.Values{

		"Operator": []string{"zhangsanfeng"},
		"UserCode": []string{"wangwu"},
		"Name":     []string{"lisilisi"},
	}
	err := SetDefault(ss, &a)
	if err != nil {
		t.Fatal(err)
	}
	if a.Operator != "zhangsanfeng" {
		t.Fail()
	}
	if a.UserCode != "wangwu" {
		t.Fail()
	}
	if a.Name != "lisilisi" {
		t.Fail()
	}
}
