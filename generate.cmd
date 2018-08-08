@ECHO OFF
rd /S /Q config
rd /S /Q environment
call openapi-generator generate -DapiTests=false -DapiDocs=false -DmodelTests=false -DmodelDocs=false -g go -o libdtapiconf -DpackageName=dtconf -i https://siz65484.live.dynatrace.com/api/config/v1/spec2.json
call openapi-generator generate -DapiTests=false -DapiDocs=false -DmodelTests=false -DmodelDocs=false -g go -o lidbdtapienv -DpackageName=dtenv -i https://siz65484.live.dynatrace.com/api/v1/spec2.json
