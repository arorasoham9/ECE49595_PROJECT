package queue

import (
	"time"

	"github.com/go-redis/redis"

)
const(

	HARD_EXIT = 555
	SOFT_EXIT = 556
	QUEUE_UNAVAILABLE = 557;
	QUEUE_CONTAINERISED_UNAVAILABLE = 558
	QUEUE_CONTAINER_UNAVAILABLE = 559
	QUEUE_CREATED = 560 
	REDIS_AVAILABLE = 561
	REDIS_UNAVAILABLE = 562 
	SET = "true"
	UNSET = "false"
	API_Q_CLI = "API_Q_CLI"
	SSH_Q_CLI = "SSH_Q_CLI"
	EXIT_SUCCESS = 0
	EXIT_FAILURE = 1
)



type Request  struct {
	NAME string
	EMAIL string
	CURRENT_IP string
	LOCATION Location
	TIME time.Time
	LASTSEEN time.Time

}


type Location struct{



}

type Queue struct{
	API_CLI	*redis.Client
	SSH_SERV_CLI *redis.Client
	CREATED time.Time
	ONLINE bool
	SSH_SERV_CONN_OPTIONS *redis.Options
	API_CONN_OPTIONS *redis.Options
}