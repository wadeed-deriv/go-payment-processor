const express = require('express');

const jsonRouter = express.Router();

// Middleware to parse JSON requests
jsonRouter.use(express.json());

// JSON route
jsonRouter.post('/deposit', (req, res) => {
    const { amount, clientID } = req.body;
    // Process deposit
    console.log(`Deposited ${amount} for client ${clientID}`);
    res.json({ message: `Deposited ${amount} successfully` });
});

jsonRouter.post('/withdrawal', (req, res) => {
    const { amount, clientID } = req.body;
    // Process withdrawal
    console.log(`Withdrawn ${amount} for client ${clientID}`);
    res.json({ message: `Withdrew ${amount} successfully` });
});

module.exports = jsonRouter;