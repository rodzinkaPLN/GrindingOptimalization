import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import React, { useEffect, useState } from 'react';
import UnstyledSelectsMultiple from './StyledMultipicker';

import './App.css';
import Chart from './Chart';

function App() {
  const [data, setData] = useState([]);
  const [pickedParams, setPickedParams] = useState([])
  useEffect(() => {
    const dataFetch = async () => {
      const data = await (
        await fetch(
          "http://localhost:8080/api/v1/data?dataset=probne"
        )
      ).json();

      setData(data);
    };

    dataFetch();
  }, []);
  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Admin panel
          </Typography>
          <UnstyledSelectsMultiple
            parameters={data?.parameters}
            pickedParams={pickedParams}
            setPickedParams={setPickedParams}
          />
        </Toolbar>
      </AppBar>
      <div className="App">
        <Chart data={data} />
      </div>
    </>
  );
}

export default App;
