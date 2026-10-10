# Pistas · Interfaces y method sets

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Define `Sizer` con `Size() int` y dos tipos que lo satisfagan por valor, `Folder` y `Envelope`, con un campo de tamaño.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
consumer → interface capability ← concrete type
interface value: dynamic type + dynamic value
(nil, nil) differs from (*T, nil)
```

3. Evita este error de la lección: Interfaces enormes. Escribe una predicción para una frontera antes de programar.
