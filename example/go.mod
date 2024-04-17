module gomailtesting

go 1.22.0

require (
    gomail v1.4.0
)

// add module from src folder
replace gomail => ../