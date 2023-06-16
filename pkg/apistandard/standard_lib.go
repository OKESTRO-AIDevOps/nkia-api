package apistandard

import (
	"fmt"
	"strings"

	pkgutils "github.com/seantywork/x0f_npia/pkg/utils"
)

func (asgi API_STD) LegacyInputTranslate(legacy_in string) (API_INPUT, error) {

	ret_api_input := make(API_INPUT)

	legacy_list := strings.SplitN(legacy_in, ":", 2)

	matching_key := 0

	new_key := legacy_list[0]

	c_list := strings.Split(legacy_list[1], ",")

	if new_key == "SUBMIT" {

	} else if new_key == "CALLME" {

	} else if new_key == "GITLOG" {

	} else if new_key == "PIPEHIST" {

	} else if new_key == "PIPE" {

	} else if new_key == "PIPELOG" {

	} else if new_key == "BUILD" {

	} else if new_key == "BUILDLOG" {

	} else if new_key == "DELND" {

	} else if new_key == "EXIT" {

	}

	obj, resized_c_list := pkgutils.PopFromSliceByIndex[string](c_list, 1)

	std_key_obj := ""

	if new_key == "SETTING" && obj == "CRTNS" {

		std_key_obj = new_key + "-" + obj

		_ = pkgutils.InsertToSliceByIndex[string](resized_c_list, 0, std_key_obj)

	}

	if matching_key == 0 {

		return ret_api_input, fmt.Errorf("key not found")
	}

	return ret_api_input, nil

}

func (asgi API_STD) LegacyTranslationBuildHelper(std_keys []string, legacy_c_list []string) (API_INPUT, error) {

	ret_api_std := make(API_INPUT)

	if len(std_keys) != len(legacy_c_list) {

		return ret_api_std, fmt.Errorf("unsupported translation format")

	}

	for i := 0; i < len(std_keys); i++ {

		ret_api_std[std_keys[i]] = legacy_c_list[i]

	}

	return ret_api_std, nil
}

func (asgi API_STD) LegacyOutputTranslate() {

}

func (asgi API_STD) Verify() {

}

func (asgi API_STD) PrintPrettyDefinition() {

	fmt.Println(API_DEFINITION)
}

func (asgi API_STD) PrintRawDefinition() {

	fmt.Println(asgi)

}
