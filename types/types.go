package types

import "gitlab.com/tslocum/cview"

const (
	FrontPage        = 0
	New              = 1
	Ask              = 2
	Show             = 3
	SubmissionsPanel = "0"
	HelpScreenPanel  = "help"
	ErrorScreenPanel = "error"
)

type ScreenController struct {
	Application      *cview.Application
	List             *cview.List
	MainView         *MainView
	Submissions      []*Submissions
	ApplicationState *ApplicationState
}

type Submission struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Points        int    `json:"points"`
	Author        string `json:"user"`
	Time          string `json:"time_ago"`
	CommentsCount int    `json:"comments_count"`
	URL           string `json:"url"`
	Domain        string `json:"domain"`
	Type          string `json:"type"`
}

type Submissions struct {
	MappedSubmissions  int
	MappedPages        int
	StoriesListed      int
	PageToFetchFromAPI int
	MaxPages           int
	Entries            []*Submission
}

type ApplicationState struct {
	SubmissionsToShow         int
	CurrentCategory           int
	ScreenHeight              int
	ScreenWidth               int
	CurrentPage               int
	IsOffline                 bool
	IsReturningFromSuspension bool
	IsOnHelpScreen            bool
}

type MainView struct {
	Grid          *cview.Grid
	Header        *cview.TextView
	LeftMargin    *cview.TextView
	Panels        *cview.Panels
	StatusBar     *cview.TextView
	PageIndicator *cview.TextView
}
