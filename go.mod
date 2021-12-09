module github.com/TharpHuang/gofar

go 1.16

require (
	github.com/AlecAivazis/survey/v2 v2.3.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.6.0
	github.com/spf13/cobra v1.2.1
	golang.org/x/sys v0.0.0-20211110154304-99a53858aa08 // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	xorm.io/xorm v1.0.3
)

replace github.com/TharpHuang/gofar => ./
