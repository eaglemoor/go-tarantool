package tarantool

import (
	"testing"
	"time"
)

var server = "127.0.0.1:3013"
var spaceNo = uint32(512)
var spaceName = "test"
var indexNo = uint32(0)
var indexName = "primary"
var opts = Opts{
	Timeout: 500 * time.Millisecond,
	User:    "test",
	Pass:    "test",
	//Concurrency: 32,
	//RateLimit: 4*1024,
}

func TestConnect(t *testing.T) {
	conn, err := Connect(server, opts)
	if err != nil {
		t.Error(err)
	}
	t.Log(conn.Schema.Version)

	conn.Schema.Version--
	resp, err := conn.Insert("test", []interface{}{3})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", resp.Data)

	resp, err = conn.Insert("test", []interface{}{4})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", resp.Data)

	resp, err = conn.Select(spaceNo, indexNo, 0, 1, IterEq, []interface{}{uint(1)})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Data)
}
