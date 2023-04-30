# Solution

1. El primer paso para implementar el control de acceso a la sala de conferencias en Go es definir la estructura y las variables necesarias para el programa.
Esto incluye crear una estructura llamada `ConferenceRoom` que represente la sala de conferencias, con campos para la capacidad máxima, el contador de personas y un mutex para garantizar el acceso seguro.

Además, se deben declarar dos canales, `enterCh` y `exitCh`, para sincronizar las operaciones de entrada y salida. Estos canales permitirán que las personas intenten ingresar o salir de la sala, y serán utilizados junto con la instrucción `select` para manejar las operaciones de manera no bloqueante.
..
2. El siguiente paso consiste en implementar las funciones `enterConferenceRoom()` y `exitConferenceRoom()` para controlar el acceso a la sala de conferencias.

Dentro de la función `enterConferenceRoom()`, se debe enviar un valor al canal `enterCh` para intentar permitir la entrada a la sala. Luego, utilizando un bloque `select`, se puede verificar si hay espacio disponible en la sala.

En caso afirmativo, se adquiere el bloqueo del mutex, se incrementa el contador de personas en la sala y se muestra un mensaje indicando que una persona ingresó.
Si la sala está llena, se muestra un mensaje indicando que la persona debe esperar.

La función `exitConferenceRoom()` sigue un enfoque similar, donde se envía un valor al canal `exitCh` para intentar permitir la salida de la sala.

Se utiliza el bloque `select` para verificar si hay personas en la sala y, en ese caso, se adquiere el bloqueo del mutex, se decrementa el contador de personas y se muestra un mensaje indicando que una persona salió de la sala.

Si la sala está vacía, se muestra un mensaje indicando que no hay personas para salir. Una vez que se haya completado la implementación de estas funciones, se puede continuar con los siguientes pasos para finalizar el control de acceso a la sala de conferencias.
..
3. El siguiente paso consiste en crear una goroutine para controlar el acceso a la sala de conferencias.
Esta goroutine se encargará de escuchar los canales `enterCh` y `exitCh` utilizando la instrucción `select`.

En cada caso, se adquirirá el bloqueo del mutex correspondiente y se realizarán las operaciones necesarias para permitir la entrada o salida de una persona de la sala de conferencias.

Por ejemplo, cuando se recibe un valor en `enterCh`, se verificará si hay espacio disponible en la sala. En caso afirmativo, se incrementará el contador de personas y se mostrará un mensaje indicando que alguien ha ingresado a la sala. De manera similar, cuando se recibe un valor en `exitCh`, se verificará si hay personas en la sala y, en ese caso, se decrementará el contador y se mostrará un mensaje indicando que alguien ha salido de la sala.

Esta goroutine se ejecutará de forma continua, escuchando los canales y respondiendo adecuadamente a las operaciones de entrada y salida.
Una vez que se haya creado esta goroutine, se habrá completado la implementación del control de acceso a la sala de conferencias.
..
4. Además, es importante incluir una simulación de personas intentando ingresar y salir de la sala. Esto se puede hacer mediante un bucle o ciclo en el que se llamen a las funciones `enterConferenceRoom()` y `exitConferenceRoom()` en un orden específico.

Puedes definir la lógica de simulación que mejor se adapte a tu proyecto, como permitir que un número determinado de personas intente ingresar y luego realizar una combinación de entradas y salidas en diferentes momentos. La simulación de personas intentando ingresar y salir te permitirá probar el funcionamiento del control de acceso y verificar si se están cumpliendo los requisitos establecidos, como la capacidad máxima de la sala. Una vez que hayas realizado la simulación y estés satisfecho con los resultados, puedes considerar el proyecto completado.
..

## Solutions Step by Step

### ConferenceRoom

Define la estructura ConferenceRoom, que representa la sala de conferencias en el proyecto. Una estructura es una colección de campos que representan un conjunto de datos relacionados. La palabra clave type se utiliza para definir un nuevo tipo en Go. Después de type, se especifica el nombre del nuevo tipo, que en este caso es ConferenceRoom. Luego, se utiliza la palabra clave struct para indicar que ConferenceRoom es una estructura.

Segudiamente y dentro de {} especificamos los campos de la estructura. Los campos son variables que representan los datos que se desean almacenar en la estructura.

- **capacity int:** Es un campo entero que representa la capacidad máxima de la sala de conferencias. Este valor indica cuántas personas pueden estar presentes en la sala al mismo tiempo.

- **count int:** Es un campo entero que representa el contador de personas actualmente en la sala de conferencias. Este valor se actualiza cada vez que una persona entra o sale de la sala.

- **mutex sync.Mutex:** Es un campo del tipo **sync.Mutex**, que se utiliza como un mecanismo de bloqueo para garantizar el acceso seguro a la variable count. El mutex se adquiere y libera para asegurar que solo una goroutine tenga acceso a la vez a la variable compartida.

- **enterCh chan struct{}:** Es un canal utilizado para permitir la entrada a la sala de conferencias. Las goroutines que desean ingresar a la sala enviarán un valor vacío **(struct{})** a este canal para solicitar el acceso.

- **exitCh chan struct{}:** Es un canal utilizado para permitir la salida de la sala de conferencias. Las goroutines que desean salir de la sala enviarán un valor vacío **(struct{})** a este canal para solicitar la salida.

