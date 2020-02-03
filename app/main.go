package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	result := DoSomething()
	fmt.Println(result.Glossary.Title)

	bodyJson, err := json.Marshal(result)
	if err != nil {
		return
	}

	var r struct {
		Glossary struct {
			Title string
			GlossDiv struct {
				Title string
				GlossList struct {
					GlossEntry struct {
						ID string
						SortAs string
						GlossTerm string
						Acronym string
						Abbrev string
						GlossDef struct {
							Para string
							GlossSeeAlso []string

						}
						GlossSee string
					}
				}
			}
		}
	}
	err = DecodeJSONResponse(bodyJson, &r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	fmt.Println(r)

}

func DoSomething() (result struct {
	Glossary struct {
		Title string
		GlossDiv struct {
			Title string
			GlossList struct {
				GlossEntry struct {
					ID string
					SortAs string
					GlossTerm string
					Acronym string
					Abbrev string
					GlossDef struct {
						Para string
						GlossSeeAlso []string

					}
					GlossSee string
				}
			}
		}
	}
}) {
	a := `{"glossary":{"title":"example glossary","GlossDiv":{"title":"S","GlossList":{"GlossEntry":{"ID":"SGML","SortAs":"SGML","GlossTerm":"Standard Generalized Markup Language","Acronym":"SGML","Abbrev":"ISO 8879:1986","GlossDef":{"para":"A meta-markup language, used to create markup languages such as DocBook.","GlossSeeAlso":["GML","XML"]},"GlossSee":"markup"}}}}}`
	err := DecodeJSONResponse([]byte(a), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func DecodeJSONResponse(a []byte, result interface{}) error {
	return json.Unmarshal(a, &result)
}
