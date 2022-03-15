import React, { useState } from 'react';
import Typography from '@mui/material/Typography';
import FormControl from '@mui/material/FormControl';
import Button from '@mui/material/Button';

import { ApiRequest } from '../../utils/apiRequest';

import { StyledPaper, StyledTextField } from './styled';

function Login(): React.ReactElement {
    const [username, setUsername] = useState<string>('')
    const [password, setPassword] = useState<string>('')

    const handleChange = (evt: any, setState: any) => {
        const input = evt.target.value || '';
        setState(input);
    }

    const handleLogin = (evt: any) => {
        if(username && password) {
            const loginData = {
                credentials: {
                    username,
                    password,
                }
            };
    
            ApiRequest('/login', (response) => { console.log(response) }, {
                method: 'POST',
                body: JSON.stringify(loginData),
                headers: {
                    'Content-Type': 'application/json'
                }
            });
        }
    }

    return (
        <>
            <StyledPaper>
                <Typography variant="h5" component="div">
                    Login
                </Typography>
                <FormControl fullWidth sx={{ m: 1 }} variant="standard">
                    <StyledTextField 
                        label="Username"
                        variant="standard"
                        value={username}
                        onChange={(evt) => handleChange(evt, setUsername)}
                        
                    />
                    <StyledTextField 
                        label="Password"
                        variant="standard"
                        value={password}
                        type="password"
                        onChange={(evt) => handleChange(evt, setPassword)}
                    />
                    <Button variant="outlined" onClick={handleLogin}>Login</Button>
                </FormControl>


            </StyledPaper>
        </>
    );
}

export default Login;
