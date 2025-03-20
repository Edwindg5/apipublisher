package entities

type NotificacionResumida struct {
	Producto string `json:"producto"`
	Cantidad int    `json:"cantidad"`
	Estado   string `json:"estado"`
	Fecha    string `json:"fecha"`
}
