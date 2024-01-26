package db

import "encoding/json"

type SliceString []string

func NewSliceString(v []string) *SliceString {
	r := SliceString(v)
	return &r
}

func (d *SliceString) MarshalText() (text []byte, err error) {
	return json.Marshal(*d)
}

func (d *SliceString) UnmarshalText(v []byte) error {
	slice := make([]string, 0)
	err := json.Unmarshal(v, &slice)
	if err != nil {
		return err
	}

	*d = slice

	return nil

}
func (d *SliceString) UnmarshalBinary(v []byte) error {
	return json.Unmarshal(v, d)
}

func (d *SliceString) MarshalBinary() (data []byte, err error) {
	return json.Marshal(*d)
}

func (d *SliceString) Get() []string {
	if len(*d) == 0 {
		return []string{}
	}

	return *d
}
