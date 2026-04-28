package main

import (
	"errors"
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {

	//Crear repo usando la interfaz
	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()

	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Cristina", Carrera: "Arquitectura", Saldo: 80})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Lina", Carrera: "Enfermería", Saldo: 30})

	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Café", Precio: 5, Stock: 10, Categoria: "Bebida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Sándwich", Precio: 10, Stock: 5, Categoria: "Comida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Galleta", Precio: 2, Stock: 20, Categoria: "Snack"})

	// Obtener Cliente que SI existe

	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}


	//Obtener Cliente que NO existe
	c, err = repo.ObtenerCliente(99)
	if err != nil {
		if errors.Is(err, cafeteria.ErrClienteNoEncontrado) {
			fmt.Println("Cliente no existe:", err)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}

	// 5. Listar productos
	fmt.Println("\n=== LISTA DE PRODUCTOS ===")
	productos := repo.ListarProductos()
	for _, p := range productos {
		fmt.Printf("[%d] %s - $%.2f (Stock: %d)\n",
			p.ID, p.Nombre, p.Precio, p.Stock)
	}

	// 6. Mostrar Pedido con Cliente y Producto completos
	cliente, _ := repo.ObtenerCliente(1)
	producto, _ := repo.ObtenerProducto(1)

	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  cliente,
		Producto: producto,
		Cantidad: 2,
		Total:    float64(2) * producto.Precio,
		Fecha:    "2026-04-27",
	}

	fmt.Println("\n=== PEDIDO ===")
	fmt.Printf("Pedido ID: %d\n", pedido.ID)
	fmt.Printf("Cliente: %s (%s)\n", pedido.Cliente.Nombre, pedido.Cliente.Carrera)
	fmt.Printf("Producto: %s\n", pedido.Producto.Nombre)
	fmt.Printf("Cantidad: %d\n", pedido.Cantidad)
	fmt.Printf("Total: $%.2f\n", pedido.Total)
}

//Preguntas de Reflexión
//¿Tuviste que poner Cliente, Producto y Pedido en el mismo paquete? ¿Por qué sí o por qué no?
// R. Sí, los puse en el mismo paquete porque todos pertenecen al mismo sistema de la cafetería y están relacionados entre sí. Esto facilita organizar el código y usar las estructuras sin tener que importar varios paquetes.

//¿Qué problema aparecería si intentaras separar Producto en un paquete aparte cuando Pedido lo tiene anidado?
// R. Podría aparecer un problema porque Pedido usa Producto dentro de su estructura. Si Producto estuviera en otro paquete, habría que importarlo para poder usarlo. Si no se importa correctamente, el programa daría error porque no reconocería qué es Producto.

//Comparando con el Día A (donde usamos IDs): ¿qué ventaja tiene el modelo con IDs para organizar el código en paquetes? 
// R.Una ventaja de usar IDs es que no es necesario tener todas las estructuras juntas. En lugar de guardar todo el objeto, solo se guarda su identificador. Esto permite separar mejor el código en distintos paquetes y mantener el proyecto más organizado.