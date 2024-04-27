

// Eliminar comentarios de una sola línea
func quitarComentariosPython(contenido string) string {
	re := regexp.MustCompile("#[^\n]*")
	contenido = re.ReplaceAllString(contenido, "")

	// Eliminar comentarios de varias líneas
	re = regexp.MustCompile("\"\"\"[\\s\\S]*?\"\"\"")
	contenido = re.ReplaceAllString(contenido, "")

	re = regexp.MustCompile("'''[\\s\\S]*?'''")
	contenido = re.ReplaceAllString(contenido, "")

	return contenido
}