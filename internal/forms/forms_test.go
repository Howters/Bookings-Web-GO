package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST","/whatever",nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("Form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a","a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Has("test", r)
	if form.Valid(){
		t.Error("Form shows valid when test is missing")
	}

	r, _ =  http.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("test", "hello")
	postedData.Add("name", "thisisname")

	r.PostForm = postedData
	form.Has("test",r)
	form = New(r.PostForm)
	if !form.Valid(){
		t.Error("Form should be valid since there are values and not empty")
	}
} 

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("name", "a")
	r.PostForm = postedData

	form := New(r.PostForm)

	form.MinLength("name", 12, r)

	if form.Valid(){
		t.Error("Form has length less than 12 but still submited.")
	}

	r, _ =  http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("test", "asdfasdfa")
	r.PostForm = postedData
	form = New(r.PostForm)

	form.MinLength("test", 2, r)
	if !form.Valid(){
		t.Error("Form should be valid since min length has met")
	}
} 

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("email", "he")
	r.PostForm = postedData
	form := New(r.PostForm)

	form.IsEmail("email",r)

	if form.Valid(){
		t.Error("Form has invalid email but still valid.")
	}

	r, _ =  http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("email", "hello@gmail.com")
	r.PostForm = postedData

	form = New(r.PostForm)

	form.IsEmail("email", r)
	if !form.Valid(){
		t.Error("Form should be valid since email is vaalid")
	}
} 
