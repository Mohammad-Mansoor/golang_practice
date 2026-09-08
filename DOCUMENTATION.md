#to start and initialize a go project we need to enter this command
go mod init github.com/githubusername/projectname


# for running a project through cmd in windows use below command
go run main.go // this will build and run the project but it is only used in development Environment 

# to build the project for production like 
go build -o bin/main main.go // main.go is the name of file from which the build will ba taken and the main is the name of build binary file 


#to have something like pakcage.json in express or nestjs app we can use make tool to add our commands so with the help of that we able to use it and automate it 

build:
    go build -o bin/api cmd/api/main.go

run: bin/api
