package main

import "github.com/mholt/archiver"

func archiverRar(fileName, dir, password string) error {
	rar := archiver.NewRar()
	rar.Password = password
	rar.OverwriteExisting = true
	return rar.Unarchive(fileName, dir)
}
