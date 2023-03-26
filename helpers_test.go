package russian_name_generator

import (
	"testing"
)

func TestSeed(t *testing.T) {
	Seed(0)
}

func TestTransliterate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single letter",
			input:    "а",
			expected: "a",
		},
		{
			name:     "Multiple letters",
			input:    "привет, мир",
			expected: "privet, mir",
		},
		{
			name:     "Non-Russian characters",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "Mixed Russian and non-Russian characters",
			input:    "Привет, world!",
			expected: "Privet, world!",
		},
	}

	for _, test := range tests {
		localTest := test
		t.Run(localTest.name, func(t *testing.T) {
			t.Parallel()

			result := Transliterate(localTest.input)
			if result != localTest.expected {
				t.Errorf("Expected %q, but got %q", localTest.expected, result)
			}
		})
	}
}

func TestGetPatronymicFromName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		isFeminine bool
		expected   string
	}{
		{
			name:     "Иван",
			expected: "Иванович",
		},
		{
			name:     "Виктор",
			expected: "Викторович",
		},
		{
			name:     "Пётр",
			expected: "Петрович",
		},
		{
			name:     "Лев",
			expected: "Львович",
		},
		{
			name:     "Зосима",
			expected: "Зосимович",
		},
		{
			name:     "Акакий",
			expected: "Акакиевич",
		},
		{
			name:     "Андрей",
			expected: "Андреевич",
		},
		{
			name:     "Кондратий",
			expected: "Кондратьевич",
		},
		{
			name:     "Прокофий",
			expected: "Прокофьевич",
		},
		{
			name:     "Дмитрий",
			expected: "Дмитриевич",
		},
		{
			name:     "Димитрий",
			expected: "Димитриевич",
		},
		{
			name:     "Георгий",
			expected: "Георгиевич",
		},
		{
			name:     "Николай",
			expected: "Николаевич",
		},
		{
			name:     "Кирилл",
			expected: "Кириллович",
		},
		{
			name:     "Владимир",
			expected: "Владимирович",
		},
		{
			name:     "Глеб",
			expected: "Глебович",
		},
		{
			name:     "Гаврила",
			expected: "Гаврилович",
		},
		{
			name:     "Иона",
			expected: "Ионович",
		},
		{
			name:     "Фома",
			expected: "Фомич",
		},
		{
			name:     "Фока",
			expected: "Фокич",
		},
		{
			name:     "Никита",
			expected: "Никитич",
		},
		{
			name:     "Жорж",
			expected: "Жоржевич",
		},
		{
			name:       "Фока",
			isFeminine: true,
			expected:   "Фокична",
		},
		{
			name:       "Фёдор",
			isFeminine: true,
			expected:   "Фёдоровна",
		},
		{
			name:       "Фома",
			isFeminine: true,
			expected:   "Фоминична",
		},
		{
			name:     "Егор",
			expected: "Егорович",
		},
		{
			name:     "Фрол",
			expected: "Фролович",
		},
		{
			name:     "Фаддей",
			expected: "Фаддеевич",
		},
		{
			name:     "Эмиль",
			expected: "Эмилевич",
		},
		{
			name:     "Игорь",
			expected: "Игоревич",
		},
		{
			name:     "Лазарь",
			expected: "Лазаревич",
		},
		{
			name:     "Яков",
			expected: "Яковлевич",
		},
		{
			name:     "Ярослав",
			expected: "Ярославович",
		},
		{
			name:     "Савва",
			expected: "Саввич",
		},
		{
			name:     "Илья",
			expected: "Ильич",
		},
		{
			name:     "Павел",
			expected: "Павлович",
		},
		{
			name:     "Валерий",
			expected: "Валерьевич",
		},
		{
			name:     "Захар",
			expected: "Захарович",
		},
		{
			name:     "Захарий",
			expected: "Захарьевич",
		},
		{
			name:     "Юрий",
			expected: "Юрьевич",
		},
		{
			name:     "Василий",
			expected: "Васильевич",
		},
		{
			name:       "Андрей",
			isFeminine: true,
			expected:   "Андреевна",
		},
		{
			name:       "Семён",
			isFeminine: true,
			expected:   "Семёновна",
		},
		{
			name:       "Пётр",
			isFeminine: true,
			expected:   "Петровна",
		},
		{
			name:       "Лев",
			isFeminine: true,
			expected:   "Львовна",
		},
		{
			name:       "Илья",
			isFeminine: true,
			expected:   "Ильинична",
		},
		{
			name:       "Савва",
			isFeminine: true,
			expected:   "Саввична",
		},
		{
			name:       "Василий",
			isFeminine: true,
			expected:   "Васильевна",
		},
		{
			name:       "Павел",
			isFeminine: true,
			expected:   "Павловна",
		},
		{
			name:       "Никита",
			isFeminine: true,
			expected:   "Никитична",
		},
		{
			name:       "Кузьма",
			isFeminine: true,
			expected:   "Кузьминична",
		},
		{
			name:       "Евгений",
			isFeminine: true,
			expected:   "Евгеньевна",
		},
		{
			name:       "Виктор",
			isFeminine: true,
			expected:   "Викторовна",
		},
		{
			name:       "Михаил",
			isFeminine: true,
			expected:   "Михайловна",
		},
		{
			name:       "Данила",
			isFeminine: true,
			expected:   "Даниловна",
		},
		{
			name:       "Всеволод",
			isFeminine: true,
			expected:   "Всеволодовна",
		},
		{
			name:       "Гаврила",
			isFeminine: true,
			expected:   "Гавриловна",
		},
		{
			name:       "Даниил",
			isFeminine: true,
			expected:   "Даниловна",
		},
		{
			name:       "Менея",
			isFeminine: true,
			expected:   "Менеевна",
		},
		{
			name:       "Пров",
			isFeminine: true,
			expected:   "Провна",
		},
		{
			name:       "Вилли",
			isFeminine: true,
			expected:   "Виллиевна",
		},
		{
			name:     "Вяйне",
			expected: "Вяйневич",
		},
		{
			name:     "Василько",
			expected: "Василькович",
		},
		{
			name:     "Мина",
			expected: "Минович"},
	}

	for _, test := range tests {
		localTest := test
		t.Run(localTest.name, func(t *testing.T) {
			t.Parallel()

			result := getPatronymicFromName(localTest.name, localTest.isFeminine)
			if result != localTest.expected {
				t.Errorf("Expected %q, but got %q", localTest.expected, result)
			}
		})
	}
}

func TestGetRandValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"person", "notfound"}} {
		if getRandValue(globalFaker.Rand, test) != "" {
			t.Error("You should have gotten no value back")
		}
	}
}
