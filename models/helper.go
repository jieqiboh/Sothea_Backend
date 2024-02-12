package models

// PtrTo is a generic method for creating a pointer to a primitive type, used when manually creating a models object
func PtrTo[T any](v T) *T {
	return &v
}

// SafeDeref is a generic method used in String(), to safely dereference a pointer that might be nil
// Returns the null version of a type, instead of nil, which cannot be printed
func SafeDeref[T any](p *T) T {
	if p == nil {
		var v T
		return v
	}
	return *p
}

///**
//sql.NullString and other nullable types provided by the database/sql library don't provide support for
//marshalling to and from json for our use case
//*/
//// Nullable Bool that overrides sql.NullBool
//type NullBool struct {
//	sql.NullBool
//}
//
//func (nb NullBool) MarshalJSON() ([]byte, error) {
//	if nb.Valid {
//		return json.Marshal(nb.Bool)
//	}
//	return json.Marshal(nil)
//}
//
//func (nb *NullBool) UnmarshalJSON(data []byte) error {
//	var b *bool
//	if err := json.Unmarshal(data, &b); err != nil {
//		return err
//	}
//	if b != nil {
//		nb.Valid = true
//		nb.Bool = *b
//	} else {
//		nb.Valid = false
//	}
//	return nil
//}
//
//// Nullable Float64 that overrides sql.NullFloat64
//type NullFloat64 struct {
//	sql.NullFloat64
//}
//
//func (nf NullFloat64) MarshalJSON() ([]byte, error) {
//	if nf.Valid {
//		return json.Marshal(nf.Float64)
//	}
//	return json.Marshal(nil)
//}
//
//func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
//	var f *float64
//	if err := json.Unmarshal(data, &f); err != nil {
//		return err
//	}
//	if f != nil {
//		nf.Valid = true
//		nf.Float64 = *f
//	} else {
//		nf.Valid = false
//	}
//	return nil
//}
//
//// Nullable Int64 that overrides sql.NullInt64
//type NullInt64 struct {
//	sql.NullInt64
//}
//
//func (ni NullInt64) MarshalJSON() ([]byte, error) {
//	if ni.Valid {
//		return json.Marshal(ni.Int64)
//	}
//	return json.Marshal(nil)
//}
//
//func (ni *NullInt64) UnmarshalJSON(data []byte) error {
//	var i *int64
//	if err := json.Unmarshal(data, &i); err != nil {
//		return err
//	}
//	if i != nil {
//		ni.Valid = true
//		ni.Int64 = *i
//	} else {
//		ni.Valid = false
//	}
//	return nil
//}
//
//// Nullable String that overrides sql.NullString
//type NullString struct {
//	sql.NullString
//}
//
//func (ns NullString) MarshalJSON() ([]byte, error) {
//	if ns.Valid {
//		return json.Marshal(ns.String)
//	}
//	return json.Marshal(nil)
//}
//
//func (ns *NullString) UnmarshalJSON(data []byte) error {
//	var s *string
//	if err := json.Unmarshal(data, &s); err != nil {
//		return err
//	}
//	if s != nil {
//		ns.Valid = true
//		ns.String = *s
//	} else {
//		ns.Valid = false
//	}
//	return nil
//}
//
//// Nullable String that overrides sql.NullString
//type NullTime struct {
//	sql.NullTime
//}
//
//func (nt NullTime) MarshalJSON() ([]byte, error) {
//	if nt.Valid {
//		return json.Marshal(nt.Time)
//	}
//	return json.Marshal(nil)
//}
//
//func (nt *NullTime) UnmarshalJSON(data []byte) error {
//	var t *time.Time
//	if err := json.Unmarshal(data, &t); err != nil {
//		return err
//	}
//	if t != nil {
//		nt.Valid = true
//		nt.Time = *t
//	} else {
//		nt.Valid = false
//	}
//	return nil
//}
