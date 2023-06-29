package main

import (
	"flag"
	"log"
	repository "soat1-challenge1/internal/repositories"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type category struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Category  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func main() {
	rollback := flag.Bool("rollback", false, "use when migration should rollback")
	flag.Parse()

	db := repository.NewRepository()

	migration := []*gormigrate.Migration{
		// create category table
		{
			ID: "202306291_createcategory",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&category{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&category{})
			},
		},
	}

	m := gormigrate.New(db.Conn, gormigrate.DefaultOptions, migration)

	if *rollback {
		if err := m.RollbackMigration(migration[0]); err != nil {
			log.Fatalf("Could not rollback: %v", err)
		}
		log.Printf("Migration rollback did run successfully")
		return
	}

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
