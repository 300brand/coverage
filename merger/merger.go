package merger

import (
	"reflect"
)

type Merger interface {
	Changelog() Changelog
}

type m struct {
	Struct reflect.Value
	Log    Changelog
}

func Merge(dst, src Merger) (madeChanges bool) {
	d := m{
		Struct: reflect.ValueOf(dst).Elem(),
		Log:    dst.Changelog(),
	}
	s := m{
		Struct: reflect.ValueOf(src).Elem(),
		Log:    src.Changelog(),
	}

	for _, f := range src.Changelog().Fields() {
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
