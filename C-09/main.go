package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		errs := ValidateEmployee(emp)
		for newErr := range errs {
			if errors.Is(err, invalidID) {
				fmt.Printf("the employee has an: %v", newErr)
				continue
			}
			var empFErr emptyFieldErr
			if ok := errors.As(err, &empFErr); ok {
				fmt.Printf("the employee's %v", newErr)
				continue
			}
		}
		if errs != nil {
			fmt.Printf("record %d: %+v error: %v\n", count, emp, errors.Join(errs...))
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID   = regexp.MustCompile(`\w{4}-\d{3}`)
	invalidID = errors.New("invalid ID")
)

type emptyFieldErr struct {
	emptyFieldName string
}

func (e emptyFieldErr) Error() string {
	return fmt.Sprintf("%v:field is missing", e.emptyFieldName)
}

func ValidateEmployee(e Employee) []error {
	var errorsOut []error
	if len(e.ID) == 0 {
		errorsOut = append(errorsOut, emptyFieldErr{
			"ID",
		})
	}
	if !validID.MatchString(e.ID) {
		errorsOut = append(errorsOut, invalidID)
	}
	if len(e.FirstName) == 0 {
		errorsOut = append(errorsOut, emptyFieldErr{
			"FirstName",
		})
	}
	if len(e.LastName) == 0 {
		errorsOut = append(errorsOut, emptyFieldErr{
			"LastName",
		})
	}
	if len(e.Title) == 0 {
		errorsOut = append(errorsOut, emptyFieldErr{
			"Title",
		})
	}
	if len(errorsOut) > 0 {
		return errorsOut
	}
	return nil
}
