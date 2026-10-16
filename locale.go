package ledger

import "golang.org/x/text/language"

// ParseLocale returns the canonical form of a BCP 47 language tag.
func ParseLocale(tag string) (string, error) {
	t, err := language.Parse(tag)
	if err != nil {
		return "", err
	}
	return t.String(), nil
}
