package handlers
import (
	"fmt"
	"net/http"
)

func quoteHandler(w http.ResponseWriter, r *http.Response) { 
	fmt.Println("Testing")
	return
} // quoteHandler