.PHONY: runServer runClient compileExp

COMPILE_SCRIPT=scripts/compile_protoc.sh
GO_RUN_SCRIPT=scripts/run_go.sh

runServer:
	cd src; \
	bash ../${GO_RUN_SCRIPT} -t gRPC/normal/server
runClient:
	cd src; \
	bash ../${GO_RUN_SCRIPT} -t gRPC/normal/client
compileExp:
	cd src; \
	bash ../${COMPILE_SCRIPT} -t gRPC/streaming
