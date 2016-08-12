package form

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

func SetDefault(ss url.Values, obj interface{}) error {
	vt := reflect.TypeOf(obj)
	vv := reflect.ValueOf(obj)

	return setDefault(vv, vt, ss)
}

func setDefault(v reflect.Value, t reflect.Type, ss url.Values) error {
	fmt.Println("\nv.CanSet():", v.CanSet(), "v.IsValid():", v.IsValid())
	if t.Kind() != reflect.Ptr && t.Kind() != reflect.Struct {
		fmt.Println("不是struct,也不是ptr")
		return errors.New("不是struct,也不是ptr")
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	v = reflect.Indirect(v)

	fmt.Println("\nv.CanSet():", v.CanSet(), "v.IsValid():", v.IsValid())
	fmt.Println("v的Type", v.Type())
	fmt.Println("v的kind", v.Type().Kind())
	fmt.Printf("struct%s共有%d个字段\n", t.Name(), t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fieldV := v.Field(i)
		fieldT := t.Field(i)

		fmt.Println(fieldV.Type())
		if fieldV.Kind() == reflect.Ptr {
			//todo fmt.Println("如何处理")
			continue
		}
		fieldV = reflect.Indirect(fieldV)

		fmt.Println("\n本轮", fieldT.Name, fieldV.Type(), fieldT.Type, fieldV.Kind())
		fmt.Println("\nfieldV.CanSet():", fieldV.CanSet(), "fieldV.IsValid():", fieldV.IsValid())

		if !fieldV.CanSet() {
			continue
		}

		var value string

		switch fieldV.Kind() {
		case reflect.Bool:
			value = ss.Get(fieldT.Name)
			if len(value) == 0 {
				continue
			}

			a, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldV.SetBool(a)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = ss.Get(fieldT.Name)
			if len(value) == 0 {
				continue
			}

			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = ss.Get(fieldT.Name)
			if len(value) == 0 {
				continue
			}

			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			value = ss.Get(fieldT.Name)
			if len(value) == 0 {
				continue
			}
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			value = ss.Get(fieldT.Name)
			if len(value) == 0 {
				continue
			}
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			value = ss.Get(fieldT.Name)
			if len(value) == 0 {
				continue
			}
			fieldV.SetString(value)
		case reflect.Struct:
			if fieldV.Type().String() == "time.Time" {
				value = ss.Get(fieldT.Name)
				if len(value) == 0 {
					continue
				}
				format := time.RFC3339
				t, err := time.ParseInLocation(format, value, time.Local)
				if err != nil {
					return err
				}
				fieldV.Set(reflect.ValueOf(t))

			} else {
				fmt.Println("进入下一轮")
				setDefault(fieldV.Addr(), fieldV.Type(), ss)
			}
		case reflect.Ptr:
			fmt.Println("todo怎么处理?什么情况能到这里?")

		default:

			fmt.Printf("Unsupported kind: %v\n", v.Type().String())

		}
	}

	return nil

}
