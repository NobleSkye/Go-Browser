@echo off

:: Set Go project directory (change this to your project's directory if needed)
SET PROJECT_DIR=%cd%

:: Output directory
SET OUTPUT_DIR=%PROJECT_DIR%\..\builds

:: Make sure the output directory exists
echo Creating output directory at %OUTPUT_DIR%...
mkdir %OUTPUT_DIR%

:: Clean up old builds
echo Cleaning old builds...
del /f /q %OUTPUT_DIR%\SkyeBrowser-linux
del /f /q %OUTPUT_DIR%\SkyeBrowser-macos
del /f /q %OUTPUT_DIR%\SkyeBrowser-windows.exe

:: Build for Linux (x86_64)
echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o %OUTPUT_DIR%\SkyeBrowser-linux main.go

:: Build for macOS (x86_64)
echo Building for macOS...
set GOOS=darwin
set GOARCH=amd64
go build -o %OUTPUT_DIR%\SkyeBrowser-macos main.go

:: Build for Windows (x86_64)
echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o %OUTPUT_DIR%\SkyeBrowser-windows.exe main.go

echo Build completed!

:: Optionally, create a zip archive of the output files
:: echo Creating zip archive...
:: tar -czf %OUTPUT_DIR%\SkyeBrowser.tar.gz %OUTPUT_DIR%\*

:: End of script
