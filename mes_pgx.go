package fecha

import (
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/rs/zerolog/log"
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
		d.Time = time.Date(src.año, time.Month(src.mes), 1, 0, 0, 0, 0, nil)
	}
	return d.EncodeBinary(ci, buf)
}

// func (t *Mes) DecodeText(ci *pgtype.ConnInfo, src []byte) error {

// 	if src == nil {
// 		return nil
// 	}
// 	if len(src) == 0 {
// 		return nil
// 	}
// 	return nil
// 	// Es time?
// 	// t, ok := src.(time.Time)
// 	// if ok {
// 	// 	*f = NewFechaFromTime(t)
// 	// 	return nil
// 	// }

// 	// // Tiene interface Time?
// 	// tInterface, ok := value.(Time)
// 	// if ok {
// 	// 	*f = NewFechaFromTime(tInterface.Time())
// 	// 	return nil
// 	// }

// 	// return nil
// }

// func (src *Mes) EncodeText(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
// 	d := pgtype.Date{}

// 	if src == nil {
// 		d.Status = pgtype.Null
// 	} else {
// 		d.Status = pgtype.Present
// 		d.Time = src.Time()
// 	}
// 	return d.EncodeText(ci, buf)
// }

// TypeName returns the PostgreSQL name of this type.
func (Mes) TypeName() string {
	log.Error().Msgf("Returning typename: %v", "date")
	return "date"
}

func (t *Mes) NewTypeValue() pgtype.Value {
	log.Error().Msgf("NewTypeValue: returning %v", *t)
	return t
}

func (t *Mes) Set(src interface{}) error {
	log.Error().Msgf("Set: src was %v", src)
	return fmt.Errorf("cannot convert %v to Point", src)
}
func (t *Mes) Get() interface{} {
	log.Error().Msgf("Get: returning %v", *t)
	return *t
}
func (t *Mes) AssignTo(dst interface{}) error {
	log.Error().Msgf("Assigning to, dst was %v", dst)
	return nil
}
