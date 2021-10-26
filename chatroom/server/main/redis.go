package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// local varibal pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // maximum number of idle connect
		MaxActive:   maxActive,   // maximum number of connect to the database， 0: no limit
		IdleTimeout: idleTimeout, // maximum number of timeout
		Dial: func() (redis.Conn, error) { // init connect
			return redis.Dial("tcp", address)
		},
	}
}
