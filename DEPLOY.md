# DEPLOY.md

## Deployment Instructions

### Prerequisites
- Go 1.26 or higher
- Node.js (for building the UI submodule)
- Make (optional, but recommended)

### Build Instructions

1. **Clone the repository and submodules:**
   ```bash
   git clone --recursive https://github.com/router-for-me/CLIProxyAPIPlus.git
   cd CLIProxyAPIPlus
   ```

2. **Build the Management UI:**
   Ensure you have installed npm.
   ```bash
   cd ui
   npm install
   npm run build
   cd ..
   ```
   The `ui/dist/index.html` file will be generated. The Go compiler will automatically embed this file into the binary.

3. **Build the Proxy Binary:**
   ```bash
   go build -o cli-proxy-api ./cmd/server
   ```

### Running the Server
Use the example config as a base:
```bash
cp config.example.yaml config.yaml
```

Run the compiled executable:
```bash
./cli-proxy-api
```
The Management Center UI will now be available automatically on `http://localhost:8317/`.
