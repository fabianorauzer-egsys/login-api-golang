package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("O usuário escolhido já existe.")
	}

	if strings.Contains(err, "email") {
		return errors.New("O e-mail escolhido já está em uso.")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Senha incorreta. Por favor, tente novamente.")
	}
	return errors.New("Incorrect Details")
}
