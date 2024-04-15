package testx

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// AnyTime 表示任何时间的结构，只验证是否是时间类型，不匹配时间的值
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type MockDB struct {
	*sql.DB
	sqlmock.Sqlmock
}

type MockGorm struct {
	*MockDB
	DB *gorm.DB
}

func NewSqlmock() *MockDB {
	db, mock, err := sqlmock.New() // mock db
	if err != nil {
		panic(fmt.Errorf("failed to create sqlmock: %v", err))
	}
	return &MockDB{DB: db, Sqlmock: mock}
}

func (m *MockDB) Gorm() *MockGorm {
	// 创建 gorm.DB
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      m.DB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to open gorm connection: %v", err))
	}
	return &MockGorm{MockDB: m, DB: db}
}
