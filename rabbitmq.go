package commons

import (
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"math"
	"os"
	"time"
)

/*
amqp_URI       = "amqp://" amqp_authority [ "/" vhost ] [ "?" query ]
amqp_authority = [ amqp_userinfo "@" ] host [ ":" port ]
amqp_userinfo  = username [ ":" password ]
username       = *( unreserved / pct-encoded / sub-delims )
password       = *( unreserved / pct-encoded / sub-delims )
vhost          = segment
*/

// "amqp://guest:guest@rabbitmq"
const (
	RabbitMqProtocolKey  string = "RABBITMQ_PROTOCOL"
	RabbitMqUserKey             = "RABBITMQ_USER"
	RabbitMqPasswordKey         = "RABBITMQ_PASSWORD"
	RabbitMqHostKey             = "RABBITMQ_HOST"
	RabbitMqUrlTemplate         = "%s://%s:%s@%s"
	maxConnectionRetries int    = 5
)

var DefaultValues = map[string]string{
	RabbitMqProtocolKey: "amqp",
	RabbitMqUserKey:     "guest",
	RabbitMqPasswordKey: "guest",
	RabbitMqHostKey:     "rabbitmq",
}

func getDefinedEnvOrDefault(key string) string {
	envVal := os.Getenv(key)

	if len(envVal) == 0 {
		return DefaultValues[key]
	}
	return envVal
}

func getConnectionStringFromEnv() string {
	rabbitMqProtocol := getDefinedEnvOrDefault(RabbitMqProtocolKey)
	rabbitMqUser := getDefinedEnvOrDefault(RabbitMqUserKey)
	rabbitMqPassword := getDefinedEnvOrDefault(RabbitMqPasswordKey)
	rabbitMqHost := getDefinedEnvOrDefault(RabbitMqHostKey)

	isAnyEmpty := IsEmpty(rabbitMqProtocol, rabbitMqUser, rabbitMqPassword, rabbitMqHost)

	if isAnyEmpty {
		ExitOnError(errors.New("at least one env variable for rabbitmq was empty"))
	}

	return fmt.Sprintf(RabbitMqUrlTemplate, rabbitMqProtocol, rabbitMqUser, rabbitMqPassword, rabbitMqHost)
}

func ConnectToRabbitMq() *amqp.Connection {
	connString := getConnectionStringFromEnv()
	LogInfo(fmt.Sprintf("connecting to rabbitmq (%s)", connString))

	retries := 0
	backOff := 1 * time.Second
	for {

		// "amqp://guest:guest@rabbitmq"
		c, err := amqp.Dial(connString)

		if err != nil {
			LogInfo("rabbitmq not yet ready")
			retries++
		} else {
			LogInfo("connected to rabbitmq")
			return c
		}

		if retries > maxConnectionRetries {
			ExitOnError(errors.New(fmt.Sprintf("exceed max reconnect attempts(%d) to rabbitmq", maxConnectionRetries)))
		}

		backOff = time.Duration(math.Pow(float64(retries), 2)) * time.Second
		LogInfo(fmt.Sprintf("retry in %s", backOff))
		time.Sleep(backOff)
		continue
	}
}
