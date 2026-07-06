local Players = game:GetService("Players")
local RunService = game:GetService("RunService")
local UIS = game:GetService("UserInputService")
local TweenService = game:GetService("TweenService")
local HttpService = game:GetService("HttpService")
local LP = Players.LocalPlayer




-- // THEME (ENVY STYLE - BLACK + WHITE)
local BG         = Color3.fromRGB(0, 0, 0)
local CARD_BG    = Color3.fromRGB(14, 14, 14)
local CARD_HOV   = Color3.fromRGB(24, 24, 24)
local BORDER     = Color3.fromRGB(40, 40, 40)
local BORDER2    = Color3.fromRGB(70, 70, 70)
local WHITE      = Color3.fromRGB(255, 255, 255)
local DIM        = Color3.fromRGB(160, 160, 160)
local KB_BG      = Color3.fromRGB(10, 10, 10)

-- // STATE
local State = {
    autoBatToggled = false,
    hittingCooldown = false,
    guiVisible = true,
    mobileMode = false,
}

local Keys = {
    autoBat = Enum.KeyCode.X,
    autoBatType = "Keyboard", -- "Keyboard" or "Gamepad"
    guiHide = Enum.KeyCode.LeftControl,
}

local h, hrp = nil, nil

-- // CONFIG
local function saveConfig()
    local cfg = {
        autoBatKey = Keys.autoBat.Name,
        autoBatKeyType = Keys.autoBatType,
        mobileMode = State.mobileMode,
    }
    pcall(function() 
        writefile("EnvyAutoBatDesyncConfig.json", HttpService:JSONEncode(cfg)) 
    end)
end

local function loadConfig()
    local hasFile = false
    pcall(function() hasFile = isfile("EnvyAutoBatDesyncConfig.json") end)
    if not hasFile then return end
    
    local ok, cfg = pcall(function() 
        return HttpService:JSONDecode(readfile("EnvyAutoBatDesyncConfig.json")) 
    end)
    if not ok or not cfg then return end
    
    if cfg.autoBatKey and Enum.KeyCode[cfg.autoBatKey] then
        Keys.autoBat = Enum.KeyCode[cfg.autoBatKey]
    end
    
    if cfg.autoBatKeyType then
        Keys.autoBatType = cfg.autoBatKeyType
    end
    
    if cfg.mobileMode ~= nil then 
        State.mobileMode = cfg.mobileMode 
    end
end

loadConfig()

-- // CLEANUP OLD GUI
for _, name in pairs({"EnvyAutoBatDesyncGUI", "MwvaneNewaBatDesyncGUI", "PhazeAutoBatDesyncGUI"}) do
    local old = game:GetService("CoreGui"):FindFirstChild(name)
    if old then old:Destroy() end
end

-- // GUI
local gui = Instance.new("ScreenGui")
gui.Name="EnvyAutoBatDesyncGUI"; gui.ResetOnSpawn=false; gui.DisplayOrder=10
gui.IgnoreGuiInset=true; gui.Parent=LP:WaitForChild("PlayerGui")

local main = Instance.new("Frame",gui)
main.Name="Main"; main.Size=UDim2.new(0,240,0,80); main.Position=UDim2.new(0.5,-120,0.5,-40)
main.BackgroundColor3=BG; main.BorderSizePixel=0; main.Active=true; main.ClipsDescendants=true
Instance.new("UICorner",main).CornerRadius=UDim.new(0,12)
local mainStroke = Instance.new("UIStroke",main); mainStroke.Color=BORDER; mainStroke.Thickness=1

