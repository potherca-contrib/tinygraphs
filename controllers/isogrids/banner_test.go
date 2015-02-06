package isogrids

import (
	"net/http"
	"testing"

	"github.com/taironas/route"
	tgTesting "github.com/taironas/tinygraphs/testing"
)

func TestBannerRandom(t *testing.T) {

	r := new(route.Router)
	r.HandleFunc("/isogrids/banner/random", BannerRandom)

	test := tgTesting.GenerateHandlerFunc(t, Hexa)
	for _, p := range tgTesting.GoodParams {
		recorder := test("/isogrids/banner/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}

	for _, p := range tgTesting.BadParams {
		recorder := test("/isogrids/banner/random", "GET", p, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
	}
}
