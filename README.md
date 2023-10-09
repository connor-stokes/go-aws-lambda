# summary of go lambda function

```bash
go test -v
```

```bash
GOARCH=amd64 GOOS=linux go build  -o main main.go
```

```bash
sam package --template-file template.yaml --s3-bucket <your-s3-bucket> --output-template-file packaged-template.yaml
```

```bash
sam deploy --template-file packaged-template.yaml --stack-name <your-stack-name> --capabilities CAPABILITY_IAM
```
