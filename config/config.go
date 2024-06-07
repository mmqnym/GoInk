package config

/*
	Discription: You can add your own config here.
*/

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type serverConfig struct {
	Mode  string `yaml:"mode"`
	Port  string `yaml:"port"`
	Http2 bool   `yaml:"http2"`
}

type logConfig struct {
	Level  string `yaml:"level"`
	Output struct {
		Stdout bool `yaml:"stdout"`
		File   bool `yaml:"file"`
	} `yaml:"output"`
	FileName  string `yaml:"fileName"`
	MaxSize   int    `yaml:"maxSize"`
	MaxAge    int    `yaml:"maxAge"`
	LocalTime bool   `yaml:"localTime"`
	Compress  bool   `yaml:"compress"`
}

type databaseConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type middlewareConfig struct {
	Cors struct {
		Enabled          bool     `yaml:"enabled"`
		AllowOrigins     []string `yaml:"allowOrigins"`
		AllowMethods     string   `yaml:"allowMethods"`
		AllowHeaders     string   `yaml:"allowHeaders"`
		ExposeHeaders    string   `yaml:"exposeHeaders"`
		MaxAge           string   `yaml:"maxAge"`
		AllowCredentials string   `yaml:"allowCredentials"`
	} `yaml:"cors"`

	Auth struct {
		JwtSecret        string `yaml:"jwtSecret"`
		JwtSigningMethod string `yaml:"jwtSigningMethod"`
		JwtExpires       int64  `yaml:"jwtExpires"`
		XAuthToken       string `yaml:"xAuthToken"`
	} `yaml:"auth"`
}

type gatewayConfig struct {
	Example struct {
		Url    string `yaml:"url"`
		ApiKey string `yaml:"apiKey"`
	} `yaml:"example"`
}

var rootPath string

var Server serverConfig
var Log logConfig
var Database databaseConfig
var Middleware middlewareConfig
var Gateway gatewayConfig

func init() {
	setProjectRoot()

	loadConfig("server.yml", &Server)
	loadConfig("log.yml", &Log)
	loadConfig("database.yml", &Database)
	loadConfig("middleware.yml", &Middleware)
	loadConfig("gateway.yml", &Gateway)
	// add more config here

	validateConfig()
}

func loadConfig(fileName string, target interface{}) {
	file, err := os.ReadFile(rootPath + "/config/" + fileName)
	checkError(err)

	err = yaml.Unmarshal(file, target)
	checkError(err)
}

// setProjectRoot() will finds the root of this project and sets it to rootPath
func setProjectRoot() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		gomod := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(gomod); err == nil {
			rootPath = currentDir
			break
		}

		nextDir := filepath.Dir(currentDir)
		if nextDir == currentDir {
			panic(errors.New("RootNotFoundError"))
		}

		currentDir = nextDir
	}
}

func validateConfig() {
	validateServerConfig()
	validateLogConfig()

	// add more validation here
}

func validateServerConfig() {
	if Server.Mode != "dev" && Server.Mode != "release" {
		panic(errors.New("InvalidModeError"))
	}

	if Server.Port == "" || Server.Port[0] != ':' || len(Server.Port) < 2 {
		panic(errors.New("InvalidPortError"))
	}
}

func validateLogConfig() {
	if Log.Level != "debug" &&
		Log.Level != "info" &&
		Log.Level != "warn" &&
		Log.Level != "error" &&
		Log.Level != "fatal" {
		panic(errors.New("InvalidLevelError"))
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
