import * as React from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/material/styles';
import { Button, Card, CardActions, CardContent, FormControlLabel, Switch, Typography } from '@mui/material';
import { alignProperty } from '@mui/material/styles/cssUtils';

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

export default function Recommendations(props) {
    return (
        <Stack spacing={2}>
            <Typography variant='h5'>Rekomendacje</ Typography> <SwitchXD />
            {Object.entries(props.predictions).
                map(([key, value]) =>
                    <Card>
                        <CardContent>
                            <Typography textAlign='left'>{`Typ  : ${key}`}</Typography>
                            <Typography textAlign='left'>{`Wartość: ${value}`}</Typography>
                        </CardContent>
                        <CardActions sx={{
                            justifyContent: 'center',
                        }}>
                            <Button
                                onClick={() => { alert("TODO: dodać obsługe sterownika") }}
                                variant='outlined'
                            >

                                APLIKUJ
                            </Button>
                        </CardActions>
                    </Card>
                )
            }
        </Stack >
    );
}


