//Package numeros donde se hace la conversion
package numeros

//Convertir es una funcion para convertir un numero a letras
func Convertir(n int) string {
	var respuesta string
	if n >= 10 {
		respuesta = otrosdiez[n%10]
	} else {
		respuesta = unidades[n]
	}
	return respuesta
}
