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

func TestNewDate(t *testing.T) {
	// Creo fecha
	f, err := NewFecha("2016-02-19")
	if err != nil {
		t.Error(err)
	}

	if f != Fecha(20160219) {
		t.Error()
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
