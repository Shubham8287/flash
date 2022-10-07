module flash

go 1.19

replace flash/app => ./app

require (
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20220927170352-d9d178bc13c6 // indirect
)

require github.com/google/uuid v1.3.0 // indirect
