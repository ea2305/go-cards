package util

import "fmt"

var RegexPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"
var ApiIdFormat = fmt.Sprintf("/api/v1/decks/{id:%v}", RegexPattern)
