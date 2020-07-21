package queue_server

import "algori/queue/model"

type Context struct {
	Data *model.MetaData
}

func InitServer() *Context {
	return &Context{Data:model.NewMetaData()}
}

