package main

import (
	"fmt"
	"math"
	"strings"
)

type sampah struct {
	tanggal int
	Or      float64
	An      float64
}

const nmax = 30

type dataSampah [nmax]sampah

var jumData int

/*
	IS : Program menerima bulan, data harian (tanggal, organik, anorganik), status daur ulang.
	FS: Program menyediakan menu untuk edit, hapus, cari, urut data; menghitung persentase daur ulang; menampilkan ringkasan sampah.
*/

func main() {
	/*
		IS:  Tidak ada input langsung dari fungsi ini, tetapi menerima input tidak langsung dari user melalui fungsi lain.
		FS:  Mengelola alur utama program:
			-   Menerima input bulan dan data sampah harian.
			-   Menghitung total sampah.
			-   Memproses status daur ulang.
			-   Menampilkan menu dan memanggil fungsi lain berdasarkan pilihan pengguna.
	*/
	var data dataSampah
	var i int = 0
	var bulan, status string
	var organik, anorganik, po, pa, pt float64

	fmt.Print("Halo kamu mau merekap sampahmu pada bulan apa : ")
	fmt.Scan(&bulan)
	fmt.Println(" ")

	for {
		fmt.Print("Masukkan tanggal , kalo mau berhenti ketik 0 ya : ")
		fmt.Scan(&data[i].tanggal)

		if data[i].tanggal == 0 {
			break
		}

		fmt.Print("Masukkan jumlah sampah organik (Kg): ")
		fmt.Scan(&data[i].Or)
		fmt.Print("Masukkan jumlah sampah anorganik (Kg): ")
		fmt.Scan(&data[i].An)

		organik += data[i].Or
		anorganik += data[i].An

		i++
		fmt.Println(" ")
	}

	jumData = i

	fmt.Println(" ")
	fmt.Print("Apakah kamu mendaur ulang sampah bulan ini ? ")
	fmt.Scan(&status)
	fmt.Println(" ")

	switch strings.ToLower(status) {
	case "ya", "iya", "udah":
		po, pa, pt = daurUlang(organik, anorganik)
	}

	po = math.Round(po*10) / 10
	pa = math.Round(pa*10) / 10
	pt = math.Round(pt*10) / 10

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
			editData(&data)
		case 2:
			hapusData(&data)
		case 3:
			cariData(data)
		case 4:
			urutkanData(&data)
		case 5:
			tampilData(organik, anorganik, po, pa, pt, status)
		case 0:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan kamu gaada di menu.")
		}
	}
}

func daurUlang(organik, anorganik float64) (float64, float64, float64) {
	/*
		IS:  Menerima total sampah organik dan anorganik. Meminta input jumlah sampah yang didaur ulang.
		FS:  Menghitung persentase sampah organik, anorganik, dan total yang didaur ulang.
			Mengembalikan persentase tersebut.
	*/
	var daurOrganik, daurAnorganik, totalDaur, total, persenOrganik, persenAnorganik, persenTotal float64

	fmt.Print("berapa jumlah sampah organik yang udah kamu daur ulang ? tolong diisi dengan satuan kilogram ya : ")
	fmt.Scan(&daurOrganik)
	fmt.Println(" ")
	fmt.Print("kalo gitu berapa jumlah sampah anorganik yang udah kamu daur ulang ? yang ini juga pakai kilogram ya : ")
	fmt.Scan(&daurAnorganik)
	fmt.Println(" ")

	total = organik + anorganik
	totalDaur = daurOrganik + daurAnorganik

	persenOrganik = (daurOrganik / organik) * 100
	persenAnorganik = (daurAnorganik / anorganik) * 100
	persenTotal = (totalDaur / total) * 100

	return persenOrganik, persenAnorganik, persenTotal
}

func tampilData(organik, anorganik, po, pa, pt float64, status string) {
	/*
		IS:  Menerima total sampah organik dan anorganik, persentase daur ulang, dan status daur ulang.
		FS:  Menampilkan total sampah.
			Jika ada daur ulang, menampilkan persentase dan pesan motivasi.
			Jika tidak ada daur ulang, menampilkan pesan negatif.
	*/

	if po >= 100 {
		po = 100
	}
	if po >= 100 {
		po = 100
	}
	if pt > 100 {
		pt = 100
	}

	fmt.Printf("Selama %d hari kamu sudah mengeluarkan sampah sebanyak : %.1f Kg\n", jumData, organik+anorganik)

	switch strings.ToLower(status) {
	case "ya", "iya", "udah":
		fmt.Printf("Kamu sudah mendaur ulang sebanyak %.1f%% dari keseluruhan sampah yang kamu keluarkan,\n", pt)
		fmt.Printf("Untuk sampah organik kamu telah mendaur ulang %.1f%% dan untuk sampah anorganik %.1f%%\n", po, pa)
		if pt >= 50 {
			fmt.Println("Wow luar biasa, kamu harus tetap konsisten ya demi bumi yang lebih baik ! ")
		} else if pt < 50 && pt >= 9 {
			fmt.Println("Ayo kamu pasti bisa lebih dari ini, mari bersama mewujudkan lingkungan hidup yang lebih baik ")
		} else if pt < 9 && pt >= 1 {
			fmt.Println("Kamu harus lebih baik dari ini, untuk menjadikan bumi menjadi tempat yang lebih baik")
		} else {
			fmt.Println("kamu engga mendaur ulang sama sekali. Parah banget lu bang")
		}
	}

	switch strings.ToLower(status) {
	case "ngga", "tidak", "belom":
		fmt.Println("kamu engga mendaur ulang sama sekali. Parah banget lu bang")
	}

}

