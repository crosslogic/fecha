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
