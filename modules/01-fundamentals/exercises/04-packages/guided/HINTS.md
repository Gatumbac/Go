# Pistas · Packages y modules

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea un package `labels` dentro de guided y un subdirectorio `cmd` con main. Define `Title(name string) string` que devuelva `Topic: ` seguido del nombre.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
module example.com/go-from-zero (go.mod)
  ├─ package main: example A
  ├─ package main: example B
  └─ package reusable: exported names
```

3. Evita este error de la lección: Crear go.mod en cada carpeta sin necesidad. Escribe una predicción para una frontera antes de programar.
