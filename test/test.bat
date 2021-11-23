@echo off
setlocal
pushd "%~dp0"


cd ..
type "%~f1"|go run main.go > test\res.txt

popd
endlocal
exit /b