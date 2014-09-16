package services

import (
	"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

// Service is the result of a list or information query.
type Service struct {
	Description *string `json:"description,omitempty"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
}

// ServicePage is a single page of Service results.
type ServicePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (p ServicePage) IsEmpty() (bool, error) {
	services, err := ExtractServices(p)
	if err != nil {
		return true, err
	}
	return len(services) == 0, nil
}

// ExtractServices extracts a slice of Services from a Collection acquired from List.
func ExtractServices(page pagination.Page) ([]Service, error) {
	var response struct {
		Services []Service `mapstructure:"services"`
	}

	err := mapstructure.Decode(page.(ServicePage).Body, &response)
	return response.Services, err
}
