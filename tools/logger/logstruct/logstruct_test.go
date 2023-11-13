package logstruct

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestFields(t *testing.T) {
	r := require.New(t)

	type EmailBox struct {
		Email       string `log:"email"`
		HiddenEmail string `log:"-"`
	}

	type Data struct {
		Contract    int     `db:"c_id" log:"c_id"`
		OperLogin   string  `json:"oper_login" log:"oper_login"`
		Balance     float64 `log:"balance"`
		Exists      bool    `log:"exists"`
		NotExists1  bool    `log:"-"` // пропускается
		NotExists2  bool    // тоже самое что и выше, если нет log тэга - не попадет в logrus.Fields
		TestPointer *int    `log:"test_pointer"` // поинтер значения или структуры пропускаются из лога
		EmailBox
	}

	i := 10

	d := Data{
		Contract:    1234,
		OperLogin:   "m.zarif@sarkor.uz",
		Balance:     12345.6,
		Exists:      true,
		NotExists1:  true,
		NotExists2:  true,
		TestPointer: &i,
		EmailBox: EmailBox{
			Email:       "test@email.uz",
			HiddenEmail: "hiddentest@email.uz",
		},
	}

	expected := logrus.Fields{
		"c_id":       1234,
		"oper_login": "m.zarif@sarkor.uz",
		"balance":    12345.6,
		"exists":     true,
		"email":      "test@email.uz",
	}

	// умеет работать как со значением
	result := Fields(d)
	r.Equal(expected, result)

	// умеет работать так и по поинтеру структуры
	result = Fields(&d)
	r.Equal(expected, result)

	// при передаче не структуры = результат будет пустым
	result = Fields("test")
	r.Empty(result)

}
