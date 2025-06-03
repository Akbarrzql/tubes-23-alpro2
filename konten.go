package main

import "fmt"

const maxContent int = 100

type Konten struct {
	Judul    string
	Platform string
	Kategori string
	Tanggal  string
	Jam      string
	Like     int
	Komentar int
	Share    int
}

type kontenArray [maxContent]Konten

func main() {
	var pilihan int
	var data kontenArray
	var n int

	menu(pilihan, data, n)
}

func menu(pilihan int, data kontenArray, n int) {
	fmt.Println()

	fmt.Println("=== Aplikasi Manajemen Konten ===")
	fmt.Println("1. Ide Konten")
	fmt.Println("2. Tampilkan Semua Ide Konten")
	fmt.Println("3. Cari Konten")
	fmt.Println("4. Urutkan daftar konten")
	fmt.Println("5. Urutkan berdasarkan Engagement")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu: ")

	fmt.Scan(&pilihan)
	fmt.Println()

	switch pilihan {
	case 1:
		kustomisasiDataKonten(pilihan, data, n)
	case 2:
		tampilkanSemuaKonten(data, n)
	case 3:
		cariKonten(data, n)
	case 4:
		urutkanDaftarKonten(data, n)
	case 5:
		urutkanBerdasarkanEngagement(data, n)
	case 6:
		endProgram()
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func endProgram() {
	fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
	fmt.Println("Program selesai.")
	fmt.Println()
}

func kustomisasiDataKonten(pilihan int, data kontenArray, n int) {

	fmt.Println("=== Menu Kustomisasi Data ===")
	fmt.Println("1. Tambah Konten")
	fmt.Println("2. Mengubah Konten")
	fmt.Println("3. Menghapus Konten")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilih menu: ")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		// menambahkan konten
		createKonten(&data, &n)
	case 2:
		// mengubah konten
		updateKonten(&data, n)
	case 3:
		// menghapus konten
		deleteKonten(pilihan, &data, &n)
	case 4:
		// kembali ke menu utama
		menu(pilihan, data, n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
	fmt.Println()
}

func createKonten(data *kontenArray, n *int) {
	var i int

	if *n >= maxContent {
		fmt.Println("Data sudah penuh.")
		kustomisasiDataKonten(0, *data, *n)
	}

	fmt.Print("Masukkan Judul Konten: ")
	fmt.Scanln(&data[*n].Judul)

	fmt.Print("Masukkan Platform: ")
	fmt.Scanln(&data[*n].Platform)

	fmt.Print("Masukkan Kategori: ")
	fmt.Scanln(&data[*n].Kategori)

	fmt.Print("Masukkan Tanggal (format: YYYY-MM-DD, contoh: 2025-02-01): ")
	fmt.Scanln(&data[*n].Tanggal)

	fmt.Print("Masukkan Jam: ")
	fmt.Scanln(&data[*n].Jam)

	fmt.Print("Masukkan Jumlah Like (Masukkan dalam jumlah angka): ")
	fmt.Scanln(&data[*n].Like)

	fmt.Print("Masukkan Jumlah Komentar (Masukkan dalam jumlah angka): ")
	fmt.Scanln(&data[*n].Komentar)

	fmt.Print("Masukkan Jumlah Share (Masukkan dalam jumlah angka): ")
	fmt.Scanln(&data[*n].Share)

	*n++
	fmt.Println("Konten berhasil ditambahkan.")
	fmt.Println("=== Konten yang telah ditambahkan ===")
	for i = 0; i < *n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	fmt.Println()
	kustomisasiDataKonten(0, *data, *n)
}

func updateKonten(data *kontenArray, n int) {
	var judul, konfirmasi string
	var i, pilihan int

	fmt.Println("\n=== Daftar Konten Tersedia ===")

	for i = 0; i < n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}

	fmt.Print("Masukkan judul konten yang ingin diubah: ")
	fmt.Scan(&judul)

	for i = 0; i < n; i++ {
		if data[i].Judul == judul {
			fmt.Println("Konten ditemukan.")
			fmt.Printf("Judul: %s\n", data[i].Judul)
			fmt.Printf("Platform: %s\n", data[i].Platform)
			fmt.Printf("Kategori: %s\n", data[i].Kategori)
			fmt.Printf("Tanggal: %s\n", data[i].Tanggal)
			fmt.Printf("Jam: %s\n", data[i].Jam)
			fmt.Printf("Like: %d\n", data[i].Like)
			fmt.Printf("Komentar: %d\n", data[i].Komentar)
			fmt.Printf("Share: %d\n", data[i].Share)

			fmt.Print("Apakah anda ingin mengubah konten ini? (y/n): ")
			fmt.Scan(&konfirmasi)
			if konfirmasi == "Y" || konfirmasi == "y" {
				fmt.Println("=== Menu Ubah Konten ===")
				fmt.Println("1. Ubah Judul")
				fmt.Println("2. Ubah Platform")
				fmt.Println("3. Ubah Kategori")
				fmt.Println("4. Ubah Tanggal")
				fmt.Println("5. Ubah Jam")
				fmt.Println("6. Ubah Like")
				fmt.Println("7. Ubah Komentar")
				fmt.Println("8. Ubah Share")
				fmt.Println("9. Kembali ke Kustomisasi Data")
				fmt.Print("Pilih menu: (1/2/3/4/5/6/7/8/9) : ")
				fmt.Scan(&pilihan)

				switch pilihan {
				case 1:
					editJudulKonten(data, n, i)
				case 2:
					editPlatformKonten(data, n, i)
				case 3:
					editKategoriKonten(data, n, i)
				case 4:
					editTanggalKonten(data, n, i)
				case 5:
					editJamKonten(data, n, i)
				case 6:
					editLikeKonten(data, n, i)
				case 7:
					editKomentarKonten(data, n, i)
				case 8:
					editShareKonten(data, n, i)
				case 9:
					fmt.Println("Kembali ke Kustomisasi Data")
					kustomisasiDataKonten(0, *data, n)
				default:
					fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
					updateKonten(data, n)
				}
				fmt.Println()
			} else if konfirmasi == "N" || konfirmasi == "n" {
				kustomisasiDataKonten(0, *data, n)
				fmt.Println()
			} else {
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
				kustomisasiDataKonten(0, *data, n)
				fmt.Println()
			}
		}

		if i == n-1 && data[i].Judul != judul {
			fmt.Println("Konten tidak ditemukan.")
			fmt.Println("Silakan coba lagi.")
			kustomisasiDataKonten(0, *data, n)
		}
	}
	fmt.Println()
	menu(0, *data, n)
}

func editJudulKonten(data *kontenArray, n int, i int) {
	var konfirmasi, judul string

	fmt.Print("Masukkan judul baru: ")
	fmt.Scan(&judul)

	fmt.Printf("Apakah anda yakin ingin mengubah judul dari %s menjadi %s? (y/n): ", data[i].Judul, judul)
	fmt.Scan(&konfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Judul = judul
		fmt.Println("Judul berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func editPlatformKonten(data *kontenArray, n int, i int) {
	var inputPlatform, inputKonfirmasi string

	fmt.Print("Masukkan platform baru: ")
	fmt.Scan(&inputPlatform)

	fmt.Printf("Apakah anda yakin ingin mengubah platform dari %s menjadi %s? (y/n): ", data[i].Platform, inputPlatform)
	fmt.Scan(&inputKonfirmasi)

	if inputKonfirmasi == "Y" || inputKonfirmasi == "y" {
		data[i].Platform = inputPlatform
		fmt.Println("Platform berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func editKategoriKonten(data *kontenArray, n int, i int) {
	var inputKategori, inputKonfirmasi string

	fmt.Print("Masukkan kategori baru: ")
	fmt.Scan(&inputKategori)

	fmt.Printf("Apakah anda yakin ingin mengubah kategori dari %s menjadi %s? (y/n): ", data[i].Kategori, inputKategori)
	fmt.Scan(&inputKonfirmasi)

	if inputKonfirmasi == "Y" || inputKonfirmasi == "y" {
		data[i].Kategori = inputKategori
		fmt.Println("Kategori berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func editTanggalKonten(data *kontenArray, n int, i int) {
	var inputTanggal, inputKonfirmasi string

	fmt.Print("Masukkan tanggal baru: ")
	fmt.Scan(&inputTanggal)

	fmt.Printf("Apakah anda yakin ingin mengubah tanggal dari %s menjadi %s? (y/n): ", data[i].Tanggal, inputTanggal)
	fmt.Scan(&inputKonfirmasi)

	if inputKonfirmasi == "Y" || inputKonfirmasi == "y" {
		data[i].Tanggal = inputTanggal
		fmt.Println("Tanggal berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func editJamKonten(data *kontenArray, n int, i int) {
	var inputJam, inputKonfirmasi string

	fmt.Print("Masukkan jam baru: ")
	fmt.Scan(&inputJam)

	fmt.Printf("Apakah anda yakin ingin mengubah jam dari %s menjadi %s? (y/n): ", data[i].Jam, inputJam)
	fmt.Scan(&inputKonfirmasi)

	if inputKonfirmasi == "Y" || inputKonfirmasi == "y" {
		data[i].Jam = inputJam
		fmt.Println("Jam berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func editLikeKonten(data *kontenArray, n int, i int) {
	var like int
	var inputKonfirmasi string

	fmt.Print("Masukkan jumlah like baru: ")
	fmt.Scan(&like)

	fmt.Printf("Apakah anda yakin ingin mengubah jumlah like dari %d menjadi %d? (y/n): ", data[i].Like, like)
	fmt.Scan(&inputKonfirmasi)

	if inputKonfirmasi == "Y" || inputKonfirmasi == "y" {
		data[i].Like = like
		fmt.Println("Jumlah like berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func editKomentarKonten(data *kontenArray, n int, i int) {
	var komentar int
	var inputKonfirmasi string

	fmt.Print("Masukkan jumlah komentar baru: ")
	fmt.Scan(&komentar)

	fmt.Printf("Apakah anda yakin ingin mengubah jumlah komentar dari %d menjadi %d? (y/n): ", data[i].Komentar, komentar)
	fmt.Scan(&inputKonfirmasi)

	if inputKonfirmasi == "Y" || inputKonfirmasi == "y" {
		data[i].Komentar = komentar
		fmt.Println("Jumlah komentar berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func editShareKonten(data *kontenArray, n int, i int) {
	var share int
	var inputKonfirmasi string

	fmt.Print("Masukkan jumlah share baru: ")
	fmt.Scan(&share)

	fmt.Printf("Apakah anda yakin ingin mengubah jumlah share dari %d menjadi %d? (y/n): ", data[i].Share, share)
	fmt.Scan(&inputKonfirmasi)

	if inputKonfirmasi == "Y" || inputKonfirmasi == "y" {
		data[i].Share = share
		fmt.Println("Jumlah share berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
	fmt.Println()
}

func deleteKonten(pilihan int, data *kontenArray, n *int) {
	var i, j int
	var judul string

	fmt.Println("=== Menu Hapus Konten ===")
	fmt.Println("1. Hapus Konten")
	fmt.Println("2. Kembali ke Kustomisasi Data")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		if *n == 0 {
			fmt.Println("Tidak ada konten yang tersedia.")
		}

		fmt.Println("Daftar Konten:")
		for i = 0; i < *n; i++ {
			fmt.Printf("%d. %s\n", i+1, data[i].Judul)
		}

		fmt.Print("Masukkan judul konten yang ingin dihapus: ")
		fmt.Scan(&judul)

		for i = 0; i < *n; i++ {
			if data[i].Judul == judul {
				fmt.Printf("Konten yang dihapus: %s\n", data[i].Judul)
				for j = i; j < *n-1; j++ {
					data[j] = data[j+1]
				}
				*n--
				fmt.Println("Konten berhasil dihapus.")
				fmt.Println()
			} else {
				fmt.Println("Konten tidak ditemukan.")
				fmt.Println()
			}
		}

	} else if pilihan == 2 {
		fmt.Println("Kembali ke Kustomisasi Data")
		kustomisasiDataKonten(0, *data, *n)
		fmt.Println()
	} else {
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		deleteKonten(pilihan, data, n)
		fmt.Println()
	}

	fmt.Println("=== Daftar Konten Saat Ini ===")
	for i = 0; i < *n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	kustomisasiDataKonten(0, *data, *n)
}

func tampilkanSemuaKonten(data kontenArray, n int) {
	var i int

	if n == 0 {
		fmt.Println("Tidak ada konten yang tersedia.")
		menu(0, data, n)
	}

	fmt.Println("\n=== Daftar Konten ===")
	for i = 0; i < n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	fmt.Println()
	menu(0, data, n)
}

func cariKonten(data kontenArray, n int) {
	var pilihan int

	fmt.Println("=== Menu Cari Konten ===")
	fmt.Println("1. Cari Kata Kunci (Judul)")
	fmt.Println("2. Cari Kategori")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilih menu: ")

	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		cariKontenJudul(data, n)
	case 2:
		cariKontenKategori(data, n)
	case 3:
		menu(0, data, n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		cariKonten(data, n)
	}
}

func cariKontenJudul(data kontenArray, n int) {
	var pilihan int
	var hasil string
	var kembaliKeMenuUtama bool

	fmt.Println()

	for !kembaliKeMenuUtama {
		fmt.Println("=== Menu Cari Konten ===")
		fmt.Println("1. Cari Konten (Judul) (sequential search)")
		fmt.Println("2. Cari Konten (Judul) (binary search)")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			hasil = sequentialSearchJudul(data, n)
			fmt.Println(hasil)
		case 2:
			hasil = binarySearchJudul(data, n)
			fmt.Println(hasil)
		case 3:
			kembaliKeMenuUtama = true
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
	menu(0, data, n)
}

func cariKontenKategori(data kontenArray, n int) {
	var pilihan int
	var hasil string
	var kembaliKeMenuUtama bool

	fmt.Println()

	for !kembaliKeMenuUtama {
		fmt.Println("=== Menu Cari Kategori ===")
		fmt.Println("1. Cari Konten (Kategori) (sequential search)")
		fmt.Println("2. Cari Konten (Kategori) (binary search)")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			hasil = sequentialSearchKategori(data, n)
			fmt.Println(hasil)
		case 2:
			hasil = binarySearchKategori(data, n)
			fmt.Println(hasil)
		case 3:
			kembaliKeMenuUtama = true
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}

	menu(0, data, n)
}

func sequentialSearchJudul(data kontenArray, n int) string {
	var i, found int
	var judul string

	fmt.Println()

	fmt.Print("Masukkan judul konten yang ingin dicari: ")
	fmt.Scanln(&judul)

	found = -1
	i = 0
	for i < n && found == -1 {
		if data[i].Judul == judul {
			return "Konten berhasil ditemukan (Sequential Search):" + data[i].Judul
		}
		i++
	}
	return "Konten tidak ditemukan."
}

func binarySearchJudul(data kontenArray, n int) string {
	var left, right, middle, found int
	var judul string

	fmt.Println()

	fmt.Print("Masukkan judul konten yang ingin dicari: ")
	fmt.Scan(&judul)

	left = 1
	right = n
	found = -1
	for left <= right && found == -1 {
		middle = (left + right) / 2
		if judul < data[middle].Judul {
			right = middle - 1
		} else if judul > data[middle].Judul {
			left = middle + 1
		} else {
			found = middle
			return "Konten berhasil ditemukan (Binary Search):" + data[middle].Judul
		}
	}
	fmt.Println()
	return "Konten tidak ditemukan."
}

func sequentialSearchKategori(data kontenArray, n int) string {
	var i, found int
	var kategori string

	fmt.Println()

	fmt.Print("Masukkan kategori konten yang ingin dicari: ")
	fmt.Scan(&kategori)

	found = -1
	i = 0
	for i < n && found == -1 {
		if data[i].Kategori == kategori {
			return "Konten berhasil ditemukan (Sequential Search):" + data[i].Kategori
		}
		i++
	}
	return "Konten tidak ditemukan."
}

func binarySearchKategori(data kontenArray, n int) string {
	var left, right, middle, found int
	var kategori string

	fmt.Println()

	fmt.Print("Masukkan kategori konten yang ingin dicari: ")
	fmt.Scan(&kategori)

	left = 1
	right = n
	found = -1
	for left <= right && found == -1 {
		middle = (left + right) / 2
		if kategori < data[middle].Kategori {
			right = middle - 1
		} else if kategori > data[middle].Kategori {
			left = middle + 1
		} else {
			found = middle
			return "Konten berhasil ditemukan (Binary Search):" + data[middle].Kategori
		}
	}

	return "Konten tidak ditemukan."
}

func urutkanDaftarKonten(data kontenArray, n int) {
	var pilihan, subPilihan int
	var kembaliKeMenuUtama bool

	fmt.Println()

	for !kembaliKeMenuUtama {
		fmt.Println("=== Menu Urutkan Daftar Konten ===")
		fmt.Println("1. Urutkan berdasarkan tanggal posting")
		fmt.Println("2. Urutkan berdasarkan engagement")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println()
			fmt.Println("=== Urutkan berdasarkan tanggal posting ===")
			fmt.Println("1. Urutkan berdasarkan tanggal posting (Selection Sort)")
			fmt.Println("2. Urutkan berdasarkan tanggal posting (Insertion Sort)")
			fmt.Println("3. Kembali")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				selectionSortTanggal(&data, n)
			} else if subPilihan == 2 {
				insertionSortTanggal(data, n)
			} else if subPilihan == 3 {
				kembaliKeMenuUtama = true
			} else {
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		} else if pilihan == 2 {
			fmt.Println()
			fmt.Println("=== Urutkan berdasarkan engagement ===")
			fmt.Println("1. Urutkan berdasarkan engagement (Selection Sort)")
			fmt.Println("2. Urutkan berdasarkan engagement (Insertion Sort)")
			fmt.Println("3. Kembali")
			fmt.Print("Pilih menu: ")
			fmt.Scan(&subPilihan)

			if subPilihan == 1 {
				selectionSortEngagement(data, n)
			} else if subPilihan == 2 {
				insertionSortEngagement(data, n)
			} else if subPilihan == 3 {
				kembaliKeMenuUtama = true
			} else {
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		} else if pilihan == 3 {
			kembaliKeMenuUtama = true
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		fmt.Println()
	}

	menu(0, data, n)
}

func selectionSortTanggal(data *kontenArray, n int) {
	var pass, idx, i int
	var temp Konten

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if data[i].Tanggal < data[idx].Tanggal {
				idx = i
			}
			i = i + 1
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass = pass + 1
	}
	fmt.Println("Konten berhasil diurutkan (Insertion Sort) berdasarkan tanggal posting.")
	fmt.Println("=== Konten yang telah diurutkan ===")
	for i = 0; i < n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	fmt.Println()
}

// prosedur insertionSortTanggal
func insertionSortTanggal(data kontenArray, n int) {
	var pass, i int
	var temp Konten

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.Tanggal < data[i-1].Tanggal {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}

	fmt.Println("Konten berhasil diurutkan (Insertion Sort) berdasarkan tanggal posting.")
	fmt.Println("=== Konten yang telah diurutkan ===")
	for i = 0; i < n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	fmt.Println()
}

func selectionSortEngagement(data kontenArray, n int) {
	var pass, idx, i int
	var temp Konten

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if data[i].Like > data[idx].Like {
				idx = i
			}
			i = i + 1
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass = pass + 1
	}
	fmt.Println("Konten berhasil diurutkan (Selection Sort) berdasarkan engagement (Like).")
	fmt.Println("=== Konten yang telah diurutkan ===")
	for i = 0; i < n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	fmt.Println()
}

func insertionSortEngagement(data kontenArray, n int) {
	var pass, i int
	var temp Konten

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.Like > data[i-1].Like {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}

	fmt.Println("Konten berhasil diurutkan (Selection Sort) berdasarkan engagement (Like).")
	fmt.Println("=== Konten yang telah diurutkan ===")
	for i = 0; i < n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	fmt.Println()
}

func urutkanBerdasarkanEngagement(data kontenArray, n int) {
	var tglAwal, tglAkhir string
	var i, j, maxIdx, m int
	var adaKonten bool
	var tempKonten, k Konten
	var filtered [maxContent]Konten
	var filteredEngagement [maxContent]float64

	fmt.Println()

	fmt.Print("Masukkan tanggal awal (format: YYYY-MM-DD, contoh: 2025-02-01): ")
	fmt.Scan(&tglAwal)
	fmt.Print("Masukkan tanggal akhit (format: YYYY-MM-DD, contoh: 2025-02-01): ")
	fmt.Scan(&tglAkhir)

	if n == 0 {
		fmt.Println()
		fmt.Println("Tidak ada konten yang tersedia.")
		menu(0, data, n)
	}

	m = 0

	for i = 0; i < n; i++ {
		if data[i].Tanggal >= tglAwal && data[i].Tanggal <= tglAkhir {
			filtered[m] = data[i]
			filteredEngagement[m] = float64(data[i].Like+data[i].Komentar+data[i].Share) / 3.0
			m++
			adaKonten = true
		}
	}

	if !adaKonten {
		fmt.Println()
		fmt.Println("Tidak ada konten yang ditemukan dalam rentang tanggal tersebut.")
		menu(0, data, n)
	}
	for i = 0; i < m-1; i++ {
		maxIdx = i
		for j = i + 1; j < m; j++ {
			if filteredEngagement[j] > filteredEngagement[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			tempKonten = filtered[i]
			filtered[i] = filtered[maxIdx]
			filtered[maxIdx] = tempKonten

			tempEng := filteredEngagement[i]
			filteredEngagement[i] = filteredEngagement[maxIdx]
			filteredEngagement[maxIdx] = tempEng
		}
	}

	fmt.Println("\n=== Daftar Konten Berdasarkan Engagement (Tinggi ke Rendah) ===")
	fmt.Printf("%-3s %-16s %-10s %-10s %-15s %-6s %-9s %-6s %-10s\n",
		"No", "Judul", "Platform", "Kategori", "Tanggal & Jam", "Like", "Komentar", "Share", "Engagement")
	for j = 0; j < 100; j++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i = 0; i < m; i++ {
		k = filtered[i]
		fmt.Printf("%-3d %-16s %-10s %-10s %-15s %-6d %-9d %-6d %-10.2f\n",
			i+1, k.Judul, k.Platform, k.Kategori,
			k.Tanggal+" "+k.Jam, k.Like, k.Komentar, k.Share, filteredEngagement[i])
	}

	menu(0, data, n)
}
