# fecha

Este package permite trabajar con **fechas** y **períodos mensuales**.

```go

    _, _ := fecha.NewFechaFromString("2020-08-23")
    _ := fecha.NewFechaFromInts(2020, 8, 23)
    _ := fecha.NewFechaFromLayout("02/01/2006", "23/08/2020")
    _ := fecha.NewFechaFromTime(time.Now())

```

Utiliza con un int como tipo subyacente, por lo que las comparaciones son fáciles:

```go
    f1, _ := fecha.NewFechaFromString("2020-08-23")
    f1, _ := fecha.NewFechaFromString("2020-08-24")

    _ := f1 > f0 // true
    _ := f1 == f2 // false

    f3, _ := fecha.NewFechaFromString("2020-08-23")
    _ := f1 == f3 // true
```

Operaciones con fechas:

```go

_ = f.AgregarDias(5)  // 2020-08-26

_ = f.AgregarDiasHabiles(5) // 2020-08-28

// Diferencia dias
f0 := fecha.Fecha(20190823)
dias := fecha.Diff(f0, f1)   // 366
```
