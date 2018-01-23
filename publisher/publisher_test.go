package publisher

import (
	"os"
	"testing"

	"github.com/NeowayLabs/wabbit/amqptest/server"
	"github.com/stretchr/testify/assert"
)

func TestValidateEnv(t *testing.T) {
	// error rabbit_path
	pathresp := ValidateEnv()
	assert.Equal(t, ErrEnv, pathresp)

	// error rabbit_path
	os.Setenv("RABBIT_PATH", "path")
	chresp := ValidateEnv()
	assert.Equal(t, ErrEnvChannel, chresp)

	// success
	os.Setenv("RABBIT_CHANNEL", "rabbit ch")
	resp := ValidateEnv()
	assert.Equal(t, nil, resp)

}

func TestSendMessage(t *testing.T) {
	os.Setenv("RABBIT_PATH", "amqp://localhost:5672/%2f")
	os.Setenv("RABBIT_CHANNEL", "channel")

	fakeServer := server.NewServer("amqp://localhost:5672/%2f")
	fakeServer.Start()

	resp := SendMessage("send a message")
	assert.Equal(t, nil, resp)

}
