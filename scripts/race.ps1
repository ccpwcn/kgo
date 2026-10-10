#Requires -Version 5.1
#
# ============================================================================
# 重要说明（务必先读）：
#   * Visual Studio 自带的 clang 编译器【不能】用于 go test -race。
#     原因：Go 在 Windows 上的 -race 走 GNU 风格外部链接，会向编译器传 GNU ld 参数
#     （-Wl,-T,...ld、-lsynchronization、--compress-debug-sections 等），VS clang 是 MSVC
#     模式，链接阶段必然失败（LNK4044 / LNK1143）；MSVC 的 cl.exe 同样不可用。
#   * 必须使用 MSYS2 中的 mingw-w64 gcc，并且在【MSYS2 终端】或【PowerShell 7 (pwsh)】中调用，
#     以保证 UTF-8 环境与正确的 GNU 工具链。
#       - 安装： winget install MSYS2.MSYS2
#                然后在 MSYS2 终端执行： pacman -S mingw-w64-ucrt-x86_64-gcc
#       - gcc 路径通常为： C:\msys64\ucrt64\bin\gcc.exe
#   * 因此本脚本只自动探测 mingw-w64 gcc，不再探测 VS clang（避免误导）。
#   * 不推荐用 Windows 自带的 PowerShell 5.1 运行本脚本：它既通常没有 gcc，也不是目标环境。
# ============================================================================
<#
.SYNOPSIS
    自动探测本机可用的 mingw-w64 gcc，启用 cgo 并运行 Go 竞态检测（go test -race）。

.DESCRIPTION
    Go 的 -race 依赖 cgo，且在 Windows 上只支持 GNU 工具链（mingw-w64 gcc）。
    本脚本按以下优先级探测编译器：
        1) -Compiler 显式指定的路径/名称
        2) PATH 中的 gcc
        3) 常见 mingw 安装位置的 gcc（MSYS2 / TDM-GCC / w64devkit / chocolatey 等）
    找到后设置 CC 与 CGO_ENABLED=1，然后运行 go test -race。

    说明：Visual Studio 自带的 clang（MSVC 模式）与 cl.exe 都无法用于 -race，故本脚本不探测它们。
    如确有可用的 GNU 目标 clang，请用 -Compiler 显式指定。

    本文件为 UTF-8（带 BOM）编码，可在 Windows PowerShell 5.1 与 PowerShell 7 下正确解析。

.PARAMETER Compiler
    强制指定编译器（可执行文件完整路径，或 PATH 中的命令名，如 gcc）。

.PARAMETER TestTarget
    传给 go test 的包目标，默认 ./...

.PARAMETER ExtraGoTestArgs
    追加到 go test 的额外参数（位于包目标之前），例如 -v、-run、-count=1。

.PARAMETER NoRun
    只探测并设置/打印环境变量，不运行测试。
    若希望环境变量在当前会话保留，请用点源方式运行：  . .\scripts\race.ps1 -NoRun

.EXAMPLE
    # 直接跑全量竞态检测（自动探测 gcc）
    .\scripts\race.ps1

.EXAMPLE
    # 只跑本包的测试并显示详细输出
    .\scripts\race.ps1 -ExtraGoTestArgs -v -TestTarget .

.EXAMPLE
    # 指定编译器
    .\scripts\race.ps1 -Compiler 'C:\msys64\ucrt64\bin\gcc.exe'

.EXAMPLE
    # 在当前会话中导出 CC/CGO_ENABLED，之后自行运行 go 命令
    . .\scripts\race.ps1 -NoRun
#>
[CmdletBinding()]
param(
    [string]$Compiler,
    [string[]]$TestTarget = @('./...'),
    [string[]]$ExtraGoTestArgs = @(),
    [switch]$NoRun
)

$ErrorActionPreference = 'Stop'

function Write-Info([string]$m) { Write-Host "[race] $m" -ForegroundColor Cyan }
function Write-Ok([string]$m)   { Write-Host "[race] $m" -ForegroundColor Green }
function Write-Warn2([string]$m){ Write-Host "[race] $m" -ForegroundColor Yellow }

