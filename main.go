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
	}

	app.Action = func(c *cli.Context) error {
		database.DatabaseInfo = model.Database{
			Name:     c.String("database_name"),
			User:     c.String("database_user"),
			Password: c.String("database_password"),
			IP:       c.String("database_IP"),
		}

		fmt.Printf("Database: %v", database.DatabaseInfo)

		r := router.New()

		//user := model.User{}
		/*

			rows, err := db.Queryx("SELECT email, first_name, last_name, zip_code, password, last_updated, date_created FROM user ORDER BY date_created ASC")
			if err != nil {
				log.Fatal(err)
			}
			for rows.Next() {
				err := rows.StructScan(&user)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Printf("%#v\n", user.DateCreated)
			}
		*/

		return r.Run(":8110")
	}

	app.Run(os.Args)
}
