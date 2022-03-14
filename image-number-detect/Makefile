FUNC_NAME=image-number-detect
FUNC_IMAGE=355501594325.dkr.ecr.ap-northeast-3.amazonaws.com/func-$(FUNC_NAME)

build:
	docker build -t $(FUNC_IMAGE) --no-cache --progress plain .

run: build
	docker run -v ~/.aws-lambda-rie:/aws-lambda --entrypoint /aws-lambda/aws-lambda-rie -e AZ=Asia/Shanghai -p 9000:8080 $(FUNC_IMAGE) /main

test:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name": "Bob"}'

login:
	aws ecr get-login-password --region ap-northeast-3 | docker login --username AWS --password-stdin 355501594325.dkr.ecr.ap-northeast-3.amazonaws.com

deploy: login build
	docker push $(FUNC_IMAGE)
	aws lambda update-function-code --function-name $(FUNC_NAME) --image-uri $(FUNC_IMAGE):latest

invoke:
	aws lambda invoke --function-name $(FUNC_NAME) --payload '{ "name": "value" }' --cli-binary-format raw-in-base64-out test.log