# --- 探测：PATH 或常见位置的 mingw-w64 gcc ---
function Find-Gcc {
    $cmd = Get-Command gcc -ErrorAction SilentlyContinue
    if ($cmd) { return $cmd.Source }
    $paths = @(
        'C:\msys64\ucrt64\bin\gcc.exe',
        'C:\msys64\mingw64\bin\gcc.exe',
        'C:\TDM-GCC-64\bin\gcc.exe',
        'C:\mingw64\bin\gcc.exe',
        'C:\w64devkit\bin\gcc.exe',
        'C:\ProgramData\chocolatey\bin\gcc.exe'
    )
    foreach ($p in $paths) { if (Test-Path $p) { return $p } }
    return $null
}

# ======================= 主流程 =======================

$ccPath = $null

if ($Compiler) {
    if (Test-Path $Compiler) {
        $ccPath = $Compiler
    } else {
        $cmd = Get-Command $Compiler -ErrorAction SilentlyContinue
        if ($cmd) { $ccPath = $cmd.Source }
    }
    if (-not $ccPath) { throw "指定的编译器未找到：$Compiler" }
} else {
    $ccPath = Find-Gcc
}

if (-not $ccPath) {
    Write-Warn2 "未检测到 mingw-w64 gcc（Go 的 -race 在 Windows 上只能用 gcc，VS clang/cl 均不可用）。"
    Write-Host @'
请安装 mingw-w64 gcc 后重试（任选其一）：
  1) MSYS2（推荐）：
       winget install MSYS2.MSYS2
       然后在 MSYS2 终端执行： pacman -S mingw-w64-ucrt-x86_64-gcc
       gcc 路径通常为 C:\msys64\ucrt64\bin\gcc.exe
  2) w64devkit： 解压到 C:\w64devkit 即可
安装后直接运行： .\scripts\race.ps1
或手动指定：     .\scripts\race.ps1 -Compiler 'C:\msys64\ucrt64\bin\gcc.exe'
提示：请在 MSYS2 终端或 PowerShell 7 (pwsh) 中运行本脚本。
'@ -ForegroundColor Yellow
    exit 2
}

$env:CC = $ccPath
$env:CGO_ENABLED = '1'
# MSYS2/mingw 的 gcc 运行时依赖同目录下的 DLL（libwinpthread、libgcc 等），
# 若其 bin 目录不在 PATH 中，cgo 调用会失败（cgo.exe: exit status 2），故在此加入 PATH。
$ccDir = Split-Path -Parent $ccPath
if ($ccDir -and (($env:Path -split ';') -notcontains $ccDir)) {
    $env:Path = $ccDir + ';' + $env:Path
    Write-Info "已将编译器目录加入 PATH：$ccDir"
}
# cgo 会按空格拆分 CC 的值，因此含空格的路径必须内嵌双引号，
# 否则 "C:\Program Files\..." 会被拆成编译器 "C:\Program" 而找不到。
if ($ccPath -match '\s') {
    $env:CC = '"' + $ccPath + '"'
}
Write-Ok "CC = $env:CC"
Write-Ok "CGO_ENABLED = 1"

try {
    $ver = (& $ccPath --version 2>$null | Select-Object -First 1)
    if ($ver) { Write-Info "编译器版本：$ver" }
} catch {
    Write-Warn2 "无法执行 '$ccPath --version'，将继续尝试运行测试。"
}

# 快速确认 cgo 是否真的启用
$envCheck = (& go env CGO_ENABLED 2>$null)
Write-Info "go env CGO_ENABLED = $envCheck"

if ($NoRun) {
    Write-Info "已设置环境变量（-NoRun）。要在当前会话保留，请用点源方式运行： . .\scripts\race.ps1 -NoRun"
    return
}

$goArgs = @('test', '-race') + $ExtraGoTestArgs + $TestTarget
Write-Info ("运行： go " + ($goArgs -join ' '))
& go @goArgs
exit $LASTEXITCODE
