package logbook

import (
	"time"

	"gopkg.in/guregu/null.v3/zero"
)

type LogBook struct {
	Log_ID             string    `db:"Log_ID" json:"log_id"`
	Log_Date           time.Time `db:"Log_Date" json:"log_date"`
	Log_CabAms         string    `db:"Log_CabAms" json:"log_cabams"`
	Log_PLNum          string    `db:"Log_PLNum" json:"log_plnum"`
	Log_DONum          string    `db:"Log_DONum" json:"log_donum"`
	Log_OutDest        string    `db:"Log_OutDest" json:"log_outdest"`
	Log_OutName        string    `db:"Log_OutName" json:"log_outname"`
	Log_TotBerat       float64   `db:"Log_TotBerat" json:"log_totberat"`
	Log_TotProcod      int       `db:"Log_TotProcod" json:"log_totprocod"`
	Log_LastUpdate     time.Time `db:"Log_LastUpdate" json:"log_lastupdate"`
	Log_ActiveYN       string    `db:"Log_ActiveYN" json:"log_activeyn"`
	Scan_OutNumber     int       `db:"Scan_OutNumber" json:"scan_outnumber"`
	Scan_OutDate       time.Time `db:"Scan_OutDate" json:"scan_outdate"`
	Scan_OutCabDist    string    `db:"Scan_OutCabDist" json:"scan_outcabdist"`
	Scan_OutBeratKoli  int       `db:"Scan_OutBeratColi" json:"scan_outberatcoli"`
	Scan_OutPLNum      string    `db:"Scan_OutPLNum" json:"scan_outplnum"`
	Scan_OutDONum      string    `db:"Scan_OutDONum" json:"scan_outdonum"`
	Scan_OutOutletDest string    `db:"Scan_OutOutletDest" json:"scan_outoutletdest"`
	Scan_OutOutletName string    `db:"Scan_OutOutletName" json:"scan_outoutletname"`
	Scan_OutNIP        string    `db:"Scan_OutNIP" json:"scan_outnip"`
	Scan_OutActiveYN   string    `db:"Scan_OutActiveYN" json:"scan_outactiveyn"`
	Scan_OutLastUpdate time.Time `db:"Scan_OutLastUpdate" json:"scan_outlastupdate"`
	Scan_InNumber      int       `db:"Scan_InNumber" json:"scan_innumber"`
	Scan_InDate        time.Time `db:"Scan_InDate" json:"scan_indate"`
	Scan_InActiveYN    string    `db:"Scan_InActiveYN" json:"scan_inactiveyn"`
	Scan_InLastUpdate  time.Time `db:"Scan_InLastUpdate" json:"scan_inlastupdate"`
}

type Test struct {
	THP_NoPL           zero.String `db:"THP_NoPL" json:"thp_nopl"`
	THP_TglPL          zero.Time   `db:"THP_TglPL" json:"thp_tglpl"`
	THP_DistCode       zero.String `db:"THP_DistCode" json:"thp_distcode"`
	THP_DistName       zero.String `db:"THP_DistName" json:"thp_distname"`
	THP_PemesanID      zero.String `db:"THP_PemesanID" json:"thp_pemesanid"`
	THP_TujuanID       zero.String `db:"THP_TujuanID" json:"thp_tujuanid"`
	THP_Via            zero.String `db:"THP_Via" json:"thp_via"`
	THP_NIPPelaksana   zero.String `db:"THP_NIPPelaksana" json:"thp_nippelaksana"`
	THP_BeratTotalMax  zero.Float  `db:"THP_BeratTotalMax" json:"thp_berattotalmax"`
	THP_BeratTotalMin  zero.Float  `db:"THP_BeratTotalMin" json:"thp_berattotalmin"`
	THP_BeratTotalSHI  zero.Float  `db:"THP_BeratTotalSHI" json:"thp_berattotalshi"`
	THP_BeratTotalReal zero.Float  `db:"THP_BeratTotalReal" json:"thp_berattotalreal"`
	THP_Status         zero.String `db:"THP_Status" json:"thp_status"`
	THP_PrintKe        zero.Int    `db:"THP_PrintKe" json:"thp_printke"`
	THP_TimbanganID    zero.Int    `db:"THP_TimbanganID" json:"thp_timbanganid"`
	THP_AktifYN        zero.String `db:"THP_AktifYN" json:"thp_aktifyn"`
	THP_UpdateID       zero.String `db:"THP_UpdateID" json:"thp_updateid"`
	THP_UpdateDate     zero.Time   `db:"THP_UpdateDate" json:"thp_updatedate"`
	THP_NIPAngkat      zero.String `db:"THP_NIPAngkat" json:"thp_nipangkat"`
	THP_NoPOD          zero.String `db:"THP_NoPOD" json:"thp_nopod"`
	THP_Flag           zero.Int    `db:"THP_Flag" json:"thp_flag"`
	Dir_kode           zero.String `db:"Dir_kode" json:"dir_kode"`
	Dir_nama           zero.String `db:"Dir_nama" json:"dir_nama"`
	Dir_alamat         zero.String `db:"Dir_alamat" json:"dir_alamat"`
	Dir_daerah         zero.String `db:"Dir_daerah" json:"dir_daerah"`
	Dir_kota           zero.String `db:"Dir_kota" json:"dir_kota"`
	Dir_KDPOS          zero.String `db:"Dir_KDPOS" json:"dir_kdpos"`
	Dir_TELP1          zero.String `db:"Dir_TELP1" json:"dir_telp1"`
	Dir_TELP2          zero.String `db:"Dir_TELP2" json:"dir_telp2"`
	Dir_FAX            zero.String `db:"Dir_FAX" json:"dir_fax"`
	Dir_email          zero.String `db:"Dir_email" json:"dir_email"`
	Dir_UPDATEID       zero.String `db:"Dir_UPDATEID" json:"dir_updateid"`
	Dir_UPDATETIME     zero.Time   `db:"Dir_UPDATETIME" json:"dir_updatetime"`
	Dir_Lintang        zero.Float  `db:"Dir_Lintang" json:"dir_lintang"`
	Dir_Bujur          zero.Float  `db:"Dir_Bujur" json:"dir_bujur"`
	Dir_Distributor    zero.String `db:"Dir_Distributor" json:"dir_distributor"`
	Dir_AktifYN        zero.String `db:"Dir_AktifYN" json:"dir_aktifyn"`
	Dir_KdCabDist      zero.String `db:"Dir_KdCabDist" json:"dir_kdcabdist"`
	Dir_KdDistributor  zero.String `db:"Dir_KdDistributor" json:"dir_kddistributor"`
	DownloadDate       zero.Time   `db:"DownloadDate" json:"downloaddate"`
}
