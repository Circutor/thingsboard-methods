// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// SaveCustomerBody body for request saveCustomer.
type SaveCustomerBody struct {
	ID             EntityID              `json:"id"`
	Country        string                `json:"country"`
	State          string                `json:"state"`
	Phone          string                `json:"phone"`
	Name           string                `json:"name"`
	Email          string                `json:"email"`
	Title          string                `json:"title"`
	CreatedTime    int64                 `json:"createdTime"`
	AdditionalInfo additionalInfoRequest `json:"additionalInfo"`
}

// additionalInfoRequest we will save the data that does not have THingsBoard.
type additionalInfoRequest struct {
	CompanyName     string `json:"companyName"`
	CompanyActivity string `json:"companyActivity"`
	CompanyPosition string `json:"companyPosition"`
	Language        string `json:"language"`
}

// NewSaveCustomerBody create a new saveCustomerRequest for a Login.
func NewSaveCustomerBody(entityType, id, country, state, phone, name, email, title,
	companyName, companyActivity, companyPosition, language string, createTime int64) SaveCustomerBody {
	saveCustomerBody := SaveCustomerBody{
		ID: EntityID{
			EntityType: entityType,
			ID:         id,
		},
		Country:     country,
		State:       state,
		Phone:       phone,
		Name:        name,
		Email:       email,
		Title:       title,
		CreatedTime: createTime,
		AdditionalInfo: additionalInfoRequest{
			CompanyName:     companyName,
			CompanyActivity: companyActivity,
			CompanyPosition: companyPosition,
			Language:        language,
		},
	}

	return saveCustomerBody
}
