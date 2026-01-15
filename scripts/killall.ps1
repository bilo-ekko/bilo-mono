# killall.ps1 - Kill all processes from bilo-mono apps by port (Windows)
# This script kills processes running on ports used by the monorepo apps

Write-Host "================================================" -ForegroundColor Blue
Write-Host "  Bilo Mono - Kill All Running Apps" -ForegroundColor Blue
Write-Host "================================================" -ForegroundColor Blue
Write-Host ""

# Define ports used by apps
$ports = @{
    "api-golang" = 8080
    "api-nestjs" = 3000
    "web-dashboard" = 4000
    "web-sdks-apps" = 4001
}

function Kill-Port {
    param(
        [string]$AppName,
        [int]$Port
    )
    
    Write-Host "Checking port $Port ($AppName)..." -ForegroundColor Yellow
    
    try {
        # Find processes using the port
        $connections = Get-NetTCPConnection -LocalPort $Port -ErrorAction SilentlyContinue
        
        if ($null -eq $connections) {
            Write-Host "  ✓ No process found on port $Port" -ForegroundColor Green
            return
        }
        
        foreach ($conn in $connections) {
            $process = Get-Process -Id $conn.OwningProcess -ErrorAction SilentlyContinue
            
            if ($null -ne $process) {
                Write-Host "  ✗ Found process: $($process.Name) (PID: $($process.Id))" -ForegroundColor Red
                
                try {
                    Stop-Process -Id $process.Id -Force
                    Write-Host "  ✓ Killed PID $($process.Id)" -ForegroundColor Green
                }
                catch {
                    Write-Host "  ✗ Failed to kill PID $($process.Id)" -ForegroundColor Red
                }
            }
        }
    }
    catch {
        Write-Host "  ✓ No process found on port $Port" -ForegroundColor Green
    }
}

# Kill all app processes
foreach ($app in $ports.Keys) {
    Kill-Port -AppName $app -Port $ports[$app]
    Write-Host ""
}

Write-Host "================================================" -ForegroundColor Blue
Write-Host "All processes checked and killed!" -ForegroundColor Green
Write-Host "================================================" -ForegroundColor Blue
Write-Host ""

# Additional cleanup
Write-Host "Cleaning up additional processes..." -ForegroundColor Yellow

# Kill node processes related to nest, next, vite
Get-Process node -ErrorAction SilentlyContinue | Where-Object {
    $_.Path -like "*bilo-mono*"
} | ForEach-Object {
    Stop-Process -Id $_.Id -Force
    Write-Host "  ✓ Killed node process (PID: $($_.Id))" -ForegroundColor Green
}

Write-Host ""
Write-Host "✓ All cleanup complete!" -ForegroundColor Green
Write-Host ""
