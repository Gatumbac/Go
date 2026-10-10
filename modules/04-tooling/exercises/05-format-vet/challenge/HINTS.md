# Pistas · Gofmt y go vet

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe checks.md clasificando tres fallos: tipo incompatible, formato incorrecto, algoritmo incorrecto. Salida: qué herramienta puede detectarlos y un límite por herramienta. Aceptación: no atribuir corrección semántica a gofmt.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
source → gofmt: layout
       → build: types/syntax
       → vet: suspicious patterns
       → tests: behavior scenarios
```

3. Evita este error de la lección: Vet como prueba completa. Escribe una predicción para una frontera antes de programar.
