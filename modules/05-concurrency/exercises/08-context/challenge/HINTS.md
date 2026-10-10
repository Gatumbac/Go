# Pistas · Context y cancelación

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa envío cancelable `Send(ctx context.Context,ch chan<- int,n int) error`.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
parent context
  └─ child timeout/cancel → Done closes
                               ↓
                       cooperative work exits
caller join confirms exit
```

3. Evita este error de la lección: Context en struct global. Escribe una predicción para una frontera antes de programar.
