version := 1.0.0

build-m1:
	docker buildx build --platform linux/amd64 --push -t ghcr.io/kla7016/pod-burner:$(version) -t ghcr.io/kla7016/pod-burner:latest .