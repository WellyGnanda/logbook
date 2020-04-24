package packinglist

import (
	"time"
)

//PackingList Object Model
type PackingList struct {
	Bri_FlagTransaksi  string    `firebase:"bri_FlagTransaksi" json:"bri_flagtransaksi"`
	Bri_LastUpdate     time.Time `firebase:"bri_LastUpdate" json:"bri_lastupdate"`
	Bri_NoDO           string    `firebase:"bri_NoDO" json:"bri_nodo"`
	Bri_NoPL           string    `firebase:"bri_NoPL" json:"bri_nopl"`
	Bri_NoPO           string    `firebase:"bri_NoPO" json:"bri_nopo"`
	Bri_POPharmanetYN  string    `firebase:"bri_POPharmanetYN" json:"bri_popharmanetYN"`
	Bri_PesananID      string    `firebase:"bri_PesananID" json:"bri_pesananid"`
	Bri_PharmanetID    int       `firebase:"bri_PharmanetID" json:"bri_pharmentid"`
	Bri_PharmanetCode  string    `firebase:"bri_PharmanetCode" json:"bri_pharmanetcode"`
	Bri_RekapYN        string    `firebase:"bri_RekapYN" json:"bri_rekapyn"`
	TDPD_NoDO          string    `firebase:"tdpd_NoDO" json:"tdpd_nodo"`
	TDPD_NoPL          int       `firebase:"tdpd_NoPL" json:"tdpd_nopl"`
	TDPD_UpdateDate    time.Time `firebase:"tdpd_UpdateDate" json:"tdpd_updatedate"`
	TDOD_UpdateID      time.Time `firebase:"tdpd_UpdateID" json:"tdpd_updateid"`
	THP_AktifYN        string    `firebase:"thp_AktifYN" json:"thp_aktifyn"`
	THP_BEratTotalMax  float64   `firebase:"thp_BeratTotalMax" json:"thp_berattotalmax"`
	THP_BeratTotalMin  float64   `firebase:"thp_BeratTotalMin" json:"thp_berattotalmin"`
	THP_BeratTotalReal float64   `firebase:"thp_BeratTotalReal" json:"thp_berattotalreal"`
	THP_BeratTotalSHI  float64   `firebase:"thp_BeratTotalSHI" json:"thp_berattotalshi"`
	THP_DistCode       string    `firebase:"thp_DistCode" json:"thp_distcode"`
	THP_DistName       string    `firebase:"thp_DistName" json:"thp_distname"`
	THP_Flag           int       `firebase:"thp_Flag" json:"thp_flag"`
	THP_NIPPelaksana   string    `firebase:"thp_NIPPelaksana" json:"thp_nippelaksana"`
	THP_NipAngkat      string    `firebase:"thp_NipAngkat" json:"thp_nipangkat"`
	THP_NoPL           string    `firebase:"thp_NoPL" json:"thp_nopl"`
	THP_NoPOD          string    `firebase:"thp_NoPOD" json:"thp_nopod"`
	THP_PemesanID      string    `firebase:"thp_PemesanID" json:"thp_pemesanid"`
	THP_PrintKe        int       `firebase:"thp_PrintKe" json:"thp_printke"`
	THP_Status         string    `firebase:"thp_Status" json:"thp_status"`
	THP_TglPL          time.Time `firebase:"thp_TglPL" json:"thp_tglpl"`
	THP_TimbanganID    int       `firebase:"thp_TimbanganID" json:"thp_timbanganid"`
	THP_TujuanID       string    `firebase:"thp_TujuanID" json:"thp_tujuanid"`
	THP_UpdateDate     time.Time `firebase:"thp_UpdateDate" json:"thp_updatedate"`
	THP_UpdateID       string    `firebase:"thp_UpdateID" json:"thp_updateid"`
	THP_Via            string    `firebase:"thp_Via" json:"thp_via"`
}

type UpdatePackingList struct {
	Bri_NoPL           string  `firebase:"bri_NoPL" json:"bri_nopl"`
	THP_BeratTotalReal float64 `firebase:"thp_BeratTotalReal" json:"thp_berattotalreal"`
}
