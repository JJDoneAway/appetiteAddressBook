package models

import (
	_ "embed"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"golang.org/x/exp/slog"

	"github.com/morkid/gocache"
	"github.com/morkid/paginate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

// see https://pkg.go.dev/github.com/go-playground/validator
// for validation
type Address struct {
	ID        uint      `gorm:"primarykey ->" json:"id"`
	CreatedAt time.Time `gorm:"<-:create" json:"-"`
	UpdatedAt time.Time `json:"-"`
	FirstName string    `gorm:"column:firstName" json:"firstName" binding:"required,gt=1"`
	LastName  string    `gorm:"column:lastName" json:"lastName" binding:"required,gt=1"`
	Email     string    `json:"email" binding:"required,email"`
	Phone     string    `json:"phone" binding:"required,e164"`
}

// for creation or update we must not provide an id
type AddressChange struct {
	FirstName string `json:"firstName" binding:"required,gt=1"`
	LastName  string `json:"lastName" binding:"required,gt=1"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,e164"`
}

// model for the pagination
type Page struct {
	Page       int64     `json:"page"`
	TotalPages int64     `json:"totalPages"`
	Total      int64     `json:"total"`
	Items      []Address `json:"items"`
}

type Entity interface {
	GetAllUsers() []*Address
	InsertUser() error
	GetUserByID() (*Address, error)
	UpdateUser() error
	DeleteUserByID() error
	DeleteAllUsers() error
	PaginateAddresses() []*Address
}

var (

	//error to be thrown in case of inconsistencies
	ErrIdMustBeZero  = errors.New("if you insert new users the ID must be zero")
	ErrUnknownID     = errors.New("id is unknown")
	ErrDuplicatedKey = errors.New("the key we generated already exists")

	db *gorm.DB
	pg *paginate.Pagination
)

func init() {
	if ok := godotenv.Load(); ok != nil {
		slog.Error("ENVIRONMENT", "Message", "Error loading .env file", "Error", ok)
	}

	dsn := os.Getenv("ODJ_DEP_POSTGRESQL_URL")

	level := logger.Error
	switch os.Getenv("ODJ_DEP_BACKEND_LOGGING") {
	case "DEBUG":
		level = logger.Silent
	case "INFO":
		level = logger.Silent
	case "ERROR":
		level = logger.Error
	default:
		slog.Error("ENV", "MESSAGE", "ODJ_DEP_BACKEND_LOGGING is not set as environment")
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		slog.Error("DB", "Message", "Not able to init DB connection: "+dsn, "Error", err)
	}

	if err := db.Migrator().DropTable(&Address{}); err != nil {
		slog.Error("DB", "Message", "Not able to drop Tables", "Error", err)
	}
	if err := db.AutoMigrate(&Address{}); err != nil {
		slog.Error("DB", "Message", "Not able create Tables", "Error", err)
	}

	tempDir, _ := os.MkdirTemp("", "myCache")
	adapterConfig := gocache.DiskCacheConfig{
		Directory: tempDir,
		ExpiresIn: 1 * time.Hour,
	}

	pg = paginate.New(&paginate.Config{
		DefaultSize:  15,
		CacheAdapter: gocache.NewDiskCache(adapterConfig),
	})

}

func CheckDB() error {
	return db.Exec("select 1").Error
}

// To avoid long running queries we will return not more than 20 items
func GetAllAddresses() ([]*Address, error) {
	var ret []*Address
	err := db.Limit(20).Find(&ret).Error
	return ret, err
}

func PaginateAddresses(req *http.Request) Page {
	model := db.Where("id > ?", 0).Model(&Address{})
	result := pg.With(model).Request(req).Cache("article").Response(&[]Address{})

	data, _ := json.Marshal(result)
	ret := Page{}
	json.Unmarshal(data, &ret)
	//I hate myself for this but I need camel case in the result json
	ret.TotalPages = result.TotalPages
	return ret
}

func (u *Address) InsertAddress() error {
	pg.ClearCache("article")
	res := db.Create(&u)
	return res.Error
}

func (u *Address) GetAddressByID() *Address {
	var ret *Address

	if ok := db.First(&ret, u.ID).Error; ok != nil {
		return nil
	}
	return ret

}

func (u *Address) UpdateAddress() error {

	if old := u.GetAddressByID(); old == nil {
		return ErrUnknownID
	}
	pg.ClearCache("article")
	if ok := db.Save(&u).Error; ok != nil {
		return ok
	}

	return nil
}

func (u *Address) DeleteAddressByID() error {

	if old := u.GetAddressByID(); old == nil {
		return ErrUnknownID
	}
	pg.ClearCache("article")
	if ok := db.Delete(&u).Error; ok != nil {
		slog.Error("DB", "Message", "Not able to clear the cache", "Error", ok)
		return ErrUnknownID
	}

	return nil
}

func DeleteAllAddresses() error {
	if ok := db.Where("1 = 1").Delete(&Address{}).Error; ok != nil {
		return ok
	}
	pg.ClearCache("article")
	return nil
}
