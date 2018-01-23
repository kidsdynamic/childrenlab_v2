package database

import (
	"fmt"
	"time"

	"crypto/sha256"
	"io"

	"log"

	"strings"

	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kidsdynamic/childrenlab_v2/config"
	"github.com/kidsdynamic/childrenlab_v2/model"
)

type FinalTest struct {
	ID              int       `db:"id"`
	MacID           string    `db:"mac_id"`
	FirmwareVersion string    `db:"firmware_version"`
	ProductVersion  int64     `db:"product_version"`
	BatteryLevel    string    `db:"battery_level"`
	DateCreated     time.Time `db:"date_created"`
	Result          bool      `db:"result"`
	UVMax           string    `gorm:"column:uv_max"`
	UVMin           string    `gorm:"column:uv_min" db:"uv_min"`
	XMax            string    `gorm:"column:x_max" db:"x_max"`
	XMin            string    `gorm:"column:x_min" db:"x_min"`
	YMax            string    `gorm:"column:y_max" db:"y_max"`
	YMin            string    `gorm:"column:y_min" db:"y_min"`
	Company         *string   `db:"company"`
	Language        string    `db:"language"`
	Converted       bool      `db:"converted"`
}

func (FinalTest) TableName() string {
	return "Final_Test"
}

var DatabaseInfo model.Database

func NewGORM() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		DatabaseInfo.User, DatabaseInfo.Password, DatabaseInfo.IP, DatabaseInfo.Name))

	if err != nil {
		panic(err)
	}

	if config.ServerConfig.Debug {
		db.LogMode(true)
	}

	db.SingularTable(true)
	return db
}

func NewTestRecordGORM() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		DatabaseInfo.User, DatabaseInfo.Password, DatabaseInfo.IP, "swing_test_record"))

	if err != nil {
		panic(err)
	}

	if config.ServerConfig.Debug {
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
		&model.HourlyActivity{},
		&model.InitialDeviceFirmware{},
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

	// Copy swing test record final result to the database
	sdb := NewTestRecordGORM()
	defer sdb.Close()

	var finalTest []FinalTest
	if err := sdb.Where("converted = false").Find(&finalTest).Error; err != nil {
		fmt.Printf("Error on retrieve final test. %#v", err)
	} else {
		for _, f := range finalTest {
			macId := f.MacID
			macId = strings.Replace(macId, ":", "", -1)
			device := &model.InitialDeviceFirmware{
				MacId:           macId,
				FirmwareVersion: f.FirmwareVersion,
				Language:        f.Language,
				ProductVersion:  f.ProductVersion,
			}
			db.Save(device)

			f.Converted = true
			sdb.Save(f)
		}

	}

}

func EncryptPassword(password string) string {
	h := sha256.New()
	io.WriteString(h, password)
	fmt.Printf("\n%x\n", h.Sum(nil))

	return fmt.Sprintf("%x", h.Sum(nil))

}
