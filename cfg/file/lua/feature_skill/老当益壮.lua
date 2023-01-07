

--属性随着年龄增长而增长
function inCalAttr(u)
    t = json.decode(u)
    print(t["User"])
    str_json = json.encode(u)
    return str_json
end