// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
)

type AddressReferenceInputType string

const (
	AddressReferenceInputTypeID       AddressReferenceInputType = "id"
	AddressReferenceInputTypeExplicit AddressReferenceInputType = "explicit"
)

type AddressReferenceInput struct {
	AddressReferenceID            *AddressReferenceID
	AddressReferenceExplicitInput *AddressReferenceExplicitInput

	Type AddressReferenceInputType
}

func CreateAddressReferenceInputID(id AddressReferenceID) AddressReferenceInput {
	typ := AddressReferenceInputTypeID

	typStr := AddressReferenceIDTag(typ)
	id.DotTag = typStr

	return AddressReferenceInput{
		AddressReferenceID: &id,
		Type:               typ,
	}
}

func CreateAddressReferenceInputExplicit(explicit AddressReferenceExplicitInput) AddressReferenceInput {
	typ := AddressReferenceInputTypeExplicit

	typStr := AddressReferenceExplicitTag(typ)
	explicit.DotTag = typStr

	return AddressReferenceInput{
		AddressReferenceExplicitInput: &explicit,
		Type:                          typ,
	}
}

func (u *AddressReferenceInput) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string `json:".tag"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "id":
		addressReferenceID := new(AddressReferenceID)
		if err := utils.UnmarshalJSON(data, &addressReferenceID, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AddressReferenceID = addressReferenceID
		u.Type = AddressReferenceInputTypeID
		return nil
	case "explicit":
		addressReferenceExplicitInput := new(AddressReferenceExplicitInput)
		if err := utils.UnmarshalJSON(data, &addressReferenceExplicitInput, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AddressReferenceExplicitInput = addressReferenceExplicitInput
		u.Type = AddressReferenceInputTypeExplicit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AddressReferenceInput) MarshalJSON() ([]byte, error) {
	if u.AddressReferenceID != nil {
		return utils.MarshalJSON(u.AddressReferenceID, "", true)
	}

	if u.AddressReferenceExplicitInput != nil {
		return utils.MarshalJSON(u.AddressReferenceExplicitInput, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
