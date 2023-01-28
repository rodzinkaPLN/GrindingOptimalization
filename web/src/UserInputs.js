import { Button, Card, CardActions, CardContent, CircularProgress, Slider, TextField, Typography } from '@mui/material';
import Stack from '@mui/material/Stack';
import * as React from 'react';

const SingleInput = (props) => {
    const [val, setVal] = React.useState(props.input.default_value)
    const v = props.input;
    return <Card raised sx={{ padding: '1em' }}>
        <CardContent>
            <TextField
                fullWidth
                id=""
                label={v.name}
                type="number"
                onChange={(e) => {
                    setVal(parseFloat(e.target.value));
                }}
                value={val}
                onBlur={() => {
                    if (val > v.max) {
                        setVal(v.max)
                    } else if (val < v.min) {
                        setVal(v.min);
                    }
                }}
            />
        </CardContent>
        <CardActions sx={{
            justifyContent: 'center',
        }}>
            <Slider
                key={`slider-${v.name}`}
                value={val}
                onChange={(_, vx) => {
                    setVal(vx);
                }}
                valueLabelDisplay="auto"
                step={v.step}
                min={v.min}
                max={v.max}
            />
        </CardActions>
    </Card>
}


export default function UserInputs(props) {
    if (props.inputs == undefined) {
        return <CircularProgress />
    }
    return (
        <Stack spacing={2}>
            <Typography variant='h5'>Parametry</ Typography>
            {props.inputs.
                map((v) => <SingleInput input={v} />)}
        </Stack >
    );
}

