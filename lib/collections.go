package lib

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// # Реализуйте функцию `make_censored()`, которая заменяет каждое вхождение указанных слов в предложении на последовательность `$#%!` и возвращает полученную строку. Аргументы:

// # * Текст
// # * Набор стоп слов

// # Словом считается любая непрерывная последовательность символов, включая любые спецсимволы (без пробелов).

// # ```ruby
// # sentence = 'When you play the game of thrones, you win or you die'
// # puts make_censored(sentence, ['die', 'play'])
// # # => 'When you $#%! the game of thrones, you win or you $#%!'

// # sentence2 = 'chicken chicken? chicken! chicken'
// # puts make_censored(sentence2, ['?', 'chicken'])
// # # => '$#%! chicken? chicken! $#%!'
// # ```

func include(arr []string, el string) bool {
	for _, e := range arr {
		if e == el {
			return true
		}
	}
	return false
}

const (
	sensoredSymbol = "$#%!"
)

func MakeSensored(sentense string, stopWords []string) string {
	var resultARR []string
	arr := strings.Split(sentense, " ")
	for _, s := range arr {
		if include(stopWords, s) {
			resultARR = append(resultARR, sensoredSymbol)
		} else {
			resultARR = append(resultARR, s)
		}
	}
	return strings.Join(resultARR, " ")
}

// # Реализуйте функцию `compare_versions()`, которая сравнивает переданные версии version1 и version2. Если version1 > version2, то функция должна вернуть 1, если version1 < version2, то -1, если же version1 = version2, то 0.

// # Версия – это строка, в которой два числа (мажорная и минорные версии) разделены точкой, например, 12.11. Важно понимать, что версия – это не число с плавающей точкой, а несколько чисел, не связанных между собой. Проверка на больше/меньше производится сравнением каждого числа независимо. Поэтому версия 0.12 больше версии 0.2.

// # ```ruby
// # puts compare_versions("0.1", "0.2")
// # # => -1

// # puts compare_versions("0.2", "0.1")
// # # => 1

// # puts compare_versions("4.2", "4.2")
// # # => 0
// # ```

func CompareVersion(v1, v2 string) int8 {
	var err error
	var versions []string
	versions = strings.Split(v1, ".")
	major1, err := strconv.Atoi(versions[0])
	if err != nil {
		panic("Error: error when parse from str to i")
	}
	minor1, err := strconv.Atoi(versions[1])
	if err != nil {
		panic("Error: error when parse from str to i")
	}
	// strconv.Atoi(1)

	versions = strings.Split(v2, ".")
	// major2 := versions[0]
	// minor2 := versions[1]

	major2, err := strconv.Atoi(versions[0])
	if err != nil {
		panic("Error: error when parse from str to i")
	}
	minor2, err := strconv.Atoi(versions[1])
	if err != nil {
		panic("Error: error when parse from str to i")
	}

	if major1 > major2 {
		return 1
	} else if major1 == major2 {
		if minor1 == minor2 {
			return 0
		} else if minor1 > minor2 {
			return 1
		} else {
			return -1
		}
	} else {
		return -1
	}
}

// # frozen_string_literal: true

// require_relative 'test_helper'
// require_relative '../lib/build_query_string'

// ### lib/build_query_string.rb

// # Реализуйте функцию `build_query_string()`, которая принимает на вход список параметров и возвращает сформированный query string из этих параметров. Имена параметров в выходной строке должны располагаться в алфавитном порядке (то есть их нужно отсортировать).

// # ```ruby
// # puts build_query_string({ per: 10, page: 1 })
// # # => 'page=1&per=10'

// # puts build_query_string({ param: 'all', param1: true  })
// # # => 'param=all&param1=true'
// # ```

// # rubocop:disable Style/For
// # BEGIN

// def build_query_string(parametrs)
//   return '' if parametrs.empty?

//   result = []
//   sorted_params = parametrs.sort
//   sorted_params.each do |el|
//     key, value = el
//     result << "#{key}=#{value}"
//   end
//   result.join('&')
// end

// # END
// # rubocop:enable Style/For

func BuildQueryString(params map[string]string) string {
	var names []string
	var strs []string

	for k, _ := range params {
		names = append(names, k)
	}

	sort.Strings(names)
	for _, v := range names {
		strs = append(strs, fmt.Sprintf("%s=%s", v, params[v]))
	}
	return strings.Join(strs, "&")
}
