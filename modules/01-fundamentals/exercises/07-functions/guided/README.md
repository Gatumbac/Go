# Ejercicio guiado · Funciones y retornos

[Lección](../../../07-functions.md) · [Pistas](HINTS.md)

## Objetivo, requisitos, entradas y salidas

Implementa `DivideExact(a, b int) (int, bool)` para enteros no negativos: devuelve cociente y true solo si b>0 y división exacta; en otro caso 0,false. Entrada: (12,3),(7,2),(5,0). Salida: (4,true),(0,false),(0,false). Restricción: sin panic ni I/O. Aceptación: documentar el bool.

## Restricciones de trabajo

- Crea manualmente tus archivos en este directorio; no hay implementación ni stubs.
- Si se pide una función sin main, usa `package practice`; si se pide programa, usa `package main`. Los enunciados de varios packages indican su organización.
- Solo biblioteca estándar, salvo el laboratorio SQL opcional explícito.
- No copies la implementación del ejemplo: resuelve el problema del enunciado.
- No generes archivos de solución ni uses respuestas de autoevaluación como receta.

## Verificación

Para un programa: `go run ./modules/01-fundamentals/exercises/07-functions/guided` desde la raíz. Para funciones: escribe tests propios y ejecuta `go test ./modules/01-fundamentals/exercises/07-functions/guided`. Para entregables Markdown, contrasta el texto con todos los criterios del enunciado.

Si hay `acceptance_test.go`, usa `go test -tags exercise ./modules/01-fundamentals/exercises/07-functions/guided`; sin tu función el error de compilación es esperado. Los tests de aceptación no son un reemplazo de tu explicación ni una solución oculta.

## Criterios de aceptación adicionales

- Cumple los casos y límites definidos arriba; agrega al menos un caso de frontera justificado.
- Explica qué concepto de la lección usaste y qué podría salir mal.
- Conserva comandos y resultados en tus notas; marca [PROGRESS](../../../../../PROGRESS.md) solo cuando tú hayas completado y explicado el trabajo.
