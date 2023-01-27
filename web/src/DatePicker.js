import TextField from '@mui/material/TextField';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import * as React from 'react';

export default function DTPicker(props) {

    return (
        <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DateTimePicker
                renderInput={(props) => <TextField {...props} />}
                label={props.label}
                value={props.date}
                onChange={(newValue) => {
                    props.setDate(newValue);
                }}
            />
        </LocalizationProvider>
    );
}