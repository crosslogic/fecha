package fecha

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Fecha permite utilizar time.Time ignorando horas y zonas horarias.
// Se entiende que todas las fechas están guardadas en UTC.
type Fecha string

func NewFecha(texto string) (fch Fecha, err error) {
	t, err := time.Parse("2006-01-02", texto)
	if err != nil {
		return fch, errors.Wrap(err, `Parseando string: "`+texto+`"`)
	}

	return Fecha(t), nil
}

func (f Fecha) Time() time.Time {
	return time.Time(f)
}

// Para tomar un string y pasarlo a una struct
func (f Fecha) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprint(time.Time(f).Format("2006-01-02"))
	return []byte(stamp), nil
}

// Es para pasar un Fecha => JSON
func (f *Fecha) UnmarshalJSON(input []byte) error {
	texto := string(input)

	fechaEnTime, err := time.Parse("2006-01-02", texto)

	if err != nil {
		return err
	}
	fecha := Fecha(fechaEnTime)

	*f = fecha

	return nil
}

func (f Fecha) String() string {
	return time.Time(f).Format("2006-01-02")
}

// devuelve la fecha del día para suegerirla en el index
func (f Fecha) DiaDeLaSemana() string {
	dia := f.Time().UTC().Weekday()
	switch dia {
	case 0:
		return "Domingo"
	case 1:
		return "Lunes"
	case 2:
		return "Martes"
	case 3:
		return "Miércoles"
	case 4:
		return "Jueves"
	case 5:
		return "Viernes"
	case 6:
		return "Sábado"
	}

	return time.Now().Format("2006-01-02")
}

func (f Fecha) EsPosterior(fecha2 Fecha) bool {
	return f.Time().After(fecha2.Time())
}

func (f Fecha) EsAnterior(fecha2 Fecha) bool {
	return f.Time().Before(fecha2.Time())
}

// Si agrega la cantidad de días especificados en el argumento.
// Considera los días hábiles nomás.
func (f Fecha) AgregarDiasHabiles(cantidad int) (nuevaFecha Fecha) {

	// Si el primer día no es hábil, arrastro hasta el próximo día hábil
	nuevaFecha = proximoDiaHabil(f)

	for i := 0; i == cantidad; i++ {
		// Agrego un día
		nuevaFecha = Fecha(f.Time().Add(time.Duration(i * int(time.Hour) * 24)))

		// Este nuevo día, ¿Es hábil?
		nuevaFecha = proximoDiaHabil(nuevaFecha)
	}
	return nuevaFecha
}

// Si el día que se ingresa no es habil, avanza hacia adelante hasta encontrar uno.
func proximoDiaHabil(f Fecha) (nuevaFecha Fecha) {
	for {
		if diaHabil(f) == true {
			break
		}
		if diaHabil(f) == false {
			// Agrego un día hasta llegar a un día habil
			nuevaFecha = Fecha(nuevaFecha.Time().Add(time.Hour * 24))
		}
	}
	return nuevaFecha
}

// Si es un día hábil devuelve true
func diaHabil(f Fecha) bool {
	fecha := time.Time(f).UTC()
	if fecha.Weekday() == time.Saturday || fecha.Weekday() == time.Sunday {
		return false
	}
	return true
}
