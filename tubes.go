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
	j int
	prodi string
	matkul [NMAX]tabNilai
}

type tabMhs [NMAX]mahasiswa

func main() {
	var A tabMhs
	var n int

	inputMenu(&A, &n)
}

func mainMenu() {
	fmt.Println("============ Main Menu ============")
	fmt.Println("1. Mengisi data mahasiswa          ")
	fmt.Println("2. Mengubah data mahasiswa         ")
	fmt.Println("3. Menambah data mahasiswa         ")
	fmt.Println("4. Menghapus data mahasiswa        ")
	fmt.Println("5. Menampilkan data mahasiswa      ")
	fmt.Println("6. Exit                            ")
	fmt.Println("===================================")
}

func inputMenu(A *tabMhs, n *int) {
	mainMenu()
	var menu int
	fmt.Scan(&menu)
	switch menu {
	case 1:
		mengisiData(*&A, *&n)
	case 2:
		mengubahData(*&A, *&n)
	case 3:
		menambahData(*&A, *&n)
	case 4:
		menghapusData(*&A, *&n)
	case 5:
		menampilkanData(*&A, *&n)
	case 6:
		fmt.Println("Terima kasih. Program berhenti.")
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		inputMenu(*&A, *&n)
	}
}

func mengisiData(A *tabMhs, n *int) {
	var pilihan int
	menuIsi()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		dataMhs(*&A, *&n)
	case 2:
		nilaiMhs(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		mengisiData(*&A, *&n)
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
	fmt.Scan(n)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	for i := 0; i < *n; i++{
		fmt.Scan(&A[i].nama, &A[i].nim, &A[i].kelas, &A[i].prodi)
	}
	fmt.Println("Data berhasil ditambahkan")
	mengisiData(*&A, *&n)
}



func nilaiMhs(A *tabMhs, n *int) {
	var index int
	var idmhs string
	var matkul string

	fmt.Println("Masukkan ID mahasiswa")
	fmt.Scan(&idmhs)	
	index = SearchId(*A, *n, idmhs)
	for index == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&idmhs)
		index = SearchId(*A, *n, idmhs)
	}
	A[index].j = 0
	fmt.Println("Masukkan nama mata kuliah")
	fmt.Scan(&matkul)
	for matkul != "0"{
		A[index].matkul[A[index].j].matakuliah = matkul
		fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
		fmt.Scan(&A[index].matkul[A[index].j].quiz, &A[index].matkul[A[index].j].UTS, &A[index].matkul[A[index].j].UAS)
		A[index].matkul[A[index].j].total = A[index].matkul[A[index].j].quiz+A[index].matkul[A[index].j].UTS+A[index].matkul[A[index].j].UAS
		var rata int
		rata = A[index].matkul[A[index].j].total / 3
		if rata > 80{
			A[index].matkul[A[index].j].grade = "A"
		}else if rata <= 80 && rata > 70{
			A[index].matkul[A[index].j].grade = "AB"
		}else if rata <= 70 && rata > 65{
			A[index].matkul[A[index].j].grade = "B"
		}else if rata <= 65 && rata > 60{
			A[index].matkul[A[index].j].grade = "BC"
		}else if rata <= 60 && rata > 50{
			A[index].matkul[A[index].j].grade = "C"
		}else if rata <= 50 && rata > 40{
			A[index].matkul[A[index].j].grade = "D"
		}else{
			A[index].matkul[A[index].j].grade = "E"
		}
		A[index].j = A[index].j + 1
		fmt.Println("Masukkan nama mata kuliah atau ketik 0 jika sudah selesai")
		fmt.Scan(&matkul)
	}
	fmt.Println("Data berhasil ditambahkan")
	mengisiData(*&A, *&n)
}

func SearchId(A tabMhs, n int, x string) int {
	var i int
	var ketemu int

	ketemu = -1
	for i = 0; i < n; i++{
		if x == A[i].nim {
			ketemu = i
		}
	}
	return ketemu
}

func cetakDataMhs(A tabMhs, n int) {
	fmt.Println("Data mahasiswa: ")
	for i := 0; i < n; i++{
		fmt.Println(A[i].nama, A[i].nim, A[i].kelas, A[i].prodi)
	}
	menampilkanData(&A, &n)
}

func cetakNilaiMhs(A tabMhs, n int) {
	var x string
	var idx int
	var j int
	fmt.Println("Masukkan id mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(A, n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(A, n, x)
	}
	j = A[idx].j
	fmt.Println("Data nilai mahasiswa", A[idx].nama)
	for i:= 0; i<j; i++{
		fmt.Println(A[idx].matkul[i].matakuliah, A[idx].matkul[i].quiz, A[idx].matkul[i].UTS, A[idx].matkul[i].UAS, A[idx].matkul[i].total, A[idx].matkul[i].grade)
	}
	menampilkanData(&A, &n)
}

