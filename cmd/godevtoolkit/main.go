package main

import (
	"os"

	"github.com/nekruzvatanshoev/godevtoolkit/pkg/godevtoolkit/cmd"
)

func main() {
	app := cmd.CreateApp()
	app.Run(os.Args)
}