do
    local dragging, dragInput, dragStart, mainStart = false, nil, nil, nil
    main.InputBegan:Connect(function(inp)
        if inp.UserInputType==Enum.UserInputType.MouseButton1 or inp.UserInputType==Enum.UserInputType.Touch then
            dragging=true; dragStart=inp.Position; mainStart=main.Position
            inp.Changed:Connect(function()
                if inp.UserInputState==Enum.UserInputState.End then dragging=false end
            end)
        end
    end)
    main.InputChanged:Connect(function(inp)
        if inp.UserInputType==Enum.UserInputType.MouseMovement or inp.UserInputType==Enum.UserInputType.Touch then dragInput=inp end
    end)
    UIS.InputChanged:Connect(function(inp)
        if inp==dragInput and dragging then
            local dx = inp.Position.X - dragStart.X
            local dy = inp.Position.Y - dragStart.Y
            main.Position = UDim2.new(mainStart.X.Scale, mainStart.X.Offset+dx, mainStart.Y.Scale, mainStart.Y.Offset+dy)
        end
    end)
end

-- // HEADER
local header = Instance.new("Frame",main)
header.Size=UDim2.new(1,0,0,40); header.BackgroundColor3=Color3.fromRGB(8, 8, 8); header.BorderSizePixel=0; header.ZIndex=5
Instance.new("UICorner",header).CornerRadius=UDim.new(0,12)
local headerPatch = Instance.new("Frame",header)
headerPatch.Size=UDim2.new(1,0,0,12); headerPatch.Position=UDim2.new(0,0,1,-12)
headerPatch.BackgroundColor3=Color3.fromRGB(8, 8, 8); headerPatch.BorderSizePixel=0; headerPatch.ZIndex=4

local headerDiv = Instance.new("Frame",header)
headerDiv.Size=UDim2.new(1,0,0,1); headerDiv.Position=UDim2.new(0,0,1,-1)
headerDiv.BackgroundColor3=BORDER; headerDiv.BorderSizePixel=0; headerDiv.ZIndex=6

local titleLbl = Instance.new("TextLabel",header)
titleLbl.Size=UDim2.new(0,130,1,0); titleLbl.Position=UDim2.new(0,14,0,0)
titleLbl.BackgroundTransparency=1; titleLbl.Text="ENVY AUTO BAT"
titleLbl.TextColor3=WHITE; titleLbl.Font=Enum.Font.GothamBlack; titleLbl.TextSize=10
titleLbl.TextXAlignment=Enum.TextXAlignment.Left; titleLbl.ZIndex=6

-- Mobile/PC Toggle Button
local modeToggleBtn = Instance.new("TextButton", header)
modeToggleBtn.Size = UDim2.new(0, 50, 0, 24)
modeToggleBtn.Position = UDim2.new(1, -82, 0.5, -12)
modeToggleBtn.BackgroundColor3 = KB_BG
modeToggleBtn.BorderSizePixel = 0
modeToggleBtn.Text = State.mobileMode and "PC" or "MOBILE"
modeToggleBtn.TextColor3 = WHITE
modeToggleBtn.Font = Enum.Font.GothamBold
modeToggleBtn.TextSize = 8
modeToggleBtn.ZIndex = 10
Instance.new("UICorner", modeToggleBtn).CornerRadius = UDim.new(0, 5)
local modeToggleStroke = Instance.new("UIStroke", modeToggleBtn)
modeToggleStroke.Color = BORDER2
modeToggleStroke.Thickness = 1

local closeBtn = Instance.new("TextButton",main)
closeBtn.Size=UDim2.new(0,24,0,24); closeBtn.Position=UDim2.new(1,-32,0,8)
closeBtn.BackgroundColor3=Color3.fromRGB(18, 18, 18); closeBtn.BorderSizePixel=0
closeBtn.Text="X"; closeBtn.TextColor3=WHITE; closeBtn.Font=Enum.Font.GothamBlack
closeBtn.TextSize=12; closeBtn.ZIndex=10
Instance.new("UICorner",closeBtn).CornerRadius=UDim.new(0,6)
local closeBtnStroke = Instance.new("UIStroke",closeBtn); closeBtnStroke.Color=BORDER
closeBtn.MouseEnter:Connect(function()
    TweenService:Create(closeBtn,TweenInfo.new(0.1),{BackgroundColor3=Color3.fromRGB(40,40,40)}):Play()
end)
closeBtn.MouseLeave:Connect(function()
    TweenService:Create(closeBtn,TweenInfo.new(0.1),{BackgroundColor3=Color3.fromRGB(18,18,18)}):Play()
end)
closeBtn.MouseButton1Click:Connect(function()
    State.autoBatToggled = false
    gui:Destroy()
end)

