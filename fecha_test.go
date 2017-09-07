package fecha

import (
	"fmt"
	"testing"
	"time"
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

func TestMarshalAndUnmarshal(t *testing.T) {

	// Creo fecha
	f, err := NewFecha("2016-02-19")
	if err != nil {
		t.Error(err)
	}

	// Genero el string
	by, err := f.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	enString := string(by)
	fmt.Println("Marshaled: ", enString)

	// Demarshalizo
	var fNueva Fecha
	err = fNueva.UnmarshalJSON(by)
	if err != nil {
		t.Error(err)
	}

	//Corroboro
	if fNueva != f {
		t.Error("No se obtuvo el string esperado.",
			"\nEsperado: ", f,
			"\nObtenido: ", fNueva,
		)
	}

	// Corroboro que me lo devuelva igual
	if enString != "2016-02-19" {
		t.Error("No se obtuvo el string esperado.",
			"\nEsperado: ", "2016-02-19",
			"\nObtenido: ", enString,
		)
	}
}