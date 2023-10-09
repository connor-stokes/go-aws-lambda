.PHONY:  package deploy

all: package deploy

package:
	sam package --template-file api.yml --s3-bucket api-go-template-connorstokes --output-template-file packaged-template.yaml

deploy:
	sam deploy --template-file packaged-template.yaml --stack-name api-go-template-connorstokes --capabilities CAPABILITY_IAM
