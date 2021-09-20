package mysql_driver

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	adminRepo "auto_traveler/driver/database/admin"
	equipmentsRepo "auto_traveler/driver/database/equipments"
	eventHistoriesRepo "auto_traveler/driver/database/event_histories"
	eventsRepo "auto_traveler/driver/database/events"
	playersRepo "auto_traveler/driver/database/players"
	"auto_traveler/helper/encrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)
	Seeder(db)

	return db
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&adminRepo.Admin{},
		&playersRepo.Players{},
		&eventsRepo.Events{},
		&equipmentsRepo.Equipments{},
		&eventHistoriesRepo.EventHistories{},
	)
}

func Seeder(db *gorm.DB) {
	admin := []adminRepo.Admin{}
	events := []eventsRepo.Events{}
	equipments := []equipmentsRepo.Equipments{}
	db.Find(&admin)
	db.Find(&events)
	db.Find(&equipments)

	if len(admin) == 0 {
		password, _ := encrypt.Hash("admin")
		var admin = []adminRepo.Admin{
			{	Name: "Superadmin",
				Email: "superadmin@admin.com",
				Password: sql.NullString{String: password, Valid: true},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		db.Create(&admin)
	}

	if len(events) == 0 {
		events := []eventsRepo.Events{
			{	ID: 1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Type: "expedition",
				Name: "Windrise Exploration",
				Description: "Explore Windrise area, some said there's a strage aura around the sacred tree",
				GoldReward: 10,
				XPReward: 10,
			},
			{	ID: 2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Type: "bounty",
				Name: "Hililchurl Hunter",
				Description: "Hunt Hililchurls that unsettling the farmers",
				GoldReward: 15,
				XPReward: 15,
			},
		}
		db.Create(&events)
	}

	if len(equipments) == 0 {
		equipments := []equipmentsRepo.Equipments{
			{	ID: 1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Type: "weapon",
				Name: "Prototype Rancour",
				Description: "A rancour known as the most sharp blade in the world, thought it just its prototype",
				ATK: 10,
				DEF: 5,
				HP: 0,
			},
			{	ID: 2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Type: "bounty",
				Name: "Prototype Armor",
				Description: "A tough armor that can deny a sword hit, thought we don't know how durable is it to the second time and so on",
				ATK: 0,
				DEF: 5,
				HP: 10,
			},
		}
		db.Create(&equipments)
	}

}