func menampilkanData(A *tabMhs, n *int) {
	var pilihan int
	menuCetak()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		sortsks(*&A, *&n)
		cetakDataMhs(*A, *n)
	case 2:
		cetakNilaiMhs(*A, *n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		menampilkanData(*&A, *&n)
	}
}

func menuCetak() {
	fmt.Println("======= Menampilkan Data Mahasiswa ========")
	fmt.Println("1. Daftar mahasiswa keseluruhan            ")
	fmt.Println("2. Daftar Nilai Seorang Mahasiswa          ")
	fmt.Println("3. Daftar Mahasiswa berdasarkan matakuliah ")
	fmt.Println("4. Kembali                                 ")
	fmt.Println("===========================================")
}

func mengubahData(A *tabMhs, n *int) {
	var pilihan int
	menuUbah()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		ubahData(*&A, *&n)
	case 2:
		ubahNilai(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		mengubahData(*&A, *&n)
	}
}

func menuUbah() {
	fmt.Println("======= Ubah Data Mahasiswa ========")
	fmt.Println("1. Biodata mahasiswa                ")
	fmt.Println("2. Nilai Mahasiswa                  ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func ubahData(A *tabMhs, n *int) {
	var idx int
	var x string
	fmt.Println("Masukkan id mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	fmt.Println(A[idx].nama, A[idx].nim, A[idx].kelas, A[idx].prodi)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	fmt.Scan(&A[idx].nama, &A[idx].nim, &A[idx].kelas, &A[idx].prodi)
	fmt.Println("Data berhasil diedit")
	mengubahData(*&A, *&n)
}

func ubahNilai(A *tabMhs, n *int) {
	var x string
	var idx, i int
	var matkul string
	fmt.Println("Masukkan id mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	fmt.Println("Masukkan nama mata kuliah yang nilainya ingin diubah")
	fmt.Scan(&matkul)
	i = searchMatkul(*A, idx, A[idx].j, matkul)
	for i == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&matkul)
		i = searchMatkul(*A, idx, A[idx].j, matkul)
	}
	fmt.Println(A[idx].matkul[i].matakuliah, A[idx].matkul[i].quiz, A[idx].matkul[i].UTS, A[idx].matkul[i].UAS, A[idx].matkul[i].total, A[idx].matkul[i].grade)
	fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
	fmt.Scan(&A[idx].matkul[i].quiz, &A[idx].matkul[i].UTS, &A[idx].matkul[i].UAS)
	A[idx].matkul[i].total = A[idx].matkul[i].quiz+A[idx].matkul[i].UTS+A[idx].matkul[i].UAS
	var rata int
	rata = A[idx].matkul[i].total / 3
	if rata > 80{
		A[idx].matkul[i].grade = "A"
	}else if rata <= 80 && rata > 70{
		A[idx].matkul[i].grade = "AB"
	}else if rata <= 70 && rata > 65{
		A[idx].matkul[i].grade = "B"
	}else if rata <= 65 && rata > 60{
		A[idx].matkul[i].grade = "BC"
	}else if rata <= 60 && rata > 50{
		A[idx].matkul[i].grade = "C"
	}else if rata <= 50 && rata > 40{
		A[idx].matkul[i].grade = "D"
	}else{
		A[idx].matkul[i].grade = "E"
	}
	fmt.Println(A[idx].matkul[i].matakuliah, A[idx].matkul[i].quiz, A[idx].matkul[i].UTS, A[idx].matkul[i].UAS, A[idx].matkul[i].total, A[idx].matkul[i].grade)
	fmt.Println("Data berhasil diedit")
	mengubahData(*&A, *&n)
}

func searchMatkul(A tabMhs, n int, j int, matkul string) int {
	var ketemu, i int
	ketemu = -1
	for i = 0; i < j; i++{
		if matkul == A[n].matkul[i].matakuliah{
			ketemu = i
		}
	}
	return ketemu
}

func menghapusData(A *tabMhs, n *int) {
	var pilihan int
	menuHapus()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		hapusData(*&A, *&n)
	case 2:
		hapusNilai(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		menghapusData(*&A, *&n)
	}
}

func menuHapus() {
	fmt.Println("======= Hapus Data Mahasiswa =======")
	fmt.Println("1. Biodata mahasiswa dan nilai      ")
	fmt.Println("2. nilai mata kuliah                ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func hapusData(A *tabMhs, n *int) {
	var x string
	var idx int
	fmt.Println("Masukkan ID Mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	for i:= idx; i<*n; i++{
		A[i] = A[i+1]
	}
	*n = *n-1
	fmt.Println("Data berhasil dihapus")
	menghapusData(*&A, *&n)
}

func hapusNilai(A *tabMhs, n *int) {
	var x, matkul string
	var idx, i int
	fmt.Println("Masukkan ID Mahasiswa")
	fmt.Scan(&x)
	idx = SearchId(*A, *n, x)
	for idx == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&x)
		idx = SearchId(*A, *n, x)
	}
	fmt.Println("Masukkan Nama Matakuliah")
	fmt.Scan(&matkul)
	i = searchMatkul(*A, idx, A[idx].j, matkul)
	for i == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&matkul)
		i = searchMatkul(*A, idx, A[idx].j, matkul)
	}
	var j int
	j = A[idx].j
	for a := i; a<j-1; a++{
		A[a].matkul = A[a+1].matkul
	}
	A[idx].j = A[idx].j - 1
	fmt.Println("Data berhasil dihapus")
	menghapusData(*&A, *&n)
}

