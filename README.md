# goapi
--
    import "github.com/andrewstuart/goapi"

Goapi is an api build on top of mgo and martini intended to be multifunctional
and work out of the box.

## Usage

#### func  SendShit

```go
func SendShit(Jsn) int
```
A function to send stuff

#### type Jsn

```go
type Jsn map[string]interface{}
```

Type Jsn is a jsn type

#### type Stupid

```go
type Stupid struct {
	Silly int //The amount of sillyness
}
```
