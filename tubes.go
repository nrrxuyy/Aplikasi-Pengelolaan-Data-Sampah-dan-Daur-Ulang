package main

import (
	"fmt"
	"math"
	"strings"
)

type Sampah struct {
	Tanggal   string
	Organik   float64 // dalam Kg
	Anorganik float64 // dalam Kg
}

const nmax = 30

var dataSampah [nmax]Sampah
var jumData int

func main() {
	var i int
	var bulan, status string
	var totalOrganik, totalAnorganik float64
	var persenOrganik, persenAnorganik, persenTotal float64
	var username string

	fmt.Print("Halo, siapa namamu? ")
	fmt.Scan(&username)
	fmt.Printf("Halo %s, kamu mau merekap sampahmu pada bulan apa? ", username)
	fmt.Scan(&bulan)
	fmt.Println()

	// Input data sampah harian sampai user ketik "stop"
	for {
		fmt.Print("Masukkan tanggal, kalau mau berhenti ketik 'stop': ")
		fmt.Scan(&dataSampah[i].Tanggal)
		if strings.ToLower(dataSampah[i].Tanggal) == "stop" {
			break
		}

		fmt.Print("Masukkan jumlah sampah organik (Kg): ")
		fmt.Scan(&dataSampah[i].Organik)

		fmt.Print("Masukkan jumlah sampah anorganik (Kg): ")
		fmt.Scan(&dataSampah[i].Anorganik)

		totalOrganik += dataSampah[i].Organik
		totalAnorganik += dataSampah[i].Anorganik

		i++
		if i >= nmax {
			fmt.Println("Data sampah sudah penuh, tidak bisa input lagi.")
			break
		}
		fmt.Println()
	}

	jumData = i

	fmt.Print("Apakah kamu mendaur ulang sampah bulan ini? (ya/tidak): ")
	fmt.Scan(&status)
	fmt.Println()

	// Hitung persentase daur ulang jika jawab ya
	switch strings.ToLower(status) {
	case "ya", "iya", "udah":
		persenOrganik, persenAnorganik, persenTotal = daurUlang(totalOrganik, totalAnorganik)
	default:
		persenOrganik, persenAnorganik, persenTotal = 0, 0, 0
	}

	// Pembulatan hasil persentase ke 1 desimal
	persenOrganik = math.Round(persenOrganik*10) / 10
	persenAnorganik = math.Round(persenAnorganik*10) / 10
	persenTotal = math.Round(persenTotal*10) / 10

	// Menu interaktif
	for {
		fmt.Println("\n--- MENU LANJUTAN ---")
		fmt.Println("1. Edit Data")
		fmt.Println("2. Hapus Data")
		fmt.Println("3. Cari Data")
		fmt.Println("4. Urutkan Data")
		fmt.Println("5. Ringkasan Singkat")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			editData()
		case 2:
			hapusData()
		case 3:
			cariData()
		case 4:
			urutkanData()
		case 5:
			tampilData(totalOrganik, totalAnorganik, persenOrganik, persenAnorganik, persenTotal, status)
		case 0:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan kamu ga ada di menu.")
		}
	}
}

// Fungsi untuk menghitung persentase daur ulang
func daurUlang(organik, anorganik float64) (float64, float64, float64) {
	var daurOrganik, daurAnorganik float64

	fmt.Print("Berapa jumlah sampah organik yang sudah kamu daur ulang (Kg)? ")
	fmt.Scan(&daurOrganik)
	fmt.Println()

	fmt.Print("Berapa jumlah sampah anorganik yang sudah kamu daur ulang (Kg)? ")
	fmt.Scan(&daurAnorganik)
	fmt.Println()

	totalSampah := organik + anorganik
	totalDaur := daurOrganik + daurAnorganik

	var persenOrganik, persenAnorganik, persenTotal float64

	if organik > 0 {
		persenOrganik = (daurOrganik / organik) * 100
	}

	if anorganik > 0 {
		persenAnorganik = (daurAnorganik / anorganik) * 100
	}

	if totalSampah > 0 {
		persenTotal = (totalDaur / totalSampah) * 100
	}

	return persenOrganik, persenAnorganik, persenTotal
}

