const { Router } = require('express');
const xmlparser = require('express-xml-bodyparser');
const xmlRouter = Router();

xmlRouter.use(xmlparser());

// SOAP/XML route
xmlRouter.post('/deposit', (req, res) => {
    console.log(req.body);
    const { amount, clientid } = req.body.request;
    console.log(`Deposited ${amount[0]} for client ${clientid[0]}`);
    // Process deposit
    res.set('Content-Type', 'application/xml');
    res.send(`<response><message>Deposited ${amount[0]} successfully</message></response>`);
});

xmlRouter.post('/withdrawal', (req, res) => {
    const { amount, clientid } = req.body.request;
    console.log(`Withdrawn ${amount[0]} for client ${clientid[0]}`);
    // Process withdrawal
    res.set('Content-Type', 'application/xml');
    res.send(`<response><message>Withdrew ${amount[0]} successfully</message></response>`);
});

module.exports = xmlRouter;