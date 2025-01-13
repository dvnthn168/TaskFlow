userpb:
				protoc --proto_path=proto/ \
				--go_out=gen --go_opt=paths=source_relative \
				--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=gen --grpc-gateway_opt=paths=source_relative,allow_delete_body=true,allow_repeated_fields_in_body=false,generate_unbound_methods=true \
				proto/userpb/user.proto

usersqlc:
				docker run --rm -v "${pwd}:/src" -w /src/sqlc/user-service sqlc/sqlc generate

migration_user: 
				migrate create -ext sql -dir migrations -seq create_users_table

migrateup:
				migrate -path migrations -database "postgres://vanThinh2512:VanThinh11168@localhost:5432/local_theBarber?sslmode=disable" up