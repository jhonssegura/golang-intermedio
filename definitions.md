# Definiciones Sobre Golang

## ¿Es Go un lenguaje orientado a objetos?

Sí y no. Aunque Go tiene tipos y métodos y permite un estilo de programación orientado a objetos, no existe una jerarquía de tipos. El concepto de **interfaz** en Go proporciona un enfoque diferente. Hay formas de incrustar tipos en otros tipos para proporcionar algo análogo, pero no idéntico, a la creación de subclases. Además, los métodos en Go son más generales que en C++ o Java: se pueden definir para cualquier tipo de datos, incluso tipos integrados, como enteros "sin caja". No están restringidos a estructuras (clases).

Además, la falta de una jerarquía de tipos hace que los "objetos" en Go parezcan mucho más ligeros que en lenguajes como C++ o Java.

## ¿Cómo obtengo el envío dinámico de métodos?
La única manera de tener métodos enviados dinámicamente es a través de una interfaz. Los métodos sobre una estructura o cualquier otro tipo concreto siempre se resuelven estáticamente.

## ¿Por qué no hay herencia de tipos?
La programación orientada a objetos, al menos en los lenguajes más conocidos, implica demasiada discusión sobre las relaciones entre tipos, relaciones que a menudo podrían derivarse automáticamente. Go adopta un enfoque diferente.

En lugar de requerir que el programador declare con anticipación que dos tipos están relacionados, en Go un tipo satisface automáticamente cualquier interfaz que especifique un subconjunto de sus métodos. Además de reducir la contabilidad, este enfoque tiene ventajas reales. Los tipos pueden satisfacer muchas interfaces a la vez, sin las complejidades de la herencia múltiple tradicional. Las interfaces pueden ser muy ligeras: una interfaz con uno o incluso cero métodos puede expresar un concepto útil. Las interfaces se pueden agregar después del hecho si surge una nueva idea o para probar, sin anotar los tipos originales. Debido a que no existen relaciones explícitas entre tipos e interfaces, no existe una jerarquía de tipos para administrar o discutir.

Es posible usar estas ideas para construir algo análogo a las tuberías Unix con seguridad de tipos. Por ejemplo, vea cómo fmt.Fprintf habilita la impresión formateada en cualquier salida, no solo un archivo, o cómo el bufiopaquete puede estar completamente separado de la E/S del archivo, o cómo los imagepaquetes generan archivos de imagen comprimidos. Todas estas ideas surgen de una sola interfaz ( io.Writer) que representa un solo método ( Write). Y eso es solo rascar la superficie. Las interfaces de Go tienen una profunda influencia en cómo se estructuran los programas.

Cuesta un poco acostumbrarse, pero este estilo implícito de dependencia de tipos es una de las cosas más productivas de Go.

## Structs

Son una colección tipada de campos, útiles para agrupaciones de datos que forman registros. Al igual que cualquier tipo en Go, permiten que se le añadan métodos modelando así el comportamiento del mismo de manera similar a como funciona una clase.

```go
package main

import (
    "fmt"
)

func main() {
    var gopher = struct {
        Name string
        Age int
        Color string
    } {"Jack", 10, "Blue"}

    fmt.Println(gopher)
}
```

Esta práctica no es la más común y lo que se suele utilizar son los **typed structs** o **type** aslias **struct**, que vienen a ser como cualquier type alias salvo que en los structs suele pasar desapercibido y entender que es la manera natural de ser.

```go
package main

import (
    "fmt"
)

func main() {
    g := gopher{"Jack", 10, "Blue"}

    fmt.Println(g)
}

type gopher struct {
    Name string
    Age int
    Color string
}
```

### Structs vs Clases

* El objetivo de una clase es definir una serie de propiedades y métodos que un objeto puede usar para alcanzar diferentes objetivos. 

* Go utiliza Structs para generar **nuevos tipos** de datos que se pueden utilizar para cumplir tareas específicas.


## Métodos y Funciones


## Constructores

Permiten la instanciación de una clase a un objeto, permite definir propiedades predefinidas. En go se puede utilizar funciones que puedan crear structs con propiedades que se pasna como parámetros.

## Herencia

Permite reutilizar código e implementar el polimorfismo.

Go aplica composición sobre herencia. A diferencia de la herencia, no es una clase de hija sino que contienen los métodos de las clases indicadas.

Para alcanzar este mismo objetivo, Go aplica un principio llamado Composition Over Inheritance que utilizando un campo anónimo en un struct puede "heredar" todas las propiedades y métodos.

## Interfaces

Diferentes lenguajes de programación utilizan sintaxis explícita para decir que una clase implementa una interfaz. Go lo hace de manera implícita, lo que permite la reutilización de código y el polimorfismo.

### Abstract Factory




## Bibliografía

* https://blog.friendsofgo.tech/posts/orientacion_a_objetos_structs/


