run:
	@go run main.go

testing:
	@go test ./... -v

update:
	@git add .
	@git commit -m "$(commit)"
	@git push -u origin master
	