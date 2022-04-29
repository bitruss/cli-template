package plugin

import "github.com/coreservice-io/CliAppTemplate/plugin/reference"

//example 3 cache instance
func initReference() error {
	//default instance
	err := reference.Init()
	if err != nil {
		return err
	}

	// cache1 instance
	err = reference.Init_("ref1")
	if err != nil {
		return err
	}

	// cache2 instance
	err = reference.Init_("ref2")
	if err != nil {
		return err
	}

	return nil
}
