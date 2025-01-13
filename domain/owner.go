package domain

import "fmt"

type Owner struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address string
}

func (o Owner) Display() string {
	return fmt.Sprintf("Owner ID: %d, Name: %s, Email: %s, Phone: %s",
		o.ID, o.Name, o.Email, o.Phone)
}
