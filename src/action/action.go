package action

import (
	"net/http"
	"fmt"
)

func Output(w http.ResponseWriter, msg ...interface{}) {
	fmt.Fprint(w, msg...)
}
