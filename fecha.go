// Package fecha está pensado para simplificar el trabajo con fechas.
// Trabaja con un int subyacente, por lo que las comparaciones son fáciles.
// Además incluye las funciones de Marshal y Unmarshal para MongoDB, SQL y JSON.
package fecha

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Fecha permite utilizar time.Time ignorando horas y zonas horarias.
// Se entiende que todas las fechas están guardadas en UTC.
type Fecha int

// NewFecha parsea un texto con formato JSON.
func NewFecha(texto string) (fch Fecha, err error) {

	t, err := time.Parse("2006-01-02", texto)
	if err != nil {
		return fch, errors.Wrap(err, `Parseando string: "`+texto+`"`)
	}

	fch = deTimeAFecha(t)

	return fch, err
}

// NewFechaFromLayout parsea el texto ingresado, utilizando el layout especificado.
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

// NewFechaFromTime le corta la hora y devuelve la fecha.
func NewFechaFromTime(t time.Time) (fch Fecha) {
	fch = deTimeAFecha(t)
	return
}

// NewFechaFromInts le corta la hora y devuelve la fecha.
func NewFechaFromInts(año, mes, dia int) (fch Fecha) {
	t := time.Date(año, time.Month(mes), dia, 0, 0, 0, 0, time.UTC)
	fch = deTimeAFecha(t)
	return
}

// IsValid devuelve true si es una fecha válida.
func (f Fecha) IsValid() bool {
	_, err := time.Parse("20060102", fmt.Sprint(int(f)))
	if err != nil {
		return false
	}
	return true
}

// Time devuele la representación con el tipo time.Time
func (f Fecha) Time() (nuevaFecha time.Time) {
	enTexto := fmt.Sprint(int(f))
	nuevaFecha, err := time.Parse("20060102", enTexto)
	if err != nil {
		panic(err)
	}
	return nuevaFecha
}

// Dia devuelve el número del día
func (f Fecha) Dia() int {
	return f.Time().Day()
}

// Mes devuelve el número del mes.
func (f Fecha) Mes() int {
	return int(f.Time().Month())
}

// Año devuelve el año en formato 2006
func (f Fecha) Año() int {
	return f.Time().Year()
}

// PeriodoMes devuelve la struct Mes correspondiente a la fecha.
// Se supone que se está trabajando con una fecha válida.
func (f Fecha) PeriodoMes() Mes {
	return Mes{
		año: f.Año(),
		mes: f.Mes(),
	}
}

// AgregarDias devuelve una nueva fecha con la cantidad de días agregados
// Si el signo es negativo los resta.
func (f Fecha) AgregarDias(dias int) (NuevaFecha Fecha) {
	enTime := f.Time().Add(time.Duration(24*dias) * time.Hour)
	return deTimeAFecha(enTime)
}

// AgregarMeses suma la cantidad de meses deseados. El día siempre queda igual
// salvo que el mes destino tenga menos días. Por ejemplo, sumar 1 mes al 31/01/2017
// resulta en 28/02/2017
func (f Fecha) AgregarMeses(cantidad int) (nuevaFecha Fecha) {
	dia := f.Dia()
	mes := f.Mes()
	año := f.Año()

	añosAgregar := 0

	// Cambia año?
	if mes+cantidad > 12 {
		añosAgregar++
	}

	añosAgregar = cantidad / 12

	mesesAgregar := cantidad - añosAgregar*12
	nuevoMes := mes + mesesAgregar

	if dia == 31 {
		switch nuevoMes {
		case 2, 4, 6, 9, 11:
			dia = 30
		}
	}

	if dia == 30 && nuevoMes == 2 {
		// TODO: si es año bisiesto que ponga 29
		dia = 28
	}

	fec := time.Date(año+añosAgregar, time.Month(nuevoMes), dia, 0, 0, 0, 0, time.UTC)
	return NewFechaFromTime(fec)

}

// AgregarAños devuelve una nueva fecha con los añós agregados
func (f Fecha) AgregarAños(cantidad int) (nuevaFecha Fecha) {
	fechaT := f.Time()

	dia := fechaT.Day()
	mes := fechaT.Month()
	año := fechaT.Year()

	nuevoAño := año + cantidad
	if nuevoAño > 9999 {
		return nuevaFecha
	}

	nuevaFecha = NewFechaFromTime(time.Date(nuevoAño, mes, dia, 0, 0, 0, 0, time.UTC))
	return
}

// Menos devuelve la cantidad de días de diferencia entre dos fechas
// Se entiende que f2 es la fecha posterior.
func (f Fecha) Menos(f2 Fecha) (dias int) {
	horas := f.Time().Sub(f2.Time()).Hours()
	dias = int(math.Trunc(horas / 24))
	return dias
}

// Diff calcula la diferencia de días entre dos fechas.Diff
// Si la segunda fecha es anterior a la primera, devuelve los días en negativo.
func Diff(f1, f2 Fecha) (dias int) {

	horas := f2.Time().Sub(f1.Time()).Hours()
	dias = int(math.Trunc(horas / 24))
	return dias
}

// Agrupacion dice el intervalo que se desea para una TimeSeries
type Agrupacion string