En resumen, la estructura **ConferenceRoom** contiene los campos necesarios para representar la capacidad, el contador de personas y los mecanismos de sincronización **(mutex)** y comunicación **(canales)** utilizados para controlar el acceso a la sala de conferencias de manera segura y coordinada.

### func enterConferenceRoom()

la función enterConferenceRoom() se encarga de solicitar la entrada a la sala de conferencias enviando un valor vacío al canal enterCh. Esto permite que la goroutine responsable de controlar el acceso procese la solicitud y tome las medidas adecuadas según la capacidad de la sala y el estado actual.

### func exitConferenceRoom()

Al igual que la función anterior, la función exitConferenceRoom() se encarga de solicitar la salida de la sala de conferencias enviando un valor vacío al canal exitCh.
**cr.exitCh <- struct{}{}**: Intenta enviar un valor vacío (struct{}{}) al canal exitCh de la sala de conferencias.
En otras palabras, la función está solicitando salir de la sala.
Al enviar un valor al canal, se notifica a la goroutine que controla el acceso a la sala que alguien está intentando salir.

### func manageAccess()

La función `manageAccess(cr *ConferenceRoom)` es una goroutine que se encarga de gestionar el acceso a la sala de conferencias. A continuación se detalla su funcionalidad:

- `func manageAccess(cr *ConferenceRoom)`: Esto indica que `manageAccess()` es una función que recibe un puntero a una instancia de `ConferenceRoom` como argumento.

- `for { ... }`: Este bucle infinito permite que la goroutine se ejecute continuamente y esté lista para manejar las solicitudes de entrada y salida de la sala.

- `select { ... }`: El bloque `select` permite que la goroutine escuche simultáneamente los canales `cr.enterCh` y `cr.exitCh` y responda a la operación que esté lista para ser procesada.

- `case <-cr.enterCh: ...`: Este caso se activa cuando se recibe un valor del canal `cr.enterCh`, lo que indica un intento de entrada a la sala. En este caso, la goroutine adquiere el bloqueo del mutex (`cr.mutex.Lock()`) para garantizar el acceso exclusivo a la variable compartida `cr.count`. Luego, se verifica si hay espacio disponible en la sala. Si es así, se incrementa el contador `cr.count`, lo que representa la entrada de una persona, y se muestra un mensaje indicando que alguien ha ingresado. Si la sala está llena, se muestra un mensaje indicando que la persona debe esperar.

- `case <-cr.exitCh: ...`: Este caso se activa cuando se recibe un valor del canal `cr.exitCh`, lo que indica un intento de salida de la sala. Al igual que en el caso anterior, se adquiere el bloqueo del mutex para garantizar el acceso exclusivo a `cr.count`. Se verifica si hay personas en la sala (`cr.count > 0`). Si es así, se decrementa el contador `cr.count`, lo que representa la salida de una persona, y se muestra un mensaje indicando que alguien ha salido. Si la sala está vacía, se muestra un mensaje indicando que no hay personas para salir.

- `cr.mutex.Unlock()`: Después de realizar las operaciones de entrada o salida y actualizar `cr.count`, se libera el bloqueo del mutex con `cr.mutex.Unlock()` para permitir que otras goroutines accedan a la variable compartida `cr.count`.

En resumen, la función `manageAccess()` es una goroutine que escucha los canales de entrada y salida de la sala de conferencias y toma las acciones necesarias para permitir o denegar el acceso. Utiliza el mutex para garantizar la exclusión mutua en el acceso a `cr.count` y muestra mensajes apropiados para informar sobre los intentos de entrada y salida de la sala.

## Simulación de acceso a una sala de conferencias

Seguidamente, simulamos el acceso a la sala de conferencias de la siguiente manera:

1. Se define la estructura `ConferenceRoom` con los valores correspondientes, que representa la sala de conferencias. Contiene campos como la capacidad de la sala, el número de personas actualmente en la sala y dos canales para controlar el acceso.
Aquí Se crea una instancia de la estructura `ConferenceRoom` llamada `cr` y se inicializan sus campos, incluyendo los canales de entrada y salida.

2. Se inicia una goroutine llamada `manageAccess(cr)` que se encargará de controlar el acceso a la sala de conferencias.

3. A continuación, se simula el acceso a la sala mediante un bucle `for` en el que se alternan los intentos de entrada y salida de personas. Dependiendo de si el número es par o impar, se llama a los métodos `enterConferenceRoom()` o `exitConferenceRoom()` respectivamente.

4. Se utiliza un objeto `sync.WaitGroup` llamado `wg` para esperar a que todas las personas hayan terminado de ingresar y salir de la sala. Se inicializa con un contador de 15.

5. Se inicia otro bucle `for` que crea goroutines adicionales para simular más personas intentando ingresar y salir de la sala. Dentro de cada goroutine se llama a los métodos `enterConferenceRoom()` o `exitConferenceRoom()` según corresponda.

6. Utilizando `defer wg.Done()` se marca la finalización de cada goroutine al completar la operación de entrada o salida.

7. Después de iniciar todas las goroutines, se llama a `wg.Wait()` para esperar a que todas las goroutines terminen y se completen las operaciones de entrada y salida.

8. Finalmente, se imprime "Programa finalizado" para indicar que todas las operaciones han finalizado.
