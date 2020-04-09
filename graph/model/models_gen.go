// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Case struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Bookmarked    bool      `json:"bookmarked"`
	Integrations  []*string `json:"integrations"`
	Collaborators []*string `json:"collaborators"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type Jotter struct {
	ID           int       `json:"id"`
	Usecase      string    `json:"usecase"`
	Index        int       `json:"index"`
	Name         string    `json:"name"`
	Content      string    `json:"content"`
	Contributors []string  `json:"contributors"`
	Completed    bool      `json:"completed"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type NewCase struct {
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Bookmarked    *bool     `json:"bookmarked"`
	Integrations  []*string `json:"integrations"`
	Collaborators []*string `json:"collaborators"`
}

type NewJotter struct {
	Index        int      `json:"index"`
	Name         string   `json:"name"`
	Content      string   `json:"content"`
	Contributors []string `json:"contributors"`
	Usecase      string   `json:"usecase"`
}

type NewOrganization struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Cases     []*string `json:"cases"`
	Usecases  []*string `json:"usecases"`
	CreatedBy string    `json:"createdBy"`
	Members   []*string `json:"members"`
}

type NewUsecase struct {
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Content       *string   `json:"content"`
	Tags          *string   `json:"tags"`
	Collaborators []*string `json:"collaborators"`
	Bookmarked    *bool     `json:"bookmarked"`
	Completed     *bool     `json:"completed"`
}

type NewUser struct {
	Name              string        `json:"name"`
	Email             *string       `json:"email"`
	IsOrganization    *bool         `json:"isOrganization"`
	Password          string        `json:"password"`
	Cases             []*NewCase    `json:"cases"`
	Usecases          []*NewUsecase `json:"usecases"`
	BokmarkedCases    []*NewCase    `json:"bokmarkedCases"`
	BokmarkedUsecases []*NewUsecase `json:"bokmarkedUsecases"`
}

type Organization struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Cases     []*string `json:"cases"`
	Usecases  []*string `json:"usecases"`
	Members   []*string `json:"members"`
}

type UpdateCase struct {
	Name          string    `json:"name"`
	Bookmarked    *bool     `json:"bookmarked"`
	Integrations  []*string `json:"integrations"`
	Collaborators []*string `json:"collaborators"`
}

type UpdateJotter struct {
	Name         string   `json:"name"`
	Content      string   `json:"content"`
	Contributors []string `json:"contributors"`
}

type UpdateOrganization struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Cases    []*string `json:"cases"`
	Usecases []*string `json:"usecases"`
	Members  []*string `json:"members"`
}

type UpdateUsecase struct {
	Title         string    `json:"title"`
	Content       *string   `json:"content"`
	Tags          *string   `json:"tags"`
	Collaborators []*string `json:"collaborators"`
	Bookmarked    *bool     `json:"bookmarked"`
	Completed     *bool     `json:"completed"`
}

type UpdateUser struct {
	Email             *string       `json:"email"`
	IsOrganization    *bool         `json:"isOrganization"`
	Password          string        `json:"password"`
	Cases             []*NewCase    `json:"cases"`
	Usecases          []*NewUsecase `json:"usecases"`
	BokmarkedCases    []*NewCase    `json:"bokmarkedCases"`
	BokmarkedUsecases []*NewUsecase `json:"bokmarkedUsecases"`
}

type Usecase struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Content       *string   `json:"content"`
	Tags          []*string `json:"tags"`
	Completed     bool      `json:"completed"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Collaborators []*string `json:"collaborators"`
	Bookmarked    bool      `json:"bookmarked"`
	Cases         []*Case   `json:"cases"`
}

type User struct {
	ID                int        `json:"id"`
	Name              string     `json:"name"`
	Email             *string    `json:"email"`
	Password          string     `json:"password"`
	IsOrganization    bool       `json:"isOrganization"`
	BokmarkedCases    []*Case    `json:"bokmarkedCases"`
	BokmarkedUsecases []*Usecase `json:"bokmarkedUsecases"`
	Cases             []*Case    `json:"cases"`
	Usecase           []*Usecase `json:"usecase"`
	CreatedAt         time.Time  `json:"createdAt"`
}
