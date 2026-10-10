# Entorno local

[Inicio](../README.md) · [Método](study-method.md)

## Comprobación antes de instalar

```sh
go version
go env GOVERSION GOOS GOARCH CGO_ENABLED GOMOD
```

El entorno de creación tiene `go1.26.8-X:nodwarf5 linux/amd64`, CGO habilitado y GCC. El sufijo describe este toolchain: no es requisito del lenguaje ni una versión que debas instalar. `go.mod` fija Go 1.26. Mantén un toolchain compatible; no hace falta reinstalar el existente. Para otros equipos sigue [la instalación oficial](https://go.dev/doc/install).

`GOMOD` debe señalar el archivo de este repo. Trabaja desde su raíz, fuera de GOPATH si quieres: los modules no necesitan una ubicación especial. `go env` describe configuración; evita cambiar valores globales sin entender su efecto.

## Run, build y documentación

```sh
go run ./modules/01-fundamentals/examples/02-compilation
go build -o /tmp/go-from-zero-demo ./modules/01-fundamentals/examples/02-compilation
/tmp/go-from-zero-demo
go doc fmt.Println
```

`go run` compila y ejecuta un programa temporal. `go build -o` deja el binario donde indicas. El binario no necesita el comando `go` en el equipo destino, pero sí un sistema/arquitectura compatibles y, según cómo se enlazó, bibliotecas dinámicas. No confundir independencia del SDK con independencia del sistema operativo.

## VS Code y herramientas opcionales

La extensión Go (`golang.go`) ya existe en el entorno inspeccionado. Las recomendaciones del repo no la instalan. Al abrir un Go file, VS Code puede ofrecer instalar herramientas: esa decisión es tuya. `gopls` proporciona navegación, diagnósticos y completado; no es el compilador. El workspace desactiva formateo automático para que tus implementaciones no cambien sin tu decisión.

Si deseas instalarlas tú, revisa sus requisitos actuales:

```sh
go install golang.org/x/tools/gopls@latest
go install github.com/go-delve/delve/cmd/dlv@latest
```

Son herramientas, no dependencias de este curso. `@latest` puede necesitar un Go más nuevo; lee [gopls](https://go.dev/gopls/) y [Delve](https://github.com/go-delve/delve/tree/master/Documentation/installation) antes. No se ejecutan durante la generación. La lección de Delve también ofrece una traza de depuración que puedes estudiar sin instalarlo.

## Validación por separado

`go test ./...` comprueba ejemplos y tests normales. Para un ejercicio implementado, ejecuta el comando específico de su README con `-tags exercise`; no actives ese tag sobre todo el libro mientras haya ejercicios sin implementar. El laboratorio PostgreSQL tiene otro `go.mod`, por lo que estos comandos no descargan su driver.

`-race` requiere un toolchain/plataforma compatibles y, en este entorno, CGO y GCC. Si tu entorno carece de ellos, registra la limitación, no un resultado aprobado.
