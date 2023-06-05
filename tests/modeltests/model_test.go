package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/SyarifuanEfendi/go-test/api/controllers"
	"github.com/SyarifuanEfendi/go-test/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M)  {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting Env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database()  {
	var err error
	
	TestDbDriver := os.Getenv("TestDBDriver")

	if TestDbDriver == "postges" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {
	refreshUserTable()
	user := models.User{
		Nickname: "Efendi",
		Email: "Efendi@gmail.com",
		Password: "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() error {
	users := []models.User{
		models.User{
			Nickname: "Syarifuan",
			Email: "Syarifuan@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Efendi",
			Email: "Efendi@gmail.com",
			Password: "password",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndPostTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndPost() (models.Post, error) {
	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}

	user := models.User{
		Nickname: "Pendik",
		Email: "pendik@gmail.com",
		Password: "password",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{
		Title: "contoh title 1",
		Content: "controh content 1",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedOneUserAndPosts() ([]models.User, []models.Post, error)  {
	var err error
	if err != nil {
		return []models.User{}, []models.Post{}, err
	}

	var users = []models.User{
		models.User{
			Nickname: "Pei",
			Email: "pei@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "Pei1",
			Email: "pei1@gmail.com",
			Password: "password",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title: "Test post 1",
			Content: "Content test 1",
		},
		models.Post{
			Title: "Test post 2",
			Content: "Content test 2",
		},
	}
	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal("cannot seed users table %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatal("cannot seed posts table: %w", err)
		}
	}
	return users, posts, nil
}