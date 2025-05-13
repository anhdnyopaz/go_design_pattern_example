package factory

import "fmt"

func GetGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed ")

}
