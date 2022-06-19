APP=rental-vehicles
APP_EXE="./build/$(APP)"

serve:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=windows go build -o ${APP_EXE}

test:
	go test -cover -v ./...