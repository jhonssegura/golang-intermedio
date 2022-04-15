package main

import "fmt"

// Go preinicializa con valores 0, vacío o false, dependiendo del tipo de dato declarado.
type Employee struct {
	id       int
	name     string
	vacation bool
}

// Trabajar un struct a una función
func NewEmployee(id int, name string, vacation bool) *Employee {
	// Es importante devolver la referencia
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Forma de iniciar N°1, dejando los valores por defecto que Go genera.
	e := Employee{}
	fmt.Printf("%v\n", e)
	// Forma de iniciar N°2, asignando los valores que se desea.
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	// Forma de iniciar N°3, utilizar la palabra reservada para que devuelva un apuntador de la instancia definida
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	// Forma de iniciar N°4, trabajar el struct a partir de una función.
	e4 := NewEmployee(3, "Name large", true)
	fmt.Printf("%v\n", *e4)
}
