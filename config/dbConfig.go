package config

import (
	"github.com/jmhodges/levigo"
	"github.com/tecbot/gorocksdb"
	"log"
)


func LevelDbConn() *levigo.DB{
	opts := levigo.NewOptions()
	opts.SetCache(levigo.NewLRUCache(3<<30))
	opts.SetCreateIfMissing(true)
	db, err := levigo.Open("db/leveldb", opts)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func RocksDbConn() *gorocksdb.DB{
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "db/rocksdb")

	if err != nil {
		log.Fatal("Error when connecting to the database")
	}
	return db
}

