package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/go-magic-package/instance"
	probe "github.com/Myriad-Dreamin/go-magic-package/package-probe"
)

//go:generate go run github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path -generate_init
//go:generate go run github.com/Myriad-Dreamin/go-magic-package/package-attach-to-path -generate_register_map

func main() {
	fmt.Println(probe.RootPackage, probe.RootPath)
	fmt.Println(instance.Get(probe.RootPackage))
}

