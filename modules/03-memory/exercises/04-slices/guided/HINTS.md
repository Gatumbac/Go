# Pistas · Slices y backing arrays

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea un experimento con array [4]int y dos ventanas: una con cap suficiente y otra limitada.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
backing array [1 2 3 4]
a = [0:2] → len 2 cap 4
b = [0:2:2] → len 2 cap 2
append(b, ...) requires different storage
```

3. Evita este error de la lección: Reslice como copia. Escribe una predicción para una frontera antes de programar.
