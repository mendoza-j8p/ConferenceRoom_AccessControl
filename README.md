# Proyecto: Control de acceso a una sala de conferencias

En este proyecto, crearás un programa en Go para controlar el acceso a una sala de conferencias. El objetivo es asegurarse de que solo un número limitado de personas pueda ingresar a la sala al mismo tiempo.

## Requisitos

- La sala de conferencias tiene una capacidad máxima de 10 personas.
- Cuando una persona desea ingresar a la sala, debe llamar a una función `enterConferenceRoom()` que intentará permitir el acceso si hay espacio disponible.
- Si la sala está llena, la persona debe esperar hasta que haya espacio disponible.
- Cuando una persona sale de la sala, debe llamar a una función `exitConferenceRoom()` para liberar su lugar y permitir que otra persona entre.
- Utiliza canales para sincronizar el acceso a la sala de conferencias.
- Utiliza la instrucción `select` para manejar las operaciones de entrada y salida de la sala de conferencias de manera no bloqueante.
- Utiliza un mutex para garantizar el acceso seguro a la variable que mantiene el número de personas en la sala de conferencias.

Puedes comenzar definiendo la estructura y las variables necesarias para el programa, incluyendo el canal, el mutex y la variable para mantener el conteo de personas en la sala de conferencias.

A partir de ahí, implementa las funciones `enterConferenceRoom()` y `exitConferenceRoom()` utilizando los canales, la instrucción `select` y el mutex para garantizar el comportamiento correcto del control de acceso.

Recuerda considerar los casos en los que las personas esperan a entrar o salir de la sala cuando está llena o vacía, respectivamente.

## Puntos clave

**Canales en Go**
Los canales en Go proporcionan una forma de comunicación y sincronización entre goroutines (hilos ligeros). Permiten enviar y recibir valores entre goroutines y aseguran
que la comunicación ocurra de manera segura y sin problemas de concurrencia.

**Instrucción select**
La instrucción select en Go permite seleccionar de manera no bloqueante entre múltiples canales. Se utiliza para esperar hasta que alguna de las operaciones de envío o
recepción en los canales esté lista para continuar.

**Mutex (Mutexes)**
Un mutex (mutual exclusion) es una estructura de datos utilizada para sincronizar el acceso a recursos compartidos por múltiples goroutines. Ayuda a evitar problemas de
concurrencia y garantiza que solo una goroutine pueda acceder a un recurso compartido en un momento dado.
