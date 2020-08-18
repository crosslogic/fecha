package fecha

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {

	// Creo fecha
	f, err := NewFecha("2016-02-19")
	if err != nil {
		t.Error(err)
	}

	enTime := f.Time()
	esperado := time.Date(2016, 2, 19, 0, 0, 0, 0, time.UTC)
	if enTime != esperado {

		t.Error("Mal convertido a Time.",
			"\nEsperado: ", esperado,
			"\nObtenido: ", enTime,
		)
	}
	fmt.Println(enTime)
}

func TestDiaHabil(t *testing.T) {
	//Creo un sábado
	f, _ := NewFecha("2017-04-29")
	if diaHabil(f) == true {
		t.Error("Dijo que era habil un sábado")
	}

	//Creo un domingo
	f, _ = NewFecha("2017-04-30")
	if diaHabil(f) == true {
		t.Error("Dijo que era habil un domingo")
	}

	//Creo un lunes
	f, _ = NewFecha("2017-05-01")
	if diaHabil(f) == false {
		t.Error("Dijo que era habil un lunes")
	}

	//Creo un martes
	f, _ = NewFecha("2017-05-02")
	if diaHabil(f) == false {
		t.Error("Dijo que era habil un martes")
	}

	//Creo un miercoles
	f, _ = NewFecha("2017-05-03")
	if diaHabil(f) == false {
		t.Error("Dijo que era habil un miercoles")
	}

	//Creo un jueves
	f, _ = NewFecha("2017-05-04")
	if diaHabil(f) == false {
		t.Error("Dijo que era habil un jueves")
	}

	//Creo un viernes
	f, _ = NewFecha("2017-05-05")
	if diaHabil(f) == false {
		t.Error("Dijo que era habil un viernes")
	}
}

func TestProximoDiaHabil(t *testing.T) {

	f, _ := NewFecha("2017-04-29")
	if proximoDiaHabil(f) == f {
		t.Error("Sabado")
	}

	f, _ = NewFecha("2017-04-30")
	if proximoDiaHabil(f) == f {
		t.Error("Sabado")
	}

	f, _ = NewFecha("2017-05-01")
	if proximoDiaHabil(f) != f {
		t.Error("Sabado")
	}
}

func TestNewDate(t *testing.T) {
	// Creo fecha
	f, err := NewFecha("2016-02-19")
	if err != nil {
		t.Error(err)
	}

	if f != Fecha(20160219) {
		t.Error("La nueva fecha debería haber sido 20160219. Fue", f)
	}
}

func TestNewFechaFromInts(t *testing.T) {
	{
		f := NewFechaFromInts(2010, 12, 31)
		assert.Equal(t, Fecha(20101231), f)
	}
}

func TestAgregarMeses(t *testing.T) {
	// Creo fecha
	f, err := NewFecha("2016-12-19")
	if err != nil {
		t.Error(err)
	}

	// Agrego 1 mes
	f1 := f.AgregarMeses(1)
	if f1 != Fecha(20170119) {
		t.Error("Se esperaba", Fecha(20170119), "se obtuvo ", f1)
	}

	// Agrego 13 meses
	f2 := f.AgregarMeses(13)
	if f2 != Fecha(20180119) {
		t.Error("Se esperaba", Fecha(20180119), "se obtuvo ", f2)
	}

	// Resto 6 meses
	f3 := f.AgregarMeses(-6)
	if f3 != Fecha(20160619) {
		t.Error("Se esperaba", Fecha(20160619), "se obtuvo ", f3)
	}

	// Resto 6 meses
	f4 := f.AgregarMeses(-18)
	if f4 != Fecha(20150619) {
		t.Error("Se esperaba", Fecha(20150619), "se obtuvo ", f3)
	}
}

func TestTimeSeries(t *testing.T) {
	desde := Fecha(20170507)
	hasta := Fecha(20180729)

	ts, err := TimeSeries(desde, hasta, AgrupacionMensual)
	if err != nil {
		t.Error(err)
	}

	esperado := []Fecha{
		Fecha(20170501),
		Fecha(20170601),
		Fecha(20170701),
		Fecha(20170801),
		Fecha(20170901),
		Fecha(20171001),
		Fecha(20171101),
		Fecha(20171201),
		Fecha(20180101),
		Fecha(20180201),
		Fecha(20180301),
		Fecha(20180401),
		Fecha(20180501),
		Fecha(20180601),
		Fecha(20180701),
	}

	if reflect.DeepEqual(ts, esperado) == false {
		t.Error("TimeSeries mal creada",
			"\nEsperado: ", esperado,
			"\nObtenido: ", ts,
		)
	}
}

func TestMarshalAndUnmarshalJSON(t *testing.T) {

	// Creo fecha
	f, err := NewFecha("2016-02-19")
	assert.Nil(t, err)

	// Genero el string
	by, err := f.MarshalJSON()
	assert.Nil(t, err)

	enString := string(by)

	// Demarshalizo
	var fNueva Fecha
	err = fNueva.UnmarshalJSON(by)
	assert.Nil(t, err)

	//Corroboro
	if fNueva != f {
		t.Error("No se obtuvo el string esperado.",
			"\nEsperado: ", f,
			"\nObtenido: ", fNueva,
		)
	}

	// Corroboro que me lo devuelva igual
	if enString[1:11] != "2016-02-19" {
		t.Error("No se obtuvo el string esperado.",
			"\nEsperado: ", "2016-02-19",
			"\nObtenido: ", enString[1:11],
		)
	}
}

func TestUnmarshalJSONFormatoJavaScript(t *testing.T) {

	js := []byte(`"2020-02-04T03:00:00.000Z"`)

	// Demarshalizo
	var fNueva Fecha
	err := fNueva.UnmarshalJSON(js)
	assert.Nil(t, err)

	//Corroboro
	esperado := Fecha(20200204)
	if fNueva != esperado {
		t.Error("No se obtuvo el string esperado.",
			"\nEsperado: ", esperado,
			"\nObtenido: ", fNueva,
		)
	}

}
func TestMarshalAndUnmarshalJSONValoresVacios(t *testing.T) {

	// Creo fecha
	f := Fecha(0)

	// Genero el string
	by, err := f.MarshalJSON()
	assert.Nil(t, err)

	enString := string(by)

	// Demarshalizo
	var fNueva Fecha
	err = fNueva.UnmarshalJSON(by)
	assert.Nil(t, err)

	//Corroboro
	if fNueva != f {
		t.Error("No se obtuvo el string esperado.",
			"\nEsperado: ", f,
			"\nObtenido: ", fNueva,
		)
	}

	// Corroboro que me lo devuelva igual
	if enString != "null" {
		t.Error("No se obtuvo el string esperado.",
			"\nEsperado: ", "null",
			"\nObtenido: ", enString,
		)
	}
}
