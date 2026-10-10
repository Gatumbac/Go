# Pistas · Race detector

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Tras el módulo concurrente, crea un contador protegido y tests bajo -race.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
goroutine A write ──?── goroutine B read
                 no synchronization → data race
lock/channel establishes ordered access
```

3. Evita este error de la lección: Run limpio como prueba absoluta. Escribe una predicción para una frontera antes de programar.
