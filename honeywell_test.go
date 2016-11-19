package honeywell_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-home-iot/honeywell"
	"github.com/stretchr/testify/require"
)

// Put your mytotalconnectcomfort values here - DON'T check them in :) NOTE: This will
// execute commands that will affect your thermostat! You have been warned, if something
// gets screwed up it's not my fault.
var LOGIN = "<YOUR_LOGIN>"
var PASSWORD = "<YOUR_PASSWORD>"

// The deviceID has to be determined manually, log in to the mytotalconnectcomfort website,
// navigate to your device, then the URL will look something like
// https://mytotalconnectcomfort.com/portal/Device/CheckDataSession/123456, you need to copy the number
// that is in place of the 123456 and use that as your device ID.
var DEVICEID = 123456

func TestLogin(t *testing.T) {
	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)
	err := ts.Connect(ctx, LOGIN, PASSWORD)

	require.Nil(t, err)
}

func TestFetchStatusSuccess(t *testing.T) {
	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)

	err := ts.Connect(ctx, LOGIN, PASSWORD)
	require.Nil(t, err)

	status, err := ts.FetchStatus(ctx)
	require.Nil(t, err)

	fmt.Printf("%+v\n", status)
}

func TestFanModeOn(t *testing.T) {
	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)

	err := ts.Connect(ctx, LOGIN, PASSWORD)
	require.Nil(t, err)

	err = ts.FanMode(ctx, "on")
	require.Nil(t, err)
}

func TestFanModeAuto(t *testing.T) {
	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)

	err := ts.Connect(ctx, LOGIN, PASSWORD)
	require.Nil(t, err)

	err = ts.FanMode(ctx, "auto")
	require.Nil(t, err)
}

func TestHeatMode(t *testing.T) {
	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)

	err := ts.Connect(ctx, LOGIN, PASSWORD)
	require.Nil(t, err)

	err = ts.HeatMode(ctx, 67.0, time.Minute*30)
	require.Nil(t, err)
}

func TestCoolMode(t *testing.T) {
	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)

	err := ts.Connect(ctx, LOGIN, PASSWORD)
	require.Nil(t, err)

	err = ts.CoolMode(ctx, 65.0, time.Minute*30)
	require.Nil(t, err)
}

func TestCancel(t *testing.T) {
	ctx := context.TODO()
	ts := honeywell.NewThermostat(DEVICEID)

	err := ts.Connect(ctx, LOGIN, PASSWORD)
	require.Nil(t, err)

	err = ts.Cancel(ctx)
	require.Nil(t, err)
}
