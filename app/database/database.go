package database

import (
	"fmt"

	"crypto/sha256"
	"io"

	"log"

	"strings"

	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kidsdynamic/childrenlab_v2/app/config"
	"github.com/kidsdynamic/childrenlab_v2/app/model"
)

var DatabaseInfo model.Database

func NewGORM() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		DatabaseInfo.User, DatabaseInfo.Password, DatabaseInfo.IP, DatabaseInfo.Name))

	if err != nil {
		panic(err)
	}

	if !config.ServerConfig.Debug {
		db.LogMode(true)
	}

	db.SingularTable(true)
	return db
}

func InitDatabase() {
	db := NewGORM()

	db.AutoMigrate(
		&model.User{},
		&model.AccessToken{},
		&model.Kid{},
		&model.Device{},
		&model.Todo{},
		&model.ActivityRawData{},
		&model.Activity{},
		&model.LogUserAction{},
		&model.BatteryStatus{},
		&model.FwFile{},
	)

	if err := db.Exec("CREATE TABLE `sub_host_kid` (`sub_host_id` bigint,`kid_id` bigint, PRIMARY KEY (`sub_host_id`,`kid_id`))").Error; err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&model.SubHost{})

	if err := db.Exec("CREATE TABLE `event_kid` (`event_id` bigint,`kid_id` bigint, PRIMARY KEY (`event_id`,`kid_id`))").Error; err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&model.Event{})

	yes := db.HasTable("role")

	if !yes {
		db.AutoMigrate(&model.Role{})
		roles := []model.Role{
			{
				Authority: "ROLE_ADMIN",
			},
			{
				Authority: "ROLE_USER",
			},
			{
				Authority: "ROLE_SUPER_ADMIN",
			},
		}

		for _, role := range roles {
			db.Create(&role)
		}

	}

	var adminRole model.Role
	var userRole model.Role
	if err := db.Where("authority = ?", model.ROLE_ADMIN).First(&adminRole).Error; err != nil {
		panic(err)
	}
	if err := db.Where("authority = ?", model.ROLE_USER).First(&userRole).Error; err != nil {
		panic(err)
	}

	//initial activity raw steps
	var activityRaw []model.ActivityRawData
	if err := db.Where("indoor_steps is null OR outdoor_steps is null").Find(&activityRaw).Error; err != nil {
		panic(err)
	}
	for _, activity := range activityRaw {
		indoorData := strings.Split(activity.Indoor, ",")
		outdoorData := strings.Split(activity.Outdoor, ",")
		indoor, err := strconv.ParseInt(indoorData[2], 10, 64)
		if err != nil {
			panic(err)
		}
		outdoor, err := strconv.ParseInt(outdoorData[2], 10, 64)
		if err != nil {
			panic(err)
		}
		activity.IndoorSteps = indoor
		activity.OutdoorSteps = outdoor
		db.Save(activity)
	}
}

func EncryptPassword(password string) string {
	h := sha256.New()
	io.WriteString(h, password)
	fmt.Printf("\n%x\n", h.Sum(nil))

	return fmt.Sprintf("%x", h.Sum(nil))

}
