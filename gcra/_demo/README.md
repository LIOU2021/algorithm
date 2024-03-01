# 情境
- 模拟动能侦测与判断
- 参考 [奇摩新闻](https://tw.news.yahoo.com/)
- 根据关键字访问可以有两个level
    - 上升
    - 急升
- 在此情境中可以发挥gcra弹性配置的特色，于是撰写此范例
# 运行
1. go run main.go
2. 访问 http://127.0.0.1:8080?q=2330
    - q: 查询参数，随意输入

# 测试
- cli
    ```bash
    go test -v .
    ```
- output
    ```bash
   main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 上升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 上升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 上升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 上升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 上升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 急升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 急升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 急升
    main_test.go:38: 2024-03-01 18:06:05, query: 2330, status: 急升
    main_test.go:18: sleep 2s ...
    main_test.go:38: 2024-03-01 18:06:07, query: 2330, status: 上升
    main_test.go:38: 2024-03-01 18:06:07, query: 2330, status: 上升
    main_test.go:38: 2024-03-01 18:06:07, query: 2330, status: 急升
    main_test.go:38: 2024-03-01 18:06:07, query: 2330, status: 急升
    main_test.go:38: 2024-03-01 18:06:07, query: 2330, status: 急升
    ```