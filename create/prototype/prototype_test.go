package prototype

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 原型模式

type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

func (k *Keyword) Clone() *Keyword {
	var newKey Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKey)
	return &newKey
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(upWords []*Keyword) Keywords {
	newKeys := Keywords{}

	for k, v := range words {
		newKeys[k] = v
	}

	for _, word := range upWords {
		newKeys[word.word] = word.Clone()
	}

	return newKeys
}

func TestPrototype(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			word:      "testA",
			visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": &Keyword{
			word:      "testB",
			visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": &Keyword{
			word:      "testC",
			visit:     3,
			UpdatedAt: &updateAt,
		},
	}

	now := time.Now()
	updatedWords := []*Keyword{
		{
			word:      "testB",
			visit:     10,
			UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)

	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.NotEqual(t, updatedWords[0], got["testB"])
	assert.Equal(t, words["testC"], got["testC"])
}
