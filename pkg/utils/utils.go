package utils

func CheckIfEleInStrList(ele string, str_list []string) bool {

	hit := false

	for i := 0; i < len(str_list); i++ {

		if str_list[i] == ele {

			hit = true

			return hit
		}

	}

	return hit

}
