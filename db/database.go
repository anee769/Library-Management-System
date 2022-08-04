package db

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

type Database struct {
	Client *badger.DB
}

func Open() (*Database, error) {
	op := badger.DefaultOptions("data")

	client, err := badger.Open(op)
	if err != nil {
		return nil, fmt.Errorf("db open fail: %w", err)
	}

	return &Database{client}, nil
}

func (Db *Database) Close() {
	if err := Db.Client.Close(); err != nil {
		panic(fmt.Errorf("db close fail:%w", err))
	}
}

func (Db *Database) GetEntry(key []byte) (value []byte, err error) {
	err = Db.Client.View(func(view *badger.Txn) error {

		item, err := view.Get(key)
		if err != nil {
			return fmt.Errorf("db get on Key '%x' fail:'%w'", key, err)
		}

		if err = item.Value(func(val []byte) error {
			value = val
			return nil
		}); err != nil {
			return fmt.Errorf("db value get on the key '%x' fail:%w", key, err)
		}
		return nil

	})
	return
}

func (Db *Database) SetEntry(key, value []byte) error {
	return Db.Client.Update(func(view *badger.Txn) error {

		if err := view.Set(key, value); err != nil {
			return fmt.Errorf("db set for Key '%x' failed: %w", key, err)
		}
		return nil
	})
}
