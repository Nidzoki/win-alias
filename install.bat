@echo off
set EXE_NAME=alias.exe
set DEST_DIR=C:\Windows

echo Building %EXE_NAME%...
go build -o %EXE_NAME% ./cmd/win-alias
if %ERRORLEVEL% neq 0 (
    echo Build failed.
    exit /b %ERRORLEVEL%
)

echo Installing to %DEST_DIR%...
copy /Y %EXE_NAME% %DEST_DIR%\
if %ERRORLEVEL% neq 0 (
    echo Copy failed. Ensure you are running as Administrator.
    exit /b %ERRORLEVEL%
)

echo.
echo win-alias installed successfully!
echo.
echo Usage:
echo   alias name="command"
echo   unalias name
echo.
echo To enable persistence across sessions, run:
echo   alias --setup
echo.
pause
