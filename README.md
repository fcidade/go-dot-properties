# Golang .properties parser
## Why this exists
I made this project to learn basic lexing and parsing algorithms. Not (yet) recommended for production.

## Goals:
- [X] Parse .properties to map[string]string
- [X] Parse struct to .properties
- [X] Parse map[string]string to .properties
- [ ] Parse .properties to struct
- [ ] Handle types other than strings

## Usage

```go
    /*
        .properties example:
        website = https://en.wikipedia.org/
        language = English
    */

    // Maps
    mapped, err := goproperties.ParseToMap(textString)

    /*
        Output:
        map[string]string{
			"website":  "https://en.wikipedia.org/",
			"language": "English",
		}
    */

    // Structs
    type MyStruct struct {
        website string
        Language string `property:"language"`
    }
    
    var myStruct MyStruct
    err := goproperties.ParseToStruct(&textString)

    /*
        Output:
        MyStruct{
            website: "https://en.wikipedia.org/",
            Language: "English",
        }
    */

```
