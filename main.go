package main

import (
	"fmt"
	linkedList "std-stack/helpers"
)

func main() {
	helper := linkedList.LinkedList{}

	// * insert3
	helper.Insert(linkedList.Mahasiswa{
		NIM:   1111111,
		Name:  "imam firmansyah",
		Prodi: "Tehnik Informatika",
	})
	helper.Insert(linkedList.Mahasiswa{
		NIM:   22222222,
		Name:  "Ahmed Shipuden",
		Prodi: "Sistem Informatika",
	})

	helper.Insert(linkedList.Mahasiswa{
		NIM:   333333,
		Name:  "Sang Arjuna",
		Prodi: "Tehnik Sipil",
	})

	helper.Insert(linkedList.Mahasiswa{
		NIM:   4444444,
		Name:  "Siti Venus",
		Prodi: "Ternak Lele",
	})

	helper.Insert(linkedList.Mahasiswa{
		NIM:   5555555,
		Name:  "Agam Wazoski",
		Prodi: "Tehnik Informatika",
	})

	fmt.Println("Inisiate Data: ")
	helper.Print()

	fmt.Println("Delete Index 1:")
	helper.DeleteIndex(1) // * delete at index 1
	helper.Print()        // * print

	fmt.Println("Delete Index 3:")
	helper.DeleteIndex(3) // * delete at index 3
	helper.Print()        // * print

	fmt.Println("Insert Data: ")
	helper.Insert(linkedList.Mahasiswa{
		NIM:   6666666,
		Name:  "Ardafa Chan",
		Prodi: "Tehnik Sipil",
	})
	helper.Print() // * print
}
