# git-typo

检查github仓库的typos，不要用来水PR哦😌  

## Qucikstart

### backend

```shell
cd backend
go mod tiny
go run typo
```

### frontend

```shell
cd frontend
npm install
npm run dev
```

Then, open http://localhost:3000/  


## TODO
- config
- mutex
- 移除后端，用内存模拟文件系统
    - https://github.com/streamich/memfs  

![image-20241026150846450](./assets/image-20241026150846450.png)