package project

type Criteria struct {
	Description string
	Grade       float32
	Weight      float32 // from 0 to 1, sum of criterias = 1
}
