import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import React, { useEffect, useState } from 'react';
import UnstyledSelectsMultiple from './StyledMultipicker';

import './App.css';
import Chart from './Chart';
import DTPicker from './DatePicker';
import { createTheme, Grid, Paper, ThemeProvider } from '@mui/material';
import Recommendations from './Recommendations';
import UserPara from './UserInputs';
import UserInputs from './UserInputs';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

function daysFromNow(days) {
  const copy = new Date()
  copy.setDate(copy.getDate() + days)
  return copy
}

function App() {
  const [data, setData] = useState([]);
  const [predictedData, setPredictedData] = useState([])
  const [fromDate, setFromDate] = useState(daysFromNow(-607))
  const [toDate, setToDate] = useState(daysFromNow(-606))
  const [recommendations, setRecommendations] = useState(undefined)
  const [userInputs, setUserInputs] = useState(undefined);

  useEffect(() => {
    const dataFetch = async () => {
      const newData = await (
        await fetch(
          `http://localhost:8080/api/v1/data?dataset=probne&from=${fromDate.toISOString()}&to=${toDate.toISOString()}`
        )
      ).json();

      setData(newData);
    };

    dataFetch();
  }, [fromDate, toDate]);

  useEffect(() => {
    const dataFetch = async () => {
      const newData = await (
        await fetch(
          `http://localhost:8080/api/v1/data?dataset=predykaty&from=${fromDate.toISOString()}&to=${toDate.toISOString()}`
        )
      ).json();

      setPredictedData(newData);
    };

    dataFetch();
  }, [fromDate, toDate]);

  useEffect(() => {
    const dataFetch = async () => {
      const newData = await (
        await fetch(
          `http://localhost:8080/api/v1/data/predictions`
        )
      ).json();

      setRecommendations(newData);
    };

    dataFetch();
  }, []);

  useEffect(() => {
    const dataFetch = async () => {
      const newData = await (
        await fetch(
          `http://localhost:8080/api/v1/userinputs`
        )
      ).json();

      setUserInputs(newData.data);
    };

    dataFetch();
  }, []);
  return (

    <ThemeProvider theme={darkTheme}>

      <AppBar position="fixed">
        < Toolbar sx={{ padding: "1em" }
        } className="AppToolbar" >
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Admin panel
          </Typography>
          <DTPicker date={fromDate} setDate={setFromDate} label="od" />
          <DTPicker date={toDate} setDate={setToDate} label="do" />
        </Toolbar >
      </AppBar >
      <div className="App" >

        <Grid container spacing={6} columns={16} paddingTop="5em" paddingLeft="3em">
          <Grid item xs={8} >
            <Paper className="componentWrapper">
              <Chart
                id="xd2"
                title="Dane procesowe"
                data={data}
              />
            </Paper>
          </Grid>
          <Grid item xs={7}>
            <Paper className="componentWrapper">

              <Recommendations
                recommendations={recommendations?.data} />
              <UserInputs
                inputs={userInputs}
              />
            </Paper>

            <Paper className="componentWrapper">
              <Chart
                id="xd2"
                title="Dane spredykowane"
                data={predictedData}
              />
            </Paper>
          </Grid>
        </Grid>
      </div>
    </ThemeProvider >
  );
}

export default App;