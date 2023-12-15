package helper

// Nama Variable harus diawali huruf besar jika mau dipanggil di luar package
var version = "1.0.0"
var Application = "Golang"

func SayHello(name string) string {
	return "Hello " + name
}

// Nama Function harus diawali huruf besar untuk bisa dipanggil di luar file atau package
func SayGoodbye(name string) string {
	return "Goodbye " + name
}
