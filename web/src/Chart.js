import React from 'react';
import {
    CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis,
    YAxis
} from 'recharts';
import { useState, useEffect } from 'react';
import CircularProgress from "@mui/material/CircularProgress"
import randomColor from "randomcolor";

const Chart = () => {
    const [data, setData] = useState([]);
    const { innerWidth: width, innerHeight: height } = window;

    useEffect(() => {
        const dataFetch = async () => {
            const data = await (
                await fetch(
                    "http://localhost:8080/api/v1/data?parameters=drgania&parameters=gestosc"
                )
            ).json();

            setData(data);
        };

        dataFetch();
        console.log(data)
    }, []);


    if (data.length == 0) {
        return <CircularProgress />
    }
    return (
        <LineChart width={width / 2} height={height / 2} data={data.data}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="ts" padding={{ left: 30, right: 30 }} />
            <YAxis />
            <Tooltip />
            <Legend />
            {
                data.parameters.map(d => (
                    <Line type="natural" dataKey={d.key} name={d.name} stroke={randomColor()} activeDot={{ r: 8 }} />
                ))
            }
        </LineChart>
    );
}

export default Chart;