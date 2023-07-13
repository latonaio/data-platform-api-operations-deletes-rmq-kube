package requests

type Item struct {
	Operations		    int   `json:"Operations"`
	OperationsItem		int   `json:"OperationsItem"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
