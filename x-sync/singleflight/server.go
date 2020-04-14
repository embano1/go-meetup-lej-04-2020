package server

import (
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

type server struct {
	db *database
	sf *singleflight.Group
}

func (s *server) Get(key string) (string, error) {
	fn := func() (interface{}, error) {
		res := s.db.Query(key)
		return res, nil
	}
	v, err, _ := s.sf.Do("GetKey", fn)
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

type database struct {
	sync.RWMutex
	hits int
}

func (d *database) Query(q string) string {
	d.Lock()
	d.hits++
	d.Unlock()
	time.Sleep(2 * time.Second)
	return `"{"meetup":"golang","location":"leipzig"}"`
}

func (d *database) Hits() int {
	d.RLock()
	defer d.RUnlock()
	return d.hits
}
