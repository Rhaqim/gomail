module gomailtesting

go 1.22.1

require (
    github.com/Rhaqim/gomail v1.0.0-beta
)

// add module from src folder
replace github.com/Rhaqim/gomail => ../