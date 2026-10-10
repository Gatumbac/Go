# Pistas · Race detector

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe race-notes.md sobre dos requests que comprueban stock y luego reservan.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
goroutine A write ──?── goroutine B read
                 no synchronization → data race
lock/channel establishes ordered access
```

3. Evita este error de la lección: Run limpio como prueba absoluta. Escribe una predicción para una frontera antes de programar.