func menambahData(A *tabMhs, n *int) {
	var pilihan int
	menuTambah()
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahData(*&A, *&n)
	case 2:
		tambahMatkul(*&A, *&n)
	case 3:
		inputMenu(*&A, *&n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		menambahData(*&A, *&n)
	}
}

func menuTambah() {
	fmt.Println("====== Tambah Data Mahasiswa =======")
	fmt.Println("1. Biodata mahasiswa                ")
	fmt.Println("2. Mata kuliah dan nilai            ")
	fmt.Println("3. Kembali                          ")
	fmt.Println("====================================")
}

func tambahData(A *tabMhs, n *int) {
	var x int
	fmt.Println("Masukkan jumlah mahasiswa yang ingin ditambahkan")
	fmt.Scan(&x)
	fmt.Println("Masukkan Nama, NIM, Kelas, Prodi")
	for i := *n; i < *n+x; i++{
		fmt.Scan(&A[i].nama, &A[i].nim, &A[i].kelas, &A[i].prodi)
	}
	*n = *n+x
	fmt.Println("Data berhasil ditambahkan")
	menambahData(*&A, *&n)
}

func tambahMatkul(A *tabMhs, n *int) {
	var index int
	var idmhs string
	var matkul string

	fmt.Println("Masukkan nama mahasiswa")
	fmt.Scan(&idmhs)	
	index = SearchId(*A, *n, idmhs)
	for index == -1{
		fmt.Println("Masukkan tidak valid. Silahkan masukkan kembali.")
		fmt.Scan(&idmhs)
		index = SearchId(*A, *n, idmhs)
	}
	fmt.Println("Masukkan nama mata kuliah")
	fmt.Scan(&matkul)
	var j int
	j = A[index].j
	for matkul != "0"{
		A[index].matkul[j].matakuliah = matkul
		fmt.Println("Masukkan nilai quiz, nilai UTS dan nilai UAS")
		fmt.Scan(&A[index].matkul[j].quiz, &A[index].matkul[j].UTS, &A[index].matkul[j].UAS)
		A[index].matkul[j].total = A[index].matkul[j].quiz+A[index].matkul[j].UTS+A[index].matkul[j].UAS
		var rata int
		rata = A[index].matkul[j].total / 3
		if rata > 80{
			A[index].matkul[j].grade = "A"
		}else if rata <= 80 && rata > 70{
			A[index].matkul[j].grade = "AB"
		}else if rata <= 70 && rata > 65{
			A[index].matkul[j].grade = "B"
		}else if rata <= 65 && rata > 60{
			A[index].matkul[j].grade = "BC"
		}else if rata <= 60 && rata > 50{
			A[index].matkul[j].grade = "C"
		}else if rata <= 50 && rata > 40{
			A[index].matkul[j].grade = "D"
		}else{
			A[index].matkul[j].grade = "E"
		}
		j++
		A[index].j = A[index].j + 1
		fmt.Println("Masukkan nama mata kuliah atau ketik 0 jika sudah selesai")
		fmt.Scan(&matkul)
	}
	fmt.Println("Data berhasil ditambahkan")
	menambahData(*&A, *&n)
}

func sortsks(A *tabMhs, n *int) {
	// descending with selection
	var pass, max, i, a int

	pass = *n-1
	for i = 0; i<pass; i++{
		max = i
		for a = i + 1; a<*n; a++{
			if A[max].j < A[a].j{
				max = a
			}
		}

		A[NMAX-1] = A[i]
		A[i] = A[max]
		A[max] = A[NMAX-1]
	}
	fmt.Println("Sort Berhasil")
}

// func sortNilai(A *tabMhs, n int) {
// 	//algoritma insertion sort secara ascending
// 	var i, j int
// 	var idx int
// 	var nim string
// 	fmt.Scan(&nim)
// 	idx = SearchId(*A, n, nim)
// 	i = 1
// 	for i <= n-1{
// 		j = i
// 		A[idx].matkul[NMAX-1] = A[idx].matkul[j]
// 		for j > 0 && A[idx].matkul[NMAX-1].total < A[idx].matkul[j].total{
// 			A[idx].matkul[j] = A[idx].matkul[j-1]
// 			j = j - 1
// 		}
// 		A[idx].nilai = A[idx].matkul[NMAX-1]
// 		i = i+1
// 	}
// 	fmt.Println("Sort Berhasil")
// }
