package flags

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	Debug bool
	Dev   bool
}

var F Flags

func AddFlags(f *Flags) {
	pflag.BoolVarP(&f.Debug, "debug", "d", false, "是否开启 Gin 的 debug 模式")
}
