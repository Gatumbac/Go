# Pistas · Defer, panic y recover

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa una función `Trace() []string` que mediante defers devuelva secuencia ["body","second","first"].

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
defer A → defer B → return
                     ↓
                 execute B
                 execute A
```

3. Evita este error de la lección: Defer antes de comprobar error. Escribe una predicción para una frontera antes de programar.
