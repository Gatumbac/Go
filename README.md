# Go From Zero

Libro-curso práctico de Go para quien conoce TypeScript, Node.js, Python, Java, C, SQL, Linux y Docker. Estudia desde cero el lenguaje, su memoria, la concurrencia y los servicios backend. Tú implementas los ejercicios y el proyecto final.

## Empieza aquí

1. Abre esta carpeta en VS Code: `code .`.
2. Lee [el entorno](docs/setup.md) y [el método](docs/study-method.md).
3. Sigue [el roadmap completo](ROADMAP.md) desde [la primera lección](modules/01-fundamentals/01-philosophy.md).
4. Abre Markdown Preview con `Ctrl+Shift+V`. Los modelos ASCII y SVG se leen sin servicios externos; los bloques Mermaid son complementarios.
5. Consulta [el glosario](GLOSSARY.md) y marca tu evidencia en [PROGRESS](PROGRESS.md).

## Cómo usar el código

Todos los comandos de lecciones parten de la raíz salvo indicación contraria. Los ejemplos son programas pequeños completos con escenarios distintos de los ejercicios. Cada directorio de ejemplo tiene un solo programa y se ejecuta con `go run ./modules/.../examples/...`.

Los ejercicios no contienen implementación. Crea tú los archivos que pide el enunciado. Algunos tienen pruebas con build tag `exercise`: `go test -tags exercise ./modules/.../exercises/.../guided`. Antes de escribir las funciones requeridas fallarán por símbolos inexistentes: no es un fallo del curso.

```sh
go version
go env GOVERSION GOOS GOARCH CGO_ENABLED GOMOD
go build ./...
go test ./...
go vet ./...
python3 scripts/validate_course.py
python3 scripts/verify_examples.py
```

El recorrido usa Go 1.26 y biblioteca estándar. El laboratorio PostgreSQL se encuentra aislado en [docs/labs/postgres](docs/labs/postgres/README.md); no requiere iniciar servicios para leer el libro. [El proyecto final](final-project/README.md) tiene ocho hitos, sin solución.

## Alcance

49 lecciones, ejemplos ejecutables, 98 enunciados, pistas, autoevaluaciones y un proyecto integrador. No hay calendario ni obligación de conversar para avanzar. Compilar sirve como evidencia técnica; comprender exige explicar y predecir el comportamiento.

La especificación y la biblioteca estándar son fuentes, no sustitutos de las explicaciones. La documentación online evoluciona: las notas de compatibilidad están en [fuentes](docs/sources.md). Consulta [la verificación](docs/verification.md) para sus límites.
