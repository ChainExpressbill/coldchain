package middlewares

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig() (*os.File, logger.Config) {
	fileName := "coldchain.log"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	loggingFormat := "[${time}]${status} - ${method} ${path}\n" +
		"[${time}]RequestHeaders: ${reqHeaders}\n" +
		"[${time}]QueryParams: ${queryParams}\n" +
		"[${time}]RequestBody: ${body}\n" +
		"[${time}]ResponseBody: ${resBody}\n" +
		"[${time}]Error: ${error}\n\n"

	return file, logger.Config{
		Format:     loggingFormat,
		TimeFormat: "2006-01-02 15:04:05",
		Output:     file,
	}
}
