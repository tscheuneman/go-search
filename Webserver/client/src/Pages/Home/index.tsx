import React, { useEffect, useState } from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import {
    Link
  } from "react-router-dom";

import { ApiRequest } from '../../utils/apiRequest';

interface Index {
    uid: string;
    createdAt: string;
    updatedAt: string;
    primaryKey: string;
}

function Home(): React.ReactElement {
    const [indexes, setIndexes] = useState<Index[]>([]);

    useEffect(() => {
        ApiRequest('/admin/index', (response) => {
            setIndexes(response);
        });
    }, []);

    return (
        <>
            <Typography variant="h4" component="div">
                Indexes
            </Typography>
            <br />
            <Grid container spacing={2}>
            {
                indexes.map(index => (
                    <Grid item xs={6}>
                        <Link style={{
                            textDecoration: 'none'
                        }} to={`/index/${index?.uid}`}>
                            <Card>
                                <CardContent>
                                    <Typography variant="h5" component="div">
                                        { index?.uid || "Name" }
                                    </Typography>
                                    <Typography variant="body2" component="div">
                                        Last Updated: { new Date(index?.updatedAt).toLocaleDateString() }
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Link>
                    </Grid>
                ))
            }
            </Grid>
        </>
    );
}

export default Home;
