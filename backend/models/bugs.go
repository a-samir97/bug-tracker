package models

// Bug Model struct ...
type Bug struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Screenshot  byte   `json:"screenshot"`
	AssignedTo  User   `json:"assigned_to"` // this should be the developer will work to fix this bug
	Reporter    User   `json:"reporter"`    // will review the task after assignee finish it
	Status      string `json:"status"`
}

// Status of the bug: Opened, Closed, In Progress, To be tested
// Get assigne for this bug
func (bug *Bug) GetAssignedUser() User {
	return bug.AssignedTo
}

func (bug *Bug) GetReporter() User {
	return bug.Reporter
}

func (bug *Bug) GetBugStatus() string {
	return bug.Status
}
