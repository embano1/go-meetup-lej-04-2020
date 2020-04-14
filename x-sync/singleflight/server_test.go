package server

import (
	"sync"
	"testing"

	"golang.org/x/sync/singleflight"
)

func parallelize(t *testing.T, workers int, testFn func(t *testing.T)) {
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			testFn(t)
		}()
	}
	wg.Wait()
}

func Test_ServerGet(t *testing.T) {
	type fields struct {
		db *database
		sf *singleflight.Group
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"query without concurrency", fields{db: &database{}, sf: &singleflight.Group{}}, args{"testKey"}, `"{"meetup":"golang","location":"leipzig"}"`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				db: tt.fields.db,
				sf: tt.fields.sf,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("server.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ServerGetParallel(t *testing.T) {
	type fields struct {
		concurrency int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"GetKey with 10 concurrent workers", fields{concurrency: 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &server{
				db: &database{},
				sf: &singleflight.Group{},
			}
			testFn := func(t *testing.T) {
				t.Log("calling srv.Get()")
				res, _ := srv.Get("testKey")
				wantRes := `"{"meetup":"golang","location":"leipzig"}"`
				if res != wantRes {
					t.Errorf("srv.Get() = %v, want %v", res, wantRes)
				}
			}
			parallelize(t, tt.fields.concurrency, testFn)

			// verify we only hit the db once
			if got := srv.db.Hits(); got != tt.want {
				t.Errorf("db.Hits() = %v, want %v", got, tt.want)
			}
		})
	}
}
