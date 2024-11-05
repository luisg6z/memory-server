gen-lobby:
	@protoc \
	--proto_path=. "lobby.proto" \
	--go_out=proto/lobby --go_opt=paths=source_relative \
	--go-grpc_out=proto/lobby --go-grpc_opt=paths=source_relative

gen-game:
	@protoc \
	--proto_path=. "game.proto" \
	--go_out=proto/game --go_opt=paths=source_relative \
	--go-grpc_out=proto/game --go-grpc_opt=paths=source_relative