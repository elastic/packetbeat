REM Configure GCC for either 32 or 64 bits
set MINGW_ARCH=64
IF NOT EXIST "%PROGRAMFILES(X86)%" (
    set MINGW_ARCH=32
)
set PATH=%WORKSPACE%\bin;C:\ProgramData\chocolatey\bin;C:\tools\mingw%MINGW_ARCH%\bin;%PATH%

REM Set the USERPROFILE to the previous location to fix issues with chocolatey in windows 2019
SET PREVIOUS_USERPROFILE=%USERPROFILE%
SET USERPROFILE=%OLD_USERPROFILE%

echo "Upgrade chocolatey to latest version"
choco upgrade chocolatey -y

curl --version >nul 2>&1 && (
    echo found curl
) || (
    choco install curl -y --no-progress
)

IF NOT EXIST C:\Python38\python.exe (
    REM Install python 3.11
    choco install python311 -y -r --no-progress || exit /b 1
)
python --version
where python

WHERE /q gcc
IF %ERRORLEVEL% NEQ 0 (
    REM Install mingw 5.3.0
    choco install mingw -y -r --no-progress --version 5.3.0 || exit /b 1
)
gcc --version
where gcc

REM Reset the USERPROFILE
SET USERPROFILE=%PREVIOUS_USERPROFILE%