// Fungsi menampilkan ringkasan data dan persentase daur ulang
func tampilData(organik, anorganik, po, pa, pt float64, status string) {
	fmt.Printf("Selama %d hari kamu sudah mengeluarkan sampah sebanyak: %.1f Kg\n", jumData, organik+anorganik)

	switch strings.ToLower(status) {
	case "ya", "iya", "udah":
		fmt.Printf("Kamu sudah mendaur ulang sebanyak %.1f%% dari keseluruhan sampah yang kamu keluarkan,\n", pt)
		fmt.Printf("Untuk sampah organik kamu telah mendaur ulang %.1f%% dan untuk sampah anorganik %.1f%%\n", po, pa)

		if pt >= 50 {
			fmt.Println("Wow luar biasa, kamu harus tetap konsisten ya demi bumi yang lebih baik!")
		} else if pt >= 9 {
			fmt.Println("Ayo kamu pasti bisa lebih dari ini, mari bersama mewujudkan lingkungan hidup yang lebih baik.")
		} else if pt >= 1 {
			fmt.Println("Kamu harus lebih baik dari ini, untuk menjadikan bumi menjadi tempat yang lebih baik.")
		} else {
			fmt.Println("Kamu engga mendaur ulang sama sekali.")
		}
	default:
		fmt.Println("Kamu engga mendaur ulang sama sekali.")
	}
}

// Fungsi edit data sampah berdasarkan tanggal
func editData() {
	var tanggal string
	fmt.Print("Masukkan tanggal yang ingin diedit: ")
	fmt.Scan(&tanggal)

	for i := 0; i < jumData; i++ {
		if dataSampah[i].Tanggal == tanggal {
			fmt.Print("Masukkan jumlah sampah organik (Kg): ")
			fmt.Scan(&dataSampah[i].Organik)
			fmt.Print("Masukkan jumlah sampah anorganik (Kg): ")
			fmt.Scan(&dataSampah[i].Anorganik)
			fmt.Println("Data sudah diedit.")
			return
		}
	}

	fmt.Println("Tanggal tidak ditemukan.")
}

// Fungsi hapus data sampah berdasarkan tanggal
func hapusData() {
	var tanggal string
	fmt.Print("Masukkan tanggal yang ingin dihapus: ")
	fmt.Scan(&tanggal)

	for i := 0; i < jumData; i++ {
		if dataSampah[i].Tanggal == tanggal {
			// Geser data setelah i ke kiri
			for j := i; j < jumData-1; j++ {
				dataSampah[j] = dataSampah[j+1]
			}
			jumData--
			fmt.Println("Data sudah dihapus.")
			return
		}
	}

	fmt.Println("Tanggal tidak ditemukan.")
}

// Fungsi cari data sampah berdasarkan tanggal
func cariData() {
	var tanggal string
	fmt.Print("Masukkan tanggal yang ingin dicari: ")
	fmt.Scan(&tanggal)

	for i := 0; i < jumData; i++ {
		if dataSampah[i].Tanggal == tanggal {
			fmt.Printf("Tanggal: %s, Organik: %.2f Kg, Anorganik: %.2f Kg\n", dataSampah[i].Tanggal, dataSampah[i].Organik, dataSampah[i].Anorganik)
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

// Fungsi urutkan data sampah berdasarkan tanggal atau total sampah
func urutkanData() {
	var pilihan int

	fmt.Println("Pilih kriteria pengurutan:")
	fmt.Println("1. Berdasarkan Tanggal")
	fmt.Println("2. Berdasarkan Jumlah Total Sampah")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	if pilihan != 1 && pilihan != 2 {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	// Insertion sort
	for pass := 1; pass < jumData; pass++ {
		temp := dataSampah[pass]
		i := pass - 1

		for i >= 0 {
			var shouldSwap bool
			if pilihan == 1 {
				shouldSwap = temp.Tanggal < dataSampah[i].Tanggal
			} else {
				totalTemp := temp.Organik + temp.Anorganik
				totalI := dataSampah[i].Organik + dataSampah[i].Anorganik
				shouldSwap = totalTemp < totalI
			}

			if !shouldSwap {
				break
			}

			dataSampah[i+1] = dataSampah[i]
			i--
		}
		dataSampah[i+1] = temp
	}

	var tampilHasil string
	fmt.Print("Tampilkan hasil pengurutan? (ya/tidak): ")
	fmt.Scan(&tampilHasil)

	if strings.ToLower(tampilHasil) == "ya" {
		if pilihan == 1 {
			fmt.Println("\nHasil pengurutan berdasarkan tanggal:")
		} else {
			fmt.Println("\nHasil pengurutan berdasarkan total sampah:")
		}

		fmt.Println("+------------+--------------+----------------+--------------+")
		fmt.Println("|    Tanggal | Organik (Kg) | Anorganik (Kg) |   Total (Kg) |")
		fmt.Println("+------------+--------------+----------------+--------------+")
		for i := 0; i < jumData; i++ {
			total := dataSampah[i].Organik + dataSampah[i].Anorganik
			fmt.Printf("| %10s | %12.2f | %14.2f | %12.2f |\n", dataSampah[i].Tanggal, dataSampah[i].Organik, dataSampah[i].Anorganik, total)
		}
		fmt.Println("+------------+--------------+----------------+--------------+")
	} else {
		fmt.Println("Hasil pengurutan tidak ditampilkan.")
	}
}
