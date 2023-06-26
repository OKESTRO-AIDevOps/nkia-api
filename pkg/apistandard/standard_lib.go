package apistandard

import (
	"fmt"

	pkgutils "github.com/OKESTRO-AIDevOps/npia-api/pkg/utils"
)

func (asgi API_STD) Verify(verifiable API_INPUT) error {

	cmd_id, okay := verifiable["id"]

	var duplicate_check []string

	if !okay {
		return fmt.Errorf("verification failed: %s", "missing command id")
	}

	v_list, okay := asgi[cmd_id]

	if !okay {

		return fmt.Errorf("verification failed: %s", "invalid command id")
	}

	if len(v_list) != len(verifiable) {
		return fmt.Errorf("verification failed: %s", "invalid command structure")
	}

	for i := range verifiable {

		hit := pkgutils.CheckIfSliceContains[string](duplicate_check, i)

		if hit {
			return fmt.Errorf("verification failed: %s", "invalid command structure: duplicate key")
		}

		hit = pkgutils.CheckIfSliceContains[string](v_list, i)

		if !hit {
			return fmt.Errorf("verification failed: %s", "invalid command structure: wrong key")
		}

		duplicate_check = append(duplicate_check, i)
	}

	return nil

}

func (asgi API_STD) PrintPrettyDefinition() {

	fmt.Println(API_DEFINITION)
}

func (asgi API_STD) PrintRawDefinition() {

	fmt.Println(asgi)

}
