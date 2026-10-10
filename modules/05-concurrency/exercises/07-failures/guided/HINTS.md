# Pistas · Carreras, deadlocks y leaks

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Reproduce el contador defectuoso de docs/debugging.md en un temporal, ejecuta con -race y escribe el objeto compartido y las dos stacks relevantes.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
A waits for B
B waits for A → deadlock
producer waits for vanished consumer → potential leak
```

3. Evita este error de la lección: Buffer como arreglo universal. Escribe una predicción para una frontera antes de programar.
