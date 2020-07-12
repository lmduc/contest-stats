// +heroku goVersion go1.14.1
// +heroku install github.com/golang-migrate/migrate/v4 github.com/golang-migrate/migrate/v4/cmd/migrate

module github.com/lmduc/contest-stats

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang-migrate/migrate/v4 v4.11.0 // indirect
	github.com/jinzhu/gorm v1.9.14
	github.com/joho/godotenv v1.3.0
	github.com/sirupsen/logrus v1.6.0
	google.golang.org/genproto v0.0.0-20200702021140-07506425bd67
)
