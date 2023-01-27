import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import React, { useEffect, useState } from 'react';
import UnstyledSelectsMultiple from './StyledMultipicker';

import './App.css';
import Chart from './Chart';
import DTPicker from './DatePicker';

function App() {
  const [data, setData] = useState([]);
  const [pickedParams, setPickedParams] = useState([])
  const [fromDate, setFromDate] = useState(Date.now())
  const [toDate, setToDate] = useState(Date.now() - 2)

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
          <DTPicker date={fromDate} setDate={setFromDate} label="od" />
          <DTPicker date={toDate} setDate={setToDate} label="do" />
          <UnstyledSelectsMultiple
            parameters={data?.parameters}
            pickedParams={pickedParams}
            setPickedParams={setPickedParams}
          />
        </Toolbar>
      </AppBar>
      <div className="App">
        <Chart
          data={data}
          pickedParams={pickedParams}
        />
      </div>
    </>
  );
}

export default App;
