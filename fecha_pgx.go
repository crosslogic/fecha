package fecha

import (
	"fmt"

	"github.com/jackc/pgtype"
)

var _ pgtype.ValueTranscoder = (*Fecha)(nil)
var _ pgtype.Value = (*Fecha)(nil)
var _ pgtype.TypeValue = (*Fecha)(nil)

func (t *Fecha) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	r := &pgtype.Date{}

	err := r.DecodeBinary(ci, src)
	if err != nil {
		return err
	}

	switch r.Status {
	case pgtype.Present:
		if r.InfinityModifier != pgtype.None {
			return fmt.Errorf("cannot assign %v", src)
		}
		*t = NewFechaFromTime(r.Time)
		return nil
	case pgtype.Null:
		return pgtype.NullAssignTo(t)
	}

	return fmt.Errorf("cannot decode %#v", src)
	// type compatibility is checked by AssignTo
	// only lossless assignments will succeed
	// if err := r.AssignTo(&t); err != nil {
	// 	return err
	// }

	// return nil
}

func (src Fecha) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {

	d := pgtype.Date{}

	if src == 0 {
		d.Status = pgtype.Null
	} else {
		d.Status = pgtype.Present
		d.Time = src.Time()
	}
	return d.EncodeBinary(ci, buf)
}

func (t *Fecha) DecodeText(ci *pgtype.ConnInfo, src []byte) error {

	if src == nil {
		return nil
	}
	if len(src) == 0 {
		return nil
	}
	return nil
	// Es time?
	// t, ok := src.(time.Time)
	// if ok {
	// 	*f = NewFechaFromTime(t)
	// 	return nil
	// }

	// // Tiene interface Time?
	// tInterface, ok := value.(Time)
	// if ok {
	// 	*f = NewFechaFromTime(tInterface.Time())
	// 	return nil
	// }

	// return nil
}

func (src *Fecha) EncodeText(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
	d := pgtype.Date{}

	if src == nil {
		d.Status = pgtype.Null
	} else {
		d.Status = pgtype.Present
		d.Time = src.Time()
	}
	return d.EncodeText(ci, buf)
}

// NewTypeValue creates a TypeValue including references to internal type information. e.g. the list of members
// in an EnumType.
// func (t Fecha) NewTypeValue() Fecha {

// }

// TypeName returns the PostgreSQL name of this type.
func (Fecha) TypeName() string {
	return "date"
}

func (t *Fecha) NewTypeValue() pgtype.Value {
	return t
}

func (t *Fecha) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to Point", src)
}
func (t *Fecha) Get() interface{} {
	return *t
}
func (t *Fecha) AssignTo(dst interface{}) error {
	return nil
}
