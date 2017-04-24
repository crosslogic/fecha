package fecha

import "testing"

func TestMarshal(t *testing.T) {

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

	// Demarshalizo
	fNueva := Fecha{}
	err = fNueva.UnmarshalJSON(by)
	if err != nil {
		t.Error(err)
	}

	// Corroboro
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
