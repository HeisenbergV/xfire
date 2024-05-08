package global

import (
	"fmt"
	"log"
	"os"
	"xfire/model"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/singleflight"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt签名
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}

type Config struct {
	Addr     int    `json:"addr"`
	Env      string `json:"env"`
	Dblink   string `json:"dblink"`
	LogLevel string `json:"loglevel"`
	LogPath  string `json:"logpath"`
	Jwt      JWT
}

var (
	GVA_Concurrency_Control = &singleflight.Group{}
	DB                      *gorm.DB
	Viper                   *viper.Viper
	C                       Config
	LOG                     *zap.Logger
	material                map[string]*model.Goods
)

func InitConfig(path string) {
	if len(path) == 0 {
		panic("Not config file")
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&C); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&C); err != nil {
		panic(err)
	}
	Viper = v
}

// logpath 日志文件路径
// loglevel 日志级别
func InitLogger() {
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   C.LogPath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,        // 每个日志文件保存10M，默认 100M
		MaxBackups: 30,        // 保留30个备份，默认不限
		MaxAge:     7,         // 保留7天，默认不限
		Compress:   true,      // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)
	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	switch C.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                      // 小写编码器
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,                     //
		EncodeCaller:   zapcore.FullCallerEncoder,                          // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(write)), // 打印到控制台和文件
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	LOG = zap.New(core, caller, development)
}

func InitDb(url string) {
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	DB = db

	var m []model.Goods
	err = DB.Model(&model.Goods{}).Where("ptype=?", model.Material).Find(&m).Error
	if err != nil {
		log.Panicln("load material err:", err.Error())
	}
	material = make(map[string]*model.Goods)
	for _, v := range m {
		material[v.Name] = &v
	}
}

func GetMaterialInfo(name string) *model.Goods {
	return material[name]
}
