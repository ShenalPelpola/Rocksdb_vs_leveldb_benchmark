package rocksDb

import (
	"fmt"
	"github.com/tecbot/gorocksdb"
	"log"
	"strconv"
)

type RocksDb struct {
	db *gorocksdb.DB
}

// This function add a single value to the database
func (rocksDbModel RocksDb) AddValue() (error, bool) {
	wo := gorocksdb.NewDefaultWriteOptions()
	err := rocksDbModel.db.Put(wo, []byte("key_random"), []byte("value_random"))

	if err != nil {
		return err, false
	}
	return nil, true
}

// This function adds 100 different random values to the db
func (rocksDbModel RocksDb) AddValues() (error, bool){
	wo := gorocksdb.NewDefaultWriteOptions()
	var err error

	for i := 0; i < 10; i++ {
		err = rocksDbModel.db.Put(wo, []byte("key_"+strconv.Itoa(i)), []byte("value_"+strconv.Itoa(i)))
	}
	if err != nil {
		log.Fatal("Error occurred when writing to database", err)
		return err, false
	}
	return nil, true
}

// This function reads the key passed as an argument
func (rocksDbModel RocksDb) ReadOne(key string) (error, string) {
	ro := gorocksdb.NewDefaultReadOptions()
	data, err := rocksDbModel.db.Get(ro, []byte(key))
	if err != nil {
		log.Fatal("key not found ",  err)
		return err, ""
	}
	return nil, string(data.Data())
}

// This function reads all the keys,values from the database
func (rocksDbModel RocksDb) ReadAll(){
	ro := gorocksdb.NewDefaultReadOptions()
	ro.SetFillCache(false)
	iter := rocksDbModel.db.NewIterator(ro)
	defer iter.Close()
	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		fmt.Println(string(iter.Key().Data()), string(iter.Value().Data()))
	}
}
