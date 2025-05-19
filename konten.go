package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

// prosedur menu
func menu(pilihan int, data kontenArray, n int) {

	fmt.Println("=== Aplikasi Manajemen Konten ===")
	fmt.Println("1. Ide Konten")
	fmt.Println("2. Tampilkan Semua Ide Konten")
	fmt.Println("3. Cari Konten")
	fmt.Println("4. Urutkan daftar konten")
	fmt.Println("5. Urutkan berdasarkan Engagement")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu: ")

	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		// menambahkan konten
		kustomisasiDataKonten(pilihan, data, n)
	case 2:
		// menampilkan semua konten sequensial dan binary search
		tampilkanSemuaKonten(data, n)
	case 3:
		// mencari konten berdasarkan judul selection sort dan insertion sort
		cariKonten(data, n)
	case 4:
		// mengurutkan daftar konten
		urutkanDaftarKonten(data, n)
	case 5:
		// mengurutkan konten berdasarkan engagement
		//urutkanBerdasarkanEngagement()
	case 6:
		fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
		return
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

// prosedur tambahKonten
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
}

// prosedur createKonten
func createKonten(data *kontenArray, n *int) {
	var i int
	var judul, platform, kategori, tanggal, jam string
	var reader = bufio.NewReader(os.Stdin)

	if *n >= maxContent {
		fmt.Println("Data sudah penuh.")
		kustomisasiDataKonten(0, *data, *n)
	}

	fmt.Print("Masukkan Judul Konten: ")
	judul, _ = reader.ReadString('\n')
	data[*n].Judul = strings.TrimSpace(judul)

	fmt.Print("Masukkan Platform: ")
	platform, _ = reader.ReadString('\n')
	data[*n].Platform = strings.TrimSpace(platform)

	fmt.Print("Masukkan Kategori: ")
	kategori, _ = reader.ReadString('\n')
	data[*n].Kategori = strings.TrimSpace(kategori)

	fmt.Print("Masukkan Tanggal: ")
	tanggal, _ = reader.ReadString('\n')
	data[*n].Tanggal = strings.TrimSpace(tanggal)

	fmt.Print("Masukkan Jam: ")
	jam, _ = reader.ReadString('\n')
	data[*n].Jam = strings.TrimSpace(jam)

	fmt.Print("Masukkan Jumlah Like: ")
	fmt.Scan(&data[*n].Like)

	fmt.Print("Masukkan Jumlah Komentar: ")
	fmt.Scan(&data[*n].Komentar)

	fmt.Print("Masukkan Jumlah Share: ")
	fmt.Scan(&data[*n].Share)

	*n++
	fmt.Println("Konten berhasil ditambahkan.")
	fmt.Println("=== Konten yang telah ditambahkan ===")
	for i = 0; i < *n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	kustomisasiDataKonten(0, *data, *n)
}

// prosedur updateKonten
func updateKonten(data *kontenArray, n int) {
	var judul, konfirmasi string
	var i, pilihan int

	fmt.Println("\n=== Daftar Konten Tersedia ===")
	fmt.Printf("%-4s %-16s %-10s %-10s %-10s %-6s %-6s %-6s\n",
		"No", "Judul", "Platform", "Kategori", "Tanggal & Jam", "Like", "Komentar", "Share")
	fmt.Println(strings.Repeat("-", 80))

	for i = 0; i < n; i++ {
		tanggalJam := data[i].Tanggal + " " + data[i].Jam
		fmt.Printf("%-4d %-16s %-10s %-10s %-10s %-6d %-6d %-6d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			tanggalJam, data[i].Like, data[i].Komentar, data[i].Share)
	}

	fmt.Print("Masukkan judul konten yang ingin diubah: ")
	fmt.Scan(&judul)

	//mencari konten berdasarkan judul
	for i = 0; i < n; i++ {
		if strings.EqualFold(data[i].Judul, judul) {
			fmt.Println("Konten ditemukan.")
			fmt.Printf("Judul: %s\n", data[i].Judul)
			fmt.Printf("Platform: %s\n", data[i].Platform)
			fmt.Printf("Kategori: %s\n", data[i].Kategori)
			fmt.Printf("Tanggal: %s\n", data[i].Tanggal)
			fmt.Printf("Jam: %s\n", data[i].Jam)
			fmt.Printf("Like: %d\n", data[i].Like)
			fmt.Printf("Komentar: %d\n", data[i].Komentar)
			fmt.Printf("Share: %d\n", data[i].Share)

			// update konten
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
				fmt.Print("Pilih menu: ")
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

			} else if konfirmasi == "N" || konfirmasi == "n" {
				kustomisasiDataKonten(0, *data, n)
			} else {
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
				kustomisasiDataKonten(0, *data, n)
			}
			return
		}
	}

	menu(0, *data, n)
}

func editJudulKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var judul, konfirmasi, inputKonfirmasi string

	fmt.Print("Masukkan judul baru: ")
	inputJudul, _ := reader.ReadString('\n')
	judul = strings.TrimSpace(inputJudul)

	fmt.Printf("Apakah anda yakin ingin mengubah judul dari %s menjadi %s? (y/n): ", data[i].Judul, judul)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Judul = judul
		fmt.Println("Judul berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

func editPlatformKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var platform, konfirmasi, inputPlatform, inputKonfirmasi string

	fmt.Print("Masukkan platform baru: ")
	inputPlatform, _ = reader.ReadString('\n')
	platform = strings.TrimSpace(inputPlatform)

	fmt.Printf("Apakah anda yakin ingin mengubah platform dari %s menjadi %s? (y/n): ", data[i].Platform, platform)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Platform = platform
		fmt.Println("Platform berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

func editKategoriKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var kategori, konfirmasi, inputKategori, inputKonfirmasi string

	fmt.Print("Masukkan kategori baru: ")
	inputKategori, _ = reader.ReadString('\n')
	kategori = strings.TrimSpace(inputKategori)

	fmt.Printf("Apakah anda yakin ingin mengubah kategori dari %s menjadi %s? (y/n): ", data[i].Kategori, kategori)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Kategori = kategori
		fmt.Println("Kategori berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

func editTanggalKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var tanggal, konfirmasi, inputTanggal, inputKonfirmasi string

	fmt.Print("Masukkan tanggal baru: ")
	inputTanggal, _ = reader.ReadString('\n')
	tanggal = strings.TrimSpace(inputTanggal)

	fmt.Printf("Apakah anda yakin ingin mengubah tanggal dari %s menjadi %s? (y/n): ", data[i].Tanggal, tanggal)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Tanggal = tanggal
		fmt.Println("Tanggal berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

func editJamKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var jam, konfirmasi, inputJam, inputKonfirmasi string

	fmt.Print("Masukkan jam baru: ")
	inputJam, _ = reader.ReadString('\n')
	jam = strings.TrimSpace(inputJam)

	fmt.Printf("Apakah anda yakin ingin mengubah jam dari %s menjadi %s? (y/n): ", data[i].Jam, jam)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Jam = jam
		fmt.Println("Jam berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

func editLikeKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var like int
	var konfirmasi, inputKonfirmasi string

	fmt.Print("Masukkan jumlah like baru: ")
	fmt.Scan(&like)

	fmt.Printf("Apakah anda yakin ingin mengubah jumlah like dari %d menjadi %d? (y/n): ", data[i].Like, like)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Like = like
		fmt.Println("Jumlah like berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

func editKomentarKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var komentar int
	var konfirmasi, inputKonfirmasi string

	fmt.Print("Masukkan jumlah komentar baru: ")
	fmt.Scan(&komentar)

	fmt.Printf("Apakah anda yakin ingin mengubah jumlah komentar dari %d menjadi %d? (y/n): ", data[i].Komentar, komentar)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Komentar = komentar
		fmt.Println("Jumlah komentar berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

func editShareKonten(data *kontenArray, n int, i int) {
	var reader = bufio.NewReader(os.Stdin)
	var share int
	var konfirmasi, inputKonfirmasi string

	fmt.Print("Masukkan jumlah share baru: ")
	fmt.Scan(&share)

	fmt.Printf("Apakah anda yakin ingin mengubah jumlah share dari %d menjadi %d? (y/n): ", data[i].Share, share)
	inputKonfirmasi, _ = reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(inputKonfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		data[i].Share = share
		fmt.Println("Jumlah share berhasil diubah.")
		kustomisasiDataKonten(0, *data, n)
	} else {
		fmt.Println("Perubahan dibatalkan.")
		updateKonten(data, n)
	}
}

// prosedur deleteKonten
func deleteKonten(pilihan int, data *kontenArray, n *int) {
	var i int
	var judul string

	fmt.Println("=== Menu Hapus Konten ===")
	fmt.Println("1. Hapus Konten")
	fmt.Println("2. Kembali ke Kustomisasi Data")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		if *n == 0 {
			fmt.Println("Tidak ada konten yang tersedia.")
			return
		}

		fmt.Println("Daftar Konten:")
		for i = 0; i < *n; i++ {
			fmt.Printf("%d. %s\n", i+1, data[i].Judul)
		}

		fmt.Print("Masukkan judul konten yang ingin dihapus: ")
		fmt.Scan(&judul)

		for i = 0; i < *n; i++ {
			if strings.EqualFold(data[i].Judul, judul) {
				for j := i; j < *n-1; j++ {
					data[j] = data[j+1]
				}
				*n--
				fmt.Println("Konten berhasil dihapus.")
				return
			}
		}
		fmt.Println("Konten tidak ditemukan.")

	} else if pilihan == 2 {
		fmt.Println("Kembali ke Kustomisasi Data")
		kustomisasiDataKonten(0, *data, *n)
		return
	} else {
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		deleteKonten(pilihan, data, n)
		return
	}

	fmt.Println("=== Daftar Konten Saat Ini ===")
	for i = 0; i < *n; i++ {
		fmt.Printf("[%d] Judul: %s | Platform: %s | Kategori: %s | Tgl: %s %s | Like: %d | Komentar: %d | Share: %d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			data[i].Tanggal, data[i].Jam, data[i].Like, data[i].Komentar, data[i].Share)
	}
	kustomisasiDataKonten(0, *data, *n)
}

// prosedur tampilkanSemuaKonten
func tampilkanSemuaKonten(data kontenArray, n int) {
	var i int

	if n == 0 {
		fmt.Println("Tidak ada konten yang tersedia.")
		menu(0, data, n)
	}

	fmt.Println("\n=== Daftar Konten ===")
	fmt.Printf("%-4s %-16s %-10s %-10s %-10s %-6s %-6s %-6s\n",
		"No", "Judul", "Platform", "Kategori", "Tanggal & Jam", "Like", "Komentar", "Share")
	fmt.Println(strings.Repeat("-", 80))

	for i = 0; i < n; i++ {
		tanggalJam := data[i].Tanggal + " " + data[i].Jam
		fmt.Printf("%d. %-16s %-10s %-10s %-10s %-6d %-6d %-6d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			tanggalJam, data[i].Like, data[i].Komentar, data[i].Share)
	}

	menu(0, data, n)
}

// prosedur cariKonten
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

	fmt.Println("=== Menu Cari Konten ===")
	fmt.Println("1. Cari Konten (Judul) (sequential search)")
	fmt.Println("2. Cari Konten (Judul) (binary search)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		sequentialSearchJudul(data, n)
	case 2:
		binarySearchJudul(data, n)
	case 3:
		menu(0, data, n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		cariKonten(data, n)
	}
}

func cariKontenKategori(data kontenArray, n int) {
	var pilihan int

	fmt.Println("=== Menu Cari Kategori ===")
	fmt.Println("1. Cari Konten (Kategori) (sequential search)")
	fmt.Println("2. Cari Konten (Kategori) (binary search)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		sequentialSearchKategori(data, n)
	case 2:
		binarySearchKategori(data, n)
	case 3:
		menu(0, data, n)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		cariKonten(data, n)
	}
}

// prosedur sequentialSearchJudul
func sequentialSearchJudul(data kontenArray, n int) string {
	var i, found int
	var judul string

	fmt.Print("Masukkan judul konten yang ingin dicari: ")
	fmt.Scan(&judul)

	found = -1
	i = 0
	for i < n && found == -1 {
		if data[i].Judul == judul {
			found = i
		}
		i++
	}
	if found != -1 {
		return fmt.Sprintf("Konten ditemukan: %s", data[found].Judul)
	} else {
		return "Konten tidak ditemukan."
	}
}

// prosedur binarySearchJudul
func binarySearchJudul(data kontenArray, n int) string {
	var left, right, middle, found int
	var judul string

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
		}
	}
	if found != -1 {
		return fmt.Sprintf("Konten ditemukan: %s", data[found].Judul)
	} else {
		return "Konten tidak ditemukan."
	}
}

// prosedur sequentialSearchKategori
func sequentialSearchKategori(data kontenArray, n int) string {
	var i, found int
	var kategori string

	fmt.Print("Masukkan kategori konten yang ingin dicari: ")
	fmt.Scan(&kategori)

	found = -1
	i = 0
	for i < n && found == -1 {
		if data[i].Kategori == kategori {
			found = i
		}
		i++
	}
	if found != -1 {
		return fmt.Sprintf("Konten ditemukan: %s", data[found].Kategori)
	} else {
		return "Konten tidak ditemukan."
	}
}

// prosedur binarySearchKategori
func binarySearchKategori(data kontenArray, n int) string {
	var left, right, middle, found int
	var kategori string

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
		}
	}
	if found != -1 {
		return fmt.Sprintf("Konten ditemukan: %s", data[found].Kategori)
	} else {
		return "Konten tidak ditemukan."
	}
}

// prosedur urutkanDaftarKonten
func urutkanDaftarKonten(data kontenArray, n int) {
	var pilihan int

	fmt.Println("=== Menu Urutkan Daftar Konten ===")
	fmt.Println("1. Urutkan berdasarkan tanggal posting")
	fmt.Println("2. Urutkan berdasarkan engagement")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilih menu: ")

	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Println("=== Urutkan berdasarkan tanggal posting ===")
		fmt.Println("1. Urutkan berdasarkan tanggal posting (Selection Sort)")
		fmt.Println("2. Urutkan berdasarkan tanggal posting (Insertion Sort)")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			// selection sort
			selectionSortTanggal(&data, n)
		case 2:
			// insertion sort
			insertionSortTanggal(data, n)
		case 3:
			urutkanDaftarKonten(data, n)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			urutkanDaftarKonten(data, n)
		}

	} else if pilihan == 2 {
		fmt.Println("=== Urutkan berdasarkan engagement ===")
		fmt.Println("1. Urutkan berdasarkan engagement (Selection Sort)")
		fmt.Println("2. Urutkan berdasarkan engagement (Insertion Sort)")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			// selection sort
			selectionSortEngagement(data, n)
		case 2:
			// insertion sort
			insertionSortEngagement(data, n)
		case 3:
			urutkanDaftarKonten(data, n)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			urutkanDaftarKonten(data, n)
		}

	} else if pilihan == 3 {
		menu(0, data, n)
	} else {
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		urutkanDaftarKonten(data, n)
	}
}

// prosedur selectionSortTanggal
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
	fmt.Printf("%-4s %-16s %-10s %-10s %-10s %-6s %-6s %-6s\n",
		"No", "Judul", "Platform", "Kategori", "Tanggal & Jam", "Like", "Komentar", "Share")
	fmt.Println(strings.Repeat("-", 80))

	for i = 0; i < n; i++ {
		tanggalJam := data[i].Tanggal + " " + data[i].Jam
		fmt.Printf("%d. %-16s %-10s %-10s %-10s %-6d %-6d %-6d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			tanggalJam, data[i].Like, data[i].Komentar, data[i].Share)
	}
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
	fmt.Printf("%-4s %-16s %-10s %-10s %-10s %-6s %-6s %-6s\n",
		"No", "Judul", "Platform", "Kategori", "Tanggal & Jam", "Like", "Komentar", "Share")
	fmt.Println(strings.Repeat("-", 80))

	for i = 0; i < n; i++ {
		tanggalJam := data[i].Tanggal + " " + data[i].Jam
		fmt.Printf("%d. %-16s %-10s %-10s %-10s %-6d %-6d %-6d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			tanggalJam, data[i].Like, data[i].Komentar, data[i].Share)
	}
}

// prosedur selectionSortEngagement
func selectionSortEngagement(data kontenArray, n int) {
	//dari jumlah like
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
	fmt.Printf("%-4s %-16s %-10s %-10s %-10s %-6s %-6s %-6s\n",
		"No", "Judul", "Platform", "Kategori", "Tanggal & Jam", "Like", "Komentar", "Share")
	fmt.Println(strings.Repeat("-", 80))
	for i = 0; i < n; i++ {
		tanggalJam := data[i].Tanggal + " " + data[i].Jam
		fmt.Printf("%d. %-16s %-10s %-10s %-10s %-6d %-6d %-6d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			tanggalJam, data[i].Like, data[i].Komentar, data[i].Share)
	}
}

// prosedur insertionSortEngagement
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
	fmt.Printf("%-4s %-16s %-10s %-10s %-10s %-6s %-6s %-6s\n",
		"No", "Judul", "Platform", "Kategori", "Tanggal & Jam", "Like", "Komentar", "Share")
	fmt.Println(strings.Repeat("-", 80))
	for i = 0; i < n; i++ {
		tanggalJam := data[i].Tanggal + " " + data[i].Jam
		fmt.Printf("%d. %-16s %-10s %-10s %-10s %-6d %-6d %-6d\n",
			i+1, data[i].Judul, data[i].Platform, data[i].Kategori,
			tanggalJam, data[i].Like, data[i].Komentar, data[i].Share)
	}
}
