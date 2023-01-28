import * as React from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/material/styles';
import { Button, Card, CardActions, CardContent, Typography } from '@mui/material';
import { alignProperty } from '@mui/material/styles/cssUtils';

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
