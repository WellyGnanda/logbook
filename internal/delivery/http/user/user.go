package user

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	logbookEntity "logbook/internal/entity/logbook"
	packinglistEntity "logbook/internal/entity/packinglist"

	"logbook/pkg/response"
)

// LogBook is an interface to User Service
type LogBook interface {
	GetFirebasePackingListData(ctx context.Context) ([]packinglistEntity.PackingList, error)
	InsertFirebasePL(ctx context.Context, packinglist packinglistEntity.PackingList) error
	GetByNoPL(ctx context.Context, NoPL string) (packinglistEntity.PackingList, error)
	//UpdateData(ctx context.Context, updatepackinglist packinglistEntity.UpdatePackingList) (packinglistEntity.PackingList, error)
	InsertFirebaseLB(ctx context.Context, logbook logbookEntity.LogBook) (map[string]string, error)
	GetPL(ctx context.Context, concat string) ([]logbookEntity.Test, error)
	UpdateData(ctx context.Context, updatepackinglist []packinglistEntity.UpdatePackingList) ([]packinglistEntity.PackingList, error)
	GetByLogID(ctx context.Context, LogID string) (logbookEntity.LogBook, error)
	GetByDONum(ctx context.Context, DoNum string) ([]logbookEntity.LogBook, error)
	GetByPLNum(ctx context.Context, PlNum string) ([]logbookEntity.LogBook, error)
	GetDistSql(ctx context.Context) ([]string, error)
	GetCabDistributortSql(ctx context.Context, distributor string) ([]string, error)
}

type (
	// Handler ...
	Handler struct {
		logBook LogBook
	}
)

// New for user domain handler initialization
func New(is LogBook) *Handler {
	return &Handler{
		logBook: is,
	}
}

// UserHandler will return user data
func (h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp        *response.Response
		metadata    interface{}
		result      interface{}
		err         error
		errRes      response.Error
		packinglist packinglistEntity.PackingList
		//updatepackinglist packinglistEntity.UpdatePackingList
		updatepackinglists []packinglistEntity.UpdatePackingList
		logbook            logbookEntity.LogBook
		//logbookID 			map[string]string
	)
	// Make new response object
	resp = &response.Response{}
	body, _ := ioutil.ReadAll(r.Body)
	// Defer will be run at the end after method finishes
	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:
		if _, x := r.URL.Query()["typeGet"]; x {
			_typeGet := r.FormValue("typeGet")
			switch _typeGet {

			case "PackingListData":
				result, err = h.logBook.GetFirebasePackingListData(context.Background())

			case "GetByNoPL":
				result, err = h.logBook.GetByNoPL(context.Background(), r.FormValue("NoPL"))

			case "GetPLLB":
				result, err = h.logBook.GetPL(context.Background(), r.FormValue("dist"))

			case "GetByLogID":
				result, err = h.logBook.GetByLogID(context.Background(), r.FormValue("LogID"))

			case "GetByDONum":
				result, err = h.logBook.GetByDONum(context.Background(), r.FormValue("Log_DONum"))

			case "GetByPLNum":
				result, err = h.logBook.GetByPLNum(context.Background(), r.FormValue("Log_PLNum"))

			case "GetDistSql":
				result, err = h.logBook.GetDistSql(context.Background())

			case "GetCabDistributorSql":
				result, err = h.logBook.GetCabDistributortSql(context.Background(), r.FormValue("Dir_Distributor"))
			}
		}

	case http.MethodPost:
		if _, x := r.URL.Query()["typePost"]; x {
			_typePost := r.FormValue("typePost")
			switch _typePost {

			case "FirebasePL":
				json.Unmarshal(body, &packinglist)
				err = h.logBook.InsertFirebasePL(context.Background(), packinglist)

			case "FirebaseLB":
				json.Unmarshal(body, &logbook)
				log.Println("deliverynya ", logbook)
				result, err = h.logBook.InsertFirebaseLB(context.Background(), logbook)

			}
		}

	case http.MethodPut:
		json.Unmarshal(body, &updatepackinglists)

		if _, x := r.URL.Query()["typePut"]; x {
			_typePut := r.FormValue("typePut")
			switch _typePut {

			case "PutUpdateData":
				log.Println("oooo", updatepackinglists)
				result, err = h.logBook.UpdateData(context.Background(), updatepackinglists)

			}
		}

	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		// Logging
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	// Inserting data to response
	resp.Data = result
	resp.Metadata = metadata
	// Logging
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
