package model

import "sync"

//first version is a simple queue in memory
type MetaData struct {
	wrlock sync.RWMutex //写锁

	//默认数据以字符串的方式存储
	data map[string]string
}

func NewMetaData() *MetaData {
	return &MetaData{
		data:   map[string]string{},
	}
}

func (m *MetaData)SetData(key string, data string)  {
	m.wrlock.Lock()
	defer m.wrlock.Unlock()
	m.data[key]=data
}

func (m *MetaData) GetData(key string) string {
	m.wrlock.Lock()
	defer m.wrlock.Unlock()
	return m.data[key]
}