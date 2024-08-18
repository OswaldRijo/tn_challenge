const express = require('express');
const path = require('path');
const cors = require('cors');

const app = express();
app.use(cors());

//////////////////////////////////////////// STATICS & RESOURCES //////////////////////////////////////////////////////////
app.use('/static', express.static(path.join(__dirname, '..', 'build', 'static')));

// Maneja todas las demás rutas enviando el archivo index.html
app.get('/manifest.json', (req, res) => {
  res.sendFile(path.join(__dirname, '..', 'build', 'manifest.json'));
});

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////   APP   ///////////////////////////////////////////////////////////
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, '..', 'build', 'index.html'));
});


// Inicia el servidor
const PORT = process.env.REACT_APP_PORT || 3000;
app.listen(PORT, () => {
  console.log(`Servidor Express escuchando en el puerto ${PORT}`);
});
