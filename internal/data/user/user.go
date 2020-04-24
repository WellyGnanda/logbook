package user

import (
	"context"
	"log"

	logbookEntity "logbook/internal/entity/logbook"
	packinglistEntity "logbook/internal/entity/packinglist"

	"logbook/pkg/errors"

	"logbook/pkg/firebaseclient"

	"cloud.google.com/go/firestore"
	"github.com/jmoiron/sqlx"
	"google.golang.org/api/iterator"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
		c    *firestore.Client
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	getLogSql  = "GetLogSql"
	qGetLogSql = "SELECT DISTINCT * FROM m_cabang_dist a JOIN TH_PackingList b ON b.THP_DistCode = a.Dir_KdDistributor " +
		"WHERE Dir_AktifYN = 'Y' AND THP_NoPL NOT IN (?) AND CONCAT(Dir_kode, ' - ',Dir_nama) = ? "

	getDistSql  = "GetDistSql"
	qGetDistSql = "SELECT DISTINCT Dir_Distributor FROM m_cabang_dist"

	getCabDistributorSql  = "GetConcatSql"
	qGetCabDistributorSql = "SELECT CONCAT(Dir_kode, ' - ',Dir_nama) AS CabangDist FROM m_cabang_dist " +
		" WHERE Dir_Distributor = ? "
)

var (
	readStmt = []statement{
		{getLogSql, qGetLogSql},
		{getDistSql, qGetDistSql},
		{getCabDistributorSql, qGetCabDistributorSql},
	}
)

// New ...
func New(db *sqlx.DB, fc *firebaseclient.Client) Data {
	d := Data{
		db: db,
		c:  fc.Client,
	}
	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// GetFirebasePackingListData digunakan untuk mengambil semua data user
func (d Data) GetFirebasePackingListData(ctx context.Context) ([]packinglistEntity.PackingList, error) {
	var (
		pl  packinglistEntity.PackingList
		pla []packinglistEntity.PackingList
		err error
	)

	iter := d.c.Collection("PL").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return pla, errors.Wrap(err, "[DATA][TampilanSemuaData] Failed to iterate Document!")
		}
		err = doc.DataTo(&pl)
		if err != nil {
			return pla, errors.Wrap(err, "[DATA][TampilanSemuaData] Failed to Populate Struct!")
		}
		pla = append(pla, pl)
	}
	return pla, err
}

//InsertFirebasePL ...
func (d Data) InsertFirebasePL(ctx context.Context, packinglist packinglistEntity.PackingList) error {
	_, err := d.c.Collection("PL").Doc(packinglist.THP_DistCode).Set(ctx, packinglist)
	return err
}

//GetByNoPL ...
func (d Data) GetByNoPL(ctx context.Context, NoPL string) (packinglistEntity.PackingList, error) {
	doc, err := d.c.Collection("PL").Doc(NoPL).Get(ctx)
	var packingList packinglistEntity.PackingList
	err = doc.DataTo(&packingList)
	if err != nil {
		return packingList, err
	}
	if doc == nil {
		return packingList, errors.Wrap(err, "Data Not Exist")
	}

	return packingList, err
}

//UpdateData ...
func (d Data) UpdateData(ctx context.Context, updatepackinglist packinglistEntity.UpdatePackingList) (packinglistEntity.PackingList, error) {
	_, err := d.c.Collection("PL").Doc(updatepackinglist.Bri_NoPL).Update(ctx, []firestore.Update{{Path: "thp_BeratTotalReal", Value: updatepackinglist.THP_BeratTotalReal}})
	doc, err := d.c.Collection("PL").Doc(updatepackinglist.Bri_NoPL).Get(ctx)
	var packingList packinglistEntity.PackingList
	err = doc.DataTo(&packingList)
	if err != nil {
		return packingList, err
	}
	if doc == nil {
		return packingList, errors.Wrap(err, "Data Not Exist")
	}
	return packingList, err
}

//InsertFirebaseLB ...
func (d Data) InsertFirebaseLB(ctx context.Context, logbook logbookEntity.LogBook, maxLogID int) (map[string]string, error) {
	var logbookID map[string]string
	logbookID = make(map[string]string)
	logbookID["logbook_id"] = logbook.Log_ID
	

	_, err := d.c.Collection("nomorMax").Doc("NoLogBook").Update(ctx, []firestore.Update{{
		Path: "nomorMaxLogBook", Value: maxLogID,
	}})

	_, err = d.c.Collection("LogBook").Doc(logbook.Log_ID).Set(ctx, logbook)

	return logbookID, err
}

//LogIDIncrement ...
func (d Data) LogIDIncrement(ctx context.Context) (int, error) {

	doc, err := d.c.Collection("nomorMax").Doc("NoLogBook").Get(ctx)
	max, err := doc.DataAt("nomorMaxLogBook")
	maxLogID := int(max.(int64))
	return maxLogID, err

}

