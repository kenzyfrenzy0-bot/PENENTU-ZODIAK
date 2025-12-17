package main

import "fmt"

func main() {
	var nama, bulan, zodiak string
	var tgl int

	// ===== CEK NAMA =====
	fmt.Print("Masukkan nama anda: ")
	fmt.Scanln(&nama)

	validNama := false
	for validNama == false {
		salah := false
		for _, c := range nama {
			if c >= '0' && c <= '9' {
				salah = true
			}
		}

		if salah {
			fmt.Print("Nama tidak boleh angka, masukkan lagi: ")
			fmt.Scanln(&nama)
		} else {
			validNama = true
		}
	}

	// ===== CEK BULAN =====
	fmt.Print("Masukkan bulan lahir anda: ")
	fmt.Scanln(&bulan)

	validBulan := false
	for validBulan == false {
		adaAngka := false
		for _, c := range bulan {
			if c >= '0' && c <= '9' {
				adaAngka = true
			}
		}

		bulanValid :=
			bulan == "januari" || bulan == "februari" || bulan == "maret" ||
				bulan == "april" || bulan == "mei" || bulan == "juni" ||
				bulan == "juli" || bulan == "agustus" || bulan == "september" ||
				bulan == "oktober" || bulan == "november" || bulan == "desember"

		if adaAngka || !bulanValid {
			fmt.Print("Bulan tidak valid, masukkan lagi: ")
			fmt.Scanln(&bulan)
		} else {
			validBulan = true
		}
	}

	// ===== CEK TANGGAL =====
	fmt.Print("Masukkan tanggal lahir anda: ")
	fmt.Scanln(&tgl)

	validTanggal := false
	for validTanggal == false {
		if tgl <= 0 {
			fmt.Print("Tanggal harus > 0, masukkan lagi: ")
			fmt.Scanln(&tgl)
		} else if bulan == "februari" && tgl > 29 {
			fmt.Print("Februari maksimal 29, masukkan lagi: ")
			fmt.Scanln(&tgl)
		} else if (bulan == "april" || bulan == "juni" || bulan == "september" || bulan == "november") && tgl > 30 {
			fmt.Print("Bulan ini maksimal 30 hari, masukkan lagi: ")
			fmt.Scanln(&tgl)
		} else if tgl > 31 {
			fmt.Print("Tanggal maksimal 31, masukkan lagi: ")
			fmt.Scanln(&tgl)
		} else {
			validTanggal = true
		}
	}

	// ===== ZODIAK =====
	if (bulan == "maret" && tgl >= 21) || (bulan == "april" && tgl <= 19) {
		zodiak = "Aries"
	} else if (bulan == "april" && tgl >= 20) || (bulan == "mei" && tgl <= 20) {
		zodiak = "Taurus"
	} else if (bulan == "mei" && tgl >= 21) || (bulan == "juni" && tgl <= 20) {
		zodiak = "Gemini"
	} else if (bulan == "juni" && tgl >= 21) || (bulan == "juli" && tgl <= 22) {
		zodiak = "Cancer"
	} else if (bulan == "juli" && tgl >= 23) || (bulan == "agustus" && tgl <= 22) {
		zodiak = "Leo"
	} else if (bulan == "agustus" && tgl >= 23) || (bulan == "september" && tgl <= 22) {
		zodiak = "Virgo"
	} else if (bulan == "september" && tgl >= 23) || (bulan == "oktober" && tgl <= 22) {
		zodiak = "Libra"
	} else if (bulan == "oktober" && tgl >= 23) || (bulan == "november" && tgl <= 21) {
		zodiak = "Scorpio"
	} else if (bulan == "november" && tgl >= 22) || (bulan == "desember" && tgl <= 21) {
		zodiak = "Sagitarius"
	} else if (bulan == "desember" && tgl >= 22) || (bulan == "januari" && tgl <= 19) {
		zodiak = "Capricorn"
	} else if (bulan == "januari" && tgl >= 20) || (bulan == "februari" && tgl <= 18) {
		zodiak = "Aquarius"
	} else {
		zodiak = "Pisces"
	}

	// ===== OUTPUT =====
	fmt.Println("===============================")
	fmt.Println("Nama    :", nama)
	fmt.Println("Tanggal :", tgl, bulan)
	fmt.Println("Zodiak  :", zodiak)
}
