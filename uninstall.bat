@echo off
set EXE_NAME=alias.exe
set DEST_DIR=C:\Windows

echo win-alias Uninstaller
echo =====================

:: 1. Disable AutoRun persistence
if exist "%DEST_DIR%\%EXE_NAME%" (
    echo Disabling AutoRun persistence...
    "%DEST_DIR%\%EXE_NAME%" --disable
)

:: 2. Remove binary
if exist "%DEST_DIR%\%EXE_NAME%" (
    echo Removing %EXE_NAME% from %DEST_DIR%...
    del /F /Q "%DEST_DIR%\%EXE_NAME%"
    if %ERRORLEVEL% neq 0 (
        echo Failed to delete %EXE_NAME%. Ensure you are running as Administrator.
        pause
        exit /b %ERRORLEVEL%
      )
) else (
    echo %EXE_NAME% not found in %DEST_DIR%.
)

echo.
echo Optional: Cleaning up Registry aliases storage...
reg delete "HKCU\Software\win-alias" /f >nul 2>&1

echo.
echo win-alias uninstalled successfully!
pause
