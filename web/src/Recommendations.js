import * as React from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/material/styles';
import { Typography } from '@mui/material';

const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));

export default function Recommendations(props) {
    return (
        <Stack spacing={2}>
            <Typography variant='h5'>Rekomendacje</Typography>
            {Object.entries(props.predictions).
                map(([key, value]) => <p className="label"><b>{`${key}: ${value}`}</b></p>
                )
            }
        </Stack>
    );
}
