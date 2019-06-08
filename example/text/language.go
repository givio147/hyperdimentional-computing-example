package text

import (
    "github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"
)

type Language struct {
    Name string
    Profile *hyperdimensional.VecBinomial
    letters *Letters
    encoder *Encoder
}

func NewLanguage(name string, letters *Letters) *Language {
    return &Language{
        Name: name,
        letters: letters,
        encoder: NewEncoder(letters),
    }
}

func (l *Language) encodeLanguage(text *string) {
    l.Profile = l.encoder.encodeLanguage(text)
}
