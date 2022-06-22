package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

type Wallet struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

func main() {

	wif, err := networks["btc"].CreatePrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	pk := wif.String()

	address, err := networks["btc"].GetAddress(wif)
	if err != nil {
		log.Fatal(err)
	}

	createWalletFolder("wallets")

	fmt.Println(color.CyanString("PublicKey: "), color.YellowString(address.EncodeAddress()))
	fmt.Println(color.CyanString("PrivateKey: "), color.YellowString(pk))

	wallet := Wallet{
		PublicKey:  address.EncodeAddress(),
		PrivateKey: pk,
	}

	file, _ := json.MarshalIndent(wallet, "", " ")
	_ = ioutil.WriteFile("wallets/"+address.EncodeAddress()+".json", file, 0644)

}

func createWalletFolder(dirname string) bool {
	_, error := os.Stat(dirname)
	if os.IsNotExist(error) {
		if err := os.Mkdir(dirname, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		return true
	} else {
		return true
	}
}
