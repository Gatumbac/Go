# Ejercicio guiado · Table-driven tests y subtests

[Lección](../../../02-tables.md) · [Pistas](HINTS.md)

## Objetivo, requisitos, entradas y salidas

Implementa y prueba `Clamp(n,low,high int) int`, low<=high. Entrada: (-1,0,10),(5,0,10),(12,0,10). Salida:0,5,10. Restricción: tabla con fronteras incluidas. Aceptación: subtests con nombres únicos.

## Restricciones de trabajo

- Crea manualmente tus archivos en este directorio; no hay implementación ni stubs.
- Si se pide una función sin main, usa `package practice`; si se pide programa, usa `package main`. Los enunciados de varios packages indican su organización.
- Solo biblioteca estándar, salvo el laboratorio SQL opcional explícito.
- No copies la implementación del ejemplo: resuelve el problema del enunciado.
- No generes archivos de solución ni uses respuestas de autoevaluación como receta.

## Verificación

Para un programa: `go run ./modules/04-tooling/exercises/02-tables/guided` desde la raíz. Para funciones: escribe tests propios y ejecuta `go test ./modules/04-tooling/exercises/02-tables/guided`. Para entregables Markdown, contrasta el texto con todos los criterios del enunciado.

Si hay `acceptance_test.go`, usa `go test -tags exercise ./modules/04-tooling/exercises/02-tables/guided`; sin tu función el error de compilación es esperado. Los tests de aceptación no son un reemplazo de tu explicación ni una solución oculta.

## Criterios de aceptación adicionales

- Cumple los casos y límites definidos arriba; agrega al menos un caso de frontera justificado.
- Explica qué concepto de la lección usaste y qué podría salir mal.
- Conserva comandos y resultados en tus notas; marca [PROGRESS](../../../../../PROGRESS.md) solo cuando tú hayas completado y explicado el trabajo.
