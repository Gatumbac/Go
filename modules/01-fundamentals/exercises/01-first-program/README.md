# Ejercicio 01 · Primer programa

## Objetivo

Reconocer el punto de entrada de un programa y distinguir ejecutar mediante Go de ejecutar un binario compilado.

## Requisitos y restricciones

- Crea manualmente main.go en este directorio.
- Declara package main y una función main sin parámetros ni valores de retorno.
- Usa el package fmt de la biblioteca estándar para imprimir las dos líneas indicadas.
- No añadas dependencias, entrada interactiva ni funcionalidades adicionales.

Entrada: ninguna.

Salida exacta, con un salto de línea después de cada línea:

```text
Hello, Go!
I am learning how Go programs run.
```

## Ejecución

Desde la raíz del proyecto (`/home/gabrieltumbaco/code/Go`):

```sh
go run ./modules/01-fundamentals/exercises/01-first-program
go build -o /tmp/go-from-zero-first-program ./modules/01-fundamentals/exercises/01-first-program
/tmp/go-from-zero-first-program
```

## Criterios de aceptación

- El programa escrito por ti compila y ambas formas de ejecución producen la salida exacta.
- Puedes identificar package main, la función main y el uso de fmt en tu archivo.
- Puedes explicar qué produjo go build y qué sucede al ejecutar el binario.
- Puedes distinguir las responsabilidades del compilador y del runtime.

Cuando termines, comparte tu implementación y la salida de los comandos. Si hay un error, comparte el mensaje exacto y tu hipótesis sobre la causa.

## Preguntas para la revisión

1. ¿Qué papel cumplen package main y la función main, y por qué necesitamos ambos?
2. ¿Qué diferencia hay entre go run y go build? ¿Qué paso hace go run antes de ejecutar?
3. ¿Qué tareas realiza el compilador y cuáles realiza el runtime mientras el programa está en ejecución?

El mentor revisará tu código sin reescribirlo y esperará tus respuestas antes de acordar la siguiente lección.
