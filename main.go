////////////////////////////////////////////////////////////////////////////
// Program: Wireframe-CICD
// Purpose: Wireframe CI/CD Demo
// Authors: Tong Sun (c) 2019, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Host      string
	Port      int
	Demo      string
	Daemonize bool
	Verbose   int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "Wireframe-CICD"
	version  = "0.1.0"
	date     = "2019-04-20"

	// Opts store all the configurable options
	Opts OptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	fmt.Fprintf(os.Stderr, "%s - Wireframe CI/CD Demo\n", progname)
	fmt.Fprintf(os.Stderr, "  v%s built on %s\n", version, date)
	fmt.Fprintf(os.Stderr, "Copyright (C) 2019, Tong Sun\n\n")
	fmt.Println("Program ends successfully.")
}
