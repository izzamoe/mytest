package entities

type APIResponse struct {
	Status  string      `json:"status"`           // success atau error
	Message string      `json:"message"`          // Pesan deskriptif tentang status
	Data    interface{} `json:"data"`             // Data untuk respons sukses, nil jika error
	Code    int         `json:"code"`             // Kode HTTP (misalnya 200, 404, 500)
	Errors  interface{} `json:"errors,omitempty"` // Detail kesalahan, opsional
}
