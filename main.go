package main

import (
	"bufio"
	"fmt"
	"os"
	"retail-app/model"
	"retail-app/service"
	"retail-app/utils"
	"retail-app/view"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Meminta input login
	var session string
	for {
		fmt.Println("=== Login ===")
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		// Melakukan login
		s, err := utils.Login(username, password)
		if err != nil {
			fmt.Println("Login gagal:", err)
			continue // Jika login gagal, minta input ulang
		}
		session = s
		fmt.Println("Login berhasil!")
		break
	}

	// Verifikasi sesi setelah login berhasil
	if !utils.CheckSession(session) {
		fmt.Println("Sesi tidak valid!")
		return
	}

	for {
		// Menampilkan pilihan menu
		fmt.Println("\n=== Menu Retail ===")
		fmt.Println("1. Input Barang")
		fmt.Println("2. Jual Barang")
		fmt.Println("3. Tampilkan Daftar Barang dan Transaksi")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih opsi (1/2/3/4): ")

		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)
		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Input tidak valid, silakan masukkan nomor 1, 2, 3, atau 4.")
			continue
		}

		switch choice {
		case 1:
			// Input barang baru
			fmt.Println("\n=== Input Barang Baru ===")
			fmt.Print("Masukkan ID Barang: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Masukkan Nama Barang: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Masukkan Jumlah Barang: ")
			quantityInput, _ := reader.ReadString('\n')
			quantity, _ := strconv.Atoi(strings.TrimSpace(quantityInput))

			fmt.Print("Masukkan Harga Barang: ")
			priceInput, _ := reader.ReadString('\n')
			price, _ := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)

			newProduct := model.Product{
				ID:       id,
				Name:     name,
				Quantity: quantity,
				Price:    price,
			}

			if err := service.CreateProduct(newProduct); err != nil {
				fmt.Println("Gagal menambahkan barang:", err)
			} else {
				fmt.Println("Barang berhasil ditambahkan!")
			}

		case 2:
			// Jual barang
			fmt.Println("\n=== Jual Barang ===")
			fmt.Print("Masukkan ID Transaksi: ")
			transactionID, _ := reader.ReadString('\n')
			transactionID = strings.TrimSpace(transactionID)

			fmt.Print("Masukkan ID Barang yang Dijual: ")
			productID, _ := reader.ReadString('\n')
			productID = strings.TrimSpace(productID)

			fmt.Print("Masukkan Jumlah yang Dijual: ")
			sellQuantityInput, _ := reader.ReadString('\n')
			sellQuantity, _ := strconv.Atoi(strings.TrimSpace(sellQuantityInput))

			transaction := model.Transaction{
				ID:        transactionID,
				ProductID: productID,
				Quantity:  sellQuantity,
			}

			if err := service.CreateTransaction(transaction); err != nil {
				fmt.Println("Transaksi gagal:", err)
			} else {
				fmt.Println("Transaksi berhasil dilakukan!")
			}

		case 3:
			// Tampilkan daftar produk dan transaksi
			fmt.Println("\n=== Daftar Barang ===")
			products, _ := utils.LoadProducts()
			view.PrintJSON(products)

			fmt.Println("\n=== Daftar Transaksi ===")
			transactions, _ := utils.LoadTransactions()
			view.PrintJSON(transactions)

		case 4:
			// Keluar dari aplikasi
			fmt.Println("Terima kasih telah menggunakan aplikasi retail!")
			return

		default:
			fmt.Println("Pilihan tidak valid, silakan pilih nomor antara 1 dan 4.")
		}
	}
}
