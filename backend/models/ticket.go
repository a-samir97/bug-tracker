package models

// Customer will create this ticket, then the stuff will check it first then convert this ticket to bug
type Ticket struct {
	Base
	Customer    User // who will report this ticket to our service
	SolvedAt    string
	Description string
	Content     string
	Status      string
}

// TODO: add Message Model, for every ticket, there will be message between customer and our agent
func (ticket *Ticket) CreateTicket() error {

	return nil
}

func (ticket *Ticket) UpdateTicket(ticketId int) error {

	return nil
}

func (ticket *Ticket) GetTicket(ticketId int) error {

	return nil
}

func (ticket *Ticket) SolveTicket(ticketId int) error {

	return nil
}
