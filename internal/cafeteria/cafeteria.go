package cafeteria

import (
	"errors"
)


type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	Cliente    Cliente
	Producto   Producto
	Cantidad   int
	Total      float64
	Fecha      string
}

var (
	ErrClienteNoEncontrado = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente   = errors.New("stock insuficiente")
	ErrSaldoInsuficiente   = errors.New("saldo insuficiente")
)

type Repository interface {
	GuardarCliente(cliente Cliente) error
	ObtenerCliente(id int) (Cliente, error)
	ListarClientes() []Cliente


	GuardarProducto(producto Producto) error
	ObtenerProducto(id int) (Producto, error)	
	ListarProductos() []Producto

	RegistrarPedido(pedido Pedido) error
}

//IMplementar RepoMemoriay

type RepoMemoria struct {
	clientes []Cliente
	productos []Producto
	pedidos []Pedido
}


//CLIENTES
func (r *RepoMemoria) GuardarCliente(cliente Cliente) error {
	r.clientes = append(r.clientes, cliente)
	return nil
}

func (r *RepoMemoria) ObtenerCliente(id int) (Cliente, error) {
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}

func (r *RepoMemoria) ListarClientes() []Cliente { // Listar() ([]Cliente
	return r.clientes
}



//PRODUCTOS
func (r *RepoMemoria) GuardarProducto(producto Producto) error {
	r.productos = append(r.productos, producto)
	return nil
}

func (r *RepoMemoria) ObtenerProducto(id int) (Producto, error) {
	for _, p := range r.productos {
		if p.ID == id {	
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}


func (r *RepoMemoria) ListarProductos() []Producto {
	return r.productos
}


func (r *RepoMemoria) RegistrarPedido(pedido Pedido) error {
    r.pedidos = append(r.pedidos, pedido)
    return nil
}

func NewRepoMemoriaProductos() *RepoMemoria {
	return &RepoMemoria{clientes: []Cliente{}, productos: []Producto{}, pedidos: []Pedido{}}
}


func NewRepoMemoria() *RepoMemoria {
    return &RepoMemoria{}
}


// Verificación en tiempo de compilación:
// Si RepoMemoria NO cumple Repository, esto da error al compilar.
var _ Repository = (*RepoMemoria) (nil)
