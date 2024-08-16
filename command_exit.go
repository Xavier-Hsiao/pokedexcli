package main

import "os"

func callExit(cfg *config) error {
	os.Exit(0)
	return nil
}
