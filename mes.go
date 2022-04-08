package fecha

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// AñoMinimo determina que si se ingresa una fecha con año menor,
// el período será considerado inválido.
const AñoMinimo = 1900

// AñoMaximo determina que si se ingresa una fecha con año posterior,
// el período será considerado inválido.
const AñoMaximo = 2200

// SeparadorMes es el caracter utilizado para separar el mes de año
// al mostrarlo como string
const SeparadorMes = "/"

// MesPrimero true: 03/2020. Si es false se muestra 2020/03
const MesPrimero = true

// Mes es una estructura para representar un mes calendario de un año particular.Mes
// Reduce la ambiguedad ya que no es necesario ingresar un día.
//
// Se usaron variables locales para evitar tener una fecha en un estado
// inválido.
//
// En JSON se marshaliza con el formato "2020-08"
// En la base de datos se persiste como un DATE
type Mes struct {
	año int
	mes int
}

func NewMesMust(año, mes int) Mes {
	m := Mes{año, mes}
	if m.mes < 1 || m.mes > 12 {
		panic(errors.Errorf("mes '%v' es inválido", m.mes))
	}
	if m.año < AñoMinimo || m.año > AñoMaximo {
		panic(errors.Errorf("año '%v' fuera del rango permitido", m.año))
	}
	return m
}

func NewMes(año, mes int) (out Mes, err error) {
	m := Mes{año, mes}
	if m.mes < 1 || m.mes > 12 {
		return out, errors.Errorf("mes '%v' es inválido", m.mes)
	}
	if m.año < AñoMinimo || m.año > AñoMaximo {
		return out, errors.Errorf("año '%v' fuera del rango permitido", m.año)
	}
	return
}

// MesDelAño devuelve el número de mes.
func (m Mes) MesDelAño() int {
	return m.mes
}

// Año devuelve el número del año.
func (m Mes) Año() int {
	return m.año
}

func (m Mes) Anterior(f2 Mes) bool {
	if m.año > f2.año { // 2030 > 2021
		return false
	}
	if m.año < f2.año { // 2030 > 2021
		return true
	}
	// Son años distintos, define mes
	return m.mes < f2.mes
}

func (m Mes) AnteriorOIgual(f2 Mes) bool {
	if m == f2 {
		return true
	}
	return m.Anterior(f2)
}

func (m Mes) Posterior(f2 Mes) bool {
	if m.año > f2.año {
		return true
	}
	if m.año < f2.año {
		return false
	}
	return m.mes > f2.mes
}

func (m Mes) PosteriorOIgual(f2 Mes) bool {
	if m == f2 {
		return true
	}
	return m.Posterior(f2)
}

func (m Mes) String() (out string) {
	if !m.Valid() {
		return "mes inválido"
	}

	año := ""
	mes := ""
	switch {
	case m.año >= 1000:
		año = fmt.Sprint(m.año)
	case m.año >= 100:
		año = fmt.Sprint("0", m.año)
	case m.año >= 10:
		año = fmt.Sprint("00", m.año)
	case m.año >= 1:
		año = fmt.Sprint("000", m.año)
	}

	switch {
	case m.mes < 10:
		mes = fmt.Sprint("0", m.mes)
	default:
		mes = fmt.Sprint(m.mes)
	}

	switch MesPrimero {
	case true:
		return fmt.Sprintf("%v%v%v", mes, SeparadorMes, año)
	case false:
		return fmt.Sprintf("%v%v%v", año, SeparadorMes, mes)
	}

	return
}

// Valid devuelve true si la fecha es válida
func (m Mes) Valid() bool {
	// fmt.Println(m)
	if m.mes < 1 || m.mes > 12 {
		return false
	}
	if m.año < AñoMinimo || m.año > AñoMaximo {
		return false
	}

	return true
}

// Zero devuelve true si el día y el año son cero
func (m Mes) Zero() bool {
	if m.año == 0 && m.mes == 0 {
		return true
	}
	return false
}

// SumarMeses devuelve una nueva fecha con los meses agregados.
// Si se quiere restar, ingresar meses en negativo.
// Se supone que se está trabajando con un Mes válido no cero.
func (m Mes) SumarMeses(meses int) (out Mes) {
	out.año = m.año
	out.mes = m.mes + meses

	switch {
	case meses > 0:
		for {
			if out.mes <= 12 {
				return out
			}
			out.mes -= 12
			out.año++
		}
	case meses < 0:
		for {
			if out.mes > 0 {
				return out
			}
			out.mes += 12
			out.año--
		}
	}
	return m

}

