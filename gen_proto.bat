@echo off
chcp 65001 > nul

rem 设置颜色为黄色
color 09

rem 清屏
cls



echo 正在生成 RPC 文件...

rem 检查是否传递了参数，如果没有传递则提示并退出
if "%1"=="" (
    echo 请指定要生成的服务路径作为参数（例如：用户服务传入user）作为参数。
    pause
    exit /b 1
)

rem 设置 proto 文件所在路径
set PROTO_PATH=D:\AkitaCode\Go\project\Panda-IM\Panda-IM-Server\service\%1\rpc

rem 遍历所有的 .proto 文件
for %%i in (%PROTO_PATH%\*.proto) do (

    echo 正在处理 %%i...

    rem 使用 goctl rpc protoc 命令生成代码，并指定输出路径为 rpc 文件夹
    goctl rpc protoc %%i --proto_path=%PROTO_PATH% --go_out=%PROTO_PATH% --go-grpc_out=%PROTO_PATH% --zrpc_out=%PROTO_PATH% --style=goZero >nul 2>nul

    rem 检查命令执行的返回值
    if errorlevel 1 (
        rem 如果返回值为 1，表示生成过程中出现错误
        echo 生成 %%i 时出错！
        pause
        exit /b 1
    ) else (
        rem 如果返回值为 0，表示生成成功
        echo %%~nxi 生成成功。
    )
)

echo.
echo pb 文件生成完成。
echo.
exit /b
