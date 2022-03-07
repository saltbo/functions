
run:
	docker run -v ~/.aws-lambda-rie:/aws-lambda --entrypoint /aws-lambda/aws-lambda-rie  -p 9000:8080 myfunction:latest /main
test:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name": "Bob"}'

build:
	docker build -t 355501594325.dkr.ecr.ap-northeast-3.amazonaws.com/func-todo-reminder .
	docker push 355501594325.dkr.ecr.ap-northeast-3.amazonaws.com/func-todo-reminder

deploy: build
	aws lambda update-function-code --function-name todo-reminder --image-uri 355501594325.dkr.ecr.ap-northeast-3.amazonaws.com/func-todo-reminder:latest

invoke:
	aws lambda invoke --function-name todo-reminder --payload '{ "name": "value" }' --cli-binary-format raw-in-base64-out test.log
