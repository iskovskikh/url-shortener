package save

import "testing"

func TestSaveHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		url       string
		respError string
		mockError error
	}{
		{
			name:  "success",
			alias: "test_alias",
			url:   "https://google.com",
		},
		{
			name:  "empty alias",
			alias: "",
			url:   "https://google.com",
		},

	}

}
