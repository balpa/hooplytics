const express = require('express');
const router = express.Router();
const helloWorldController = require('../../controllers/helloWorldController');

router.get('/helloWorld', helloWorldController.helloWorld);

module.exports = router;