# What is this
A WIP wordle clone, playable either on console or in the browser.

## Start
This wordle game can be played either on console or in the webbrowser.

```shell
# for web:
./go-wordle -modus=web 
# for console:
./go-wordle -modus=console
```

## Build
```shell
go build
```

## Test
```shell
go test
```

## Create word database

Create a SQLite database in the project directory, like so
```shell
touch wordle.db
```

And then fill it with numerically ascending ids and five-letter words, e.g.:

```
1,Abart
2,Abbau
3,Abend
4,Abgas
5,Abort
6,Abruf
7,Absud
8,Abtei
9,Abweg
10,Abzug
```
