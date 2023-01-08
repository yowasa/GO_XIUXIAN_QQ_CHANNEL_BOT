local json = require("json")
--创建角色时敏捷低于60则重新roll一个60-100的随机数
function inCreate(u)
    t = json.decode(u)
    minjie=t["User"]["MinJie"]
    math.randomseed(os.time())
    if minjie<60
    then
        t["User"]["MinJie"]=math.random(60,100)
    end
    str_json = json.encode(t)
    return str_json
end

--计算属性时 速度*1.1倍
function inCalAttr(u)
    t = json.decode(u)
    t["BattleInfo"]["SPD"]=math.ceil(t["BattleInfo"]["SPD"]*1.1)
    str_json = json.encode(t)
    return str_json
end