const (
	// AgrupacionMensual devolverá el primer día de cada mes
	AgrupacionMensual Agrupacion = "Mensual"
	// AgrupacionSemanal devolverá las semanas agrupando el lunes como primer día
	AgrupacionSemanal = "Semanal"
)

// TimeSeries devuelve todos los intervalos de fechas entre las dos fechas especificadas.
// Por ejemplo: entre 01/05/2017 y 04/01/2018 = [01/05/2017, 01/06/2017, 01/08/2018, 01/09/2018...]
func TimeSeries(desde, hasta Fecha, agrupacion Agrupacion) (fechas []Fecha, err error) {
	switch agrupacion {
	case AgrupacionMensual:
		// Desde
		mes := desde.Mes()
		año := desde.Año()
		f1Temp := time.Date(año, time.Month(mes), 1, 0, 0, 0, 0, time.UTC)
		f1 := NewFechaFromTime(f1Temp)
		// La primer fecha va seguro
		fechas = append(fechas, f1)

		// Hasta
		mesHasta := hasta.Mes()
		añoHasta := hasta.Año()
		fnTemp := time.Date(añoHasta, time.Month(mesHasta), 1, 0, 0, 0, 0, time.UTC)
		fn := NewFechaFromTime(fnTemp)

		fSiguiente := f1
		for {
			fSiguiente = fSiguiente.AgregarMeses(1)
			fechas = append(fechas, fSiguiente)

			// Verifico si llegué a la última fecha
			if fSiguiente == fn {
				break
			}
		}
	default:
		return fechas, errors.New("Agrupacion no implementada")
	}
	return
}

// MarshalJSON es para tomar un string y pasarlo a una fecha.Fecha
func (f Fecha) MarshalJSON() (by []byte, err error) {
	if f == 0 {
		by = []byte("null")
		return by, nil
	}
	if f.IsValid() == false {
		return by, errors.New(fmt.Sprint("No se puede marshalizar la fecha ", int(f), " . No es válida"))
	}
	enTime := f.Time()
	enString := enTime.Format("2006-01-02")
	by = []byte(`"` + enString + `"`)
	return by, nil
}

// UnmarshalJSON Es para pasar un Fecha => JSON
func (f *Fecha) UnmarshalJSON(input []byte) error {
	texto := string(input)

	if texto == "null" || texto == `""` {
		*f = Fecha(0)
		return nil
	}

	// Quito las comillas
	texto = strings.Replace(texto, `"`, "", -1)

	// Si la fecha viene en formato Date de Javascript(), tomo la primera parte nomás
	// 2020-02-04T03:00:00.000Z
	if len(texto) == 24 {
		texto = texto[:10]
	}
	fechaEnTime, err := time.Parse("2006-01-02", texto)

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

// JSONString devuelve el la fecha en forato 2016-02-19
func (f *Fecha) JSONString() string {
	if f == nil {
		return "null"
	}
	if *f == 0 {
		return "null"
	}
	return f.Time().Format("2006-01-02")
}

func (f Fecha) String() string {
	// si es zero, devuelvo el valor cero
	if f.IsZero() {
		return "01/01/0001"
	}

	enTexto := fmt.Sprint(int(f))
	fecha, err := time.Parse("20060102", enTexto)

	// Si es inválida
	if err != nil {
		return "N/A"
	}

	// Está ok
	return fecha.Format("02/01/2006")
}

// IsZero devuelve true si la fecha es el número 0.
func (f Fecha) IsZero() bool {
	return int(f) == 0
}

// DiaDeLaSemana devuelve la fecha del día para suegerirla en el index
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

// AgregarDiasHabiles suma la cantidad de días especificados en el argumento.
// Considera los sábados y domingos (no tiene en cuenta feriados)
func (f Fecha) AgregarDiasHabiles(cantidad int) (nuevaFecha Fecha) {

	// Si el primer día no es hábil, arrastro hasta el próximo día hábil
	nuevaFecha = proximoDiaHabil(f)

	for i := 0; i == cantidad; i++ {
		// Agrego un día
		nuevoTime := f.Time().Add(time.Duration(int64(i) * int64(time.Hour) * 24))
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

var _ driver.Valuer = (*Fecha)(nil)

// Value satisface la interface de package sql.
// En la base de datos la guarda como el tipo DATE que es un int64
func (f Fecha) Value() (driver.Value, error) {
	if f.IsZero() {
		return nil, nil
	}
	if !f.IsValid() {
		return nil, errors.Errorf("invalid date %v", int(f))
	}
	return f.JSONString(), nil
}

type Time interface {
	Time() time.Time
}

var _ sql.Scanner = (*Fecha)(nil)

// Scan satisface la interface de package sql.
// En la base de datos la guarda como el tipo DATE que es un int64
func (f *Fecha) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if value == 0 {
		return nil
	}

	// Es time?
	t, ok := value.(time.Time)
	if ok {
		*f = NewFechaFromTime(t)
		return nil
	}

	// Tiene interface Time?
	tInterface, ok := value.(Time)
	if ok {
		*f = NewFechaFromTime(tInterface.Time())
		return nil
	}

	return nil
}
