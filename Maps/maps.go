package maps

import (
	"fmt"
)

func Map() {
	user := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Pedro",
			"segundo":  "Lucas",
		},
		"curso": {
			"nome":      "Engenharia",
			"faculdade": "UNICAMP",
		},
	}
	fmt.Println(user2)
	delete(user2, "nome")
	fmt.Println(user2)

	user2["idade"] = map[string]string{
		"idade": "23",
	}

	fmt.Println(user2)
}
