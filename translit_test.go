package translit

import "testing"

func TestTranslit(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"Привет, это тестовое сообщение для транслита.",
			"Privet, jeto testovoe soobshhenie dlja translita."},
		{"Не следует, однако забывать, что консультация с широким активом позволяет оценить значение новых предложений.",
			"Ne sleduet, odnako zabyvat, chto konsultacija s shirokim aktivom pozvoljaet ocenit znachenie novyh predlozhenij."},
	}

	for _, test := range tests {
		if got := Translit(test.input); got != test.want {
			t.Errorf("Translit(%q) = %q", test.input, got)
		}
	}
}

func BenchmarkTranslit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Translit("Не следует, однако забывать, что консультация с широким активом позволяет оценить значение новых предложений.")
	}
}
