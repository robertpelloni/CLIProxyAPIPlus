@echo off
setlocal
title CLI Proxy API+
cd /d "%~dp0"

echo [CLI Proxy API+] Starting...
where go >nul 2>nul
if errorlevel 1 (
    echo [CLI Proxy API+] go not found. Please install it.
    pause
    exit /b 1
)

go run ./cmd/mcpdebug

if errorlevel 1 (
    echo [CLI Proxy API+] Exited with error code %errorlevel%.
    pause
)
endlocal
