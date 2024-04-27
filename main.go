package main

import (
	"fmt"
	"io/ioutil"

	"github.com/culturadevops/chatgpt-golang/gpt"
)

func LeerArchivo(rutaArchivo string) (string, error) {
	contenido, err := ioutil.ReadFile(rutaArchivo)
	if err != nil {
		return "", err
	}
	texto := string(contenido)
	return texto, nil
}
func escribirArchivo(nombreArchivo string, contenido string) error {
	// Escribe el contenido en el archivo
	err := ioutil.WriteFile(nombreArchivo, []byte(contenido), 0644)
	if err != nil {
		return err
	}
	return nil
}
func Traductor(idioma, Origen string) (string, error) {
	resultado := fmt.Sprintf("traduce a %s lo siguiente '%s'", idioma, Origen)
	return gpt.VarGpt.ChatWithTextDavinci002(resultado)
}
func GeneradordeContenido(texto, archivo string) (string, error) {
	resultado := fmt.Sprintf("crea un articulo sobre '%s'", texto)
	codigonuevo, err := gpt.VarGpt.ChatWithTextDavinci003(resultado)
	if err == nil {
		escribirArchivo(archivo, codigonuevo)
	}
	return codigonuevo, err
}
func ResumeContenido(texto, archivo string) (string, error) {
	resultado := fmt.Sprintf("creame un resumen corto sobre '%s'", texto)
	codigonuevo, err := gpt.VarGpt.ChatWithTextDavinci003(resultado)
	if err == nil {
		escribirArchivo(archivo, codigonuevo)
	}
	return codigonuevo, err
}
func AnalizadorCodigo(archivo string) (string, error) {
	file, _ := LeerArchivo(archivo)
	resultado := fmt.Sprintf("analizar el codigo y crea un resumen corto '%s'", file)
	return gpt.VarGpt.ChatWithTextDavinci002(resultado)
}
func Ortografia(archivo string) (string, error) {
	file, _ := LeerArchivo(archivo)
	resultado := fmt.Sprintf("Corrige los errores ortograficos de '%s'", file)
	return gpt.VarGpt.ChatWithTextDavinci002(resultado)
}
func TraductorDeCodigo(archivo string, lenguaje string) (string, error) {
	file, _ := LeerArchivo(archivo)
	resultado := fmt.Sprintf("dame el siguiente codigo en lenguaje '%s' '%s'", lenguaje, file)
	codigonuevo, err := gpt.VarGpt.ChatWithTextDavinci003(resultado)
	if err == nil {
		escribirArchivo(archivo+".go", codigonuevo)
	}
	return codigonuevo, err
}
func main() {
	gpt.VarGpt = &gpt.Gpt{}
	gpt.VarGpt.SetApi("config", "yml", 500, "")

	//respuesta, Mesajeerror := Traductor("ingles", "hola como estas?")
	//respuesta, Mesajeerror := Ortografia("erroresortograficos.txt")
	//respuesta, Mesajeerror := TraductorDeCodigo("python.py", "golang")
	//respuesta, Mesajeerror := AnalizadorCodigo("python.py")
	//respuesta, Mesajeerror := GeneradordeContenido("programcion orientada a objetos", "oop.txt")
	respuesta, Mesajeerror := ResumeContenido("oop.txt", "resumido.txt")
	if Mesajeerror != nil {
		fmt.Println(Mesajeerror.Error())
	} else {
		fmt.Println(respuesta)
	}

}
