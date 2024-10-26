package settings

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)


// загрузка переменных окружения
var _ error = godotenv.Load("./.env")


// распаковка переменных окружения для Go app
var Port string = os.Getenv("GO_PORT")
var SecretForJWT string = os.Getenv("SECRET")

var CorsAllowedOrigins []string = strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
var CorsAllowedMethods []string = strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")
var CorsAllowCredentials bool = false

// распаковка переменных окружения для MySQL DB
var dbName string = os.Getenv("DB_NAME")
var dbUser string = os.Getenv("DB_USER")
var dbPassword string = os.Getenv("DB_PASSWORD")
var dbHost string = os.Getenv("DB_HOST")
var dbPort string = os.Getenv("DB_PORT")

// строка для подключения к БД из GORM
var DSN string = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	dbUser,
	dbPassword,
	dbHost,
	dbPort,
	dbName,
)

// формат логов
var LogFmt string = "[${time_rfc3339}] -- ${status} -- from ${remote_ip} to ${host} (${method} ${uri}) [time: ${latency_human}] | ${bytes_in} ${bytes_out} | error: ${error} | -> User-Agent: ${user_agent}\n"
// формат времени
var TimeFmt string = "06-01-02 15:04:05 -07"

// логеры
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
var ErrorLog *log.Logger = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

// функция для обработки критических ошибок
func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}
