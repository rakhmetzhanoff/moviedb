package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"moviedb/internal/models"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := "disable"
	databaseDriver := "postgres"

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	fmt.Println("Connecting to:", dbURL)

	sqlDB, err := sql.Open(databaseDriver, dbURL)
	if err != nil {
		log.Fatal("sql.Open error:", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("migratepg.WithInstance error:", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working dir: %v", err)
	}
	log.Println("Current working dir:", cwd)

	migrationsPath := filepath.ToSlash(filepath.Join(cwd, "internal", "db", "migrations"))
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)
	log.Println("Migrations path:", migrationsURL)

	m, err := migrate.NewWithDatabaseInstance(migrationsURL, databaseDriver, driver)
	if err != nil {
		log.Fatal("migrate.NewWithDatabaseInstance error:", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration failed:", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm.Open error:", err)
	}

	DB = gormDB

	var admin models.User
	if err := DB.Where("role = ?", "admin").First(&admin).Error; err != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Ошибка при хэшировании пароля админа:", err)
			return
		}

		admin = models.User{
			Username: "admin",
			Password: string(hashedPassword),
			Role:     "admin",
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Println("Не удалось создать администратора:", err)
		} else {
			log.Println("Администратор создан: username=admin password=admin123")
		}
	} else {
		log.Println("Администратор уже существует, пропускаем создание")
	}

}
