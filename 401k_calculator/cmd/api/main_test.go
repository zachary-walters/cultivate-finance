package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"

	"github.com/nats-io/nats-server/v2/server"
	natsserver "github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
)

func TestRequestError_Error(t *testing.T) {
	err := errors.New("test error")
	requestError := &RequestError{Err: err}

	assert.Equal(t, err.Error(), requestError.Error())
}

func TestReqError(t *testing.T) {
	err := errors.New("test error")
	expected, _ := json.Marshal(RequestError{Err: err})

	assert.Equal(t, expected, reqError(err))
}

func RunServerOnPort(port int) *server.Server {
	opts := natsserver.DefaultTestOptions
	opts.Port = port
	return RunServerWithOptions(&opts)
}

func RunServerWithOptions(opts *server.Options) *server.Server {
	return natsserver.RunServer(opts)
}

func TestMain(t *testing.T) {
	os.Setenv("NATS_URI", "nats://localhost:4222")
	os.Setenv("PORT", "4222")

	p, _ := strconv.Atoi(os.Getenv("PORT"))

	s := RunServerOnPort(p)
	defer s.Shutdown()

	go main()

	// Test connection
	var nc *nats.Conn
	var err error
	for {
		nc, err = nats.Connect(os.Getenv("NATS_URI"))
		if err == nil {
			break
		}

		log.Println("Waiting before connecting to NATS at:", os.Getenv("NATS_URI"))
		time.Sleep(2 * time.Second)
	}

	time.Sleep(2 * time.Second)

	assert.NoError(t, err)

	resp, err := nc.Request("calculate_all_401k", []byte("test"), 2*time.Second)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = nc.Request("calculate_401k_by_datakey", []byte("test"), 2*time.Second)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCalculateModel(t *testing.T) {
	input := calculator.Input{}
	inputBytes, _ := json.Marshal(input)
	modelMap, err := calculateModel(inputBytes)

	assert.NoError(t, err)

	assert.NotNil(t, modelMap)
	assert.NotEmpty(t, modelMap)

	// test json unmarshal error
	badInput := make([]byte, 4)
	rand.Read(badInput)

	badModelMap, err := calculateModel(badInput)
	assert.Error(t, err)
	assert.Nil(t, badModelMap)
}

func TestCalculateDatakey(t *testing.T) {
	input := calculator.Input{}
	inputBytes, _ := json.Marshal(input)

	data, err := calculateDatakey(inputBytes)
	assert.NoError(t, err)

	expectedData := struct {
		Datakey                    string `json:"datakey"`
		TraditionalValue           any    `json:"traditional_value,omitempty"`
		TraditionalRetirementValue any    `json:"traditional_retirement_value,omitempty"`
		RothValue                  any    `json:"roth_value,omitempty"`
		RothRetirementValue        any    `json:"roth_retirement_value,omitempty"`
	}{}
	assert.Equal(t, expectedData, data)

	// test json unmarshal error
	badInput := make([]byte, 4)
	rand.Read(badInput)

	badData, err := calculateDatakey(badInput)
	assert.Error(t, err)
	assert.Nil(t, badData)
}
