package merger

import (
	"errors"
	"reflect"
	"strings"
)

// Used to store frequently accessed information during merges
type structInfo struct {
	Struct reflect.Value
	Log    Changelog
}

// Merges two structs with a Changelog at the *top-level*. dst and src must be
// the same type. dst must always be passed by reference or a panic will occur
func Merge(dst, src interface{}) (madeChanges bool, err error) {
	var s, d *structInfo

	if !typesMatch(dst, src) {
		return false, errors.New("Cannot merge mismatching types")
	}

	if s, err = newInfo(src); err != nil {
		return
	}
	if d, err = newInfo(dst); err != nil {
		return
	}

	for _, f := range s.Log.Fields() {
		dChange := d.Log.LastChangeTo(f)
		if !s.Log.LastChangeTo(f).After(dChange) {
			// Source value is up-to-date
			continue
		}
		// Update source value for field
		dVal, sVal := d.Struct.FieldByName(f), s.Struct.FieldByName(f)
		dVal.Set(sVal)
		d.Log.Add(f, dChange)
		madeChanges = true
	}
	return
}

func getChangelog(i interface{}) (cl Changelog, err error) {
	s := reflect.ValueOf(i)
	// Dereference pointers
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	if s.Kind() != reflect.Struct {
		return nil, errors.New("Cannot extract Changelog from non-struct type")
	}
	for idx := 0; idx < s.NumField(); idx++ {
		f := s.Field(idx)
		if c, ok := f.Interface().(Changelog); ok {
			cl = c
			return
		}
	}
	return nil, errors.New("Could not find Changelog in struct top-level")
}

func newInfo(i interface{}) (si *structInfo, err error) {
	si = &structInfo{}
	si.Struct = reflect.ValueOf(i)
	if si.Struct.Kind() == reflect.Ptr {
		si.Struct = si.Struct.Elem()
	}
	si.Log, err = getChangelog(i)
	return
}

func typesMatch(a, b interface{}) bool {
	return strings.TrimPrefix(reflect.TypeOf(a).String(), "*") == strings.TrimPrefix(reflect.TypeOf(b).String(), "*")
}