-- // CONTENT AREA
local content = Instance.new("Frame",main)
content.Size=UDim2.new(1,-20,1,-50); content.Position=UDim2.new(0,10,0,44)
content.BackgroundTransparency=1; content.BorderSizePixel=0

-- // AUTO BAT TOGGLE ROW
local toggleRow = Instance.new("Frame",content)
toggleRow.Size=UDim2.new(1,0,0,32); toggleRow.BackgroundColor3=CARD_BG; toggleRow.BorderSizePixel=0
Instance.new("UICorner",toggleRow).CornerRadius=UDim.new(0,7)
local toggleRowStroke = Instance.new("UIStroke",toggleRow); toggleRowStroke.Color=BORDER
local toggleLbl = Instance.new("TextLabel",toggleRow)
toggleLbl.Size=UDim2.new(0,80,1,0); toggleLbl.Position=UDim2.new(0,10,0,0)
toggleLbl.BackgroundTransparency=1; toggleLbl.Text="Auto Bat"; toggleLbl.TextColor3=WHITE
toggleLbl.Font=Enum.Font.GothamBold; toggleLbl.TextSize=11; toggleLbl.TextXAlignment=Enum.TextXAlignment.Left

-- Keybind button
local keyBtn = Instance.new("TextButton",toggleRow)
keyBtn.Size=UDim2.new(0,44,0,20); keyBtn.Position=UDim2.new(1,-90,0.5,-10)
keyBtn.BackgroundColor3=KB_BG; keyBtn.BorderSizePixel=0; keyBtn.Text=Keys.autoBat.Name
keyBtn.TextColor3=WHITE; keyBtn.Font=Enum.Font.GothamBold; keyBtn.TextSize=8; keyBtn.ZIndex=11
Instance.new("UICorner",keyBtn).CornerRadius=UDim.new(0,5)
local ks = Instance.new("UIStroke",keyBtn); ks.Color=BORDER2; ks.Thickness=1
local kListening = false; local kConn

local function stopListen(key)
    kListening = false
    if kConn then 
        kConn:Disconnect()
        kConn = nil 
    end
    TweenService:Create(ks, TweenInfo.new(0.1), {Color = BORDER2}):Play()
    keyBtn.TextColor3 = WHITE
    if key then 
        keyBtn.Text = key.Name
        Keys.autoBat = key
        saveConfig()
    end
end

keyBtn.MouseButton1Click:Connect(function()
    if kListening then 
        stopListen(nil)
        return 
    end
    
    kListening = true
    keyBtn.Text = "..."
    keyBtn.TextColor3 = WHITE
    TweenService:Create(ks, TweenInfo.new(0.1), {Color = WHITE}):Play()
    
    kConn = UIS.InputBegan:Connect(function(inp)
        if not kListening then return end
        
        -- Accept both keyboard and gamepad inputs
        if inp.UserInputType == Enum.UserInputType.Keyboard then
            if inp.KeyCode == Enum.KeyCode.Escape then 
                stopListen(nil)
                return 
            end
            Keys.autoBatType = "Keyboard"
            stopListen(inp.KeyCode)
        elseif inp.UserInputType == Enum.UserInputType.Gamepad1 then
            Keys.autoBatType = "Gamepad"
            stopListen(inp.KeyCode)
        end
    end)
end)

