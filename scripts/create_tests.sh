find . -type f -name '*_test.go' -exec rm {} +
gotests -all -w ./*
