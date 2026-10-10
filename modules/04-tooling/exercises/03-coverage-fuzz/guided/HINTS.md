# Pistas · Coverage y fuzzing

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe tests con coverage para `SumPositive(values []int) int`, dominio -100..100 y longitud <=100.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
seeds → mutated inputs → property check
                          ↓ violation
                    minimal regression input
```

3. Evita este error de la lección: Coverage como garantía. Escribe una predicción para una frontera antes de programar.
