package handler

import (
	"log"
	"mpg/entry/internal/injections"
	"sync"
)

type EntryHandler struct {
	logger *log.Logger
	instances *injections.Instances
}

func NewEntryHandler(logger *log.Logger) *EntryHandler {
	return &EntryHandler{logger: logger, instances: injections.NewInstances(logger)}
}

var instance *EntryHandler
var once sync.Once

func GetEntryHandler(logger *log.Logger) *EntryHandler {
	once.Do(func() {
		instance = NewEntryHandler(logger)
	})
	return instance
}

func (e *EntryHandler) GetInstance() *injections.Instances {
	return e.instances
}