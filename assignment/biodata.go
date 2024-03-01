package main

import (
	"fmt"
	"os"
)

type temanKelas struct {
	nama string
	alamat string
	pekerjaan string

}

func showTemanKelas(absen int, teman []temanKelas){

	if absen < 1 || absen > len(teman){
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan")
		return
	}

	temanku := teman[absen-1]

	fmt.Println("Siswa dengan nomor absen", absen)
	fmt.Println("Nama", temanku.nama)
	fmt.Println("alamat", temanku.alamat)
	fmt.Println("pekerjaan", temanku.pekerjaan)
}

func main (){
	temanku :=[]temanKelas{
		{"anggit", "jalan adi sucipto", "programmer"},
		{"wahyu", "jalan adi sucipto", "designer"},
		{"romadhon", "jalan adi sucipto", "writer"},
	}
	
	args := os.Args

	absen := 0
	_, err := fmt.Sscanf(args[1], "%d", & absen)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	showTemanKelas(absen, temanku)
}