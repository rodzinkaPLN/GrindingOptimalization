import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import React, { useEffect, useState } from 'react';
import UnstyledSelectsMultiple from './StyledMultipicker';

import './App.css';
import Chart from './Chart';
import DTPicker from './DatePicker';

function daysFromNow(days) {
  const copy = new Date()
  copy.setDate(copy.getDate() + days)
  return copy
}

function App() {
  const [data, setData] = useState([]);
  const [pickedParams, setPickedParams] = useState([])
  const [fromDate, setFromDate] = useState(daysFromNow(-607))
  const [toDate, setToDate] = useState(daysFromNow(-606))

  useEffect(() => {
    const dataFetch = async () => {
      const data = await (
        await fetch(
          `http://localhost:8080/api/v1/data?dataset=probne&from=${fromDate.toISOString()}&to=${toDate.toISOString()}`
        )
      ).json();

      setData(data);
    };

    dataFetch();
  }, [fromDate, toDate]);
  return (
    <>
      <AppBar position="fixed">
        <Toolbar sx={{ padding: "1em" }} className="AppToolbar">
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