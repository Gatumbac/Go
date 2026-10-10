# Pistas · Goroutines y scheduler

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea un programa con tres goroutines que calculen el doble de 1,2,3 y entreguen resultados por channel.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
main → launch G → receive result → return
             ↓
       runtime schedules on threads
```

3. Evita este error de la lección: Sleep para esperar. Escribe una predicción para una frontera antes de programar.
