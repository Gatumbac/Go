# Pistas · Concurrencia y paralelismo

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Diseña un presupuesto para 10000 requests con límite de 20 operaciones activas.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
one core: A → B → A → B (concurrent scheduling)
two cores: core1 AAAAA
           core2 BBBBB (possible parallel execution)
```

3. Evita este error de la lección: Orden de logs como prueba de paralelismo. Escribe una predicción para una frontera antes de programar.
