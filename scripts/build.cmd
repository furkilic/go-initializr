@echo off
title %0
set BASE_DIR=%~dp0..

set EXTENSION=".exe"
if "%GOOS%"==""  goto build
if NOT "%GOOS%" == "windows" set EXTENSION=

:build
"%BASE_DIR%\gow.cmd" build -o "%BASE_DIR%\bin\go-initializr%EXTENSION%" "%BASE_DIR%\cmd\go-initializr\main.go"