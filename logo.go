package logo

import (
	"reflect"
	"strings"

	"github.com/naxmefy/logo/utils"
)

func Index(vs interface{}, t interface{}) (index int) {
	vsVal := reflect.ValueOf(vs)
	vsKind := reflect.TypeOf(vs).Kind()

	defer func() {
		if e := recover(); e != nil {
			index = -1
		}
	}()

	if vsKind == reflect.String {
		tVal := reflect.ValueOf(t)
		return strings.Index(vsVal.String(), tVal.String())
	}

	if vsKind == reflect.Map {
		mapKeys := vsVal.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if utils.ObjectsAreEqual(mapKeys[i].Interface(), t) {
				return i
			}
		}
		return -1
	}

	for i := 0; i < vsVal.Len(); i++ {
		if utils.ObjectsAreEqual(vsVal.Index(i).Interface(), t) {
			return i
		}
	}

	return -1
}

func Include(vs interface{}, t interface{}) bool {
	return Index(vs, t) >= 0
}

func Any(vs interface{}, f func(v interface{}) bool) (any bool) {
	vsVal := reflect.ValueOf(vs)
	vsKind := reflect.TypeOf(vs).Kind()

	defer func() {
		if e := recover(); e != nil {
			any = false
		}
	}()

	if vsKind == reflect.String {
		for _, v := range vsVal.String() {
			if f(string(v)) {
				return true
			}
		}

		return false
	}

	if vsKind == reflect.Map {
		mapKeys := vsVal.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if f(mapKeys[i].Interface()) {
				return true
			}
		}
		return false
	}

	for i := 0; i < vsVal.Len(); i++ {
		if f(vsVal.Index(i).Interface()) {
			return true
		}
	}

	return false
}

func All(vs interface{}, f func(v interface{}) bool) (all bool) {
	vsVal := reflect.ValueOf(vs)
	vsKind := reflect.TypeOf(vs).Kind()

	defer func() {
		if e := recover(); e != nil {
			all = false
		}
	}()

	if vsKind == reflect.String {
		for _, v := range vsVal.String() {
			if !f(string(v)) {
				return false
			}
		}

		return true
	}

	if vsKind == reflect.Map {
		mapKeys := vsVal.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if !f(mapKeys[i].Interface()) {
				return false
			}
		}
		return true
	}

	for i := 0; i < vsVal.Len(); i++ {
		if !f(vsVal.Index(i).Interface()) {
			return false
		}
	}

	return true
}

// Filter - not implemented yet TODO: implement
func Filter(vs interface{}, f func(v interface{}) bool) interface{} {
	return nil
}

// Map - not implemented yet TODO: implement
func Map(vs interface{}, f func(v interface{}) interface{}) []interface{} {
	return nil
}
