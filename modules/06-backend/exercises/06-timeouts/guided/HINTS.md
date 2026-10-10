# Pistas · Context y timeouts

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa handler que llame una dependencia `Lookup(ctx)` inyectada; un error de deadline se traduce según contrato a 504 y cancelación a 408 en tests locales.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
request ctx → child deadline → dependency
service shutdown → stop admission → bounded wait → cancel → join
```

3. Evita este error de la lección: Request ctx para jobs duraderos. Escribe una predicción para una frontera antes de programar.
