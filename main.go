package main

import "fmt"

type Penghuni interface {
	lakukanHal() string
}

type RumahKos struct {
	nama string
}

func (r *RumahKos) lakukanHal() string {
	return "Halo, saya " + r.nama
}

// concrete RumahKos bernama PemilikKos, implement RumahKos interface (PemilikKos)
type PemilikKos struct {
	*RumahKos
}

// encapsulate RumahKos creation
func NewPemilikKos() Penghuni {
	return &PemilikKos{
		RumahKos: &RumahKos{
			nama: "PemilikKos",
		},
	}
}

// concrete RumahKos bernama AnakKos, implement RumahKos interface (AnakKos)
type AnakKos struct {
	*RumahKos
}

// encapsulate RumahKos creation
func NewAnakKos() Penghuni {
	return &AnakKos{
		RumahKos: &RumahKos{
			nama: "AnakKos",
		},
	}
}

// factory
func CreateRumahKos(nama string) Penghuni {
	if nama == "PemilikKos" {
		return NewPemilikKos()
	} else if nama == "AnakKos" {
		return NewAnakKos()
	}
	return nil
}

func main() {
	pemilik := CreateRumahKos("PemilikKos")
	fmt.Println(pemilik.lakukanHal())
	anak := CreateRumahKos("AnakKos")
	fmt.Println(anak.lakukanHal())
}
