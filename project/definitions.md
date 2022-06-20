# Project

## Definiendo Workers, Jobs y Dispatchers

Golang permite utilizar conceptos como los worker pools y, en combinación con buffered channels, la creación de job queues que utilizando concurrencia nos permitirán tener un alto rendimiento a la hora de la creación de muchas tareas.

* El **dispatcher** recibe todos los jobs, se puede decir que es como el componente global

* Cada **worker** tiene su canal de jobs, y saben cual es el canal del disptacher, es decir el workerpool es el mismo canal para todos los workers.

* Cada worker esta enviando su canal al canal del dispatcher

* En la medida que el dispatcher recibe jobs este los va repartiendo entre los workers a través de sus canales

## Creando Web Server para Procesar Jobs