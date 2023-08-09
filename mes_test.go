package fecha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "mes inválido", Mes{}.String())
	assert.Equal(t, "08/2020", Mes{2020, 8}.String())
	assert.Equal(t, "mes inválido", Mes{AñoMinimo - 1, 8}.String())
	assert.Equal(t, "mes inválido", Mes{AñoMaximo + 1, 8}.String())
}

func TestZero(t *testing.T) {
	assert.True(t, Mes{}.Zero())
	assert.False(t, Mes{1, 0}.Zero())
	assert.False(t, Mes{2020, 8}.Zero())
}

func TestValid(t *testing.T) {
	assert.False(t, Mes{}.Valid())
	assert.False(t, Mes{1, 0}.Valid())
	assert.False(t, Mes{AñoMinimo - 1, 1}.Valid())
	assert.False(t, Mes{AñoMaximo + 1, 1}.Valid())
	assert.False(t, Mes{2020, 0}.Valid())
	assert.False(t, Mes{2020, 13}.Valid())
	assert.True(t, Mes{2020, 1}.Valid())
	assert.True(t, Mes{2020, 2}.Valid())
	assert.True(t, Mes{2020, 3}.Valid())
	assert.True(t, Mes{2020, 4}.Valid())
	assert.True(t, Mes{2020, 5}.Valid())
	assert.True(t, Mes{2020, 6}.Valid())
	assert.True(t, Mes{2020, 7}.Valid())
	assert.True(t, Mes{2020, 8}.Valid())
	assert.True(t, Mes{2020, 9}.Valid())
	assert.True(t, Mes{2020, 10}.Valid())
	assert.True(t, Mes{2020, 11}.Valid())
	assert.True(t, Mes{2020, 12}.Valid())
}

func TestUltimoDia(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(31, ultimoDia(1, 2023))
	assert.Equal(28, ultimoDia(2, 2023))
	assert.Equal(31, ultimoDia(3, 2023))
	assert.Equal(30, ultimoDia(4, 2023))
	assert.Equal(31, ultimoDia(5, 2023))
	assert.Equal(30, ultimoDia(6, 2023))
	assert.Equal(31, ultimoDia(7, 2023))
	assert.Equal(31, ultimoDia(8, 2023))
	assert.Equal(30, ultimoDia(9, 2023))
	assert.Equal(31, ultimoDia(10, 2023))
	assert.Equal(30, ultimoDia(11, 2023))
	assert.Equal(31, ultimoDia(12, 2023))

	assert.Equal(31, ultimoDia(1, 2024))
	assert.Equal(29, ultimoDia(2, 2024))
	assert.Equal(31, ultimoDia(3, 2024))
	assert.Equal(30, ultimoDia(4, 2024))
	assert.Equal(31, ultimoDia(5, 2024))
	assert.Equal(30, ultimoDia(6, 2024))
	assert.Equal(31, ultimoDia(7, 2024))
	assert.Equal(31, ultimoDia(8, 2024))
	assert.Equal(30, ultimoDia(9, 2024))
	assert.Equal(31, ultimoDia(10, 2024))
	assert.Equal(30, ultimoDia(11, 2024))
	assert.Equal(31, ultimoDia(12, 2024))
}

func TestNeMesFromJSON(t *testing.T) {

	{
		str := "2021-7"
		_, err := NewMesFromJSON(str)
		assert.NotNil(t, err)
	}
	{
		str := "2021-13"
		_, err := NewMesFromJSON(str)
		assert.NotNil(t, err)
	}
	{
		str := "202113"
		_, err := NewMesFromJSON(str)
		assert.NotNil(t, err)
	}
	{
		str := "2021-13"
		_, err := NewMesFromJSON(str)
		assert.NotNil(t, err)
	}
	{
		str := "2021-00"
		_, err := NewMesFromJSON(str)
		assert.NotNil(t, err)
	}
	{
		str := "0000-01"
		_, err := NewMesFromJSON(str)
		assert.NotNil(t, err)
	}
	{
		str := "2021-01"
		_, err := NewMesFromJSON(str)
		assert.Nil(t, err)
	}
	{
		f := NewMesMust(2021, 1)
		assert.Equal(t, Mes{año: 2021, mes: 1}, f)
	}
	{
		f := NewMesMust(2007, 12)
		assert.Equal(t, Mes{año: 2007, mes: 12}, f)
	}
}

