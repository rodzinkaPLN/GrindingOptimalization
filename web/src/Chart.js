import { Typography } from "@mui/material";
import CircularProgress from "@mui/material/CircularProgress";
import React, { useState } from 'react';
import {
    Line, LineChart, ResponsiveContainer, Tooltip, XAxis,
    YAxis
} from 'recharts';
import UnstyledSelectsMultiple from "./StyledMultipicker";

String.prototype.toProperCase = function () {
    return this.replace(/\w\S*/g, function (txt) { return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase(); });
};

const Chart = (props) => {
    const { innerWidth: width, innerHeight: height } = window;
    const [lastHovered, setLastHovered] = useState("")
    const [pickedParams, setPickedParams] = useState([])

    const CustomTooltip = ({ active, payload, _, keyX }) => {
        if (active && payload && payload.length && lastHovered == keyX) {
            const date = new Date(payload[0].payload.Data)
            return (
                < div className="custom-tooltip" >
                    <p className="label"><b>{` Data: ${date.toLocaleDateString("pl")} ${date.toLocaleTimeString("pl")}`}</b></p>
                    {
                        Object.entries(payload[0].payload).
                            map(([key, value]) => {
                                if (pickedParams.includes(key))
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
            <Typography variant='h5'>{props.title}
            </Typography>
            <UnstyledSelectsMultiple
                parameters={props.data?.parameters}
                pickedParams={pickedParams}
                setPickedParams={setPickedParams}
            />
            {
                props.data.parameters.
                    filter(v => pickedParams.includes(v.key)).
                    map(param => (
                        <>
                            <Typography variant='h6' >{param.key.toProperCase()}</Typography>
                            <ResponsiveContainer width="100%" height={height / 4}>
                                <LineChart
                                    data={props.data.data.filter(d => param.key in d)}
                                    syncId={props.id}
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
                                    <Tooltip content={<CustomTooltip keyX={param.key} />} />
                                </LineChart>
                            </ResponsiveContainer>
                        </>
                    ))
            }
        </>)
}

export default Chart;