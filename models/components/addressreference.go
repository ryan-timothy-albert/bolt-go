// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
)

type AddressReferenceType string

const (
	AddressReferenceTypeID       AddressReferenceType = "id"
	AddressReferenceTypeExplicit AddressReferenceType = "explicit"
)

type AddressReference struct {
	AddressReferenceID       *AddressReferenceID
	AddressReferenceExplicit *AddressReferenceExplicit

	Type AddressReferenceType
}

func CreateAddressReferenceID(id AddressReferenceID) AddressReference {
	typ := AddressReferenceTypeID

	typStr := AddressReferenceIDTag(typ)
	id.DotTag = typStr

	return AddressReference{
		AddressReferenceID: &id,
		Type:               typ,
	}
}

func CreateAddressReferenceExplicit(explicit AddressReferenceExplicit) AddressReference {
	typ := AddressReferenceTypeExplicit

	typStr := AddressReferenceExplicitTag(typ)
	explicit.DotTag = typStr

	return AddressReference{
		AddressReferenceExplicit: &explicit,
		Type:                     typ,
	}
}

func (u *AddressReference) UnmarshalJSON(data []byte) error {

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
		u.Type = AddressReferenceTypeID
		return nil
	case "explicit":
		addressReferenceExplicit := new(AddressReferenceExplicit)
		if err := utils.UnmarshalJSON(data, &addressReferenceExplicit, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AddressReferenceExplicit = addressReferenceExplicit
		u.Type = AddressReferenceTypeExplicit
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AddressReference) MarshalJSON() ([]byte, error) {
	if u.AddressReferenceID != nil {
		return utils.MarshalJSON(u.AddressReferenceID, "", true)
	}

	if u.AddressReferenceExplicit != nil {
		return utils.MarshalJSON(u.AddressReferenceExplicit, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
