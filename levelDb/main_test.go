package levelDb

import (
	"github.com/ShenalPelpola/rocksdb_vs_leveldb_benchmark/config"
	"testing"
)

func BenchmarkAddValue(b *testing.B) {
	db := config.LevelDbConn()

	dbModel := LevelDB{
		db: db,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = dbModel.AddValue()
	}
	b.StopTimer()
	defer db.Close()
}

func BenchmarkAddValues(b *testing.B) {
	db := config.LevelDbConn()

	dbModel := LevelDB{
		db: db,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = dbModel.AddValues()
	}
	b.StopTimer()
	defer db.Close()
}


func BenchmarkReadOne(b *testing.B) {
	db := config.LevelDbConn()

	dbModel := LevelDB{
		db: db,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		key := "key_1"
		_, _ = dbModel.ReadOne(key)
	}
	b.StopTimer()
	defer db.Close()
}

func BenchmarkReadAll(b *testing.B) {
	db := config.LevelDbConn()
	dbModel := LevelDB{
		db: db,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dbModel.ReadAll()
	}
	b.StopTimer()
	defer db.Close()
}

