package radius

import (
	"errors"
)

var (
	AttributeNotFound = errors.New("Attribute not found.")
	AttributeNotAdded = errors.New("Couldn't add attribute.")
	AttributeExists   = errors.New("Attribute already exists.")
)

type Dictionary interface {
	GetAttribute(vendorId uint32, typeId uint8) (Attribute, error)
	AddAttribute(Attribute) error
}

type MemoryDictionary map[uint16]map[uint8]Attribute

func (m MemoryDictionary) GetAttribute(vId uint16, tId uint8) (Attribute, error) {
	attr, ok := m[vId][tId]

	if !ok {
		return Attribute{}, AttributeNotFound
	}

	return attr, nil
}

func (m MemoryDictionary) AddAttribute(a Attribute) error {
	_, ok := m[a.VendorId][a.TypeId]

	if ok {
		return AttributeExists
	}

	m[a.VendorId][a.TypeId] = a
	return nil
}

type DictionaryParser interface {
	ParseAttribute() Attribute
	Done() bool
}

func LoadDictionary(p *DictionaryParser, d *Dictionary) error {

}
