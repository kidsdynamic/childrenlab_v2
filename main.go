package main

import (
	"fmt"
	"os"

	"github.com/kidsdynamic/childrenlab_v2/model"
	"github.com/kidsdynamic/childrenlab_v2/router"

	"github.com/kidsdynamic/childrenlab_v2/config"
	"github.com/kidsdynamic/childrenlab_v2/database"
	"github.com/kidsdynamic/childrenlab_v2/global"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "childrenlab"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "Debug",
			Value:  "false",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_USER",
			Name:   "database_user",
			Usage:  "Database user name",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_PASSWORD",
			Name:   "database_password",
			Usage:  "Database password",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_IP",
			Name:   "database_IP",
			Usage:  "Database IP address with port number",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "DATABASE_NAME",
			Name:   "database_name",
			Usage:  "Database name",
			Value:  "swing_test_record",
		},
		cli.StringFlag{
			EnvVar: "SUPER_ADMIN_TOKEN",
			Name:   "super_admin_token",
			Value:  "1",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "BASE_URL",
			Name:   "base_url",
			Value:  "http://localhost:8110",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "EMAIL_AUTH_NAME",
			Name:   "email_auth_name",
			Value:  "no-reply@kidsdynamic.com",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "EMAIL_AUTH_PASSWORD",
			Name:   "email_auth_password",
			Value:  "",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "EMAIL_SERVER",
			Name:   "email_server",
			Value:  "smtp.gmail.com",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "EMAIL_PORT",
			Name:   "email_port",
			Value:  "587",
			Usage:  "",
		},
		cli.StringFlag{
			EnvVar: "ERROR_LOG_EMAIL",
			Name:   "error_log_email",
			Value:  "jay@kidsdynamic.com",
			Usage:  "",
		},
	}

	app.Action = func(c *cli.Context) error {
		database.DatabaseInfo = model.Database{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		}

		config.ServerConfig = config.ServerConfiguration{
			BaseURL:           c.String("base_url"),
			EmailAuthName:     c.String("email_auth_name"),
			EmailAuthPassword: c.String("email_auth_password"),
			EmailServer:       c.String("email_server"),
			EmailPort:         c.Int("email_port"),
			ErrorLogEmail:     c.String("error_log_email"),
			Debug:             c.Bool("debug"),
		}

		global.SuperAdminToken = c.String("super_admin_token")

		fmt.Printf("Database: %v", database.DatabaseInfo)

		database.InitDatabase()

		r := router.New()

		return r.Run(":8111")

	}

	app.Run(os.Args)
}
