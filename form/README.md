# utils

SetDefault
将url.Values 对应转换到struct


`
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
	Operator  string    `gorm:"-" json:"operator"`
	Third
}
type Client struct {
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
}`

	   `a := Client{}

	ss := url.Values{

		"Operator": []string{"zhangsanfeng"},
		"UserCode": []string{"wangwu"},
		"Name":     []string{"lisilisi"},
	}
	err := SetDefault(ss, &a)`
