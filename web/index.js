const electron = require('electron');
const {app} = electron;

const {BrowserWindow} = require('electron')

app.on('ready', function() {
  var mainWindow = new BrowserWindow({
    title: 'Spring Hackathon 2019',
    backgroundColor: '#000',
    width: 800,
    height: 600,
    minWidth: 800,
    minHeight: 600
  })
  mainWindow.loadURL('file://' + __dirname + '/index.html')
});