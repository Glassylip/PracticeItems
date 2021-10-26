package model

import (
	"encoding/json"
	"fmt"
	"moudle/common/message"

	"github.com/garyburd/redigo/redis"
)

// After the server is started, we initialize a userDao instance,
// Make it a global variable, and use it directly when you need to operate with redis
var (
	MyUserDao *UserDao
)

// Complete various operations on the User structure.

type UserDao struct {
	pool *redis.Pool
}

// Using the factory pattern, create an instance of UserDao
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

//1. Return a User instance + err according to the user id
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	// Go to redis to view this user with a given id
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //In the users hash, no corresponding id was found
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	// Deserialize res into a User instance
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// Login verification
//1. Authenticate user
//2. If the user's id and pwd are correct, a user instance is returned
//3. If the user's id or pwd is wrong, the corresponding error message will be returned
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	// Take out a connection from UserDao's connection pool
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {

	// Take out a connection from UserDao's connection pool
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	// ID is not yet in redis, you can complete the registration
	data, err := json.Marshal(user) // Serialization
	if err != nil {
		return
	}
	// Write in redis
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("Save registered user error", err)
		return
	}
	return
}
