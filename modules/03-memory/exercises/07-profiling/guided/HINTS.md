# Pistas · Asignaciones y profiling

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe un programa que compare AllocsPerRun al construir un slice de tamaño 100 con y sin capacidad inicial.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
representative workload → measurement → profile
                              ↓
                         hypothesis → change → compare
```

3. Evita este error de la lección: CPU profile para latencia de red. Escribe una predicción para una frontera antes de programar.
