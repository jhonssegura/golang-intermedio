# Testing

Ejecutar comando para test

    $ go test

## Code Coverage

Es una métrica que nos permite identificar las partes de nuestro código que no han sido probadas y que potencialmente tienen el riesgo de contener un bug.

**Comandos útiles para coverage:**

Comando que indica el porcentaje de código que está cubierto por pruebas

    $ go test -cover

Crea un archivo el cual permitirá tener métricas más claras y especificas de lo que pasa con las pruebas en el código

    $ go test -coverprofile=coverage.out

Obtener un mayor detalle del archivo coverage.out

    $ go tool cover -func=coverage.out

Muestra en nuestro navegador todas las funciones que tienen y no tienen las pruebas aplicadas 

    $ go tool cover -html=coverage.out

## Profiling

El profiling nos ayudará a encontrar en nuestro código aquellas partes críticas que influyen en una ejecución lenta o con mucho consumo de memoria. Con estas técnicas se podrá saber en qué enfocarse a la hora de mejorar nuestro código

**Comandos útiles para profilin:**

Crea un archivo el cual almacenará el detalle de la ejecución del test

$ go test -cpuprofile=cpu.out

Permite visualizar el detalle ocurrido que se almacenó en el archivo cpu.out, una vez terminada la ejecución se debe escribir top + enter.

$ go tool pprof cpu.out

Con el comando list + función se podrá ver un mayor detalle de ese proceso. Con los comandos web o pdf se genera un reporte del código para poder darle mayor seguimiento. Para ello es necesario tener instalado **graphviz**.

## Mocks

Cuando escribimos un test unitario que tiene dependencias en diferentes servicios, nos vemos en la necesidad de crear "mocks" que emulen comportamientos de esos servicios con el objetivo que nuestro test se encargue de probar la funcionalidad que nos interesa y no la de las dependencias.

### Implementando Mocks

Es muy importante guardar los valores originales de las funciones que se están utilizando para realizar los test en caso de que en test posteriores se decida evaluar otras partes del código que también dependen de esas funciones.