func editData(a *dataSampah) {
	/*
		IS:  Menerima pointer ke array data sampah. Meminta input tanggal data yang akan diedit dan data sampah baru.
		FS:  Mencari data berdasarkan tanggal.
			Jika data ditemukan, memperbarui data.
			Jika data tidak ditemukan, menampilkan pesan error.
	*/
	var tanggal, i int
	fmt.Print("Masukkan tanggal yang pengen kamu edit: ")
	fmt.Scan(&tanggal)

	for i = 0; i < jumData; i++ {
		if a[i].tanggal == tanggal {
			fmt.Print("Masukkan jumlah sampah organik (Kg): ")
			fmt.Scan(&a[i].Or)
			fmt.Print("Masukkan jumlah sampah anorganik (Kg): ")
			fmt.Scan(&a[i].An)
			fmt.Println("Data udah diedit")
			return
		}
	}
	fmt.Println("Tanggalnya engga ketemu")
}

func hapusData(a *dataSampah) {
	/*
		IS:  Menerima pointer ke array data sampah. Meminta input tanggal data yang akan dihapus.
		FS:  Mencari data berdasarkan tanggal.
			Jika data ditemukan, menghapus data.
			Jika data tidak ditemukan, menampilkan pesan error.
	*/
	var tanggal, i int
	fmt.Print("Masukkan tanggal yang ingin dihapus: ")
	fmt.Scan(&tanggal)

	for i = 0; i < jumData; i++ {
		if a[i].tanggal == tanggal {
			for j := i; j < jumData-1; j++ {
				a[j] = a[j+1]
			}
			jumData--
			fmt.Println("Data udah dihapus")
			return
		}
	}
	fmt.Println("Tanggalnya engga ketemu")
}

func cariData(a dataSampah) {
	/*
		IS:  Menerima array data sampah. Meminta input tanggal data yang akan dicari.
		FS:  Mencari data berdasarkan tanggal.
			Jika data ditemukan, menampilkan data tersebut.
			Jika data tidak ditemukan, menampilkan pesan error.
	*/
	var tanggal, i int
	var ditemukan bool
	fmt.Print("Masukkan tanggal yang pengen kamu cari: ")
	fmt.Scan(&tanggal)

	ditemukan = false
	for i = 0; i < jumData; i++ {
		if a[i].tanggal == tanggal {
			fmt.Printf("Tanggal: %d, Organik: %.2f, Anorganik: %.2f\n", a[i].tanggal, a[i].Or, a[i].An)
			ditemukan = true
			break
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

func urutkanData(a *dataSampah) {
	/*
		IS:  Menerima pointer ke array data sampah. Meminta input kriteria pengurutan dan opsi pencarian setelah pengurutan.
		FS:  Mengurutkan data berdasarkan kriteria yang dipilih (tanggal atau total sampah).
			Menampilkan data yang diurutkan.
			Jika pengguna memilih, melakukan pencarian pada data yang diurutkan.
	*/
	var i, pass, l, r, m, f, x int
	var temp sampah
	var bs string

	fmt.Println("Pilih kriteria pengurutan:")
	fmt.Println("1. Berdasarkan Tanggal")
	fmt.Println("2. Berdasarkan Jumlah Total Sampah")
	fmt.Print("Masukkan pilihan: ")

	var pilihan int
	fmt.Scan(&pilihan)

	for pass = 1; pass < jumData; i++ {
		temp = a[pass]
		i = pass
		for i > 0 {
			if pilihan == 1 {
				if temp.tanggal < a[i-1].tanggal {
					a[i] = a[i-1]
				} else {
					break
				}
			} else if pilihan == 2 {
				if (temp.Or + temp.An) < (a[i-1].Or + a[i-1].An) {
					a[i] = a[i-1]
				} else {
					break
				}
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
			i--
		}
		a[i] = temp
		pass += 1
	}

	if pilihan == 1 {
		fmt.Println("ini hasil pengurutan berdasar tanggal")
	} else if pilihan == 2 {
		fmt.Println("ini hasil pengurutan berdasar tanggal")
	}

	fmt.Println("\nData setelah diurutkan:")
	fmt.Println("+------------+--------------+----------------+--------------+")
	fmt.Println("|   Tanggal  | Organik (Kg) | Anorganik (Kg) |  Total (Kg)  |")
	fmt.Println("+------------+--------------+----------------+--------------+")
	for i = 0; i < jumData; i++ {
		total := a[i].Or + a[i].An
		fmt.Printf("| %10d | %12.2f | %14.2f | %12.2f |\n", a[i].tanggal, a[i].Or, a[i].An, total)
	}
	fmt.Println("+------------+--------------+----------------+--------------+")

	fmt.Println("mau cari data dari hasil yang udah diurutkan ?")
	fmt.Scan(&bs)

	l = 1
	r = jumData
	f = -1

	switch strings.ToLower(bs) {
	case "ya", "iya", "boleh", "mau":
		fmt.Println("mau cari data kamu pas tanggal berapa ?")
		fmt.Scan(&x)
		for l <= r && f == -1 {
			m = (l + r) / 2
			if x < a[m].tanggal {
				r = m - 1
			} else if x > a[m].tanggal {
				l = m + 1
			} else {
				f = m
			}
		}
		if f != -1 {
			fmt.Printf("Tanggal: %d, Organik: %.2f, Anorganik: %.2f, Total: %.2f\n", a[f].tanggal, a[f].Or, a[f].An, a[f].Or+a[f].An)
		} else {
			fmt.Print("kamu engga masukin data pas tanggal itu")
		}

	case "engga", "tidak", "ga":
		fmt.Println("okee")
	}

}
