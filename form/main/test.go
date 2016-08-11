package main

import (
	"reflect"
	"fmt"
	"time"
)

var
	timeType = reflect.TypeOf(time.Time{})

type Third struct {
	Name string
	Aget uint
}

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
	Operator  string `gorm:"-" json:"operator"` //记录本次操作的user_code
	Third
}
type Client struct {
	                                    //todo uint time 不可以
	ID        uint `gorm:"primary_key"`
	UserCode  string `json:"user_code"` //who
	Operation string `json:"operation"` //did what
	Result    string  `json:"result"`
	M         string `json:"m"`
	C         string `json:"c"`
	A         string `json:"a"`
	P         string `json:"p" gorm:"type:TEXT"`
	IP        string `json:"ip"`
	BaseModel
}

func SetDefault1(v reflect.Value, t reflect.Type) error {
	//
	//t := reflect.TypeOf(s)
	//v := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		fmt.Println("1是指针")

		t = t.Elem()
		v = v.Elem()

	}
	v = reflect.Indirect(v)



	//fmt.Println("shuchu", v.String(), v.Type(), v.Kind(), t.Field(0).Name(), t.String())
	switch v.Kind() {
	case reflect.Uint:
		fmt.Println("")
	case reflect.Int:
		v.SetInt(42)
	case reflect.String:
		v.SetString("Foo")
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Struct:
		if v.Type().String() == "time.Time" {
			fmt.Println("111时间设置初始值")
		} else {

			fmt.Printf("struct%s共有%d个字段\n", t.Name(), t.NumField())
			for i := 0; i < t.NumField(); i++ {
				fieldV := v.Field(i)
				fieldT := t.Field(i)
				fmt.Println("下一轮", fieldT.Name, fieldV.Type())
				SetDefault1(fieldV.Addr(), fieldV.Type())
			}

		}
	default:
		fmt.Println("Unsupported kind: " + v.Kind().String())
	}

	return nil

}
func SetDefault(v reflect.Value, t reflect.Type) error {

	if t.Kind() == reflect.Ptr {
		fmt.Println("1是指针")

		t = t.Elem()
		v = v.Elem()

	}
	v = reflect.Indirect(v)

	if v.Kind() == reflect.Struct {
		fmt.Printf("struct%s共有%d个字段\n", t.Name(), t.NumField())

		for i := 0; i < t.NumField(); i++ {
			fieldV := v.Field(i)
			fieldT := t.Field(i)
			fmt.Println("本轮", fieldT.Name, fieldV.Type())
			fieldV = reflect.Indirect(fieldV)

			switch fieldV.Kind() {
			case reflect.Uint:
				fmt.Println("unit设置0")
			case reflect.Int:
				fieldV.SetInt(42)
			case reflect.String:
				fieldV.SetString("Foo")
			case reflect.Bool:
				fieldV.SetBool(true)
			case reflect.Struct:
				if fieldV.Type().String() == "time.Time" {
					fmt.Println("111时间设置初始值")
				} else {
					fmt.Println("进入下一轮")
					SetDefault(fieldV.Addr(), fieldV.Type())

				}
			default:
				fmt.Println("Unsupported kind: " + v.Kind().String())
			}
		}

	}
	return nil

}

func main() {
	a := Client{}

	t := reflect.TypeOf(&a)
	v := reflect.ValueOf(&a)
	err := SetDefault(v, t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)

}
