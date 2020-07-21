package test

import (
	queue_server "algori/queue/queue-server"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

var(
	ctx=queue_server.InitServer()
)

type KVData struct {
	Key string
	Val string
}
func TestQueueServer(t *testing.T) {
	http.HandleFunc("/put", putData)
	http.HandleFunc("/get",getData)
	t.Fatal(http.ListenAndServe(":8080",nil))
}

func putData(w http.ResponseWriter,r *http.Request)  {
	data,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Print("putdata read data err:",err)
		w.WriteHeader(400)
	}
	log.Print("putData read data:",string(data))
	kv:=&KVData{}
	_=json.Unmarshal(data,kv)
	ctx.Data.SetData(kv.Key,kv.Val)
	w.WriteHeader(200)
}

func getData(w http.ResponseWriter, r *http.Request) {
	data,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		log.Print("getData read err:",err)
		return
	}
	log.Print("getdata read data:",string(data))
	val:=ctx.Data.GetData(string(data))
	w.WriteHeader(200)
	_,err=w.Write([]byte(val))
	if err!=nil{
		log.Print("getData write resp err:",err)
	}
}

func TestQueueClient(t *testing.T) {
	cli:=&http.Client{}
	kv:=&KVData{"name","cc"}
	data,err:=json.Marshal(kv)
	if err!=nil{
		log.Print("json marshal data err:",err)
	}
	req,err:=http.NewRequest("post","http://localhost:8080/put",strings.NewReader(string(data)))
	if err!=nil{
		t.Fatal("err create cli:",err)
		return
	}
	res,err:=cli.Do(req)
	if err!=nil{
		t.Fatal("err read res:",err)
	}
	data,err=ioutil.ReadAll(res.Body)
	if err!=nil{
		t.Fatal("read res err:",err)
		return
	}
	t.Log("res data:",string(data))
	req,err=http.NewRequest("post","http://localhost:8080/get",strings.NewReader("name"))
	if err!=nil{
		t.Fatal("err create cli:",err)
		return
	}

	res,err=cli.Do(req)
	if err!=nil{
		t.Fatal("err do http req:",err)
	}

	data,err=ioutil.ReadAll(res.Body)
	if err!=nil{
		t.Fatal("get read res data err:",err)
		return
	}
	t.Log("read data:",string(data))
}