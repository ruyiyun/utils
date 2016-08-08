package form

import (
	"errors"
	"net/url"
	"reflect"
	"time"
)

const (
	blank              = ""
	namespaceSeparator = '.'
	ignore             = "-"
	fieldNS            = "Field Namespace:"
	errorText          = " ERROR:"
)

var (
	timeType = reflect.TypeOf(time.Time{})
)

//自定义类型处理函数
type DecodeCustomTypeFunc func([]string) (interface{}, error)

type key struct {
	ivalue      int
	value       string
	searchValue string
}

type recursiveData struct {
	alias    string
	sliceLen int
	keys     []key
}

func NewDecoder() *decoder {

	d := &decoder{
		tagName: "form",
	}

	return d
}

//设置model的tag字段值,默认为form
func (d *decoder) SetTagName(tagName string) {
	d.tagName = tagName
}

//注册自定义类型的转换
//decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
//	return time.Parse("2006-01-02", vals[0])
//}, time.Time{})
func (d *decoder) RegisterCustomTypeFunc(fn DecodeCustomTypeFunc, types ...interface{}) {

	if d.customTypeFuncs == nil {
		d.customTypeFuncs = map[reflect.Type]DecodeCustomTypeFunc{}
	}

	for _, t := range types {
		d.customTypeFuncs[reflect.TypeOf(t)] = fn
	}
}

//将url.Values值转换为struct
func (d *decoder) Decode(values url.Values, obj interface{}) (err error) {

	if len(values) == 0 {
		return errors.New("url.Values 不能为空")
	}
	d.values = values

	val := reflect.ValueOf(obj)

	//检测是否空,是否指针类型
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return &InvalidDecoderError{reflect.TypeOf(obj)}
	}

	val = val.Elem()
	typ := val.Type()

	if val.Kind() == reflect.Struct && typ != timeType {
		d.traverseStruct(val, typ, d.namespace[0:0])
	} else {
		d.setFieldByType(val, d.namespace[0:0], 0)
	}

	if len(d.errs) > 0 {
		err = d.errs
	}

	return
}
