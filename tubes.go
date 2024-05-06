package main

import "fmt"

const NMAX int = 1000

type tabNilai struct{
	quiz int
	tugas int
	UTS int
	UAS int
}

type nilai [NMAX]tabNilai

type mahasiswa struct{
	nama string
	nim string
	kelas string
	semester int
	prodi string
	nilai tabNilai
}


func main() {
	mainMenu()
}

func mainMenu() {
	fmt.Println("=== Main Menu ===")
	fmt.Println("1. Mengisi data mahasiswa")
	fmt.Println("2. Menghapus data mahasiswa")
	fmt.Println("=== Main Menu ===")
	fmt.Println("=== Main Menu ===")
}
