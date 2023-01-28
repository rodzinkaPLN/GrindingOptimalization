import { Button, Card, CardActions, CardContent, CircularProgress, Slider, Typography } from '@mui/material';
import Stack from '@mui/material/Stack';
import * as React from 'react';

export default function UserInputs(props) {
    if (props.inputs == undefined) {
        return <CircularProgress />
    }
    return (
        <Stack spacing={2}>
            <Typography variant='h5'>Parametry</ Typography>
            {Object.entries(props.inputs).
                map(([key, value]) =>
                    <Card raised>
                        <CardContent>
                            <Typography textAlign='left'>{`Typ  : ${key}`}</Typography>
                        </CardContent>
                        <CardActions sx={{
                            justifyContent: 'center',
                        }}>
                            <Slider
                                defaultValue={50}
                                aria-label="Default"
                                valueLabelDisplay="auto"

                            />
                        </CardActions>
                    </Card>
                )
            }
        </Stack >
    );
}


