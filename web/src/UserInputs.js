import { Button, Card, CardActions, CardContent, CircularProgress, FormControlLabel, Grid, Slider, Switch, TextField, Typography } from '@mui/material';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/material/styles';
import * as React from 'react';


const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));

const IOSSwitch = styled((props) => (
    <Switch focusVisibleClassName=".Mui-focusVisible" disableRipple {...props} />
))(({ theme }) => ({
    width: 42,
    height: 26,
    padding: 0,
    '& .MuiSwitch-switchBase': {
        padding: 0,
        margin: 2,
        transitionDuration: '300ms',
        '&.Mui-checked': {
            transform: 'translateX(16px)',
            color: '#fff',
            '& + .MuiSwitch-track': {
                backgroundColor: theme.palette.mode === 'dark' ? '#2ECA45' : '#65C466',
                opacity: 1,
                border: 0,
            },
            '&.Mui-disabled + .MuiSwitch-track': {
                opacity: 0.5,
            },
        },
        '&.Mui-focusVisible .MuiSwitch-thumb': {
            color: '#33cf4d',
            border: '6px solid #fff',
        },
        '&.Mui-disabled .MuiSwitch-thumb': {
            color:
                theme.palette.mode === 'light'
                    ? theme.palette.grey[100]
                    : theme.palette.grey[600],
        },
        '&.Mui-disabled + .MuiSwitch-track': {
            opacity: theme.palette.mode === 'light' ? 0.7 : 0.3,
        },
    },
    '& .MuiSwitch-thumb': {
        boxSizing: 'border-box',
        width: 22,
        height: 22,
    },
    '& .MuiSwitch-track': {
        borderRadius: 26 / 2,
        backgroundColor: theme.palette.mode === 'light' ? '#E9E9EA' : '#39393D',
        opacity: 1,
        transition: theme.transitions.create(['background-color'], {
            duration: 500,
        }),
    },
}));
const SwitchXD = () => {
    return (
        <FormControlLabel
            control={<IOSSwitch
                sx={{ m: 1 }}
                defaultChecked={false} />}
            label="Automatyczna aplikacja"
            disabled
        />)
}

const SingleInput = (props) => {
    const [val, setVal] = React.useState(props.input.default_value)
    const v = props.input;
    React.useEffect(() => {
        console.log(v.name)
        if (v.name == 'woda') {
            props.setter(v => ({
                ...v,
                ['water']: val,
            }))
        } else if (v.name == 'ruda') {
            props.setter(v => ({
                ...v,
                ['cu']: val,
            }))
        } else {
            props.setter(v => ({
                ...v,
                ['speed']: val,
            }))
        }
    }, [val])
    return <Card raised sx={{ padding: '1em' }}>
        <CardContent>
            <Stack spacing={2} direction='row'>
                <TextField
                    fullWidth
                    id=""
                    label={v.name}
                    type="number"
                    disabled={props.disabled}
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
                <Button
                    onClick={() => { alert("TODO: dodać obsługe sterownika") }}
                    variant='outlined'
                    disabled={props.disabled}
                >
                    APLIKUJ
                </Button>
            </Stack>
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
                disabled={props.disabled}
                valueLabelDisplay="auto"
                step={v.step}
                min={v.min}
                max={v.max}
            />
        </CardActions>
    </Card>
}


const Statz = (props) => {
    if (props.usage == undefined) {
        return <CircularProgress />
    }

    return <Stack spacing={2}>
        <Card raised>
            <Typography variant="h6" > Energia </Typography>
            <Button> {props.usage.energy}</Button>
        </Card>
        <Card raised>
            <Typography variant="h6" > Granulacja</Typography>
            <Button> {props.usage.granulate}</Button>
        </Card>
    </Stack>
}
export default function UserInputs(props) {
    const [data, setData] = React.useState({
        cu: 800,
        speed: 900,
        water: 30,
    })

    const [usage, setUsage] = React.useState(undefined);
    React.useEffect(() => {
        const dataFetch = async () => {
            setUsage(undefined)
            const newData = await (
                await fetch(
                    `http://localhost:8080/api/v1/data/usage?cu=${data.cu}&speed=${data.speed}&water=${data.water}`
                )
            ).json();

            setUsage(newData.data);
            console.log(newData.data)
        };
        dataFetch();
    }, [data])
    if (props.inputs == undefined) {
        return <CircularProgress />
    }
    return (
        <Grid container columns={16} spacing={2}>
            <Grid item xs={16}>
                <Typography variant='h5'>Parametry</ Typography> <SwitchXD />
            </Grid>
            <Grid item xs={12}>
                <Stack spacing={2}>
                    {props.inputs.
                        map((v) => <SingleInput input={v} setter={setData} disabled={usage == undefined} />)}
                </Stack >
            </Grid>
            <Grid item xs={4}>
                <Statz usage={usage} />
            </Grid>
        </Grid>
    );
}

