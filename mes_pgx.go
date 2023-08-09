package fecha

import (
	"fmt"
	"time"

	"github.com/jackc/pgtype"
)

var _ pgtype.ValueTranscoder = (*Fecha)(nil)
var _ pgtype.Value = (*Fecha)(nil)
var _ pgtype.TypeValue = (*Fecha)(nil)

func (t *Mes) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
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
		*t = Mes{año: r.Time.Year(), mes: int(r.Time.Month())}
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

func (src Mes) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {

	d := pgtype.Date{}

	if src == NilMes {
		d.Status = pgtype.Null
	} else {
		d.Status = pgtype.Present
		d.Time = time.Date(src.año, time.Month(src.mes), 1, 0, 0, 0, 0, time.UTC)
	}
	return d.EncodeBinary(ci, buf)
}

// TypeName returns the PostgreSQL name of this type.
func (Mes) TypeName() string {
	return "date"
}

func (t *Mes) NewTypeValue() pgtype.Value {
	return t
}

func (t *Mes) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to Point", src)
}
func (t *Mes) Get() interface{} {
	return *t
}
func (t *Mes) AssignTo(dst interface{}) error {
	return nil
}
