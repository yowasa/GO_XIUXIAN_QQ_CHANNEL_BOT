local json = require("json")
--创建角色时体质低于80则重新roll一个80-100的随机数
function inCreate(u)
    t = json.decode(u)
    tizhi=t["User"]["TiZhi"]
    math.randomseed(os.time())
    if tizhi<80
    then
        t["User"]["TiZhi"]=math.random(80,100)
    end
    str_json = json.encode(t)
    return str_json
end

--计算属性时 血量*1.1倍
function inCalAttr(u)
    t = json.decode(u)
    t["BattleInfo"]["HP"]=math.ceil(t["BattleInfo"]["HP"]*1.1)
    str_json = json.encode(t)
    return str_json
end