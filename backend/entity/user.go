package entity
 
import (
  "time"
  "gorm.io/gorm"
  . "github.com/onsi/gomega"
)
 
type User struct {
  gorm.Model
  FirstName        string
  LastName         string
  Email            string
  Age              uint8
  BirthDay         time.Time
}