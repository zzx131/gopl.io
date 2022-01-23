package main

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"io/ioutil"
	"log"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 使用 bucket
	key := []byte("key001")
	val := []byte("val001")
	bucket001 := "bucket001"
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(bucket001, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
	bucket002 := "bucket002"
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(bucket002, key, val, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
	// 读取
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			key := []byte("key001")
			if e, err := tx.Get("bucket001", key); err != nil {
				return err
			} else {
				fmt.Println(string(e.Value)) // "val1-modify"
			}
			return nil
		}); err != nil {
		log.Println(err)
	}
	// 获取文件数据
	bytes, e := ioutil.ReadFile("/home/zzx/go-proj/GOPATH/src/gopl.io/ch15/nutsdb/doc/go学习.txt")
	if e != nil {
		fmt.Println(e)
	}
	// 放到数据库中
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte("go-learn")
			if err := tx.Put("learn-go.txt", key, bytes, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
	// 从数据库中取出文件
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			key := []byte("go-learn")
			if e, err := tx.Get("learn-go.txt", key); err != nil {
				return err
			} else {
				fmt.Println(string(e.Value))
				err := ioutil.WriteFile("./doc/1.txt", e.Value, 0644)
				if err != nil {
					fmt.Println(err)
				}
			}
			return nil
		}); err != nil {
		log.Println(err)
	}

}
