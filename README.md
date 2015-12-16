# logCommons

A wrapper around the official log package.

* Supports log level
* Supports log color

A simple usage is:

```
package main

import (
    "github.com/Golang-Commons/logCommons"
)

type T struct {
    A string
    B string
}

func main() {
    a := T{"tt", "nn"}
    logCommons.SetLevel(logCommons.INFO)
    logCommons.SetFlag(logCommons.Lcolor | logCommons.Lfuncname | logCommons.Llevelname)
    logCommons.Debug("6666") //this log will not be recorded.
    logCommons.Info("6666")
    logCommons.Warn(a)
    logCommons.Error("6666")
    logCommons.Panic("6666")
    logCommons.Fatal("6666")
}
```
