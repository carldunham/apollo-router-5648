// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Item interface {
	IsItem()
	GetID() string
}

type Post interface {
	IsPost()
	GetID() string
}

type PostInfo interface {
	IsPostInfo()
	GetID() string
}

type UserInfo interface {
	IsUserInfo()
	GetID() string
}

type ASimplePost struct {
	ID string `json:"id"`
}

func (ASimplePost) IsPostInfo()        {}
func (this ASimplePost) GetID() string { return this.ID }

func (ASimplePost) IsPost() {}

func (ASimplePost) IsEntity() {}

type Identity struct {
	ID string `json:"id"`
}

func (Identity) IsEntity() {}

type ItemConnection struct {
	Edges []*ItemEdge `json:"edges"`
}

type ItemEdge struct {
	Node Item `json:"node,omitempty"`
}

type Query struct {
}

type UnavailableUser struct {
	ID                     string `json:"id"`
	IsPermanentlySuspended bool   `json:"isPermanentlySuspended"`
}

func (UnavailableUser) IsUserInfo()        {}
func (this UnavailableUser) GetID() string { return this.ID }

func (UnavailableUser) IsEntity() {}

type User struct {
	ID string `json:"id"`
}

func (User) IsUserInfo()        {}
func (this User) GetID() string { return this.ID }

func (User) IsEntity() {}

type TimeRange string

const (
	TimeRangeAll   TimeRange = "ALL"
	TimeRangeHour  TimeRange = "HOUR"
	TimeRangeDay   TimeRange = "DAY"
	TimeRangeWeek  TimeRange = "WEEK"
	TimeRangeMonth TimeRange = "MONTH"
	TimeRangeYear  TimeRange = "YEAR"
)

var AllTimeRange = []TimeRange{
	TimeRangeAll,
	TimeRangeHour,
	TimeRangeDay,
	TimeRangeWeek,
	TimeRangeMonth,
	TimeRangeYear,
}

func (e TimeRange) IsValid() bool {
	switch e {
	case TimeRangeAll, TimeRangeHour, TimeRangeDay, TimeRangeWeek, TimeRangeMonth, TimeRangeYear:
		return true
	}
	return false
}

func (e TimeRange) String() string {
	return string(e)
}

func (e *TimeRange) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimeRange(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TimeRange", str)
	}
	return nil
}

func (e TimeRange) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}