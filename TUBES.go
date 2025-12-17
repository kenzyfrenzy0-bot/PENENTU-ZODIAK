package main

import "fmt"

func main() {
	var nama, bulan, zodiak string
	var tgl int

	// CEK NAMA 
	validNama := false
	for !validNama {
		fmt.Print("Masukkan nama anda: ")
		fmt.Scanln(&nama)

		if nama != "" {
			validNama = true
		} else {
			fmt.Println("Nama tidak boleh kosong!")
		}
	}

	// CEK BULAN 
	validBulan := false
	for !validBulan {
		fmt.Print("Masukkan bulan lahir anda: ")
		fmt.Scanln(&bulan)

		if bulan == "januari" || bulan == "februari" || bulan == "maret" ||
			bulan == "april" || bulan == "mei" || bulan == "juni" ||
			bulan == "juli" || bulan == "agustus" || bulan == "september" ||
			bulan == "oktober" || bulan == "november" || bulan == "desember" {
			validBulan = true
		} else {
			fmt.Println("Bulan tidak valid dan tidak boleh mengandung angka!")
		}
	}

	// ===== CEK TANGGAL =====
	validTanggal := false
	for !validTanggal {
		fmt.Print("Masukkan tanggal lahir anda: ")
		n, _ := fmt.Scanln(&tgl)

		if n == 0 {
			fmt.Println("Tanggal harus berupa angka!")
			fmt.Scanln()
		} else if tgl < 1 || tgl > 31 {
			fmt.Println("Tanggal harus antara 1 sampai 31!")
		} else if bulan == "februari" && tgl > 29 {
			fmt.Println("Februari maksimal 29 hari!")
		} else if (bulan == "april" || bulan == "juni" ||
			bulan == "september" || bulan == "november") && tgl > 30 {
			fmt.Println("Bulan ini maksimal 30 hari!")
		} else {
			validTanggal = true
		}
	}

	// ZODIAK 
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

	//OUTPUT
	fmt.Println("===============================")
	fmt.Println("Nama    :", nama)
	fmt.Println("Tanggal :", tgl, bulan)
	fmt.Println("Zodiak  :", zodiak)
}
