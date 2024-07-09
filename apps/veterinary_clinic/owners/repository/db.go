package repository

import (
	"fmt"
	"sync"
	"vet/owners/pkg"
)

var db map[string]pkg.Owner
var mutex sync.Mutex

func init() {
	db = make(map[string]pkg.Owner, 0)
	populate()
}

func Create(e pkg.Owner) {
	mutex.Lock()
	defer mutex.Unlock()
	db[e.Id] = e
}

func Update(e pkg.Owner) {
	mutex.Lock()
	defer mutex.Unlock()
	db[e.Id] = e
}

func GetOne(id string) (pkg.Owner, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	d, o := db[id]
	return d, o
}

func GetAll() []pkg.Owner {
	mutex.Lock()
	defer mutex.Unlock()
	r := make([]pkg.Owner, 0, len(db))
	for _, v := range db {
		r = append(r, v)
	}
	return r
}

func Delete(id string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := db[id]; !ok {
		return fmt.Errorf("could not find entity %s during delete operation", id)
	}
	delete(db, id)
	return nil
}