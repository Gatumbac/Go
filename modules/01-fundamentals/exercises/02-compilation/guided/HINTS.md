# Pistas · Compilador, binario y runtime

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea manualmente `main.go`, package main, que imprima `build experiment`.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
source .go → type checking → object code → linker → executable
                                                       |
                                                       v
                                             main + Go runtime
```

3. Evita este error de la lección: Intentar ejecutar un binario de otro OS. Escribe una predicción para una frontera antes de programar.
