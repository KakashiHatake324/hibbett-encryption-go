package hibbettencryptiongo

import "github.com/amenzhinsky/go-memexec"

// encrypt card number for hibbett
func EncryptCard(cardNum string) (string, error) {
	exe, err := memexec.New(program)
	if err != nil {
		return "", err
	}

	out, err := exe.Command(cardNum).Output()
	if err != nil {
		return "", err
	}
	return string(out), err
}
