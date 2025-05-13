package factory

type ak47 struct {
	Gun
}

func newAk47() IGun {
	return &ak47{
		Gun: Gun{
			name:  "Ak47 gun",
			power: 4,
		},
	}

}
