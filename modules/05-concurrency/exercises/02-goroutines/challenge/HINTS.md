# Pistas · Goroutines y scheduler

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe scheduler-notes.md explicando una carga con 1000 goroutines y GOMAXPROCS=2.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
main → launch G → receive result → return
             ↓
       runtime schedules on threads
```

3. Evita este error de la lección: Sleep para esperar. Escribe una predicción para una frontera antes de programar.
