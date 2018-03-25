package redis

import (
//redigo "github.com/garyburd/redigo/redis"
)

// var redis = make(map[string]*Redis)

// type Redis struct {
// 	redigo.Pool
// }

// func GetRedis(server string) (*Redis, error) {
// 	var err error

// 	if isMock {
// 		return getRedisMock(server)
// 	}

// 	if myConfig.Redis[server] == nil {
// 		return redis[server], fmt.Errorf("Server %s not found", server)
// 	}
// 	if redis[server] == nil {
// 		newRedis := redigo.Pool{
// 			MaxIdle:   myConfig.Redis[server].MaxIdle,
// 			MaxActive: 16,
// 			Dial: func() (redigo.Conn, error) {
// 				return redigo.Dial(
// 					"tcp", myConfig.Redis[server].EndPoint,
// 					redigo.DialConnectTimeout(myConfig.Redis[server].Timeout*time.Second),
// 				)
// 			},
// 		}

// 		redis[server] = &Redis{
// 			newRedis,
// 			nil,
// 		}
// 	}
// 	return redis[server], err
// }
