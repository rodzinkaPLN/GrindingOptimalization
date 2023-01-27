import CircularProgress from "@mui/material/CircularProgress";
import React, { useEffect, useState } from 'react';
import {
    Area, AreaChart, Brush, CartesianGrid, Line, LineChart, ResponsiveContainer, Tooltip, XAxis,
    YAxis
} from 'recharts';

const Chart = () => {
    const [data, setData] = useState([]);
    const { innerWidth: width, innerHeight: height } = window;

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
        console.log(data)
    }, []);


    if (data.length == 0) {
        return <CircularProgress />
    }
    return (
        <div>
            {
                data.parameters.map(param => (
                    <LineChart
                        width={width / 2}
                        height={height / 4}
                        data={data.data}
                        syncId="anyId"
                    >
                        <XAxis dataKey="ts" />
                        <YAxis />
                        <Tooltip />

                        <Line type="monotone" dataKey={param.key} stroke="#8884d8" activeDot={{ r: 8 }} />
                    </LineChart>
                ))
            }
        </div>)
}

export default Chart;