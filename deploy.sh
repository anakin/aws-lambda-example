#!/usr/bin/env bash

if [ -z $1 ] ; then
    echo "usage: ./deploy.sh func_name"
    exit
fi

DIR=$1
case $DIR in
    getsource)
        FUNC="PROD-OAUTH-SG-LAMBDA_GetSourceFunction"
        ;;
    *)
        echo "invalid func_name"
        ;;
esac

cd $DIR
ROLE="your role arn here"
GOARCH=amd64 GOOS=linux go build -o $DIR main.go
zip deployment.zip $DIR

aws lambda create-function --region ap-southeast-1 --function-name $FUNC --zip-file fileb://./deployment.zip --runtime go1.x --tracing-config Mode=Active --role $ROLE --handler $DIR
rm $DIR
rm deployment.zip
cd ..
