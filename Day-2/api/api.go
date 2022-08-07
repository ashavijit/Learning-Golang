package api

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mycodesmells/mongo-go-api/db"
	"github.com/mycodesmells/mongo-go-api/models"
	"strconv"
)

func handleError(err error,message string ,w http.ResponseWriter){
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(message))
		return
	}
}


func GetAllItems(w http.ResponseWriter, r *http.Request){
	/* items, err := db.GetAllItems()
	handleError(err,"Error getting all items",w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
} */

    rs, err := db.GetAllItems()
	if err !=nil {
		handleError(err,"Error getting all items",w)
		return
	}
	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err,"Error marshalling items",w)
		return
	}
	//We can use the Marshal function to convert Go objects to JSON
   w.Write(bs)
}

func PostItems(w http.ResponseWriter, r *http.Request){
	ID := r.FormValue("id")
	valueStr := r.FormValue("value")
	value, err := strconv.Atoi(valueStr)
	// Atoi() function which is equivalent to ParseInt(str string, base int, bitSize int) is used to convert string type into int type. To access Atoi() function you need to import strconv Package in your program.
	if err!=nil{
		handleError(err,"Error converting value",w)
		return
	}
	Item := db.item{ID:ID,Value:value}
	if err := db.CreateItem(Item); err != nil {
		handleError(err,"Error creating item",w)
		return
	}
   w.Write([]byte("Item created"))
}

func DleteItem(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if err := db.DeleteItem(id); err != nil {
		handleError(err,"Error deleting item",w)
		return
	}
	w.Write([]byte("Item deleted"))
	
}


