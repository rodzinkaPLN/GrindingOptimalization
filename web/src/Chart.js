import CircularProgress from "@mui/material/CircularProgress";
import React, { useEffect, useState } from 'react';
import {
    Line, LineChart, Tooltip, XAxis,
    YAxis
} from 'recharts';

const Chart = () => {
    const [data, setData] = useState([]);
    const { innerWidth: width, innerHeight: height } = window;
    const [lastHovered, setLastHovered] = useState("")


    const CustomTooltip = ({ active, payload, label }) => {
        if (active && payload && payload.length) {
            return (
                <div className="custom-tooltip">
                    {Object.entries(payload[0].payload).map(([key, value]) => {
                        return key == lastHovered ?
                            <p className="label"><b>{`> ${key} : ${value}`}</b></p> :
                            <p className="label">{`${key} : ${value}`}</p>
                    })}
                </div >
            );
        }

        return null;
    };

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
                        data={data.data.filter(d => param.key in d)}
                        syncId="anyId"
                        onMouseEnter={() => {
                            setLastHovered(param.key)
                        }}
                    >
                        <XAxis dataKey="ts" name="Data" />
                        <YAxis />
                        <Line type="monotone" dataKey={param.key} stroke="#8884d8" activeDot={{ r: 2 }} />
                        {lastHovered == param.key &&
                            <Tooltip content={<CustomTooltip key={param.key} />} />}
                    </LineChart>
                ))
            }
        </div>)
}

export default Chart;