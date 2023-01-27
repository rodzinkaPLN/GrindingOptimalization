import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import './App.css';
import Chart from './Chart';

function App() {
  return (
    <>
      <AppBar position="static">
        <Toolbar>
          {/* <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
            >
            <MenuIcon />
          </IconButton> */}
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Admin panel
          </Typography>
        </Toolbar>
      </AppBar>
      <div className="App">
        <Chart />
      </div>
    </>
  );
}

export default App;
