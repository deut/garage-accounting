package config

import (
	"path"
	"path/filepath"
	"runtime"
)

type C struct {
	DBName         string
	DBFileFolder   string
	DBFileLocation string
}

var Conf *C

func (c *C) Defaults() {
	c.DBName = "garage.db"
	_, b, _, _ := runtime.Caller(1)
	c.DBFileFolder = filepath.Dir(b)
	c.DBFileLocation = path.Join(c.DBFileFolder, c.DBName)
}
