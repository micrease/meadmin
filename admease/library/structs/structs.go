package structs

import (
	"errors"
	"giftcard/library/datetime"
	"reflect"
)

//结构体copy只能copy相同属性的值,不同类型需要手动转换赋值
func Copy(src, dst interface{}) error {
	srcV, err := srcFilter(src)
	if err != nil {
		return err
	}
	dstV, err := dstFilter(dst)
	if err != nil {
		return err
	}
	srcKeys := make(map[string]bool)
	for i := 0; i < srcV.NumField(); i++ {
		srcKeys[srcV.Type().Field(i).Name] = true
	}
	for i := 0; i < dstV.Elem().NumField(); i++ {
		fName := dstV.Elem().Type().Field(i).Name
		if _, ok := srcKeys[fName]; ok {
			v := srcV.FieldByName(dstV.Elem().Type().Field(i).Name)
			if v.CanInterface() {
				if v.Type().String() == dstV.Elem().Field(i).Type().String() {
					//只能copy相同类型的值
					dstV.Elem().Field(i).Set(v)
				} else {
					//constraints.Integer()
					//fmt.Println("src.type.string != dest.type.string", v.Type().String(), dstV.Elem().Field(i).Type().String())
				}
			}
		}
	}
	return nil
}

func ToMap(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if value, ok := f.Tag.Lookup("json"); ok {
			if "datetime.DBTime" == f.Type.String() {
				v, ok := v.Field(i).Interface().(datetime.DBTime)
				if ok {
					m[value] = v.String()
				}
			} else {
				m[value] = v.Field(i).Interface()
			}
		}
	}
	return m
}

func srcFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not a struct or a pointer to struct")
	}
	return v, nil
}

func dstFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() != reflect.Ptr {
		return reflect.Zero(v.Type()), errors.New("src type error: not a pointer to struct")
	}
	if v.Elem().Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not point to struct")
	}
	return v, nil
}
