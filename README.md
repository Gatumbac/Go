# Go desde cero

Espacio de aprendizaje práctico para un desarrollador backend con experiencia en TypeScript, Node.js, C, Java y Python. El objetivo es comprender, explicar y construir programas en Go de forma independiente.

## Método

Cada tema sigue esta secuencia: concepto → modelo mental → implementación manual → revisión → verificación de comprensión.

El mentor explica y orienta; el estudiante escribe los ejercicios. La ayuda progresa desde una pregunta orientadora hasta una pista conceptual, pseudocódigo o ejemplo análogo y una explicación más explícita. Las soluciones completas requieren una petición expresa.

No hay fechas límite. Compilar no demuestra por sí solo comprensión. El avance requiere comprensión razonable y acuerdo del estudiante; nunca es automático.

## Itinerario

| Módulo | Objetivo |
| --- | --- |
| 01 · Fundamentos | Comprender programas, tipos, funciones y colecciones. |
| 02 · Diseño | Modelar datos, dependencias y errores con las abstracciones de Go. |
| 03 · Herramientas | Evaluar corrección, rendimiento y calidad del código. |
| 04 · Concurrencia | Coordinar trabajo y gestionar su ciclo de vida. |
| 05 · Backend | Construir servicios HTTP con la biblioteca estándar. |
| 06 · Integrador | Diseñar una cola de tareas y un worker pool con cancelación y observación. |

Los README de los módulos describen su alcance. La teoría y los ejercicios se presentan cuando corresponde.

## Entorno y primer paso

La primera inspección encontró Go `go1.26.8-X:nodwarf5 linux/amd64`. El módulo usa Go 1.26 y no tiene dependencias externas.

Trabaja desde este directorio. Empieza por [el primer ejercicio](modules/01-fundamentals/exercises/01-first-program/README.md). Su archivo `main.go` debes crearlo tú.

La raíz del proyecto y del repositorio Git es `/home/gabrieltumbaco/code/Go`. El módulo conserva el nombre `example.com/go-from-zero`. [PROGRESS.md](PROGRESS.md) registra el avance observado, no un nivel de dominio supuesto.
