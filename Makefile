MOCKS_DESTINATION=mocks
.PHONY: mocks
mocks: internal/service/interface.go internal/controller/interface.go
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done

.PHONY: cover
cover:
	go test --short -count=1 -race --coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

install_swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

update_swagger:
	swag init -g swagger_info.go --dir ./internal --parseInternal --parseDependency
