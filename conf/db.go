package conf

import "os"

var POSTGRES_URI string = os.Getenv("POSTGRES_URI")
