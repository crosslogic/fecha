package fecha

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/pkg/errors"
)

// Fecha permite utilizar time.Time ignorando horas y zonas horarias.
// Se entiende que todas las fechas están guardadas en UTC.
type Fecha int

func NewFecha(texto string) (fch Fecha, err error) {

	t, err := time.Parse("2006-01-02", texto)
	if err != nil {
		return fch, errors.Wrap(err, `Parseando string: "`+texto+`"`)
	}

	fch = deTimeAFecha(t)

	return fch, err
}

func NewFechaFromLayout(layout, texto string) (fch Fecha, err error) {

	if layout == "" {
		layout = "02/01/2006"
	}
	t, err := time.Parse(layout, texto)
	if err != nil {
		return fch, errors.Wrap(err, `Parseando string: "`+texto+`"`)
	}

	fch = deTimeAFecha(t)

	return fch, err
}

func NewFechaFromTime(t time.Time) (fch Fecha) {
	fch = deTimeAFecha(t)
	return
}

func (f Fecha) IsValid() bool {
	_, err := time.Parse("20060102", fmt.Sprint(int(f)))
	if err != nil {
		return false
	}
	return true
}

// Devuele la representación en formato time.Time
func (f Fecha) Time() (nuevaFecha time.Time) {
	enTexto := fmt.Sprint(int(f))
	nuevaFecha, err := time.Parse("20060102", enTexto)
	if err != nil {
		panic(err)
	}
	return nuevaFecha
}

// Devuelve una nueva fecha con la cantidad de días agregados
func (f Fecha) AgregarDias(dias int) (NuevaFecha Fecha) {
	enTime := f.Time().Add(time.Duration(24*dias) * time.Hour)
	return deTimeAFecha(enTime)
}

// Funciones para serializar y desserializar las fechas usando una base de datos MongoDB
func (f *Fecha) SetBSON(raw bson.Raw) (err error) {

	valor := int32(0)
	err = raw.Unmarshal(&valor)
	if err != nil {
		return err
	}
	*f = Fecha(valor)
	if f.IsValid() == false {
		//		panic(fmt.Sprint("No se pudo desserializar la fecha", raw))
		return errors.New(fmt.Sprint("No se pudo deserializar la fecha: ", valor))
	}
	return nil
}

func (f Fecha) GetBSON() (rtdo interface{}, err error) {
	return int32(f), nil
}

// Para tomar un string y pasarlo a una struct
func (f Fecha) MarshalJSON() (by []byte, err error) {
	if f.IsValid() == false {
		return by, errors.New(fmt.Sprint("No se puede marhsalizar la fecha ", int(f), " . No es válida"))
	}
	enTime := f.Time()
	enString := enTime.Format("2006-01-02")
	by = []byte(enString)
	return by, nil
}

// Es para pasar un Fecha => JSON
func (f *Fecha) UnmarshalJSON(input []byte) error {
	texto := string(input)

	log.Println("Unmarshalling input: ", texto)
	fechaEnTime, err := time.Parse("2006-01-02", texto)

	log.Println("Pasado a fecha", texto)
	if err != nil {
		return err
	}
	*f = NewFechaFromTime(fechaEnTime)

	return nil
}

// Transforma a Fecha un time
func deTimeAFecha(f time.Time) (fecha Fecha) {

	// Convierto la fecha a texto
	enTexto := f.Format("20060102")

	// Convierto el texto al int
	enInt, err := strconv.Atoi(enTexto)
	if err != nil {
		panic(err)
	}

	return Fecha(enInt)
}

// JsonString devuelve el la fecha en forato 2016-02-19
func (f Fecha) JsonString() string {
	return f.Time().Format("2006-01-02")
}
func (f Fecha) String() string {
	enTexto := fmt.Sprint(int(f))
	fecha, err := time.Parse("20060102", enTexto)
	if err != nil {
		panic(err)
	}

	return fecha.Format("02/01/2006")
}

// devuelve la fecha del día para suegerirla en el index
func (f Fecha) DiaDeLaSemana() string {
	dia := f.Time().Weekday()
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
	return ""
}

// Si agrega la cantidad de días especificados en el argumento.
// Considera los días hábiles nomás.
func (f Fecha) AgregarDiasHabiles(cantidad int) (nuevaFecha Fecha) {

	// Si el primer día no es hábil, arrastro hasta el próximo día hábil
	nuevaFecha = proximoDiaHabil(f)

	for i := 0; i == cantidad; i++ {
		// Agrego un día
		nuevoTime := f.Time().Add(time.Duration(i * int(time.Hour) * 24))
		nuevaFecha = deTimeAFecha(nuevoTime)

		// Este nuevo día, ¿Es hábil?
		nuevaFecha = proximoDiaHabil(nuevaFecha)
	}
	return nuevaFecha
}

// Si el día que se ingresa no es habil, avanza hacia adelante hasta encontrar uno.
func proximoDiaHabil(f Fecha) (nuevaFecha Fecha) {
	nuevaFecha = f
	for {
		if diaHabil(nuevaFecha) == true {
			break
		}
		if diaHabil(nuevaFecha) == false {
			// Agrego un día hasta llegar a un día habil
			nuevaFecha = deTimeAFecha(nuevaFecha.Time().Add(time.Hour * 24))
		}
	}
	return nuevaFecha
}

// Si es un día hábil devuelve true
func diaHabil(f Fecha) bool {
	fecha := f.Time().UTC()
	if fecha.Weekday() == time.Saturday || fecha.Weekday() == time.Sunday {
		return false
	}
	return true
}
