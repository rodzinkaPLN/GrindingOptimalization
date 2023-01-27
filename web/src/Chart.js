import CircularProgress from "@mui/material/CircularProgress";
import React, { useEffect, useState } from 'react';
import {
    Line, LineChart, Tooltip, XAxis,
    YAxis
} from 'recharts';

const Chart = (props) => {
    const { innerWidth: width, innerHeight: height } = window;
    const [lastHovered, setLastHovered] = useState("")

    const CustomTooltip = ({ active, payload, label }) => {
        if (active && payload && payload.length) {
            const date = new Date(payload[0].payload.Data)
            return (
                < div className="custom-tooltip" >
                    <p className="label"><b>{` Data: ${date.toLocaleDateString("pl")} ${date.toLocaleTimeString("pl")}`}</b></p>
                    {
                        Object.entries(payload[0].payload).
                            map(([key, value]) => {
                                if (props.pickedParams.includes(key))
                                    return key == lastHovered ?
                                        <p className="label"><b>{`> ${key}: ${value} ${payload[0].payload}`}</b></p> :
                                        <p className="label">{`${key}: ${value}`}</p>
                            })
                    }
                </div >
            );
        }

        return null;
    };


    if (props?.parameters == undefined || props?.data?.length == 0) {
        return <CircularProgress />
    }


    return (
        <div>
            {
                props.data.parameters.
                    filter(v => props.pickedParams.includes(v.key)).
                    map(param => (
                        <LineChart
                            width={width * 0.9}
                            height={height / 4}
                            data={props.data.data.filter(d => param.key in d)}
                            syncId="anyId"
                            onMouseEnter={() => {
                                setLastHovered(param.key)
                            }}
                        >
                            <XAxis dataKey="ts" name="Data" interval={props.data.data.length / 20} />
                            <YAxis />
                            <Line type="monotone" dataKey={param.key} stroke="#8884d8" />
                            {lastHovered == param.key &&
                                <Tooltip content={<CustomTooltip key={param.key} />} />}
                        </LineChart>
                    ))
            }
        </div>)
}

export default Chart;