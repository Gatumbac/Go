# Laboratorio de bugs intencionales

[Inicio](../README.md) · [Delve](../modules/04-tooling/06-delve.md) · [Fallos concurrentes](../modules/05-concurrency/07-failures.md)

Estos fragmentos son deliberadamente defectuosos. No son ejemplos normales ni soluciones de ejercicios. Cópialos manualmente a un archivo temporal para investigar. Predice, reproduce con timeout y escribe tu hipótesis; no esperes una reparación automática.

## 1. Máximo negativo

Entrada: [-8,-3,-5]. Esperado por contrato: -3. El programa devuelve otro valor.

```go
package main
import "fmt"
func maximum(values []int) int {
 best := 0
 for _, n := range values { if n > best { best = n } }
 return best
}
func main(){fmt.Println(maximum([]int{-8,-3,-5}))}
```

Pregunta: ¿qué invariante supone la inicialización y cuándo deja de representar un elemento? Pista: compara todos negativos con una lista que contiene cero. Verifica por ti mismo una reparación y decide un contrato para lista vacía.

## 2. Contador con data race

```go
package main
import("fmt";"sync")
func main(){n:=0;var wg sync.WaitGroup;for range 2{wg.Go(func(){for range 1000{n++}})};wg.Wait();fmt.Println(n)}
```

Ejecuta `timeout 20s go run -race /tmp/race-demo.go` tras crear el archivo. No hay salida numérica garantizada; sí hay accesos no sincronizados. Pista: Wait une terminación pero no serializa incrementos. Investiga los accesos en el reporte; no compares solo el total con 2000.

## 3. Recepción sin productor: deadlock

```go
package main
func main(){ch:=make(chan int);<-ch}
```

Ejecuta con `timeout 10s go run /tmp/deadlock-demo.go`. Pregunta: ¿quién podría producir el evento? El runtime normalmente diagnostica este bloqueo global, pero no generalices a un servidor que tenga otras actividades vivas.

## 4. Leak de productor abandonado

```go
package main
import "fmt"
func main(){ch:=make(chan int);go func(){ch<-1;ch<-2}();fmt.Println(<-ch)}
```

El proceso puede terminar normalmente y aun así el productor no pudo completar antes de main. Añadir Sleep no define terminación ni corrige el contrato. Pista: dibuja el segundo envío y decide cómo el productor conoce que no habrá receptor. Registra una política de cancelación y prueba un join en tu alternativa.

## Entrega de investigación

Anota contrato roto, reproducción, hipótesis, evidencia de stack/variables y escenarios de tu reparación. Las salidas defectuosas son evidencia de laboratorio, no progreso del estudiante registrado automáticamente.
