package global

import (
	"time"

	"github.com/spf13/viper"
)

var (
	ServerSetting   ServerSettingS
	AppSetting      AppSettingS
	DatabaseSetting DatabaseSettingS
	RedisSetting    RedisSettingS
	LogConfig       Log
)

type Setting struct {
	vp *viper.Viper
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

// 初始配置方法，读取指定配置文件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	// 设置读取配置文件的位置和文件
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

// 初始化配置，将配置文件按读取到global变量当中。
func SetupSetting() error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Log", &LogConfig)
	if err != nil {
		return err
	}

	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	return nil
}

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	TokenPeriod     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	SecretKey       string
}

type DatabaseSettingS struct {
	Source   string
	Uri      string
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
}

type RedisSettingS struct {
	Host     string
	Password string
	DB       int
}

type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	File   string `mapstructure:"file"`
}
