import { Button, Card, CardActions, CardContent, CircularProgress, Slider, Typography } from '@mui/material';
import Stack from '@mui/material/Stack';
import * as React from 'react';

const SingleInput = (props) => {
    const [val, setVal] = React.useState(props.input.defaultValue)
    const v = props.input;
    return <Card raised>
        <CardContent>
            <Typography textAlign='left'>{`Typ  : ${v.name}`}</Typography>
        </CardContent>
        <CardActions sx={{
            justifyContent: 'center',
        }}>
            <Slider
                defaultValue={v.default_value}
                aria-label="Default"
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

