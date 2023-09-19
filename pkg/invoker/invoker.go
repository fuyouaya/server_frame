package invoker

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Log    *logrus.Logger
	cfg    *viper.Viper
	MainDB *gorm.DB
	Redis  *redis.Client
	Gin    *gin.Engine
)

func Init() (err error) {
	InitLogger()

	cfg = viper.New()
	cfg.AddConfigPath("././config/")
	cfg.SetConfigName("local")
	cfg.SetConfigType("toml")
	if err = cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			Log.Errorf("找不到配置文件")
		} else {
			Log.Errorf("配置文件解析出错: %s\n", err.Error())
		}
	}

	MainDB, err = gorm.Open(mysql.Open(cfg.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		Log.Errorf("连接 MySQL 失败: %s", err.Error())
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     cfg.GetString("redis.addr"),
		Password: cfg.GetString("redis.password"),
		DB:       cfg.GetInt("redis.db"),
	})

	if _, err = Redis.Ping(context.Background()).Result(); err != nil {
		Log.Errorf("连接 Redis 失败: %s", err.Error())
	}

	Gin = gin.Default()
	return nil
}

type LogFormatter struct{}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	Log = logrus.New()                //新建一个实例
	Log.SetOutput(os.Stdout)          //设置输出类型
	Log.SetReportCaller(true)         //开启返回函数名和行号
	Log.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	Log.SetLevel(logrus.DebugLevel)   //设置最低的Level
}
