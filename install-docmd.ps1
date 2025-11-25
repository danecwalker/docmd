# Colors
$Green = "`e[32m"
$Red   = "`e[31m"
$NC    = "`e[0m"

# Spinner function
function Start-Spinner {
    param([string]$Message)

    Write-Host -NoNewline "$Message "
    $spinner = "/|\-"
    $i = 0

    while ($script:spinRunning) {
        $char = $spinner[$i % $spinner.Length]
        Write-Host -NoNewline "`b$char"
        Start-Sleep -Milliseconds 100
        $i++
    }
    Write-Host "`b "
}

# Detect OS (Windows only)
$OS = "windows"

# Detect architecture
$ARCH = $(if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" })

# Form GitHub download URL
$URL = "https://github.com/danecwalker/docmd/releases/latest/download/docmd-$OS-$ARCH.exe"

# Download path
$Output = "docmd.exe"

Write-Host "Downloading docmd for $OS-$ARCH..."

$script:spinRunning = $true
Start-Job { Start-Spinner -Message "Downloading..." } | Out-Null

try {
    Invoke-WebRequest -Uri $URL -OutFile $Output -UseBasicParsing -ErrorAction Stop
    $script:spinRunning = $false
    Start-Sleep -Milliseconds 200
    Write-Host "${Green}Download complete!$NC"
}
catch {
    $script:spinRunning = $false
    Write-Host "${Red}Failed to download docmd.$NC"
    exit 1
}

# Install directory
$InstallDir = "C:\Program Files\docmd"

if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir | Out-Null
}

# Move binary
Write-Host "${Green}Installing docmd...$NC"
Move-Item -Force $Output "$InstallDir\docmd.exe"

# Add to PATH (if missing)
$CurrentPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)

if ($CurrentPath -notlike "*$InstallDir*") {
    Write-Host "Updating system PATH..."
    $NewPath = "$CurrentPath;$InstallDir"
    setx /M PATH "$NewPath" | Out-Null
    Write-Host "${Green}PATH updated! (Restart terminal to apply)$NC"
}

Write-Host "${Green}Installation complete. You can now use 'docmd'.$NC"
