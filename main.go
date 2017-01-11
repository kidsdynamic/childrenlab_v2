package main

import (
	"fmt"
	"os"

	"github.com/kidsdynamic/childrenlab_v2/model"
	"github.com/kidsdynamic/childrenlab_v2/router"

	"github.com/kidsdynamic/childrenlab_v2/database"
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
			EnvVar: "AWS_BUCKET",
			Name:   "aws_bucket",
			Usage:  "bucket name",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "AWS_REGION",
			Name:   "aws_region",
			Usage:  "AWS region",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "AWS_ACCESS_KEY_ID",
			Name:   "aws_access_key",
			Usage:  "bucket name",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "AWS_SECRET_ACCESS_KEY",
			Name:   "aws_secret_acess_key",
			Usage:  "bucket name",
			Value:  "",
		},
	}

	app.Action = func(c *cli.Context) error {
		database.DatabaseInfo = model.Database{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		}
		//fmt.Println(c.String("aws_bucket"))
		c.Set("aws_bucket", c.String("aws_bucket"))

		model.AwsConfig = model.AwsConfiguration{
			Bucket:          c.String("aws_bucket"),
			Region:          c.String("aws_region"),
			AccessKey:       c.String("aws_access_key"),
			SecretAccessKey: c.String("aws_secret_acess_key"),
		}

		fmt.Printf("Database: %v", database.DatabaseInfo)

		r := router.New()

		if c.Bool("debug") {
			return r.Run(":8111")
		} else {
			return r.RunTLS(":8111", ".ssh/childrenlab.chained.crt", ".ssh/childrenlab.com.key")
		}

	}

	app.Run(os.Args)
}
