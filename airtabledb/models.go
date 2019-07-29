package airtabledb

import (
	"encoding/json"
	"time"

	"github.com/brianloveswords/airtable"
)

// Unfortunately, the order matters here.
// We must first compute Records which are referenced by other Records...
const (
	ProviderIndex = iota
	LabelIndex
	AccountIndex
	RepositoryIndex
	MilestoneIndex
	IssueIndex
	NumTables
)

var (
	TableNameToIndex = map[string]int{
		"provider":   ProviderIndex,
		"label":      LabelIndex,
		"account":    AccountIndex,
		"repository": RepositoryIndex,
		"milestone":  MilestoneIndex,
		"issue":      IssueIndex,
	}
)

//
// provider
//

type ProviderRecord struct {
	State State `json:"-"` // internal

	airtable.Record // provides ID, CreatedTime
	Fields          struct {
		// base
		Base

		// specific
		URL    string `json:"url"`
		Driver string `json:"driver"`

		// relationship
		// n/a
	} `json:"fields,omitempty"`
}

func (r ProviderRecord) String() string {
	out, _ := json.Marshal(r)
	return string(out)
}

//
// label
//

type LabelRecord struct {
	State State `json:"-"` // internal

	airtable.Record // provides ID, CreatedTime
	Fields          struct {
		// base
		Base

		// specific
		URL         string `json:"url"`
		Name        string `json:"name"`
		Color       string `json:"color"`
		Description string `json:"description"`

		// relationship
		// n/a
	} `json:"fields,omitempty"`
}

func (r LabelRecord) String() string {
	out, _ := json.Marshal(r)
	return string(out)
}

//
// account
//

type AccountRecord struct {
	State State `json:"-"` // internal

	airtable.Record // provides ID, CreatedTime
	Fields          struct {
		// base
		Base

		// specific
		URL       string `json:"url"`
		Login     string `json:"login"`
		FullName  string `json:"fullname"`
		Type      string `json:"type"`
		Bio       string `json:"bio"`
		Location  string `json:"location"`
		Company   string `json:"company"`
		Blog      string `json:"blog"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar-url"`

		// relationships
		Provider []string `json:"provider"`
	} `json:"fields,omitempty"`
}

func (r AccountRecord) String() string {
	out, _ := json.Marshal(r)
	return string(out)
}

//
// repository
//

type RepositoryRecord struct {
	State State `json:"-"` // internal

	airtable.Record // provides ID, CreatedTime
	Fields          struct {
		// base
		Base

		// specific
		URL         string    `json:"url"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Homepage    string    `json:"homepage"`
		PushedAt    time.Time `json:"pushed-at"`
		IsFork      bool      `json:"is-fork"`

		// relationships
		Provider []string `json:"provider"`
		Owner    []string `json:"owner"`
	} `json:"fields,omitempty"`
}

func (r RepositoryRecord) String() string {
	out, _ := json.Marshal(r)
	return string(out)
}

//
// milestone
//

type MilestoneRecord struct {
	State State `json:"-"` // internal

	airtable.Record // provides ID, CreatedTime
	Fields          struct {
		// base
		Base

		// specific
		URL         string    `json:"url"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		ClosedAt    time.Time `json:"closed-at"`
		DueOn       time.Time `json:"due-on"`

		// relationships
		Creator    []string `json:"creator"`
		Repository []string `json:"repository"`
	} `json:"fields,omitempty"`
}

func (r MilestoneRecord) String() string {
	out, _ := json.Marshal(r)
	return string(out)
}

//
// issue
//

type IssueRecord struct {
	State State `json:"-"` // internal

	airtable.Record // provides ID, CreatedTime
	Fields          struct {
		// base
		Base

		// specific
		URL         string    `json:"url"`
		CompletedAt time.Time `json:"completed-at"`
		Title       string    `json:"title"`
		State       string    `json:"state"`
		Body        string    `json:"body"`
		IsPR        bool      `json:"is-pr"`
		IsLocked    bool      `json:"is-locked"`
		Comments    int       `json:"comments"`
		Upvotes     int       `json:"upvotes"`
		Downvotes   int       `json:"downvotes"`
		IsOrphan    bool      `json:"is-orphan"`
		IsHidden    bool      `json:"is-hidden"`
		Weight      int       `json:"weight"`
		IsEpic      bool      `json:"is-epic"`
		HasEpic     bool      `json:"has-epic"`

		// relationships
		Repository []string `json:"repository"`
		Milestone  []string `json:"milestone"`
		Author     []string `json:"author"`
		Labels     []string `json:"labels"`
		Assignees  []string `json:"assignees"`
		//Parents      []string    `json:"-"`
		//Children     []string    `json:"-"`
		//Duplicates   []string    `json:"-"`
	} `json:"fields,omitempty"`
}

func (r IssueRecord) String() string {
	out, _ := json.Marshal(r)
	return string(out)
}

func NewDB() DB {
	db := DB{
		Tables: make([]Table, NumTables),
	}
	db.Tables[IssueIndex].elems = &[]IssueRecord{}
	db.Tables[RepositoryIndex].elems = &[]RepositoryRecord{}
	db.Tables[AccountIndex].elems = &[]AccountRecord{}
	db.Tables[LabelIndex].elems = &[]LabelRecord{}
	db.Tables[MilestoneIndex].elems = &[]MilestoneRecord{}
	db.Tables[ProviderIndex].elems = &[]ProviderRecord{}
	if len(db.Tables) != NumTables {
		panic("missing an airtabledb Table")
	}
	return db
}