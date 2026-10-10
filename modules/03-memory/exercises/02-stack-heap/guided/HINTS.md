# Pistas · Stack y heap

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe un programa que devuelva un pointer a un struct local `Sample` con N=8 y lea N tras retornar.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
goroutine call frames → short-lived state
reachable object ─────> lifetime may outlive call
compiler decides storage; source semantics stay valid
```

3. Evita este error de la lección: Scope=vida útil. Escribe una predicción para una frontera antes de programar.