func TestSumarMeses(t *testing.T) {
	mesCero := Mes{2020, 8}

	// Cero meses
	{
		added := mesCero.SumarMeses(0)
		assert.Equal(t, Mes{2020, 8}, added)
	}

	// Positivos
	{
		added := mesCero.SumarMeses(1)
		assert.Equal(t, Mes{2020, 9}, added)
	}
	{
		added := mesCero.SumarMeses(2)
		assert.Equal(t, Mes{2020, 10}, added)
	}
	{
		added := mesCero.SumarMeses(3)
		assert.Equal(t, Mes{2020, 11}, added)
	}
	{
		added := mesCero.SumarMeses(4)
		assert.Equal(t, Mes{2020, 12}, added)
	}
	{
		added := mesCero.SumarMeses(5)
		assert.Equal(t, Mes{2021, 1}, added)
	}
	{
		added := mesCero.SumarMeses(6)
		assert.Equal(t, Mes{2021, 2}, added)
	}
	{
		added := mesCero.SumarMeses(7)
		assert.Equal(t, Mes{2021, 3}, added)
	}
	{
		added := mesCero.SumarMeses(8)
		assert.Equal(t, Mes{2021, 4}, added)
	}
	{
		added := mesCero.SumarMeses(9)
		assert.Equal(t, Mes{2021, 5}, added)
	}
	{
		added := mesCero.SumarMeses(10)
		assert.Equal(t, Mes{2021, 6}, added)
	}
	{
		added := mesCero.SumarMeses(11)
		assert.Equal(t, Mes{2021, 7}, added)
	}
	{
		added := mesCero.SumarMeses(12)
		assert.Equal(t, Mes{2021, 8}, added)
	}
	{
		added := mesCero.SumarMeses(13)
		assert.Equal(t, Mes{2021, 9}, added)
	}
	{
		added := mesCero.SumarMeses(24)
		assert.Equal(t, Mes{2022, 8}, added)
	}

	// Negativos
	{
		added := mesCero.SumarMeses(-1)
		assert.Equal(t, Mes{2020, 7}, added)
	}
	{
		added := mesCero.SumarMeses(-2)
		assert.Equal(t, Mes{2020, 6}, added)
	}
	{
		added := mesCero.SumarMeses(-3)
		assert.Equal(t, Mes{2020, 5}, added)
	}
	{
		added := mesCero.SumarMeses(-4)
		assert.Equal(t, Mes{2020, 4}, added)
	}
	{
		added := mesCero.SumarMeses(-5)
		assert.Equal(t, Mes{2020, 3}, added)
	}
	{
		added := mesCero.SumarMeses(-6)
		assert.Equal(t, Mes{2020, 2}, added)
	}
	{
		added := mesCero.SumarMeses(-7)
		assert.Equal(t, Mes{2020, 1}, added)
	}
	{
		added := mesCero.SumarMeses(-8)
		assert.Equal(t, Mes{2019, 12}, added)
	}
	{
		added := mesCero.SumarMeses(-9)
		assert.Equal(t, Mes{2019, 11}, added)
	}
	{
		added := mesCero.SumarMeses(-10)
		assert.Equal(t, Mes{2019, 10}, added)
	}
	{
		added := mesCero.SumarMeses(-11)
		assert.Equal(t, Mes{2019, 9}, added)
	}
	{
		added := mesCero.SumarMeses(-12)
		assert.Equal(t, Mes{2019, 8}, added)
	}
	{
		added := mesCero.SumarMeses(-13)
		assert.Equal(t, Mes{2019, 7}, added)
	}
	{
		added := mesCero.SumarMeses(-24)
		assert.Equal(t, Mes{2018, 8}, added)
	}
}

func TestPosterior(t *testing.T) {
	assert := assert.New(t)
	f1 := NewMesMust(2022, 7)
	f2 := NewMesMust(2022, 8)

	assert.True(f2.Posterior(f1))
	assert.False(f1.Posterior(f2))
}

func TestPosteriorOIgual(t *testing.T) {
	assert := assert.New(t)
	f1 := NewMesMust(2022, 7)
	f2 := NewMesMust(2022, 8)

	assert.True(f2.PosteriorOIgual(f1))
	assert.False(f1.PosteriorOIgual(f2))

	f2 = NewMesMust(2022, 7)
	assert.True(f1.PosteriorOIgual(f1))
	assert.True(f2.PosteriorOIgual(f2))
}

func TestAnterior(t *testing.T) {
	assert := assert.New(t)
	f1 := NewMesMust(2021, 3)
	f2 := NewMesMust(2030, 2)

	assert.True(f1.Anterior(f2))
	assert.False(f2.Anterior(f1))
}

func TestAnteriorOIgual(t *testing.T) {
	assert := assert.New(t)
	f1 := NewMesMust(2022, 7)
	f2 := NewMesMust(2022, 8)

	assert.True(f1.AnteriorOIgual(f2))
	assert.False(f2.AnteriorOIgual(f1))

	f2 = NewMesMust(2022, 7)
	assert.True(f1.AnteriorOIgual(f1))
	assert.True(f2.AnteriorOIgual(f2))
}

func TestMarshalAndUnmarshalMes(t *testing.T) {

	m := Mes{2020, 8}

	// Genero el string
	by, err := m.MarshalJSON()
	assert.Nil(t, err)

	enString := string(by)

	// Corroboro string
	assert.Equal(t, "2020-08", enString[1:8])

	// Demarshalizo
	var nuevoMes Mes
	err = nuevoMes.UnmarshalJSON(by)
	assert.Nil(t, err)

	// Corroboro Mes
	assert.Equal(t, m, nuevoMes)
}

func TestMarshalZero(t *testing.T) {

	{ // Mes cero
		m := new(Mes)

		// Genero el string
		by, err := m.MarshalJSON()
		assert.Nil(t, err)

		enString := string(by)

		// Corroboro string
		assert.Equal(t, "null", enString)
	}

	// Si estoy Mes nil
}

func TestUnmarshalMes(t *testing.T) {

	{ // Fecha válida
		m := &Mes{}

		err := m.UnmarshalJSON([]byte(`"2020-08"`))
		assert.Nil(t, err)
		assert.Equal(t, Mes{2020, 8}, *m)
	}

	{ // null
		m := &Mes{}

		err := m.UnmarshalJSON([]byte(`null`))
		assert.Nil(t, err)
		assert.Equal(t, Mes{}, *m)
	}
	{ // Fecha vacía
		m := &Mes{}

		err := m.UnmarshalJSON([]byte(`""`))
		assert.Nil(t, err)
		assert.Equal(t, Mes{}, *m)
	}

	{ // Fecha inválida
		m := &Mes{}

		err := m.UnmarshalJSON([]byte(`"9999-99"`))
		assert.NotNil(t, err)
	}

}
