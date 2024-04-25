package globalconfig

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/evcc-io/evcc/charger/eebus"
	"github.com/evcc-io/evcc/provider/mqtt"
	"github.com/evcc-io/evcc/push"
	"github.com/evcc-io/evcc/util/config"
	"github.com/evcc-io/evcc/util/modbus"
)

type All struct {
	Network      Network
	Log          string
	SponsorToken string
	Plant        string // telemetry plant id
	Telemetry    bool
	Metrics      bool
	Profile      bool
	Levels       map[string]string
	Interval     time.Duration
	Database     DB
	Mqtt         Mqtt
	ModbusProxy  []ModbusProxy
	Javascript   []Javascript
	Go           []Go
	Influx       Influx
	EEBus        *eebus.Config
	HEMS         config.Typed
	Messaging    Messaging
	Meters       []config.Named
	Chargers     []config.Named
	Vehicles     []config.Named
	Tariffs      Tariffs
	Site         map[string]interface{}
	Loadpoints   []map[string]interface{}
}

type Javascript struct {
	VM     string
	Script string
}

type Go struct {
	VM     string
	Script string
}

type ModbusProxy struct {
	Port            int
	ReadOnly        string
	modbus.Settings `mapstructure:",squash"`
}

type Mqtt struct {
	mqtt.Config `mapstructure:",squash"`
	Topic       string
}

// Influx is the influx db configuration
type Influx struct {
	URL      string
	Database string
	Token    string
	Org      string
	User     string
	Password string
}

type DB struct {
	Type string
	Dsn  string
}

type Messaging struct {
	Events   map[string]push.EventTemplateConfig
	Services []config.Typed
}

type Tariffs struct {
	Currency string
	Grid     config.Typed
	FeedIn   config.Typed
	Co2      config.Typed
	Planner  config.Typed
}

type Network struct {
	Schema string
	Host   string
	Port   int
}

func (c Network) HostPort() string {
	if c.Schema == "http" && c.Port == 80 || c.Schema == "https" && c.Port == 443 {
		return c.Host
	}
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func (c Network) URI() string {
	return fmt.Sprintf("%s://%s", c.Schema, c.HostPort())
}
