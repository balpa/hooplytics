const helloWorldService = require('../services/helloWorldService');

const helloWorld = (req, res) => {
    const message = helloWorldService.getHelloWorldMessage();

    res.send(message);
};

module.exports = {
    helloWorld
};