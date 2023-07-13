package requests

type Header struct {
	Operations		    int   `json:"Operations"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
