const express = require('express');
const app = express();
const port = 3000;

// 模拟用户数据
const users = [];

// 获取当前用户信息
app.get('/user/me', (req, res) => {
    // 假设这里有身份验证中间件，可以获取当前用户信息
    const currentUser = {
        id: 1,
        username: 'example_user',
        nickname: 'Example User',
        // 其他用户信息...
    };
    res.json(currentUser);
});

// 用户登出
app.get('/user/logout', (req, res) => {
    // 假设这里有登出逻辑...
    res.send('User logged out');
});

// 用户注册
app.post('/user/register', (req, res) => {
    const newUser = {
        id: users.length + 1,
        created_at: new Date(),
        updated_at: null,
        deleted_at: null,
        ...req.body // 假设请求体包含了需要的用户信息
    };

    users.push(newUser);
    res.json(newUser);
});

app.post('/api/v1/user/login', (req, res) => {
    // 假设这里有登录逻辑...
    res.send({
        code: 0,
        data: {
            nickname: 'example',
            role: "merchant",
        }
    })
});
const tableData = [
    {
        id: 0,
        producer: "ABC Pharmaceuticals",
        batchNumber: "B12345",
        productionDate: "2023-05-15T08:00:00",
        expirationDate: "2025-05-15T08:00:00",
        drugComposition: "Component A, Component B",
        drugFormula: "Formula X",
        qualityControlAndTestResults: "Pass",
        flow: "Production",
        transportationMethod: "Truck",
        transportationConditions: "Refrigerated",
        dealer: "XYZ Distributors",
        drugStorageConditions: "Cool and dry",
        drugStorageLocation: "Warehouse A",
        drugAcceptanceTime: "2023-06-01T10:00:00",
        productionDateStr: "[1,2, 3]",
        transportationStr: "[1,2,3]"
    },
    {
        id: 2,
        producer: "DEF Pharmaceuticals",
        batchNumber: "B54321",
        productionDate: "2023-07-20T08:00:00",
        expirationDate: "2025-07-20T08:00:00",
        drugComposition: "Component C, Component D",
        drugFormula: "Formula Y",
        qualityControlAndTestResults: "Pass",
        flow: "Production",
        transportationMethod: "Air",
        transportationConditions: "Controlled temperature",
        dealer: "LMN Distributors",
        drugStorageConditions: "Dry",
        drugStorageLocation: "Warehouse B",
        drugAcceptanceTime: "2023-08-10T09:30:00",
        productionDateStr: "[1, 2, 3]",
        transportationStr: "[1,2,3]"
    },
    // Add more data as needed
];

app.post('/api/v1/drug/list', (req, res) => {
    // 假设这里有登录逻辑...
    res.send({
        code: 0,
        data: {
            data: tableData
        }
    })
});
app.post('/api/v1/drug/buy/list', (req, res) => {
    // 假设这里有登录逻辑...
    res.send({
        code: 0,
        data: {
            data: tableData
        }
    })
});

// 获取用户列表
app.get('/users', (req, res) => {
    res.json(users);
});

// 更新当前用户信息
app.post('/user/update/me', (req, res) => {
    // 假设这里有更新当前用户信息的逻辑...
    res.send('User information updated');
});

// 更新用户信息
app.post('/user/update', (req, res) => {
    // 假设这里有更新用户信息的逻辑...
    res.send('User information updated');
});

// 用户交易处理
app.post('/user/dealer', (req, res) => {
    // 假设这里有用户交易处理的逻辑...
    res.send('User dealer');
});

// 添加药物
app.post('/drug/add', (req, res) => {
    // 假设这里有添加药物的逻辑...
    res.send('Drug added');
});

// 获取药物列表
app.get('/drug/list', (req, res) => {
    // 假设这里有获取药物列表的逻辑...
    res.send('List of drugs');
});

// 销售药物
app.post('/drug/sale', (req, res) => {
    // 假设这里有销售药物的逻辑...
    res.send('Drug sale');
});

// 药物交易处理
app.post('/drug/dealer', (req, res) => {
    // 假设这里有药物交易处理的逻辑...
    res.send('Drug dealer');
});

// 药物交易接受
app.post('/drug/accept', (req, res) => {
    // 假设这里有药物交易接受的逻辑...
    res.send('Drug accepted');
});

// 购买药物
app.post('/drug/buy', (req, res) => {
    // 假设这里有购买药物的逻辑...
    res.send('Drug bought');
});

// 获取购买药物列表
app.get('/drug/buy/list', (req, res) => {
    // 假设这里有获取购买药物列表的逻辑...
    res.send('List of drugs to buy');
});

// 启动服务器
app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});
