// Package flags provides helper functions for accessing command line flags.
//
// It also serves as a central point for creating new flags and providing
// helpers for interacting with the available information.
package flags

import (
	"github.com/go-openapi/swag"
)

var optionsGroups *[]swag.CommandLineOptionsGroup

// oshinkoOptions are the command line flags, these are formatted according
// to the documentation at https://github.com/jessevdk/go-flags
// New flags should be added to this structure, with helper functions added
// in this package.
type oshinkoOptions struct {
	LogFile string `short:"l" long:"log-file" description:"the file to write logs into, defaults to stdout"`
}

// GetLineOptionsGroups returns the CommandLineOptionsGroup structure that
// can be used to configure the command line flags for the rest server.
func GetLineOptionsGroups() []swag.CommandLineOptionsGroup {
	if optionsGroups == nil {
		newOptionsGroups := []swag.CommandLineOptionsGroup{
			{
				ShortDescription: "Oshinko REST server options",
				Options:          &oshinkoOptions{},
			},
		}
		optionsGroups = &newOptionsGroups
	}
	return *optionsGroups
}

// GetLogFile returns the log filename specified on the command line or an
// empty string in the case that no file is specified.
func GetLogFile() string {
	retval := ""
	if optionsGroups != nil {
		for _, optsGroup := range *optionsGroups {
			opts, ok := optsGroup.Options.(*oshinkoOptions)
			if ok == true {
				if opts.LogFile != "" {
					retval = opts.LogFile
				}
			}
		}
	}
	return retval
}
