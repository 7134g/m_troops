package db

import "encoding/json"

type rdsType struct {
	val any
}

func (d *rdsType) MarshalText() (text []byte, err error) {
	return json.Marshal(d.val)
}

func (d *rdsType) UnmarshalText(v []byte) error {
	err := json.Unmarshal(v, &d.val)
	if err != nil {
		return err
	}
	return nil
}

func (d *rdsType) MarshalBinary() (data []byte, err error) {
	return json.Marshal(d.val)
}

func (d *rdsType) UnmarshalBinary(v []byte) error {
	return json.Unmarshal(v, &d.val)
}

func (d *rdsType) Set(v any) {
	d.val = v
}
func (d *rdsType) Get() any {
	return d.val
}

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
