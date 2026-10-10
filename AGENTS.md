# Reglas para mantener Go From Zero

## Responsabilidad

- Escribir material en español; mantener identificadores y términos de Go en inglés.
- Este repo es un libro-curso autodidacta completo. Se puede desarrollar todo el material sin esperar una conversación por lección.
- El estudiante diseña, programa y depura los ejercicios y el proyecto final. No generar soluciones completas ni ocultas, ni stubs que resuelvan sus requisitos.
- Pedir autorización explícita antes de editar implementaciones del estudiante. La autorización para contenido educativo permite ejemplos distintos y tests de aceptación, no soluciones.
- No reescribir, formatear ni arreglar automáticamente el código del estudiante. Ante bugs, explicar evidencia y ayudar a descubrir la causa.

## Pedagogía

- Cada lección incluye los doce apartados del contrato editorial: objetivos, prerrequisitos, motivación, concepto, modelo, ejemplos, comparación, errores, guiado, desafío, preguntas, resumen y referencias.
- Usar ejemplos pequeños, ejecutables y diferentes de ejercicios. Preferir biblioteca estándar.
- Distinguir contrato del lenguaje, modelo conceptual y detalles de implementación/versiones.
- Enunciados con requisitos, restricciones, entradas, salidas y criterios verificables; pistas y respuestas conceptuales separadas. No poner respuestas de implementación en autoevaluaciones.
- El estudiante avanza por autoevaluación; no imponer fechas ni elogios automáticos ni asumir comprensión por compilación.
- Si solicita tutoría: preguntas orientadoras, pistas conceptuales, ejemplo análogo y explicación explícita progresivamente; revisar sin reescribir y hacer 2–4 preguntas conceptuales.

## Verificación y seguimiento

- Mantener Go 1.26 como base del módulo principal. No añadir dependencias externas allí.
- Validar enlaces, contenido, ejemplos, tests y vet; race detector para ejemplos concurrentes. Bugs intencionales separados y con timeout.
- Actualizar PROGRESS.md tras cada sesión: estado editorial, evidencia técnica, conceptos estudiados, ejercicios completados, dificultades observadas, refuerzo y próximo objetivo. Registrar lo no evaluado como pendiente.
- Mantener README, ROADMAP, glosario e índices sincronizados. No presentar el curso completo si faltan lecciones.
- No instalar herramientas globales, arrancar servicios persistentes, hacer commits ni pushes salvo petición expresa.
- En cualquier nuevo scaffold inicial, crear primero documentación; después ejemplos/tests autorizados. Conservar implementaciones existentes del estudiante.
