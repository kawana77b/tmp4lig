credits:
	@go install github.com/Songmu/gocredits/cmd/gocredits@latest
	@gocredits -skip-missing . > CREDITS

.PHONY: credits
