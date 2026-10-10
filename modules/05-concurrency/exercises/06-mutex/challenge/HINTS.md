# Pistas · Mutex y RWMutex

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa snapshot de un mapa protegido que devuelve una copia.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
Lock → read current → update invariant → Unlock
other goroutines wait; then observe ordered state
```

3. Evita este error de la lección: Leer sin lock porque solo se escribe poco. Escribe una predicción para una frontera antes de programar.
