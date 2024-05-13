package main

import "fmt"

const NMAX int = 1000

type tabNilai struct{
	matakuliah string
	quiz int
	total int
	UTS int
	UAS int
	grade string
}


type mahasiswa struct{
	nama string
	nim string
	kelas string
	semester int
	prodi string
	matkul [NMAX]tabNilai
}

type tabMhs [NMAX]mahasiswa

func main() {
	var A tabMhs
	var n int
	var j int
	mainMenu(&A, &n, &j)
}



func mainMenu(A *tabMhs, n *int, j *int) {
	var menu int
	fmt.Println("============ Main Menu ============")
	fmt.Println("1. Mengisi data mahasiswa          ")
	fmt.Println("2. Mengubah data mahasiswa         ")
	fmt.Println("3. Menambah data mahasiswa         ")
	fmt.Println("4. Menghapus data mahasiswa        ")
	fmt.Println("5. Menampilkan data mahasiswa      ")
	fmt.Println("6. Exit                            ")
	fmt.Println("===================================")


	fmt.Scan(&menu)
	switch menu {
	case 1:
		mengisiData(&A, &n, &j)
	case 2:
		mengubahData(&A, &n, &j)
	case 3:
		menambahData(&A, &n, &j)
	case 4:
		fmt.Println("Anda memilih untuk menghapus data mahasiswa")
	case 5:
		fmt.Println("Anda memilih untuk menampilkan data mahasiswa")
	case 6:
		fmt.Println("Terima kasih. Program berhenti.")
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
	}
}

func mengisiData(A *tabMhs, n *int, j *int) {
	var pilihan int
	menuIsi()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		dataMhs(&A, &n)
	case 2:
		nilaiMhs(&A, &n, &j)
	case 3:
		mainMenu(&A, &n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
	}
}

func menuIsi() {
	fmt.Println("======== Isi Data Mahasiswa ========")
	fmt.Println("1. Nama, NIM, Kelas, Prodi, Semester")
	fmt.Println("2. Nilai Mahasiswa                  ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func dataMhs(A *tabMhs, n *int) {
	fmt.Println("Masukkan jumlah data mahasiswa")
	fmt.Scan(&n)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	for i := 0; i < n; i++{
		fmt.Scan(&A[i].nama, &A[i].nim, &A[i].kelas, &A[i].prodi)
	}
}

func nilaiMhs(A *tabMhs, n *int, j *int) {
	var index, i int
	var nama string
	var matkul string

	fmt.Println("Masukkan nama mahasiswa")
	fmt.Scan(&nama)	
	index = SearchNama(A, n, nama)
	for index == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&nama)
		index = SearchNama(A, n, nama)
	}
	*j = 0
	fmt.Println("Masukkan nama mata kuliah")
	fmt.Scan(&matkul)
	for matkul != "0"{
		A[index].matkul[j].matakuliah = matkul
		fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
		fmt.Scan(&A[index].matkul[j].quiz, &A[index].matkul[j].UTS, &A[index].matkul[j].UAS)
		*j++
		fmt.Println("Masukkan nama mata kuliah atau ketik "0" jika sudah selesai")
		fmt.Scan(&matkul)
	}
}

func SearchNama(A tabMhs, n int, nama string) int {
	var i int
	var ketemu int

	ketemu = -1
	for i = 0; i < n; i++{
		if nama == A[i].nama {
			ketemu = i
		}
	}
	return ketemu
}

func mengubahData(A *tabMhs, n *int) {
	var pilihan int
	menuUbah()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		ubahData(&A, &n)
	case 2:
		ubahNilai(&A, &n, &j)
	case 3:
		mainMenu(&A, &n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
	}
}

func menuUbah() {
	fmt.Println("======== Ubah Data Mahasiswa ========")
	fmt.Println("1. Nama, NIM, Kelas, Prodi, Semester")
	fmt.Println("2. Nilai Mahasiswa                  ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func ubahData(A *tabMhs, n *int) {
	var index int
	var nama string
	fmt.Println("Masukkan nama mahasiswa yang datanya ingin diubah")
	fmt.Scan(&nama)
	index = SearchNama(A, n, nama)
	fmt.Println("Masukkan nama, nim, kelas dan prodi")
	fmt.Scan(&A[index].nama, &A[index].nim, &A[index].kelas, &A[index].prodi)
}

func ubahNilai(A *tabMhs, n *int, j *int) {
	var index, i int
	var nama, matkul string
	fmt.Println("Masukkan nama mahasiswa yang datanya ingin diubah")
	fmt.Scan(&nama)
	index = SearchNama(A, n, nama)
	for index == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&nama)
		index = SearchNama(A, n, nama)
	}
	fmt.Println("Masukkan nama mata kuliah yang nilainya ingin diubah")
	fmt.Scan(&matkul)
	i = searchMatkul(A, j, matkul, index)
	for i == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&nama)
		i = searchMatkul(A, j, matkul, index)
	}
	fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
	fmt.Scan(&A[index].matkul[i].quiz, &A[index].matkul[i].UTS, &A[index].matkul[i].UAS)
}

func searchMatkul(A tabMhs, j int, matkul string, index int) {
	var ketemu int
	ketemu = -1
	for i := 0; i < *j; i++{
		if matkul == A[index].matkul[i].matakuliah{
			ketemu = i
		}
	}
	return ketemu
}

func menambahData(A *tabMhs, n *int) {
	var pilihan int
	menuTambah()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahData(&A, &n)
	case 2:
		
	case 3:
		mainMenu(&A, &n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
	}
}

func menuTambah() {
	fmt.Println("======  Tambah Data Mahasiswa ======")
	fmt.Println("1. Nama, NIM, Kelas, Prodi          ")
	fmt.Println("2. Mata Kuliah Mahasiswa dan nilai  ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func tambahData(A *tabMhs, n *int) {
	var jumlah int
	fmt.Println("Masukkan jumlah mahasiswa yang ingin ditambah")
	fmt.Scan(&jumlah)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	jumlah = jumlah + *n
	for i := *n; i < jumlah; i++{
		fmt.Scan(&A[i].nama, &A[i].nim, &A[i].kelas, &A[i].prodi)
	}
	*n = *n + jumlah
}

func tambahNilai(A *tabMhs, n *int, j *int) {
	var index, i int
	var nama string
	var matkul string

	fmt.Println("Masukkan nama mahasiswa yang ingin ditambah nilainya")
	fmt.Scan(&nama)	
	index = SearchNama(A, n, nama)
	for index == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&nama)
		index = SearchNama(A, n, nama)
	}
	*j++
	fmt.Println("Masukkan nama mata kuliah")
	fmt.Scan(&matkul)
	for matkul != "0"{
		A[index].matkul[j].matakuliah = matkul
		fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
		fmt.Scan(&A[index].matkul[j].quiz, &A[index].matkul[j].UTS, &A[index].matkul[j].UAS)
		*j++
		fmt.Println("Masukkan nama mata kuliah lagi atau ketik "0" jika sudah selesai")
		fmt.Scan(&matkul)
	}

}
