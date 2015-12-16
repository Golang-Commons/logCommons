# logCommons

A wrapper around the official log package.

* Supports log level
* Supports log color

A simple usage is:

```
package main

import (
    "github.com/Golang-Commons/logCommons"
    "os"
)

type T struct {
    A string
    B string
}

func main() {
    logFile, err := os.OpenFile("./log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        logCommons.Println("error opening log file:", err)//we only copied Println mechod.
        os.Exit(1)
    }
    defer logFile.Close()

    a := T{"tt", "nn"}
    logCommons.SetLevel(logCommons.INFO)
    logCommons.SetFlag(logCommons.Lcolor | logCommons.Lfuncname | logCommons.Llevelname)
    logCommons.Debug("6666")//as level is INFO,debug would not be recorded
    logCommons.Info("6666")
    logCommons.SetFlags(logCommons.Ldate | logCommons.Ltime | logCommons.Lshortfile)
    logCommons.Warn(a)
    logCommons.SetOutput(logFile)
    logCommons.Error("6666")
    logCommons.Panic("6666")
    logCommons.Fatal("6666")
}

```


output is:

```
$ go run f.go 
2015/12/17 03:59:04 [INFO] main.main 6666
2015/12/17 03:59:04 log.go:116: [WARNING] main.main {tt nn}
$ cat log.log 
2015/12/17 03:59:04 log.go:116: [ERROR] main.main 6666
2015/12/17 03:59:04 log.go:116: [PANIC] main.main 6666
2015/12/17 03:59:04 log.go:116: [FATAL] main.main 6666
```
