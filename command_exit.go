package main

import "os"

func callExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
