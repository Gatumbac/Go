# Pistas · Benchmarks

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `JoinWords(words []string) string` con separador coma y mide dos estrategias tuyas.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
same workload → variant A / variant B
                  ↓ measure repeated runs
          correctness + meaningful improvement
```

3. Evita este error de la lección: Benchmarks sin verificar resultado. Escribe una predicción para una frontera antes de programar.
