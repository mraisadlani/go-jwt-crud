package entities

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"
)

type User struct {
	ID NullUserInt64 `json:"id"`
	Username NullUserString `json:"username"`
	Password NullUserString `json:"password"`
	Email NullUserString `json:"email"`
	Photo NullUserString `json:"photo"`
	IsActive NullUserBool `json:"is_active"`
	CreatedAt NullUserTime `json:"created_at"`
	UpdatedAt NullUserTime `json:"updated_at"`
}

type Account struct {
	Username NullUserString `json:"username"`
	Password NullUserString `json:"password"`
}

// NullUserInt64
type NullUserInt64 sql.NullInt64

func (ni *NullUserInt64) Scan(value interface{}) error {
	var i sql.NullInt64

	if err := i.Scan(value); err != nil {
		return err
	}

	*ni = NullUserInt64{i.Int64, i.Valid}

	return nil
}

// NullUserString
type NullUserString sql.NullString

func (ns *NullUserString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	*ns = NullUserString{s.String, s.Valid}

	return nil
}

// NullUserTime
type NullUserTime sql.NullTime

func (nt *NullUserTime) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	*nt = NullUserTime{t.Time, t.Valid}

	return nil
}

// NullUserBool
type NullUserBool sql.NullBool

func (nb *NullUserBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	*nb = NullUserBool{b.Bool, b.Valid}

	return nil
}

// MarshalJSON for NullUserInt64
func (ni *NullUserInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullUserInt64
func (ni *NullUserInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// MarshalJSON for NullUserBool
func (nb *NullUserBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullUserBool
func (nb *NullUserBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = (err == nil)
	return err
}

// MarshalJSON for NullUserString
func (ns *NullUserString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullUserString
func (ns *NullUserString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// MarshalJSON for NullUserTime
func (nt *NullUserTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nt.Time)
}

// UnmarshalJSON for NullUserTime
func (nt *NullUserTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	re := strings.Replace(s, "\"", "", -1)

	const layout = "2006-01-02 03:04:05"
	x, err := time.Parse(layout, re)

	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}