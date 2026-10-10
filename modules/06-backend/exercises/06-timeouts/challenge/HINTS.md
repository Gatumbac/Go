# Pistas · Context y timeouts

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe shutdown.md para una API con requests y tareas aceptadas.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
request ctx → child deadline → dependency
service shutdown → stop admission → bounded wait → cancel → join
```

3. Evita este error de la lección: Request ctx para jobs duraderos. Escribe una predicción para una frontera antes de programar.