-- Toggle pill
local pillBg = Instance.new("Frame",toggleRow)
pillBg.Size=UDim2.new(0,36,0,19); pillBg.Position=UDim2.new(1,-46,0.5,-9.5)
pillBg.BackgroundColor3=Color3.fromRGB(16, 16, 16); pillBg.BorderSizePixel=0; pillBg.ZIndex=8
Instance.new("UICorner",pillBg).CornerRadius=UDim.new(0,10)
local pStroke = Instance.new("UIStroke",pillBg); pStroke.Color=BORDER2; pStroke.Thickness=1
local dot = Instance.new("Frame",pillBg)
dot.Size=UDim2.new(0,13,0,13); dot.Position=UDim2.new(0,2,0.5,-6.5)
dot.BackgroundColor3=Color3.fromRGB(45, 45, 45); dot.BorderSizePixel=0; dot.ZIndex=9
Instance.new("UICorner",dot).CornerRadius=UDim.new(0,4)

-- Mobile button (full width)
local mobileBtn = Instance.new("TextButton", toggleRow)
mobileBtn.Size = UDim2.new(1, -12, 1, -6)
mobileBtn.Position = UDim2.new(0, 6, 0, 3)
mobileBtn.BackgroundColor3 = KB_BG
mobileBtn.BorderSizePixel = 0
mobileBtn.Text = "AUTO BAT"
mobileBtn.TextColor3 = WHITE
mobileBtn.Font = Enum.Font.GothamBold
mobileBtn.TextSize = 11
mobileBtn.Visible = State.mobileMode
mobileBtn.ZIndex = 12
Instance.new("UICorner", mobileBtn).CornerRadius = UDim.new(0, 5)
local mobileBtnStroke = Instance.new("UIStroke", mobileBtn)
mobileBtnStroke.Color = BORDER2
mobileBtnStroke.Thickness = 1

local function setToggleVisual(on)
    TweenService:Create(pillBg,TweenInfo.new(0.18),{BackgroundColor3=on and WHITE or Color3.fromRGB(16, 16, 16)}):Play()
    TweenService:Create(pStroke,TweenInfo.new(0.18),{Color=on and WHITE or BORDER2}):Play()
    TweenService:Create(dot,TweenInfo.new(0.18,Enum.EasingStyle.Back),{
        Position=on and UDim2.new(1,-15,0.5,-6.5) or UDim2.new(0,2,0.5,-6.5),
        BackgroundColor3=on and BG or Color3.fromRGB(45, 45, 45)
    }):Play()
    
    -- Update mobile button appearance
    TweenService:Create(mobileBtn,TweenInfo.new(0.18),{BackgroundColor3=on and WHITE or KB_BG}):Play()
    TweenService:Create(mobileBtnStroke,TweenInfo.new(0.18),{Color=on and WHITE or BORDER2}):Play()
    mobileBtn.TextColor3 = on and BG or WHITE
end

local function updateMode()
    if State.mobileMode then
        toggleLbl.Visible = false
        keyBtn.Visible = false
        pillBg.Visible = false
        mobileBtn.Visible = true
        modeToggleBtn.Text = "PC"
    else
        toggleLbl.Visible = true
        keyBtn.Visible = true
        pillBg.Visible = true
        mobileBtn.Visible = false
        modeToggleBtn.Text = "MOBILE"
    end
end

updateMode()

-- Mode toggle button click
modeToggleBtn.MouseButton1Click:Connect(function()
    State.mobileMode = not State.mobileMode
    updateMode()
    saveConfig()
    
    TweenService:Create(modeToggleStroke, TweenInfo.new(0.15), {Color = WHITE}):Play()
    task.wait(0.15)
    TweenService:Create(modeToggleStroke, TweenInfo.new(0.15), {Color = BORDER2}):Play()
end)

-- Mobile button click - toggles on/off
mobileBtn.MouseButton1Click:Connect(function()
    State.autoBatToggled = not State.autoBatToggled
    setToggleVisual(State.autoBatToggled)
end)

