BINARY_NAME=build
 
build:
	source .env && go build -o ${BINARY_NAME} main.go
 
run:
	source .env && go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}


clean:
	go clean
	rm ${BINARY_NAME}