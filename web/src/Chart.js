import { Typography } from "@mui/material";
import CircularProgress from "@mui/material/CircularProgress";
import React, { useState } from 'react';
import {
    Line, LineChart, ResponsiveContainer, Tooltip, XAxis,
    YAxis
} from 'recharts';

String.prototype.toProperCase = function () {
    return this.replace(/\w\S*/g, function (txt) { return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase(); });
};

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
                                        <p className="label"><b>{`> ${key.toProperCase()}: ${value}`}</b></p> :
                                        <p className="label">{`${key.toProperCase()}: ${value}`}</p>
                            })
                    }
                </div >
            );
        }

        return null;
    };


    if (props == undefined || props.data.parameters == undefined || props?.data?.length == 0) {
        return <CircularProgress />
    }


    return (
        <>
            <Typography variant='h5'>Dane procesowe</Typography>
            {
                props.data.parameters.
                    filter(v => props.pickedParams.includes(v.key)).
                    map(param => (
                        <>
                            <Typography >{param.key.toProperCase()}</Typography>
                            <ResponsiveContainer width="100%" height={height / 4}>
                                <LineChart
                                    data={props.data.data.filter(d => param.key in d)}
                                    syncId="anyId"
                                    onMouseEnter={() => {
                                        setLastHovered(param.key)
                                    }}
                                >
                                    <XAxis
                                        dataKey="ts"
                                        name="Data"
                                        interval={props.data.data.length / 20}
                                    />
                                    <YAxis />
                                    <Line type="monotone" dataKey={param.key} stroke="#8884d8" />
                                    {lastHovered == param.key &&
                                        <Tooltip content={<CustomTooltip key={param.key} />} />}
                                </LineChart>
                            </ResponsiveContainer>
                        </>
                    ))
            }
        </>)
}

export default Chart;