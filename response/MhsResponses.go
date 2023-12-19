package response

type DataMhs struct {
	Nama       string `json:"nama"`
	Nim        string `json:"nim"`
	Pembimbing string `json:"pembimbing"`
}

type KrsCourse struct {
	KodeMk   string `json:"kode_mk"`
	NamaMk   string `json:"nama_mk"`
	Sks      string `json:"sks"`
	Semester string `json:"semester"`
}
