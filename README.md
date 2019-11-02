# go-magic-package
 get package path with magic

## Get Started

```bash
>> git clone https://github.com/Myriad-Dreamin/go-magic-package
>> go generate
>> go run .
github.com/Myriad-Dreamin/go-magic-package path\to\gopath\github.com\Myriad-Dreamin\go-magic-package
path\to\gopath\github.com\Myriad-Dreamin\go-magic-package
```

## Installation

```bash
go install github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path
```

## Usage

generate runtime probe

```go

//with install
//go:generate package-attach-to-path -generate_init

//without install
//go:generate go run github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path -generate_init
```

register package->file path to map

```go
//go:generate go run github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path -generate_register_map
```

`cat main.go`

```go
package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/go-magic-package/instance"
	probe "github.com/Myriad-Dreamin/go-magic-package/package-probe"
)

//go:generate go run github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path -generate_init -path=github.com/Myriad-Dreamin/go-magic-package -package=main
//go:generate go run github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path -generate_register_map

func main() {
	fmt.Println(probe.RootPackage, probe.RootPath)
	fmt.Println(instance.Get(probe.RootPackage))
}

```

```bash
>> go run .
github.com/Myriad-Dreamin/go-magic-package path\to\gopath\github.com\Myriad-Dreamin\go-magic-package
path\to\gopath\github.com\Myriad-Dreamin\go-magic-package
```





