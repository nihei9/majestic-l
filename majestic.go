package majestic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

type Parser func(src []byte) (map[string]interface{}, error)

type Expectation map[string]MatchableValue

type Config struct {
	Parser Parser
}

type verifier struct {
	config       *Config
	expectations []Expectation
	ptr          int
}

func Expect(srcs ...interface{}) ([]Expectation, error) {
	exps := []Expectation{}
	for _, src := range srcs {
		exp, err := parseExpectation(src)
		if err != nil {
			return nil, err
		}
		exps = append(exps, exp)
	}

	return exps, nil
}

func parseExpectation(src interface{}) (Expectation, error) {
	expectation := Expectation{}
	t := reflect.TypeOf(src)
	v := reflect.ValueOf(src)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		mjTag, ok := field.Tag.Lookup("mj")
		if !ok {
			continue
		}

		mv, err := ConvertToMatchableValue(v.FieldByName(field.Name).Interface())
		if err != nil {
			return nil, fmt.Errorf("The type of `%v` is inconvertible", mjTag)
		}
		expectation[mjTag] = mv
	}

	return expectation, nil
}

func Verify(t *testing.T, config *Config, expectations []Expectation, test func()) error {
	if len(expectations) <= 0 {
		return nil
	}

	v := &verifier{
		config:       config,
		expectations: expectations,
		ptr:          0,
	}

	passed, err := v.verify(test)
	if err != nil {
		return err
	}
	if !passed {
		t.Error("Not passed")
	}

	return nil
}

func (v *verifier) verify(test func()) (bool, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return false, err
	}
	defer func() {
		r.Close()
	}()

	stdoutEscaped := os.Stdout
	os.Stdout = w

	test()

	os.Stdout = stdoutEscaped
	w.Close()

	passed := false
	s := bufio.NewScanner(r)
	for s.Scan() {
		ok := v.match(s.Bytes())
		if !ok {
			continue
		}

		haveRemaining := v.consume()
		if !haveRemaining {
			passed = true
			break
		}
	}
	err = s.Err()
	if err != nil && err != io.EOF {
		return false, err
	}

	return passed, nil
}

func (v *verifier) match(src []byte) bool {
	actual, err := v.config.Parser(src)
	if err != nil {
		return false
	}

	for eK, eV := range v.expectations[v.ptr] {
		aV, ok := actual[eK]
		if !ok {
			return false
		}
		ok = eV.Match(aV)
		if !ok {
			return false
		}
	}

	return true
}

func (v *verifier) consume() bool {
	v.ptr++
	if v.ptr >= len(v.expectations) {
		return false
	}

	return true
}
