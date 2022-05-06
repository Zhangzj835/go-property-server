package service

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	// "github.com/beego/beego/v2/client/cache"
	// "github.com/beego/beego/v2/core/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

var cc cache.Cache

func InitCache() {
	host := beego.AppConfig.String("cache::redis_host")
	passWord := beego.AppConfig.String("cache::redis_password")
	var err error
	defer func() {
		if r := recover(); r != nil {
			cc = nil
		}
	}()
	cc, err = cache.NewCache("redis", `{"conn":"`+host+`","password":"`+passWord+`"}`)
	// cc, err = cache.NewCache("redis", `{"conn":"`+host+`"}`)
	if err != nil {
		logs.Error("Connect to the redis host " + host + " failed")
		logs.Error(err)
	}
}

// SetCache
func SetCache(key string, value interface{}, timeout int) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			cc = nil
		}
	}()
	timeouts := time.Duration(timeout) * time.Second
	// err = cc.Put(context.Background(), key, data, timeouts)
	err = cc.Put(key, data, timeouts)
	if err != nil {
		logs.Error(err)
		logs.Error("SetCache失败，key:" + key)
		return err
	} else {
		return nil
	}
}

func GetCache(key string, to interface{}) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			cc = nil
		}
	}()

	// data, _ := cc.Get(context.Background(), key)
	data := cc.Get(key)
	if data == nil {
		return errors.New("Cache不存在")
	}

	err := Decode(data.([]byte), to)
	if err != nil {
		logs.Error(err)
		logs.Error("GetCache失败，key:" + key)
	}

	return err
}

// DelCache
func DelCache(key string) error {
	if cc == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	err := cc.Delete(key)
	if err != nil {
		return errors.New("Cache删除失败")
	} else {
		return nil
	}
}

// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
