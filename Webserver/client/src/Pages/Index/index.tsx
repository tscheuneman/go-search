import React, { useEffect, useState } from 'react';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Button from '@mui/material/Button';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CircularProgress from '@mui/material/CircularProgress';

import {
    useParams,
    Link
  } from "react-router-dom";

import { ConfigValues } from './types';
import { TextEditor } from '../../Components/TextEditor';

import { ApiRequest } from '../../utils/apiRequest';

function Index(): React.ReactElement {
    const [searchableFields, setSearchableFields] = useState<string[]>([]);
    const [displayFields, setDisplayFields] = useState<string[]>([]);
    const [filterableFields, setFilterableFields] = useState<string[]>([]);
    const [sortableFields, setSortableFields] = useState<string[]>([]);
    const [loading, setLoading] = useState<boolean>(false);

    const [searches, setSearches] = useState([]);

    const { id: indexId } = useParams();
    useEffect(() => {
        ApiRequest(`/admin/index/${indexId}/configure/globals`, (response) => {
            const searchableResponse: string[] = response[ConfigValues.SEARCH_CONFIG] || ["*"];
            const displayResponse: string[] = response[ConfigValues.DISPLAY_CONFIG] || ["*"];
            const filterableResponse: string[] = response[ConfigValues.FILTERABLE_CONFIG] || [];
            const sortableResponse: string[] = response[ConfigValues.SORTABLE_CONFIG] || [];

            setSearchableFields(searchableResponse)
            setDisplayFields(displayResponse)
            setFilterableFields(filterableResponse);
            setSortableFields(sortableResponse);
        });

        ApiRequest(`/admin/index/${indexId}/configure/search`, (response) => {
            setSearches(response || [])
        });
    }, []);

    const setItems = (event: React.ChangeEvent<HTMLTextAreaElement>, mutator: React.Dispatch<React.SetStateAction<any>>) => {
        const value = event.target.value;
        mutator(value.split('\n'));
    }

    const saveConfig = () => {
        const saveRequest = {
            config: {
                displayed_fields: displayFields,
                searchable_fields: searchableFields,
                filterable_fields: filterableFields,
                sortable_fields: sortableFields,
            }
        };

        ApiRequest(`/admin/index/${indexId}/configure/globals`, () => {}, {
            method: 'POST',
            body: JSON.stringify(saveRequest),
            headers: {
                'Content-Type': 'application/json'
            }
        });
    };

    const handleFileUpload = (evt: any) => {
        setLoading(true);
        const target = evt.target;
        if(target?.files?.[0]) {
            const fileReader = new FileReader();

            fileReader.addEventListener("load", () => {
                const result = fileReader?.result || '';
                const JsonResult = JSON.parse(result as string);
                ApiRequest(`/admin/index/${indexId}/document`, () => {
                    setLoading(false);
                }, {
                    method: 'POST',
                    body: JSON.stringify({ documents: JsonResult }),
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });
            }, false);

            fileReader.readAsText(target?.files?.[0]);
        }
    }

    return (
        <>
            <Grid style={{ marginBottom: '20px' }} container spacing={2}>
                <Grid item xs={6}>
                    <Typography variant="h5" component="div">
                        Configure { indexId }
                    </Typography>
                </Grid>
                <Grid style={{
                    display: 'flex',
                    justifyContent: 'end'
                }} 
                item xs={6}>
                    {
                        loading
                        ?
                            <CircularProgress size={20} />
                        :
                            <Button
                            variant="contained"
                            component="label"
                            >
                                Upload File
                                <input
                                    accept="application/json"
                                    type="file"
                                    onChange={handleFileUpload}
                                    hidden
                                />
                            </Button>
                    }

                </Grid>
            </Grid>

            <hr />
            <Grid style={{ marginBottom: '20px' }} container spacing={2}>
                <Grid item xs={6}>
                    <TextEditor setItems={setItems} mutator={setDisplayFields} state={displayFields} heading='Displayed Fields' />
                </Grid>
                <Grid item xs={6}>
                    <TextEditor setItems={setItems} mutator={setSearchableFields} state={searchableFields} heading='Searchable Fields' />
                </Grid>
            </Grid>

            <Grid style={{ marginBottom: '20px' }} container spacing={2}>
                <Grid item xs={6}>
                    <TextEditor setItems={setItems} mutator={setFilterableFields} state={filterableFields} heading='Filterable Fields' />
                </Grid>
                <Grid item xs={6}>
                    <TextEditor setItems={setItems} mutator={setSortableFields} state={sortableFields} heading='Sortable Fields' />
                </Grid>
            </Grid>
            <hr />
            <br />
            <Button onClick={saveConfig} variant="contained">Save</Button>
            <br />
            <br />
            <hr />
            <br />
            <Typography variant="h5" component="div">
                Search Indexes
            </Typography>
            <Grid style={{ marginBottom: '20px' }} container spacing={2}>
                <Grid item xs={12}>
                    {
                        searches.map((search: any) => {
                            return(
                                <Link style={{
                                    textDecoration: 'none'
                                }} to={`/index/${indexId}/search/${search?.Slug}`}>
                                <Card style={{
                                    marginTop: '20px'
                                }}>
                                    <CardContent>
                                        <Typography variant="h5" component="div">
                                             { search?.Slug || "Name" }
                                        </Typography>
                                        <Typography variant="body2" component="div">
                                            Last Updated: { new Date(search?.UpdatedAt).toLocaleDateString() }
                                        </Typography>
                                    </CardContent>
                                </Card>
                            </Link>
                            )
                        })
                    }
                </Grid>
            </Grid>
        </>
    );
}

export default Index;
