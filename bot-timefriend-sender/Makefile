FUNC_NAME=bot-timefriend-sender
FUNC_IMAGE=355501594325.dkr.ecr.ap-northeast-3.amazonaws.com/func-$(FUNC_NAME)

run:
	docker run -v ~/.aws-lambda-rie:/aws-lambda --entrypoint /aws-lambda/aws-lambda-rie  -p 9000:8080 $(FUNC_IMAGE) /main
test:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name": "Bob"}'

build:
	docker build -t $(FUNC_IMAGE) .
	docker push $(FUNC_IMAGE)

deploy: build
	aws lambda update-function-code --function-name $(FUNC_NAME) --image-uri $(FUNC_IMAGE):latest

invoke:
	aws lambda invoke --function-name $(FUNC_NAME) --payload '{ "name": "value" }' --cli-binary-format raw-in-base64-out test.log
