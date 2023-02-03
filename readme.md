# 系统服务
> 获取系统服务信息 支持 windows 和 linux


## 内置方法
> [summary](#汇总信息) = vela.service(cnd) <br />
>  cnd: 过滤条件[service](#服务信息)  返回: 汇总信息
```lua
    local s = vela.service()
    print(s[1])
    print(s.size)
    s.pipe(print)
```


## 汇总信息
> 存储[service](#服务信息)的集合

内置字段:
- size &emsp; 内容大小

内置方法:
- pipe(pipe) &emsp;遍历

```lua
    local s = vela.service("name re *ntp*")
    print(s.size)
    s.pipe(print)
```

## 服务信息
> 封装服务相关信息

相关字段:
- name
- start_type
- exec_path
- display
- description
- state
- pid
- exit_code

```lua
    local s = vela.service()
    s.pipe(function(v)
        print(name)
        print(start_type)
        print(exec_path)
        print(display)
        print(description)
        print(state)
        print(pid)
        print(exit_code)
    end)
```