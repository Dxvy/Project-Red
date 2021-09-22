package main

import (
	"fmt"
)

type Personnage struct {
	nom    string
	classe string
	niveau int
	hpMax  int
	hpAct  int
	invent []string
}

func (c *Personnage) Init(nom string, classe string, niveau int, hpMax int, hpAct int, invent []string) {
	c.nom = nom
	c.classe = classe
	c.niveau = niveau
	c.hpMax = hpMax
	c.hpAct = hpAct
	c.invent = invent
}

func (c Personnage) DisplayInfo() {
	fmt.Println("-------------------")
	fmt.Println("Nom du perso:", c.nom)
	fmt.Println("Classe du perso:", c.classe)
	fmt.Println("Niveau du perso:", c.niveau)
	fmt.Println("HpMax du perso:", c.hpMax)
	fmt.Println("HpAtuel du perso:", c.hpAct)
	fmt.Println("-------------------")
	fmt.Println("Inventaire du perso:", c.nom)
}

func (c Personnage) PrintInvent() {
	if len(c.invent) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		for i := range c.invent {
			fmt.Println("Accessoire", i+1, " : ", c.invent[i], "\n")
		}
	}
}

func (c Personnage) TakePot() {
	test := 0
	for i := range c.invent {
		if c.invent[i] == "potion" {
			test++
			if c.hpAct+50 <= c.hpMax-50 {
				c.hpAct += 50
				fmt.Println("La vie de ", c.nom, " est passée à ", c.hpAct)
			} else if c.hpMax-c.hpAct < 50 && c.hpMax-c.hpAct > 0 {
				c.hpAct = c.hpMax
				fmt.Println("La vie de ", c.nom, " est passée à ", c.hpAct)
			} else if c.hpAct == c.hpMax {
				fmt.Println(c.nom, "est déjà full vie")
			} else {
				fmt.Println("La potion ne peut être utilisée")
			}
		}
	}
	if test == 0 {
		fmt.Println(c.nom, " n'a pas de potions dans son inventaire")
	}
}
func main() {
	var c1 Personnage
	c1.Init("Hades", "Dieu", 666, 9999, 999, []string{"Casque d'invisibilité", "Cerbère"})
	c1.DisplayInfo()
	c1.PrintInvent()
	c1.TakePot()
}
