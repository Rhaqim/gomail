module gomailtesting

go 1.22.0

require (
    github.com/rhaqim/gomail v1.0.0-beta
)

// add module from src folder
replace github.com/rhaqim/gomail => ../