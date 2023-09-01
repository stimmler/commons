package commons

import (
	"errors"
	"fmt"
	"os"
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
	RabbitMqProtocolKey string = "RABBITMQ_PROTOCOL"
	RabbitMqUserKey            = "RABBITMQ_USER"
	RabbitMqPasswordKey        = "RABBITMQ_PASSWORD"
	RabbitMqHostKey            = "RABBITMQ_HOST"
	RabbitMqUrlTemplate        = "%s://%s:%s@%s"
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

func GetConnectionStringFromEnv() string {
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
