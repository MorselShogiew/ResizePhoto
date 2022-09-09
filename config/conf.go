package config

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
)

const (
	LocalEnv = "LOCAL"
)

var configs = map[string]string{ // nolint: gochecknoglobals
	LocalEnv: "config/conf_local.toml",
}

type Config struct {
	InstanceID      uuid.UUID   `json:"instance_id"`
	Environment     string      `json:"env"`
	ApplicationName string      `toml:"ApplicationName" json:"app_name"`
	PromPrefix      string      `toml:"PromPrefix" json:"prom_prefix"`
	ServerOpts      *ServerOpts `toml:"ServerOpt" json:"server_opts"`
	Logger          *Logger     `toml:"Logger" json:"logger"`
	BODB            *DB         `toml:"OfficeDB" json:"office_db"`
}
type Logger struct {
	Level     string     `toml:"Level" json:"level"`
	LoggerStd *LoggerStd `toml:"LoggerStd" json:"std"`
}

type LoggerStd struct {
	LogFile  string `toml:"LogFile" json:"file"`
	Stdout   bool   `toml:"Stdout" json:"stdout"`
	Disabled bool   `toml:"Disabled" json:"disabled"`
}

type DB struct {
	Server          string   `toml:"Server" json:"host"`
	Port            string   `toml:"Port" json:"port"`
	FailoverHost    string   `toml:"FailoverHost" json:"failover_host"`
	Database        string   `toml:"Database" json:"db_name"`
	Scheme          string   `toml:"Scheme" json:"scheme"`
	MaxIdleConns    int      `toml:"MaxIdleConns" json:"max_idle_conns"`
	MaxOpenConns    int      `toml:"MaxOpenConns" json:"max_open_conns"`
	ConnMaxLifetime Duration `toml:"ConnMaxLifetime" json:"conn_max_lifetime"`
	SSLMode         bool     `toml:"SSLMode" json:"ssl_mode"`
	Username        string   `json:"-"`
	Password        string   `json:"-"`
}

type ServerOpts struct {
	ReadTimeout  Duration `toml:"ReadTimeout" json:"read_timeout"`
	WriteTimeout Duration `toml:"WriteTimeout" json:"write_timeout"`
	IdleTimeout  Duration `toml:"IdleTimeout" json:"idle_timeout"`
	Port         string   `toml:"Port" json:"port"`
}

func LoadConfig() *Config {
	conf := &Config{}
	conf.Environment = os.Getenv("environment")
	configFile := configs[conf.Environment]
	if configFile == "" {
		conf.Environment = LocalEnv
		configFile = configs[LocalEnv]
	}

	if _, err := toml.DecodeFile(configFile, conf); err != nil {
		log.Fatal("couldn't decode config file:", err)
	}

	return conf
}

func (c *Config) String() string {
	out, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(out)
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (d *Duration) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}
