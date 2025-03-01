CPU_USAGE=2 RAM_USAGE_MB=1024 go run main.go


docker run --rm -e CPU_USAGE=2 -e RAM_USAGE_MB=1024 pod-burner:latest