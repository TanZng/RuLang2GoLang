# Ru2Go Translator

🚧 WIP

## 🚀 Run

```bash
$ docker-compose -f docker-compose-production.yaml build
$ docker-compose -f docker-compose-production.yaml up
# output
# Creating rulang2golang_goservice_1 ... done
# Creating rulang2golang_nginx_1     ... done
# Attaching to rulang2golang_goservice_1, rulang2golang_nginx_1
# ...
# nginx_1      | /docker-entrypoint.sh: Configuration complete; ready for start up
# goservice_1  | 2022/05/19 16:40:26 stdout: 🚀 Serving at  http://localhost/ru/
```

**Open http://localhost/ru/ . Enjoy!**

To stop run
```bash
$ docker-compose down
```

## 💻 Development

### 🐋 Docker
```bash
$ docker-compose build
$ docker-compose up
# output
# go-api_1  | 2022/05/19 16:35:48 Running build command!
# go-api_1  | 2022/05/19 16:35:49 Build ok.
# go-api_1  | 2022/05/19 16:35:49 Restarting the given command.
# go-api_1  | 2022/05/19 16:35:49 stdout: 🚀 Serving at  http://127.0.0.1:5555
```

On http://127.0.0.1:5555 is the app. Happy coding!

To stop run
```bash
$ docker-compose down
```

### 💾 Local

You need Go 1.18+ install in your machine. Execute
```bash
$ go run .
# output
# 🚀 Serving at  http://127.0.0.1:5555
```

## Grammar

Sentencias
- [x] Asignacion 
- [x] If
- [x] While
- [x] Log
- [x] Imprimir
- [x]   Bloque Condicional
- [x]   Bloque De Sentencia

Expresiones
- [x] MenosUnarioExpr
- [x] Not
- [x] multiplicacionExpr
- [x] aditivaExpr
- [x] relacionalExpr
- [x] igualdadExpr
- [x] andExpr
- [x] orExpr
- [x] atomExpr

Atomo
- [x] parExpr `()`
- [x] numberAtom
- [x] booleanAtom
- [x] idAtom
- [x] stringAtom
- [x] nilAtom
