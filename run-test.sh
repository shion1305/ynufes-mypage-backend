gcloud beta emulators firestore start --quiet --host-port localhost:8020 &
sleep 5
go test -v ./...