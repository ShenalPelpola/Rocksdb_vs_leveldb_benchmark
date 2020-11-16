package levelDb

import (
	"fmt"
	"github.com/jmhodges/levigo"
	_ "github.com/jmhodges/levigo"
	"log"
	"strconv"
)

type LevelDB struct {
	db *levigo.DB
}

// This function add a single value to the database
func (levelDbModel LevelDB) AddValue() (error, bool) {
	wo := levigo.NewWriteOptions()
	err := levelDbModel.db.Put(wo, []byte("key_random"), []byte("value_random"))

	if err != nil {
		return err, false
	}
	return nil, true
}

// This function adds 100 different random values to the db
func (levelDbModel LevelDB) AddValues() (error, bool){
	wo := levigo.NewWriteOptions()
	var err error

	for i := 0; i < 10; i++ {
		err = levelDbModel.db.Put(wo, []byte("key_"+strconv.Itoa(i)), []byte("value_"+strconv.Itoa(i)))
	}
	if err != nil {
		log.Fatal("Error occurred when writing to database", err)
		return err, false
	}
	return nil, true
}

// This function reads the key passed as an argument
func (levelDbModel LevelDB) ReadOne(key string) (error, string) {
	ro := levigo.NewReadOptions()
	data, err := levelDbModel.db.Get(ro, []byte(key))
	if err != nil {
		log.Fatal("key not found ",  err)
		return err, ""
	}
	return nil, string(data)
}

// This function reads all the keys,values from the database
func (levelDbModel LevelDB) ReadAll(){
	ro := levigo.NewReadOptions()
	ro.SetFillCache(false)
	iter := levelDbModel.db.NewIterator(ro)
	defer iter.Close()
	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		fmt.Println(string(iter.Key()), string(iter.Value()))
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
