package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qiniu/log"
	"os"
)

const (
	//server
	DEFAULT_READ_TIMEOUT     = 60
	DEFAULT_WRITE_TIMEOUT    = 60
	DEFAULT_MAX_HEADER_BYTES = 1 << 12 //4KB

	DEFAULT_LOG_FILE = "run.log"
)

type RtcConfig struct {
	Server ServerConfig `json:"server"`
	App    AppConfig    `json:"app"`
	Orm    OrmConfig    `json:"orm"`
}

//server config
type ServerConfig struct {
	ListenHost string `json:"listen_host"`
	ListenPort int    `json:"listen_port"`

	ReadTimeout    int `json:"read_timeout,omitempty"`
	WriteTimeout   int `json:"write_timeout,omitempty"`
	MaxHeaderBytes int `json:"max_header_bytes,omitempty"`
}

//app config
type AppConfig struct {
	AlertCriteria float32 `json:"alert_criteria"`
	AccessKey     string  `json:"access_key"`
	SecretKey     string  `json:"secret_key"`
	Hub           string  `json:"hub"`
	Prescription  int64   `json:"prescription"`
	LogFile       string  `json:"log_file,omitempty"`
	LogLevel      string  `json:"log_level,omitempty"`

	QLogLevel int
}

//orm config
type OrmConfig struct {
	DriverName string `json:"driver_name"`
	DataSource string `json:"data_source"`

	MaxIdleConn int  `json:"max_idle_conn,omitempty"`
	MaxOpenConn int  `json:"max_open_conn,omitempty"`
	DebugMode   bool `json:"debug_mode,omitempty"`
}

func LoadConfig(confFile string) (cfg *RtcConfig, err error) {
	cfgFh, openErr := os.Open(confFile)
	if openErr != nil {
		err = openErr
		return
	}
	defer cfgFh.Close()
	cfg = &RtcConfig{}

	decoder := json.NewDecoder(cfgFh)
	decodeErr := decoder.Decode(&cfg)
	if decodeErr != nil {
		err = errors.New(fmt.Sprintf("parse config error, %s", decodeErr))
		return
	}

	//check server defaults
	if cfg.Server.ReadTimeout <= 0 {
		cfg.Server.ReadTimeout = DEFAULT_READ_TIMEOUT
	}
	if cfg.Server.WriteTimeout <= 0 {
		cfg.Server.WriteTimeout = DEFAULT_WRITE_TIMEOUT
	}
	if cfg.Server.MaxHeaderBytes <= 0 {
		cfg.Server.MaxHeaderBytes = DEFAULT_MAX_HEADER_BYTES
	}

	//check app defaults
	if cfg.App.LogFile == "" {
		cfg.App.LogFile = DEFAULT_LOG_FILE
	}

	//check log level
	switch cfg.App.LogLevel {
	case "debug":
		cfg.App.QLogLevel = log.Ldebug
	case "info":
		cfg.App.QLogLevel = log.Linfo
	case "warn":
		cfg.App.QLogLevel = log.Lwarn
	case "error":
		cfg.App.QLogLevel = log.Lerror
	case "panic":
		cfg.App.QLogLevel = log.Lpanic
	case "fatal":
		cfg.App.QLogLevel = log.Lfatal
	default:
		cfg.App.QLogLevel = log.Ldebug
	}
	return
}
