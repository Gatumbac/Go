# Pistas · WaitGroups

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa dos rondas secuenciales de trabajo y registra cuándo se puede reutilizar WaitGroup.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
Add before launch → worker work → Done
                           main Wait → all finished
```

3. Evita este error de la lección: Add dentro de goroutine. Escribe una predicción para una frontera antes de programar.
