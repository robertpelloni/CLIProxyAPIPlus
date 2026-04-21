<<<<<<< HEAD
# Deployment Instructions

1. **Clone the repository (with submodules):**
   ```bash
   git clone --recursive https://github.com/router-for-me/CLIProxyAPI
   cd CLIProxyAPI
   ```

2. **Generate UI & Build Binary:**
   Make sure you have both Node.js (v18+) and Go (v1.26+) installed.
   ```bash
   make ui-all
   ```
   This command will:
   - Run `npm install` and `npm run build` inside the `ui/` submodule.
   - Copy the compiled Single Page App (`index.html`) to `internal/managementasset/ui_build/management.html`.
   - Compile the Go proxy server embedding the UI statically into the final `cli-proxy-api` binary.

3. **Run the Proxy:**
   Copy `config.example.yaml` to `config.yaml` and configure your credentials.
   ```bash
   ./cli-proxy-api
   ```

4. **Access the Web Management UI:**
   Open your browser to `http://localhost:8317/management.html`. If you configured `remote-management.allow-remote: true`, you can access it via network IP. The UI will prompt you to enter the management key defined in your `config.yaml` to authenticate.

### Deploying via Docker
You can also deploy the application using the included `Dockerfile`.
The current multi-stage Docker build handles the UI compilation and Go binary build automatically without needing Node.js or Go installed on your host.
```bash
docker build -t cli-proxy-api .
docker run -p 8317:8317 cli-proxy-api
```
=======
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
>>>>>>> origin/jules-9238706904812453426-8fd51539
