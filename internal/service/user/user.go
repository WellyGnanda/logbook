package skeleton

import (
	"context"
	"fmt"
	"log"
	"strconv"

	logbookEntity "logbook/internal/entity/logbook"
	packinglistEntity "logbook/internal/entity/packinglist"
	"logbook/pkg/errors"
)

// LogBook ...
type LogBook interface {
	GetFirebasePackingListData(ctx context.Context) ([]packinglistEntity.PackingList, error)
	InsertFirebasePL(ctx context.Context, packinglist packinglistEntity.PackingList) error
	GetByNoPL(ctx context.Context, NoPL string) (packinglistEntity.PackingList, error)
	UpdateData(ctx context.Context, updatepackinglist packinglistEntity.UpdatePackingList) (packinglistEntity.PackingList, error)
	InsertFirebaseLB(ctx context.Context, logbook logbookEntity.LogBook, maxLogID int) (map[string]string, error)
	GetNoPLLB(ctx context.Context) ([]logbookEntity.LogBook, error)
	GetLogSql(ctx context.Context, noPL []string, concat string) ([]logbookEntity.Test, error)
	GetByLogID(ctx context.Context, LogID string) (logbookEntity.LogBook, error)
	LogIDIncrement(ctx context.Context) (int, error)
	GetByDONum(ctx context.Context, DoNum string) ([]logbookEntity.LogBook, error)
	GetByPLNum(ctx context.Context, PlNum string) ([]logbookEntity.LogBook, error)
	GetDistSql(ctx context.Context) ([]string, error)
	GetCabDistributortSql(ctx context.Context, distributor string) ([]string, error)
}

// Service ...
type Service struct {
	logBook LogBook
}

// New ...
func New(logBook LogBook) Service {
	return Service{
		logBook: logBook,
	}
}

// GetFirebasePackingListData ...
func (s Service) GetFirebasePackingListData(ctx context.Context) ([]packinglistEntity.PackingList, error) {
	result, err := s.logBook.GetFirebasePackingListData(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

//InsertFirebasePL ...
func (s Service) InsertFirebasePL(ctx context.Context, packinglist packinglistEntity.PackingList) error {
	err := s.logBook.InsertFirebasePL(ctx, packinglist)
	return err
}

// InsertFirebaseLB ...
func (s Service) InsertFirebaseLB(ctx context.Context, logbook logbookEntity.LogBook) (map[string]string, error) {

	var logbookID map[string]string
	//logbookID = make(map[string]string)
	//logbookID["logbook_id"] = logbook.Log_ID

	maxLogID, err := s.logBook.LogIDIncrement(ctx)
	maxLogID = maxLogID + 1
	logbook.Log_ID = logbook.Log_OutDest + fmt.Sprintf("%09d", maxLogID)
	log.Println("servicenya ", logbook)
	logbookID, err = s.logBook.InsertFirebaseLB(ctx, logbook, maxLogID)

	return logbookID, err
}

//GetByNoPL ...
func (s Service) GetByNoPL(ctx context.Context, NoPL string) (packinglistEntity.PackingList, error) {
	packingList, err := s.logBook.GetByNoPL(ctx, NoPL)
	return packingList, err
}

//UpdateData ...
func (s Service) UpdateData(ctx context.Context, updatepackinglist []packinglistEntity.UpdatePackingList) ([]packinglistEntity.PackingList, error) {
	var (
		err          error
		packinglist  packinglistEntity.PackingList
		packinglists []packinglistEntity.PackingList
	)

	log.Println("lennya", len(updatepackinglist))

	for i := 0; i < len(updatepackinglist); i++ {
		packinglist, err = s.logBook.UpdateData(ctx, updatepackinglist[i])
		packinglists = append(packinglists, packinglist)

	}

	return packinglists, err
}

//GetPL ...
func (s Service) GetPL(ctx context.Context, concat string) ([]logbookEntity.Test, error) {

	var (
		data  []logbookEntity.Test
		noPLs []string
	)

	dataPLLB, err := s.logBook.GetNoPLLB(ctx)
	//dataPL, err := s.logBook.GetFirebasePackingListData(ctx)

	if err != nil {
		log.Fatal(err)
	}

	for _, res := range dataPLLB {
		log.Println(res)
		stringNoPL := res.Log_PLNum[3:12]
		intPL, err := strconv.Atoi(stringNoPL)
		if err != nil {
			log.Println(err)
		}
		noPL := strconv.Itoa(intPL)
		noPLs = append(noPLs, noPL)
		//log.Println(res.Log_PLNum)
	}

	log.Println("log ", noPLs)

	data, err = s.logBook.GetLogSql(ctx, noPLs, concat)

	return data, err
}

//GetByLogID ...
func (s Service) GetByLogID(ctx context.Context, LogID string) (logbookEntity.LogBook, error) {
	logBook, err := s.logBook.GetByLogID(ctx, LogID)
	return logBook, err
}

//GetByDONum ...
func (s Service) GetByDONum(ctx context.Context, DoNum string) ([]logbookEntity.LogBook, error) {
	logBook, err := s.logBook.GetByDONum(ctx, DoNum)
	return logBook, err
}

//GetByPLNum ...
func (s Service) GetByPLNum(ctx context.Context, PlNum string) ([]logbookEntity.LogBook, error) {
	logBook, err := s.logBook.GetByPLNum(ctx, PlNum)
	return logBook, err
}

// GetDistSql ...
func (s Service) GetDistSql(ctx context.Context) ([]string, error) {
	distributor, err := s.logBook.GetDistSql(ctx)
	if err != nil {
		return distributor, errors.Wrap(err, "[SERVICE][GetDistSql]")
	}
	return distributor, err
}

// GetCabDistributortSql ...
func (s Service) GetCabDistributortSql(ctx context.Context, distributor string) ([]string, error) {
	cabDistributor, err := s.logBook.GetCabDistributortSql(ctx, distributor)
	if err != nil {
		return cabDistributor, errors.Wrap(err, "[SERVICE][GetCabDistSql]")
	}
	return cabDistributor, err
}
