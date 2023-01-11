local json = require("json")

--50岁之前大部分属性下降10% 50岁之后大部分属性增加10%
function usePill(content)
    c = json.decode(content)
    u=c["User"]
    i=c["Item"]
    get_exp=50
    if i["Attr"]=="中品" then
        get_exp=200
    elseif i["Attr"]=="上品" then
        get_exp=1000
    elseif i["Attr"]=="极品" then
        get_exp=5000
    end
    u["User"]["BaseInfo"]["NowExp"]=u["User"]["BaseInfo"]["NowExp"]+get_exp
    c["Result"]=true
    c["Msg"]="你服用了"..i["ItemName"]..",修为大涨"
    str_json = json.encode(c)
    return str_json
end