package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	//TODO: ajouter les variables pour les coordonnées et la couleur du pixel
	var x, y int
	var couleur string
	var choice int

	// API endpoint
	endpoint := "http://pixelwar.peda.ydayslyon.fr/api/set-pixel-bis"

	for {
		//TODO: faire un menu pour choisir d'ajouter un pixel, une ligne de pixel, une colonne de pixel ou de quitter
		fmt.Println("1 : Ajouter un pixel")
		fmt.Println("2 : Ajouter une ligne")
		fmt.Println("3 : Ajouter une colonne de pixel")
		fmt.Println("4 : Quitter")

		fmt.Scanln(&choice)

		switch choice {
		//ajouter un pixel
		case 1:
			//TODO: récupération des coordonnées et de la couleur du pixel
			println("Entrer Pixel X")
			fmt.Scanln(&x)
			println("Entrer Pixel Y")
			fmt.Scanln(&y)
			println("Entrer Couleur")
			fmt.Scanln(&couleur)
			//TODO: formatage des données en structure url.Values
			formData := url.Values{}
			formData.Add("x", strconv.Itoa(x))
			formData.Add("y", strconv.Itoa(y))
			formData.Add("color", string(couleur))
			//TODO: envoi de la requête POST
			resp, err := http.PostForm(endpoint, formData)
			if err != nil {
				log.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				bodyBytes, _ := ioutil.ReadAll(resp.Body) // Read the response body
				bodyString := string(bodyBytes)
				log.Printf("Error: Received status code %d with message: %s", resp.StatusCode, bodyString)
			} else {
				fmt.Println("Pixel added successfully!")
			}

		//ajouter une ligne de pixel
		case 2:
			//TODO: récupération des coordonnées et de la couleur de la ligne
			println("Entrer Pixel Ligne")
			fmt.Scanln(&x)
			println("Entrer Couleur")
			fmt.Scanln(&couleur)

			//TODO: envoi de la requête POST
			formData := url.Values{}
			formData.Add("x", strconv.Itoa(1))
			formData.Add("color", string(couleur))
			for z := 0; z < 70; z++ {
				formData.Add("y", strconv.Itoa(x+z))
				//TODO: envoi de la requête POST
				resp, err := http.PostForm(endpoint, formData)
				if err != nil {
					log.Fatalf("Failed to send request: %v", err)
				}
				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					bodyBytes, _ := ioutil.ReadAll(resp.Body) // Read the response body
					bodyString := string(bodyBytes)
					log.Printf("Error: Received status code %d with message: %s", resp.StatusCode, bodyString)
				} else {
					fmt.Println("Pixel added successfully!")
				}
			}

		//ajouter une colonne de pixel
		case 3:
			//TODO: récupération des coordonnées et de la couleur de la colonne

			//TODO: envoi de la requête POST

		//quitter
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

//bonus :
//TODO: ajouter une fonction pour vérifier si les coordonnées sont valides

//TODO: ajouter une fonction pour vérifier si la ligne est valide

//TODO: ajouter une fonction pour vérifier si la colonne est valide

//TODO: ajouter une fonction pour l'envoi de la requête POST

//TODO: ajouter une fonction pour vérifier si la couleur est valide
