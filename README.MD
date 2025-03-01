# Pod Burner 🔥

**Pod Burner** is a simple tool for simulating **CPU and RAM usage**, useful for testing **auto-scaling**, **load balancing**, or **performance benchmarking** in Kubernetes or Docker environments.

---

## 🐳 Running with Docker
To run the application using Docker, use:
```sh
docker run --rm -e CPU_USAGE=2 -e RAM_USAGE_MB=1024 ghcr.io/kla7016/pod-burner:latest
```

### ⚙️ Explanation:
- `-e CPU_USAGE=2` → Uses 2 CPU cores. (Default: `1 CPU core`) **Note:** CPU usage must be an integer; decimal values are not allowed.
- `-e RAM_USAGE_MB=1024` → Allocates 1024MB (1GB) of RAM. (Default: `100MB`)
---

## 🚀 Running in Development Mode
To run the application directly using Go, use the following command:
```sh
CPU_USAGE=2 RAM_USAGE_MB=1024 go run main.go
```

### ⚙️ Explanation:
- `CPU_USAGE=2` → Uses **2 CPU cores** (Default: `1`)
- `RAM_USAGE_MB=1024` → Allocates **1024MB (1GB) of RAM** (Default: `100MB`)

---

## 📌 Default Values
If **CPU_USAGE** or **RAM_USAGE_MB** are not specified, the application will use the following default values:
- **`CPU_USAGE=1`** → Uses 1 CPU core.
- **`RAM_USAGE_MB=100`** → Allocates 100MB of RAM.


---

## 🔥 Use Cases
✅ **Testing Kubernetes Auto Scaling (HPA, VPA).**  
✅ **Simulating high CPU/RAM usage for load testing.**  
✅ **Benchmarking system performance.**

