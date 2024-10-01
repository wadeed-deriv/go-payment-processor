const express = require('express');
const jsonRouter = require('./jsonRouter');
const xmlRouter  = require('./xmlRouter');


const app = express();
const port = 3000;

app.use("/json",jsonRouter);
app.use("/xml",xmlRouter);


app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});