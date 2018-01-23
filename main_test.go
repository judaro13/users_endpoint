package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexEndpoint(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	IndexEndpoint(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("running app.."), w.Body.Bytes())
}

func TestValidateInputs(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("POST", "/users",
		strings.NewReader("name=juli&email=juli%40juli.com&password=2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824&phoneNumber=2362636&city=Any+one"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	resp := ValidateInputs(r)
	assert.Equal(t, nil, resp)

	// name error
	er, _ := http.NewRequest("POST", "/users",
		strings.NewReader(""))
	er.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	eresp := ValidateInputs(er)
	assert.Equal(t, ErrNameRequired, eresp)

	// email error
	emr, _ := http.NewRequest("POST", "/users",
		strings.NewReader("name=juli"))
	emr.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	emresp := ValidateInputs(emr)
	assert.Equal(t, ErrEmailRequired, emresp)

	// password error
	pr, _ := http.NewRequest("POST", "/users",
		strings.NewReader("name=juli&email=juli%40juli.com"))
	pr.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	presp := ValidateInputs(pr)
	assert.Equal(t, ErrPasswordRequired, presp)

	// phoneNumber error
	pnr, _ := http.NewRequest("POST", "/users",
		strings.NewReader("name=juli&email=juli%40juli.com&password=2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa74&city=Any+one"))
	pnr.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	pnresp := ValidateInputs(pnr)
	assert.Equal(t, ErrPhoneRequired, pnresp)
}

func TestReturnWithError(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()

	rwe := ReturnWithError(w, nil, 0)
	assert.Equal(t, false, rwe)

	rwee := ReturnWithError(w, ErrEmailRequired, 1)
	assert.Equal(t, true, rwee)
	assert.Equal(t, 1, w.Code)
}
