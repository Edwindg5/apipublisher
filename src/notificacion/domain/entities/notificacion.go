// api-database/src/notificacion/domain/entities/notificacion.go
package entities

type Notificacion struct {
	ID        int    `json:"id"`
	PedidoID  int    `json:"pedido_id"`
	Cliente   string `json:"cliente"`
	Producto  string `json:"producto"`
	Cantidad  int    `json:"cantidad"`
	Estado    string `json:"estado"`
	Fecha     string `json:"fecha"`
}
