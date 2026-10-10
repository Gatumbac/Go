# Verificación del material

[Inicio](../README.md) · [Estado y evidencia](../PROGRESS.md)

## Comprobaciones reproducibles

```sh
python3 scripts/validate_course.py
python3 scripts/verify_examples.py
python3 scripts/verify_examples.py --race
go build ./...
go test -count=1 ./...
go vet ./...
go test -race -count=1 ./modules/04-tooling/examples/... ./modules/05-concurrency/examples/... ./modules/06-backend/examples/...
go test -run '^$' -bench . -benchmem ./modules/04-tooling/examples/04-benchmarks
```

validate_course comprueba manifiesto, doce secciones no vacías, ejemplos, ejercicios, preguntas/respuestas, links locales y anclas, JSON y SVG. También comprueba que los bloques Go de las lecciones coincidan con source entregado. No certifica calidad pedagógica por cantidad de encabezados ni corrección semántica de Mermaid.

verify_examples construye y ejecuta cada ejemplo principal en un directorio temporal y exige finalización con timeout y comprueba salidas deterministas, orden permitido o formato de métricas variables. No recorre código del estudiante ni descarga el driver del laboratorio. Usa GOCACHE y GOMODCACHE en /tmp cuando el entorno tiene cache global restringida. El proceso principal valida escenarios finitos; no inicia servicios persistentes.

La ejecución instrumentada de los nueve ejemplos concurrentes se realiza con `python3 scripts/verify_examples.py --race`.

Los tests de ejemplos validan contratos pequeños de testing, concurrencia y handlers. Los tests de ejercicios se habilitan solo con -tags exercise y requieren implementaciones del estudiante. No se ejecutan sobre todo el repo para fingir que ejercicios pendientes están resueltos.

## Bugs intencionales

[Depuración](debugging.md) contiene programas defectuosos fuera de los ejemplos normales. La revisión del libro verifica su comportamiento en temporales con límites: máximo negativo incorrecto, carrera reportada, deadlock reportado y productor abandonado. No se incluyen sus reparaciones.

## Visuales y opcionales

Mermaid conserva fuente textual y alternativas ASCII. Si no hay parser/renderizador local, registrar «no validado automáticamente»; no instalar uno global solo para esta guía. SVG se verifica como XML y se revisa visualmente. Delve requiere instalación optativa; su sesión interactiva no se da por ejecutada si falta dlv.

El módulo PostgreSQL está excluido del recorrido raíz y requiere descargar driver/configurar DB para sus verificaciones. Docker pertenece al proyecto que escribirá el estudiante; no existe una imagen del sistema final entregada por el libro.
