type Account struct {
	gorm.Model
	GarageNumber string `validate:"required"`
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	PhoneNumber  string
	Address      string
}
