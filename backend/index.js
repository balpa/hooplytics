require('dotenv').config();

const express = require('express');
const app = express();
const PORT = process.env.PORT || 8000;

const helloWorldRouter = require('./routes/api/helloWorld');

app.use('/api', helloWorldRouter);

app.listen(PORT, () => {
    console.log(`Server running on port ${ PORT }`);
});