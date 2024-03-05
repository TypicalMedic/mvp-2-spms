package handlers

import (
	"net/http"
	"strconv"
)

func Ping(w http.ResponseWriter, r *http.Request) {

}

// move to adapters later
func GetCredentials(r *http.Request) Credntials {
	professorId, _ := strconv.ParseUint(r.Header.Get("Professor-Id"), 10, 32)
	return Credntials{
		ProfessorId: uint(professorId),
	}
}

type Credntials struct {
	ProfessorId uint
}