// PrimerDia devuelve la fecha considerando el primer día del período.
func (m Mes) PrimerDia() Fecha {
	return NewFechaFromInts(m.año, m.mes, 1)
}

// UltimoDia devuelve la fecha considerando el último día del período.
func (m Mes) UltimoDia() Fecha {
	return NewFechaFromInts(m.año, m.mes, ultimoDia(m.mes, m.año))
}

// JSONString devuelve la representación que se utiliza en JSON.
// Si es cero devuelve null
// Si no es válida devuelve "N/D"
func (m Mes) JSONString() string {

	if !m.Valid() {
		return "null"
	}
	if !m.Valid() {
		return "N/D"
	}

	año := ""
	mes := ""
	switch {
	case m.año >= 1000:
		año = fmt.Sprint(m.año)
	case m.año >= 100:
		año = fmt.Sprint("0", m.año)
	case m.año >= 10:
		año = fmt.Sprint("00", m.año)
	case m.año >= 1:
		año = fmt.Sprint("000", m.año)
	}

	switch {
	case m.mes < 10:
		mes = fmt.Sprint("0", m.mes)
	default:
		mes = fmt.Sprint(m.mes)
	}
	return fmt.Sprintf("%v-%v", año, mes)

}

// Value satisface la interface de package sql.
// En la base de datos la guarda como el tipo DATE.
// Si se intenta persistir una fecha Zero() => La guarda como null.
// Si se intenta persistir una fecha !Valid() => devuelve error.
func (m Mes) Value() (driver.Value, error) {
	if m.Zero() {
		return nil, nil
	}
	if !m.Valid() {
		return nil, errors.Errorf("se intentó persistir un Mes no válido. año=(%v), mes=(%v)", m.año, m.mes)
	}
	return m.PrimerDia().Time(), nil
}

// Scan satisface la interface de package sql.
// Si el valor que estaba persistido es menor al mínimo o mayor al máximo,
// no va a dar a error. Queda a criterio del usuario analizarla con
// el método Valid().
// Si el va
func (m *Mes) Scan(value interface{}) error {
	if value == nil {
		*m = Mes{}
	}

	enTime, ok := value.(time.Time)
	if !ok {
		return errors.Errorf("se esperaba que el tipo en la base de datos fuera time.Time, era %T", value)
	}
	*m = NewFechaFromTime(enTime).PeriodoMes()

	return nil
}

// MarshalJSON es para tomar una una struct a un string JSON.
func (m Mes) MarshalJSON() (by []byte, err error) {

	if m.Zero() {
		by = []byte("null")
		return by, nil
	}
	if !m.Valid() {
		return by, errors.Errorf("no se puede marshalizar la fecha %v, (no es válida)", m)
	}
	by = []byte(`"` + m.JSONString() + `"`)
	return by, nil
}

// UnmarshalJSON es para parsear el string a una struct Mes.
// Si llega una cadena null o vacía, se crea una struct con valor cero.
func (m *Mes) UnmarshalJSON(input []byte) error {
	texto := string(input)

	if texto == "null" || texto == `""` {
		*m = Mes{}
		return nil
	}

	// Quito las comillas
	texto = strings.Replace(texto, `"`, "", -1)

	// Si la fecha viene en formato Date de Javascript(), tomo la primera parte nomás
	// 2020-02-04T03:00:00.000Z
	if len(texto) == 24 {
		texto = texto[:7]
	}

	partes := strings.Split(texto, "-")
	if len(partes) != 2 {
		return errors.Errorf("se esperaba una cadena con el formato AAAA-MM")
	}

	año, err := strconv.Atoi(partes[0])
	if err != nil {
		return errors.Wrap(err, "convirtiendo el año")
	}

	mes, err := strconv.Atoi(partes[1])
	if err != nil {
		return errors.Wrap(err, "convirtiendo el mes")
	}
	if mes > 12 || mes < 0 {
		// Nunca puede ser válido
		return errors.Errorf("el mes debe estar entre 1 y 12 (era %v)", mes)
	}

	if m == nil {
		m = &Mes{año, mes}
		return nil
	}

	*m = Mes{año, mes}

	return nil
}

func ultimoDia(mes int, año int) (out int) {
	out = 31
	switch mes {
	case 2, 4, 6, 9, 11:
		out = 30
	}

	if mes == 2 {
		if año%4 == 0 && año%100 != 0 || año%400 == 0 {
			return 29
		} else {
			return 28
		}
	}

	return
}
