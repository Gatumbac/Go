# Pistas · Structs

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Define `Product` con Name string y PriceCents int. Implementa `Valid(p Product) bool`: nombre no vacío y precio >=0. Entradas: {"pen",100}, {"",1}, {"pen",-1}. Salidas: true,false,false. Restricción: no I/O. Aceptación: documentar si precio cero es válido (sí).

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
Book value
  Title string
  Pages int
assignment → field values copied
```

3. Evita este error de la lección: Usar literales posicionales frágiles. Escribe una predicción para una frontera antes de programar.
