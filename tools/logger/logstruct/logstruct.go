package logstruct

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

func Fields(str interface{}) logrus.Fields {
	lf := logrus.Fields{}

	v := reflect.ValueOf(str)
	var t reflect.Type

	// если это поинтер
	if v.Kind() == reflect.Pointer {
		// получить значение
		v = reflect.Indirect(v)
		t = v.Type()
	} else {
		t = reflect.TypeOf(str)
	}

	// не структуры не обрабатываются
	if v.Kind() != reflect.Struct {
		return lf
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.FieldByName(field.Name)
		kind := value.Kind()

		if kind == reflect.Struct {
			templf := Fields(value.Interface())
			for k, v := range templf {
				lf[k] = v
			}
			continue
		}

		// считываются только тэги log
		tag, ok := field.Tag.Lookup("log")
		if ok && tag != "-" {
			value := v.FieldByName(field.Name)
			kind := value.Kind()

			// если филд структуры: другая структура или поинтер - пропускаются
			if kind == reflect.Pointer {
				continue
			}

			lf[tag] = value.Interface()
		}

	}

	return lf
}
