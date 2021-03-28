package file

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {

	f1 := file{name: "temp.txt", content: "temp content"}
	f2 := file{name: "garbage.md", content: "junk data"}
	f3 := file{name: "temp.out", content: "binary data"}
	f4 := file{name: "confidential.html", content: "confidential: we will hold a party on this saturday"}

	fd1 := folder{name: "butn", component: []component{&f1, &f2}}
	home := folder{name: "home", component: []component{&fd1, &f3, &f4}}
	// home contains fd1
	var searchCase = []struct {
		name       string
		keywords   string
		folder     component
		wantResult string
		wantError  bool
	}{
		{name: "confidential", keywords: "confidential", folder: &home, wantResult: "home/confidential.html", wantError: false},
		{name: "temp", keywords: "temp", folder: &home, wantResult: "home/butn/temp.txt", wantError: false},
		{name: "password", keywords: "password", folder: &home, wantResult: "", wantError: true},
	}

	for _, sCase := range searchCase {
		t.Run(sCase.name, func(t *testing.T) {
			result, err := sCase.folder.search(sCase.keywords)
			if (err != nil) != sCase.wantError {
				t.Errorf("name: %s, wantError=%v, gotError=%v", sCase.name, sCase.wantError, err)
				return
			}
			if !reflect.DeepEqual(result, sCase.wantResult) {
				t.Errorf("name: %s, wantReuslt=%v, gotResult=%v", sCase.name, sCase.wantResult, result)
				return
			}
		})
	}
}
