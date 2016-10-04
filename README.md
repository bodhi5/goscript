# goscript

Simple wrapper around `go run` to use go as a scripting language

## CLI Usage

```
$ goscript "fmt.Println("packages will automatically be imported from your $GOPATH")
```

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Package Overview</a>


## <a name="pkg-index">Index</a>
* [type GoScript](#GoScript)
  * [func NewFromFile(sourcePath string, args ...string) (*GoScript, error)](#NewFromFile)
  * [func NewFromString(source string, args ...string) (*GoScript, error)](#NewFromString)
  * [func (gs GoScript) Clean() error](#GoScript.Clean)
  * [func (gs GoScript) FilePath() string](#GoScript.FilePath)


#### <a name="pkg-files">Package files</a>
[goscript.go](/src/github.com/bodhi5/goscript/goscript/goscript.go) 


## <a name="GoScript">type</a> [GoScript](/src/target/goscript.go?s=133:202#L4)

GoScipt Type

``` go
type GoScript struct {
    *exec.Cmd
}
```


### <a name="NewFromFile">func</a> [NewFromFile](/src/target/goscript.go?s=2206:2276#L82)

NewFromFile takes in a filePath of golang source and converts it into a GoScript type with an embedded exec.Cmd

``` go
func NewFromFile(sourcePath string, args ...string) (*GoScript, error)
```


### <a name="NewFromString">func</a> [NewFromString](/src/target/goscript.go?s=1740:1808#L67)

NewFromString takes in a string of golang source and converts into a GoScript type with an embedded exec.Cmd

``` go
func NewFromString(source string, args ...string) (*GoScript, error)
```


### <a name="GoScript.Clean">func</a> (GoScript) [Clean](/src/target/goscript.go?s=367:399#L16)

Clean deletes the generated goscript tempfile

``` go
func (gs GoScript) Clean() error
```


### <a name="GoScript.FilePath">func</a> (GoScript) [FilePath](/src/target/goscript.go?s=256:292#L11)

FilePath returns the current GoScript's file path

``` go
func (gs GoScript) FilePath() string
```