// GetNoPLLB digunakan untuk mengambil semua data user
func (d Data) GetNoPLLB(ctx context.Context) ([]logbookEntity.LogBook, error) {
	var (
		noPL  logbookEntity.LogBook
		noPLa []logbookEntity.LogBook
		err   error
	)

	iter := d.c.Collection("LogBook").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return noPLa, errors.Wrap(err, "[DATA][TampilanSemuaData] Failed to iterate Document!")
		}
		err = doc.DataTo(&noPL)
		if err != nil {
			return noPLa, errors.Wrap(err, "[DATA][TampilanSemuaData] Failed to Populate Struct!")
		}
		noPLa = append(noPLa, noPL)
	}
	return noPLa, err
}

//GetLogSql ...
func (d Data) GetLogSql(ctx context.Context, noPL []string, concat string) ([]logbookEntity.Test, error) {
	var (
		// datas []map[string]interface{}
		data  logbookEntity.Test
		datas []logbookEntity.Test
		// data  interface{}
		//noPLs string
		err      error
		noPLTemp []string
	)
	if len(noPL) == 0 {
		noPLTemp = append(noPLTemp, "0")
		query, args, err := sqlx.In(qGetLogSql, noPLTemp, concat)
		rows, err := d.db.QueryxContext(ctx, query, args...)

		for rows.Next() {
			if err := rows.StructScan(&data); err != nil {
				return datas, errors.Wrap(err, "[DATA][GetPL3] ")
			}
			//log.Printf("%+v", data)
			datas = append(datas, data)
		}

		return datas, err
	}
	query, args, err := sqlx.In(qGetLogSql, noPL, concat)
	rows, err := d.db.QueryxContext(ctx, query, args...)

	for rows.Next() {
		if err := rows.StructScan(&data); err != nil {
			return datas, errors.Wrap(err, "[DATA][GetPL3] ")
		}
		//log.Printf("%+v", data)
		datas = append(datas, data)
	}

	return datas, err
}

//GetDistSql ...
func (d Data) GetDistSql(ctx context.Context) ([]string, error) {
	var (
		listDistributor []string
		Dir_Distributor string
		err             error
	)
	query, args, err := sqlx.In(qGetDistSql)
	rows, err := d.db.QueryxContext(ctx, query, args...)

	for rows.Next() {
		if err := rows.Scan(&Dir_Distributor); err != nil {
			return listDistributor, errors.Wrap(err, "[DATA][GetDistributor] ")
		}
		log.Printf("%+v", Dir_Distributor)
		listDistributor = append(listDistributor, Dir_Distributor)
	}

	return listDistributor, err
}

//GetCabDistributorSql ...
func (d Data) GetCabDistributortSql(ctx context.Context, distributor string) ([]string, error) {
	var (
		listContact []string
		contact     string
		err         error
	)
	query, args, err := sqlx.In(qGetCabDistributorSql, distributor)
	rows, err := d.db.QueryxContext(ctx, query, args...)

	for rows.Next() {
		if err := rows.Scan(&contact); err != nil {
			return listContact, errors.Wrap(err, "[DATA][GetDistributor] ")
		}
		//log.Printf("%+v", data)
		listContact = append(listContact, contact)
	}

	return listContact, err
}

//GetByLogID ...
func (d Data) GetByLogID(ctx context.Context, LogID string) (logbookEntity.LogBook, error) {
	doc, err := d.c.Collection("LogBook").Doc(LogID).Get(ctx)
	var logBook logbookEntity.LogBook
	err = doc.DataTo(&logBook)
	if err != nil {
		return logBook, err
	}
	if doc == nil {
		return logBook, errors.Wrap(err, "Data Not Exist")
	}

	return logBook, err
}

//GetByDONum ...
func (d Data) GetByDONum(ctx context.Context, DoNum string) ([]logbookEntity.LogBook, error) {
	iter := d.c.Collection("LogBook").Where("Log_DONum", "==", DoNum).Documents(ctx)
	var (
		logbookDos []logbookEntity.LogBook
		err        error
	)
	for {
		var logbookDo logbookEntity.LogBook

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&logbookDo)
		if err != nil {
			return logbookDos, errors.Wrap(err, "[DATA][TampilanSemuaData] Failed to Populate Struct!")
		}
		logbookDos = append(logbookDos, logbookDo)
	}
	return logbookDos, err
}

//GetByPLNum ...
func (d Data) GetByPLNum(ctx context.Context, PlNum string) ([]logbookEntity.LogBook, error) {
	iter := d.c.Collection("LogBook").Where("Log_PLNum", "==", PlNum).Documents(ctx)
	var (
		logbookPls []logbookEntity.LogBook
		err        error
	)
	for {
		var logbookPl logbookEntity.LogBook

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&logbookPl)
		if err != nil {
			return logbookPls, errors.Wrap(err, "[DATA][TampilanSemuaData] Failed to Populate Struct!")
		}
		logbookPls = append(logbookPls, logbookPl)
	}
	return logbookPls, err
}