local toggleClk = Instance.new("TextButton",toggleRow)
toggleClk.Size=UDim2.new(1,0,1,0); toggleClk.BackgroundTransparency=1; toggleClk.Text=""; toggleClk.ZIndex=6
toggleClk.MouseButton1Click:Connect(function()
    if State.mobileMode then return end -- Don't trigger if in mobile mode
    State.autoBatToggled=not State.autoBatToggled
    setToggleVisual(State.autoBatToggled)
end)
toggleClk.MouseEnter:Connect(function() 
    if not State.mobileMode then
        TweenService:Create(toggleRow,TweenInfo.new(0.1),{BackgroundColor3=CARD_HOV}):Play() 
    end
end)
toggleClk.MouseLeave:Connect(function() 
    if not State.mobileMode then
        TweenService:Create(toggleRow,TweenInfo.new(0.1),{BackgroundColor3=CARD_BG}):Play() 
    end
end)

-- // BAT LOGIC
local function getBat()
    local char=LP.Character; if not char then return nil end
    local tool=char:FindFirstChild("Bat"); if tool then return tool end
    local bp2=LP:FindFirstChild("Backpack")
    if bp2 then tool=bp2:FindFirstChild("Bat"); if tool then tool.Parent=char; return tool end end
    return nil
end

local function tryHitBat()
    if State.hittingCooldown then return end; State.hittingCooldown=true
    pcall(function()
        local bat=getBat(); if bat then
            bat:Activate(); local ev=bat:FindFirstChildWhichIsA("RemoteEvent")
            if ev then ev:FireServer() end
        end
    end)
    task.delay(0.08, function() State.hittingCooldown=false end)
end

local function getClosestPlayer()
    if not hrp then return nil,math.huge end
    local cp,cd=nil,math.huge
    for _,p in pairs(Players:GetPlayers()) do
        if p~=LP and p.Character then
            local tr=p.Character:FindFirstChild("HumanoidRootPart")
            if tr then local d=(hrp.Position-tr.Position).Magnitude; if d<cd then cd=d; cp=p end end
        end
    end
    return cp,cd
end

-- // CHARACTER SETUP
local function setupChar(char)
    task.wait(0.1)
    h=char:WaitForChild("Humanoid",5); hrp=char:WaitForChild("HumanoidRootPart",5)
    if not h or not hrp then return end
end

LP.CharacterAdded:Connect(setupChar)
if LP.Character then task.spawn(function() setupChar(LP.Character) end) end

-- // BAT AIMBOT HEARTBEAT
RunService.Heartbeat:Connect(function()
    if not (State.autoBatToggled and h and hrp) then return end
    local target,dist=getClosestPlayer()
    if target and target.Character then
        local tr=target.Character:FindFirstChild("HumanoidRootPart")
        if tr then
            if sethiddenproperty then
                sethiddenproperty(hrp, "PhysicsRepRootPart", tr)
            end
            local targetPos = tr.Position + Vector3.new(0, 0.9, 0)
            if (hrp.Position - targetPos).Magnitude > 8 then
                hrp.CFrame = CFrame.new(targetPos)
            end
            local cam = workspace.CurrentCamera
            cam.CFrame = CFrame.new(cam.CFrame.Position, tr.Position)
            tryHitBat()
        end
    end
end)

-- // KEYBIND HANDLER (only active when NOT in mobile mode)
UIS.InputBegan:Connect(function(inp, gp)
    if State.mobileMode then return end -- Skip if in mobile mode
    if gp then return end
    
    -- Check if input matches configured type and key
    if Keys.autoBatType == "Keyboard" and inp.UserInputType == Enum.UserInputType.Keyboard then
        if inp.KeyCode == Keys.autoBat then
            State.autoBatToggled = not State.autoBatToggled
            setToggleVisual(State.autoBatToggled)
        elseif inp.KeyCode == Keys.guiHide then
            State.guiVisible = not State.guiVisible
            main.Visible = State.guiVisible
        end
    elseif Keys.autoBatType == "Gamepad" and inp.UserInputType == Enum.UserInputType.Gamepad1 then
        if inp.KeyCode == Keys.autoBat then
            State.autoBatToggled = not State.autoBatToggled
            setToggleVisual(State.autoBatToggled)
        end
    end
end)

print("[Envy Auto Bat Desync] Loaded!")
